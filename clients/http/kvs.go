//
// Copyright (C) 2023 IOTech Ltd
//

package http

import (
	"context"
	"net/url"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

// KVSClient is the REST client for invoking the key-value APIs(/kvs/*) from Core Keeper
type KVSClient struct {
	baseUrl string
}

// NewKVSClient creates an instance of KVSClient
func NewKVSClient(baseUrl string) interfaces.KVSClient {
	return &KVSClient{
		baseUrl: baseUrl,
	}
}

// UpdateValuesByKey updates values of the specified key and the child keys defined in the request payload.
// If no key exists at the given path, the key(s) will be created.
func (kc KVSClient) UpdateValuesByKey(ctx context.Context, key string, req requests.UpdateKeysRequest) (res responses.KeysResponse, err errors.EdgeX) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	queryParams := url.Values{}
	queryParams.Set(common.Flatten, common.ValueTrue)
	err = utils.PutRequest(ctx, &res, kc.baseUrl, path, queryParams, req)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ValuesByKey returns the values of the specified key prefix.
func (kc KVSClient) ValuesByKey(ctx context.Context, key string) (res responses.MultiKeyValueResponse, err errors.EdgeX) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	queryParams := url.Values{}
	queryParams.Set(common.Plaintext, common.ValueTrue)
	err = utils.GetRequest(ctx, &res, kc.baseUrl, path, queryParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ListKeys returns the list of the keys with the specified key prefix.
func (kc KVSClient) ListKeys(ctx context.Context, key string) (res responses.KeysResponse, err errors.EdgeX) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	queryParams := url.Values{}
	queryParams.Set(common.KeyOnly, common.ValueTrue)
	err = utils.GetRequest(ctx, &res, kc.baseUrl, path, queryParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteKey deletes the specified key.
func (kc KVSClient) DeleteKey(ctx context.Context, key string) (res responses.KeysResponse, err errors.EdgeX) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	err = utils.DeleteRequest(ctx, &res, kc.baseUrl, path)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteKeysByPrefix deletes all keys with the specified prefix.
func (kc KVSClient) DeleteKeysByPrefix(ctx context.Context, key string) (res responses.KeysResponse, err errors.EdgeX) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	queryParams := url.Values{}
	queryParams.Set("prefixMatch", common.ValueTrue)
	err = utils.DeleteRequestWithParams(ctx, &res, kc.baseUrl, path, queryParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}
