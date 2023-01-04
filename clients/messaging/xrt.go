// Copyright (C) 2023 IOTech Ltd

package messaging

import (
	"context"
	"encoding/json"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/messaging/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/xrtmodels"
)

const (
	clientName      = "xrt-go-client"
	connWaitTimeout = 5 * time.Second
)

type xrtClient struct {
	lc              logger.LoggingClient
	requestMap      utils.RequestMap
	mqttClient      mqtt.Client
	requestTopic    string
	replyTopic      string
	responseTimeout time.Duration
	qos             byte
}

func NewXrtClient(opts *mqtt.ClientOptions, requestTopic string, replyTopic string, qos byte, responseTimeout time.Duration, lc logger.LoggingClient) (interfaces.XrtClient, errors.EdgeX) {
	requestMap := utils.NewRequestMap()
	opts.OnConnect = OnConnectHandler(replyTopic, qos, requestMap, lc)
	client := mqtt.NewClient(opts)

	connectTimeout := opts.ConnectTimeout
	if connectTimeout == 0 {
		connectTimeout = time.Minute
	}
	retryInterval := opts.ConnectRetryInterval
	if retryInterval == 0 {
		retryInterval = time.Second
	}

	timeout := time.After(connectTimeout)
FOR:
	for {
		select {
		case <-timeout:
			return nil, errors.NewCommonEdgeX(errors.KindServerError, "timed out connecting MQTT broker", nil)
		default:
			token := client.Connect()
			if token.WaitTimeout(connWaitTimeout) && token.Error() != nil {
				lc.Warnf("failed to connect MQTT Broker: %v, retry it again...", token.Error())
				time.Sleep(retryInterval)
				continue
			}
			break FOR
		}
	}

	res := &xrtClient{
		lc:              lc,
		mqttClient:      client,
		requestMap:      requestMap,
		requestTopic:    requestTopic,
		replyTopic:      replyTopic,
		responseTimeout: responseTimeout,
	}

	return res, nil
}

func (c *xrtClient) DeviceByName(ctx context.Context, name string) (xrtmodels.DeviceInfo, errors.EdgeX) {
	request := xrtmodels.NewDeviceGetRequest(name, clientName)
	var response xrtmodels.DeviceResponse

	err := c.sendXrtRequest(ctx, request.RequestId, request, &response)
	if err != nil {
		return xrtmodels.DeviceInfo{}, errors.NewCommonEdgeX(errors.KindServerError, "failed to query device", err)
	}
	if response.Result.Error() != nil {
		return xrtmodels.DeviceInfo{}, errors.NewCommonEdgeXWrapper(response.Result.Error())
	}

	return response.Result.Device, nil
}

func (c *xrtClient) DeviceProfileByName(ctx context.Context, name string) (models.DeviceProfile, errors.EdgeX) {
	request := xrtmodels.NewProfileGetRequest(name, clientName)
	var response xrtmodels.ProfileResponse

	err := c.sendXrtRequest(ctx, request.RequestId, request, &response)
	if err != nil {
		return models.DeviceProfile{}, errors.NewCommonEdgeX(errors.KindServerError, "failed to query profile", err)
	}
	if response.Result.Error() != nil {
		return models.DeviceProfile{}, errors.NewCommonEdgeXWrapper(response.Result.Error())
	}

	return response.Result.Profile, nil
}

func (c *xrtClient) sendXrtRequest(ctx context.Context, requestId string, request interface{}, response interface{}) errors.EdgeX {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	// Before publishing the request, we should create responseChan to receive the response from XRT
	c.requestMap.Add(requestId)
	token := c.mqttClient.Publish(c.requestTopic, c.qos, false, jsonData)
	select {
	case <-ctx.Done():
		return nil
	case <-token.Done():
		if token.Error() != nil {
			return errors.NewCommonEdgeXWrapper(token.Error())
		}
	}

	cmdResponseBytes, err := utils.FetchXRTResponse(ctx, requestId, c.requestMap, c.responseTimeout)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	err = json.Unmarshal(cmdResponseBytes, response)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindServerError, "failed to JSON decoding command response: %v", err)
	}

	return nil
}

func OnConnectHandler(replyTopic string, qos byte, requestMap utils.RequestMap, lc logger.LoggingClient) mqtt.OnConnectHandler {
	return func(client mqtt.Client) {
		// Listen to the XRT reply topic
		token := client.Subscribe(replyTopic, qos, commandReplyHandler(requestMap, lc))
		if token.WaitTimeout(connWaitTimeout) && token.Error() != nil {
			lc.Errorf("failed to subscribe XRT reply topic: %s", token.Error())
			return
		}

		// TODO: add discovered device handler

		lc.Infof("Subscribed XRT reply topic.")
	}
}

func commandReplyHandler(requestMap utils.RequestMap, lc logger.LoggingClient) mqtt.MessageHandler {
	return func(client mqtt.Client, message mqtt.Message) {
		var response xrtmodels.BaseResponse
		err := json.Unmarshal(message.Payload(), &response)
		if err != nil {
			lc.Warnf("failed to parse XRT reply, topic: %s, message:%s, err: %v", message.Topic(), string(message.Payload()), err)
			return
		}
		resChan, ok := requestMap.Get(response.RequestId)
		if !ok {
			lc.Debugf("deprecated response from the XRT, it might be caused by timeout or unknown error, topic: %s, message:%s", message.Topic(), string(message.Payload()))
			return
		}

		resChan <- message.Payload()
	}
}
