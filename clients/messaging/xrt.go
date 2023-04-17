// Copyright (C) 2023 IOTech Ltd

package messaging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/messaging/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/xrtmodels"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	clientName      = "xrt-go-client"
	connWaitTimeout = 5 * time.Second
)

// xrtClient implements the client of MQTT management API, https://docs.iotechsys.com/edge-xrt21/mqtt-management/mqtt-management.html
type xrtClient struct {
	lc              logger.LoggingClient
	requestMap      utils.RequestMap
	mqttClient      mqtt.Client
	requestTopic    string
	replyTopic      string
	responseTimeout time.Duration
	qos             byte

	clientOptions *ClientOptions
}

type ClientOptions struct {
	*CommandOptions
	*DiscoveryOptions
	*StatusOptions
}

// CommandOptions provides the config for sending the request to manage components
type CommandOptions struct {
	CommandTopic string
}

// DiscoveryOptions provides the config for sending the discovery request like discovery:trigger, device:scan
type DiscoveryOptions struct {
	DiscoveryTopic          string
	DiscoveryMessageHandler mqtt.MessageHandler
	DiscoveryDuration       time.Duration
	DiscoveryTimeout        time.Duration
}

// StatusOptions provides the config for subscribing the XRT status
type StatusOptions struct {
	StatusTopic          string
	StatusMessageHandler mqtt.MessageHandler
}

func NewXrtClient(opts *mqtt.ClientOptions, requestTopic string, replyTopic string, qos byte,
	responseTimeout time.Duration, lc logger.LoggingClient, clientOptions *ClientOptions) (interfaces.XrtClient, errors.EdgeX) {
	client := &xrtClient{
		lc:              lc,
		requestMap:      utils.NewRequestMap(),
		requestTopic:    requestTopic,
		replyTopic:      replyTopic,
		responseTimeout: responseTimeout,
		qos:             qos,
		clientOptions:   clientOptions,
	}

	opts.OnConnect = OnConnectHandler(client, clientOptions, lc)
	mqttClient := mqtt.NewClient(opts)

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
			token := mqttClient.Connect()
			if token.WaitTimeout(connWaitTimeout) && token.Error() != nil {
				lc.Warnf("failed to connect MQTT Broker: %v, retry it again...", token.Error())
				time.Sleep(retryInterval)
				continue
			}
			break FOR
		}
	}

	client.mqttClient = mqttClient

	return client, nil
}

func NewClientOptions(commandOptions *CommandOptions, discoveryOptions *DiscoveryOptions, statusOptions *StatusOptions) *ClientOptions {
	return &ClientOptions{
		CommandOptions:   commandOptions,
		DiscoveryOptions: discoveryOptions,
		StatusOptions:    statusOptions,
	}
}

func NewCommandOptions(commandTopic string) *CommandOptions {
	return &CommandOptions{
		CommandTopic: commandTopic,
	}
}

func NewDiscoveryOptions(discoveryTopic string, discoveryMessageHandler mqtt.MessageHandler, discoveryDuration, discoveryTimeout time.Duration) *DiscoveryOptions {
	return &DiscoveryOptions{
		DiscoveryTopic:          discoveryTopic,
		DiscoveryMessageHandler: discoveryMessageHandler,
		DiscoveryDuration:       discoveryDuration,
		DiscoveryTimeout:        discoveryTimeout,
	}
}

func NewStatusOptions(statusTopic string, statusMessageHandler mqtt.MessageHandler) *StatusOptions {
	return &StatusOptions{
		StatusTopic:          statusTopic,
		StatusMessageHandler: statusMessageHandler,
	}
}

func (c *xrtClient) SetResponseTimeout(responseTimeout time.Duration) {
	c.responseTimeout = responseTimeout
}

// sendXrtRequest sends general request to XRT
func (c *xrtClient) sendXrtRequest(ctx context.Context, requestId string, request interface{}, response interface{}) errors.EdgeX {
	return c.sendXrtRequestWithTimeout(ctx, c.requestTopic, requestId, request, response, c.responseTimeout)
}

// sendXrtDiscoveryRequest sends discovery request to XRT
func (c *xrtClient) sendXrtDiscoveryRequest(ctx context.Context, requestId string, request interface{}, response interface{}) errors.EdgeX {
	if c.clientOptions == nil || c.clientOptions.DiscoveryOptions == nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "please provide DiscoveryOptions for the discovery request", nil)
	}
	timeout := time.Duration(c.responseTimeout.Nanoseconds() + c.clientOptions.DiscoveryDuration.Nanoseconds() + c.clientOptions.DiscoveryTimeout.Nanoseconds())
	return c.sendXrtRequestWithTimeout(ctx, c.requestTopic, requestId, request, response, timeout)
}

// sendXrtCommandRequest sends command request to XRT
func (c *xrtClient) sendXrtCommandRequest(ctx context.Context, requestId string, request interface{}, response interface{}) errors.EdgeX {
	if c.clientOptions == nil || c.clientOptions.CommandOptions == nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "please provide CommandOptions for the command request", nil)
	}
	return c.sendXrtRequestWithTimeout(ctx, c.clientOptions.CommandOptions.CommandTopic, requestId, request, response, c.responseTimeout)
}

func (c *xrtClient) sendXrtRequestWithTimeout(ctx context.Context, requestTopic string, requestId string, request interface{}, response interface{}, responseTimeout time.Duration) errors.EdgeX {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	// Before publishing the request, we should create responseChan to receive the response from XRT
	c.requestMap.Add(requestId)
	token := c.mqttClient.Publish(requestTopic, c.qos, false, jsonData)
	select {
	case <-ctx.Done():
		return nil
	case <-token.Done():
		if token.Error() != nil {
			return errors.NewCommonEdgeXWrapper(token.Error())
		}
	}

	cmdResponseBytes, err := utils.FetchXRTResponse(ctx, requestId, c.requestMap, responseTimeout)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	err = json.Unmarshal(cmdResponseBytes, response)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindServerError, "failed to JSON decoding command response: %v", err)
	}

	// handle error result from the XRT
	var commonResponse xrtmodels.CommonResponse
	err = json.Unmarshal(cmdResponseBytes, &commonResponse)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindServerError, "failed to JSON decoding command response: %v", err)
	}
	if commonResponse.Result.Error() != nil {
		return errors.NewCommonEdgeXWrapper(commonResponse.Result.Error())
	}
	return nil
}

func OnConnectHandler(xrtClient *xrtClient, clientOptions *ClientOptions, lc logger.LoggingClient) mqtt.OnConnectHandler {
	return func(client mqtt.Client) {
		subscriptions := createSubscriptions(xrtClient, clientOptions)
		for topic, handler := range subscriptions {
			if topic != "" && handler != nil {
				token := client.Subscribe(topic, xrtClient.qos, handler)
				if token.WaitTimeout(connWaitTimeout) && token.Error() != nil {
					lc.Errorf("failed to subscribe XRT reply topic: %s", token.Error())
					return
				}
				lc.Infof("Subscribed XRT reply topic %s.", topic)
			}
		}
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

func createSubscriptions(xrtClient *xrtClient, clientOptions *ClientOptions) map[string]mqtt.MessageHandler {
	subscriptions := make(map[string]mqtt.MessageHandler)
	subscriptions[xrtClient.replyTopic] = commandReplyHandler(xrtClient.requestMap, xrtClient.lc)
	if clientOptions == nil {
		return subscriptions
	}
	if clientOptions.DiscoveryOptions != nil {
		if clientOptions.DiscoveryOptions.DiscoveryTopic != "" && clientOptions.DiscoveryOptions.DiscoveryMessageHandler != nil {
			subscriptions[clientOptions.DiscoveryOptions.DiscoveryTopic] = clientOptions.DiscoveryOptions.DiscoveryMessageHandler
		}
	}
	if clientOptions.StatusOptions != nil {
		if clientOptions.StatusOptions.StatusTopic != "" && clientOptions.StatusOptions.StatusMessageHandler != nil {
			subscriptions[clientOptions.StatusOptions.StatusTopic] = clientOptions.StatusOptions.StatusMessageHandler
		}
	}
	return subscriptions
}

func (c *xrtClient) Close() errors.EdgeX {
	if c.mqttClient.IsConnected() {
		c.lc.Debug("Disconnect the MQTT conn")
		c.mqttClient.Disconnect(5000)
	}
	return nil
}
