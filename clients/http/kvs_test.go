//
// Copyright (C) 2023 IOTech Ltd
//

package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"

	"github.com/stretchr/testify/require"
)

const TestKey = "TestWritable"

func TestUpdateValuesByKey(t *testing.T) {
	ts := newTestServer(http.MethodPut, common.ApiKVSRoute+"/"+common.Key+"/"+TestKey, responses.KeysResponse{})
	defer ts.Close()

	client := NewKVSClient(ts.URL)
	res, err := client.UpdateValuesByKey(context.Background(), TestKey, requests.UpdateKeysRequest{})

	require.NoError(t, err)
	require.IsType(t, responses.KeysResponse{}, res)
}

func TestValuesByKey(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiKVSRoute+"/"+common.Key+"/"+TestKey, responses.MultiKeyValueResponse{})
	defer ts.Close()

	client := NewKVSClient(ts.URL)
	res, err := client.ValuesByKey(context.Background(), TestKey)

	require.NoError(t, err)
	require.IsType(t, responses.MultiKeyValueResponse{}, res)
}

func TestListKeys(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiKVSRoute+"/"+common.Key+"/"+TestKey, responses.KeysResponse{})
	defer ts.Close()

	client := NewKVSClient(ts.URL)
	res, err := client.ListKeys(context.Background(), TestKey)

	require.NoError(t, err)
	require.IsType(t, responses.KeysResponse{}, res)
}

func TestDeleteKeys(t *testing.T) {
	ts := newTestServer(http.MethodDelete, common.ApiKVSRoute+"/"+common.Key+"/"+TestKey, responses.KeysResponse{})
	defer ts.Close()

	client := NewKVSClient(ts.URL)
	res, err := client.DeleteKey(context.Background(), TestKey)

	require.NoError(t, err)
	require.IsType(t, responses.KeysResponse{}, res)
}

func TestDeleteKeysByPrefix(t *testing.T) {
	ts := newTestServer(http.MethodDelete, common.ApiKVSRoute+"/"+common.Key+"/"+TestKey, responses.KeysResponse{})
	defer ts.Close()

	client := NewKVSClient(ts.URL)
	res, err := client.DeleteKeysByPrefix(context.Background(), TestKey)

	require.NoError(t, err)
	require.IsType(t, responses.KeysResponse{}, res)
}
