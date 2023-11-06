//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type eventClient struct {
	baseUrl string
}

// NewEventClient creates an instance of EventClient
func NewEventClient(baseUrl string) interfaces.EventClient {
	return &eventClient{
		baseUrl: baseUrl,
	}
}

func (ec *eventClient) Add(ctx context.Context, req requests.AddEventRequest) (
	dtoCommon.BaseWithIdResponse, errors.EdgeX) {
	path := utils.EscapeAndJoinPath(common.ApiEventRoute, req.Event.ProfileName, req.Event.DeviceName, req.Event.SourceName)
	var br dtoCommon.BaseWithIdResponse

	bytes, encoding, err := req.Encode()
	if err != nil {
		return br, errors.NewCommonEdgeXWrapper(err)
	}

	err = utils.PostRequest(ctx, &br, ec.baseUrl, path, bytes, encoding)
	if err != nil {
		return br, errors.NewCommonEdgeXWrapper(err)
	}
	return br, nil
}

func (ec *eventClient) AllEvents(ctx context.Context, offset, limit int) (responses.MultiEventsResponse, errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiEventsResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, common.ApiAllEventRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventCount(ctx context.Context) (dtoCommon.CountResponse, errors.EdgeX) {
	res := dtoCommon.CountResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, common.ApiEventCountRoute, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventCountByDeviceName(ctx context.Context, name string) (dtoCommon.CountResponse, errors.EdgeX) {
	requestPath := utils.EscapeAndJoinPath(common.ApiEventCountRoute, common.Device, common.Name, name)
	res := dtoCommon.CountResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, requestPath, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventsByDeviceName(ctx context.Context, name string, offset, limit int) (
	responses.MultiEventsResponse, errors.EdgeX) {
	requestPath := utils.EscapeAndJoinPath(common.ApiEventRoute, common.Device, common.Name, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiEventsResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DeleteByDeviceName(ctx context.Context, name string) (dtoCommon.BaseResponse, errors.EdgeX) {
	path := utils.EscapeAndJoinPath(common.ApiEventRoute, common.Device, common.Name, name)
	res := dtoCommon.BaseResponse{}
	err := utils.DeleteRequest(ctx, &res, ec.baseUrl, path)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventsByTimeRange(ctx context.Context, start, end int64, offset, limit int) (
	responses.MultiEventsResponse, errors.EdgeX) {
	requestPath := utils.EscapeAndJoinPath(common.ApiEventRoute, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiEventsResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DeleteByAge(ctx context.Context, age int) (dtoCommon.BaseResponse, errors.EdgeX) {
	path := utils.EscapeAndJoinPath(common.ApiEventRoute, common.Age, strconv.Itoa(age))
	res := dtoCommon.BaseResponse{}
	err := utils.DeleteRequest(ctx, &res, ec.baseUrl, path)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}
