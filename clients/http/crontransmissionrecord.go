//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

type CronTransmissionRecordClient struct {
	baseUrl               string
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewCronTransmissionRecordClient creates an instance of CronTransmissionRecordClient
func NewCronTransmissionRecordClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.CronTransmissionRecordClient {
	return &CronTransmissionRecordClient{
		baseUrl:               baseUrl,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// AllCronTransmissionRecords query cron transmission records with start, end, offset, and limit
func (client *CronTransmissionRecordClient) AllCronTransmissionRecords(ctx context.Context, start, end int64, offset, limit int) (res responses.MultiCronTransmissionRecordResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, common.ApiAllCronTransmissionRecordRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// CronTransmissionRecordsByStatus queries cron transmission records with status, start, end, offset, and limit
func (client *CronTransmissionRecordClient) CronTransmissionRecordsByStatus(ctx context.Context, status string, start, end int64, offset, limit int) (res responses.MultiCronTransmissionRecordResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiCronTransmissionRecordRoute, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// CronTransmissionRecordsByJobName queries cron transmission records with jobName, start, end, offset, and limit
func (client *CronTransmissionRecordClient) CronTransmissionRecordsByJobName(ctx context.Context, jobName string, start, end int64, offset, limit int) (res responses.MultiCronTransmissionRecordResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiCronTransmissionRecordRoute, common.Job, common.Name, jobName)
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// CronTransmissionRecordsByJobNameAndStatus queries cron transmission records with jobName, status, start, end, offset, and limit
func (client *CronTransmissionRecordClient) CronTransmissionRecordsByJobNameAndStatus(ctx context.Context, jobName, status string, start, end int64, offset, limit int) (res responses.MultiCronTransmissionRecordResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiCronTransmissionRecordRoute, common.Job, common.Name, jobName, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}
