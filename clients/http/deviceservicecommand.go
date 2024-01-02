//
// Copyright (C) 2021-2023 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	"github.com/fxamacker/cbor/v2"
)

type deviceServiceCommandClient struct {
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewDeviceServiceCommandClient creates an instance of deviceServiceCommandClient
func NewDeviceServiceCommandClient(authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DeviceServiceCommandClient {
	return &deviceServiceCommandClient{
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// GetCommand sends HTTP request to execute the Get command
func (client *deviceServiceCommandClient) GetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string) (*responses.EventResponse, errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Name).SetNameFieldPath(deviceName).SetNameFieldPath(commandName).BuildPath()
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return nil, errors.NewCommonEdgeXWrapper(err)
	}
	res, contentType, edgeXerr := utils.GetRequestAndReturnBinaryRes(ctx, baseUrl, requestPath, params, client.authInjector)
	if edgeXerr != nil {
		return nil, errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	// If execute GetCommand with dsReturnEvent query parameter 'no', there will be no content returned in the http response.
	// So we can use the nil pointer to indicate that the HTTP response content is empty
	if len(res) == 0 {
		return nil, nil
	}
	response := &responses.EventResponse{}
	if contentType == common.ContentTypeCBOR {
		if err = cbor.Unmarshal(res, response); err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to decode the cbor response", err)
		}
	} else {
		if err = json.Unmarshal(res, response); err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to decode the json response", err)
		}
	}
	return response, nil
}

// SetCommand sends HTTP request to execute the Set command
func (client *deviceServiceCommandClient) SetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Name).SetNameFieldPath(deviceName).SetNameFieldPath(commandName).BuildPath()
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PutRequest(ctx, &response, baseUrl, requestPath, params, settings, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

// SetCommandWithObject invokes device service's set command API and the settings supports object value type
func (client *deviceServiceCommandClient) SetCommandWithObject(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]interface{}) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Name).SetNameFieldPath(deviceName).SetNameFieldPath(commandName).BuildPath()
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PutRequest(ctx, &response, baseUrl, requestPath, params, settings, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCommandClient) Discovery(ctx context.Context, baseUrl string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	err := utils.PostRequest(ctx, &response, baseUrl, common.ApiDiscoveryRoute, nil, "")
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}
