//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"
)

func TestCronTransmissionRecordClient_AllCronTransmissionRecords(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllCronTransmissionRecordRoute, responses.MultiCronTransmissionRecordResponse{})
	defer ts.Close()
	client := NewCronTransmissionRecordClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllCronTransmissionRecords(context.Background(), 0, 0, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiCronTransmissionRecordResponse{}, res)
}

func TestCronTransmissionRecordClient_CronTransmissionRecordsByStatus(t *testing.T) {
	status := models.Succeeded
	urlPath := path.Join(common.ApiCronTransmissionRecordRoute, common.Status, status)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiCronTransmissionRecordResponse{})
	defer ts.Close()
	client := NewCronTransmissionRecordClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.CronTransmissionRecordsByStatus(context.Background(), status, 0, 0, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiCronTransmissionRecordResponse{}, res)
}

func TestCronTransmissionRecordClient_CronTransmissionRecordsByJobName(t *testing.T) {
	jobName := TestScheduleJobName
	urlPath := path.Join(common.ApiCronTransmissionRecordRoute, common.Job, common.Name, jobName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiCronTransmissionRecordResponse{})
	defer ts.Close()
	client := NewCronTransmissionRecordClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.CronTransmissionRecordsByJobName(context.Background(), jobName, 0, 0, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiCronTransmissionRecordResponse{}, res)
}

func TestCronTransmissionRecordClient_CronTransmissionRecordsByJobNameAndStatus(t *testing.T) {
	jobName := TestScheduleJobName
	status := models.Succeeded
	urlPath := path.Join(common.ApiCronTransmissionRecordRoute, common.Job, common.Name, jobName, common.Status, status)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiCronTransmissionRecordResponse{})
	defer ts.Close()
	client := NewCronTransmissionRecordClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.CronTransmissionRecordsByJobNameAndStatus(context.Background(), jobName, status, 0, 0, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiCronTransmissionRecordResponse{}, res)
}
