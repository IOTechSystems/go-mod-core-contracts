//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	commonDTO "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

// GetRequest makes the get request and return the body
func GetRequest(ctx context.Context, returnValuePointer interface{}, baseUrl string, requestPath string, requestParams url.Values) errors.EdgeX {
	req, err := createRequest(ctx, http.MethodGet, baseUrl, requestPath, requestParams)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	// Check the response content length to avoid json unmarshal error
	if len(res) == 0 {
		return nil
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// GetRequestAndReturnBinaryRes makes the get request and return the binary response and content type(i.e., application/json, application/cbor, ... )
func GetRequestAndReturnBinaryRes(ctx context.Context, baseUrl string, requestPath string, requestParams url.Values) (res []byte, contentType string, edgeXerr errors.EdgeX) {
	req, edgeXerr := createRequest(ctx, http.MethodGet, baseUrl, requestPath, requestParams)
	if edgeXerr != nil {
		return nil, "", errors.NewCommonEdgeXWrapper(edgeXerr)
	}

	resp, edgeXerr := makeRequest(req)
	if edgeXerr != nil {
		return nil, "", errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	defer resp.Body.Close()

	bodyBytes, edgeXerr := getBody(resp)
	if edgeXerr != nil {
		return nil, "", errors.NewCommonEdgeXWrapper(edgeXerr)
	}

	if resp.StatusCode <= http.StatusMultiStatus {
		return bodyBytes, resp.Header.Get(common.ContentType), nil
	}

	// Handle error response
	var errResponse commonDTO.BaseResponse
	e := json.Unmarshal(bodyBytes, &errResponse)
	if e != nil {
		return nil, "", errors.NewCommonEdgeX(errors.KindMapping(resp.StatusCode), string(bodyBytes), e)
	}

	return nil, "", errors.NewCommonEdgeX(errors.KindMapping(errResponse.StatusCode), errResponse.Message, nil)
}

// GetRequestWithBodyRawData makes the GET request with JSON raw data as request body and return the response
func GetRequestWithBodyRawData(ctx context.Context, returnValuePointer interface{}, baseUrl string, requestPath string, requestParams url.Values, data interface{}) errors.EdgeX {
	req, err := createRequestWithRawDataAndParams(ctx, http.MethodGet, baseUrl, requestPath, requestParams, data)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// PostRequest makes the post request with encoded data and return the body
func PostRequest(
	ctx context.Context,
	returnValuePointer interface{},
	baseUrl string, requestPath string,
	data []byte,
	encoding string) errors.EdgeX {

	req, err := createRequestWithEncodedData(ctx, http.MethodPost, baseUrl, requestPath, data, encoding)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// PostRequestWithRawData makes the post JSON request with raw data and return the body
func PostRequestWithRawData(
	ctx context.Context,
	returnValuePointer interface{},
	baseUrl string, requestPath string,
	requestParams url.Values,
	data interface{}) errors.EdgeX {

	req, err := createRequestWithRawData(ctx, http.MethodPost, baseUrl, requestPath, requestParams, data)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// PutRequest makes the put JSON request and return the body
func PutRequest(
	ctx context.Context,
	returnValuePointer interface{},
	baseUrl string, requestPath string,
	requestParams url.Values,
	data interface{}) errors.EdgeX {

	req, err := createRequestWithRawData(ctx, http.MethodPut, baseUrl, requestPath, requestParams, data)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// PatchRequest makes a PATCH request and unmarshals the response to the returnValuePointer
func PatchRequest(
	ctx context.Context,
	returnValuePointer interface{},
	baseUrl string, requestPath string,
	requestParams url.Values,
	data interface{}) errors.EdgeX {

	req, err := createRequestWithRawData(ctx, http.MethodPatch, baseUrl, requestPath, requestParams, data)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// PostByFileRequest makes the post file request and return the body
func PostByFileRequest(
	ctx context.Context,
	returnValuePointer interface{},
	baseUrl string, requestPath string,
	filePath string) errors.EdgeX {

	req, err := createRequestFromFilePath(ctx, http.MethodPost, baseUrl, requestPath, filePath)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// PutByFileRequest makes the put file request and return the body
func PutByFileRequest(
	ctx context.Context,
	returnValuePointer interface{},
	baseUrl string, requestPath string,
	filePath string) errors.EdgeX {

	req, err := createRequestFromFilePath(ctx, http.MethodPut, baseUrl, requestPath, filePath)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// DeleteRequest makes the delete request and return the body
func DeleteRequest(ctx context.Context, returnValuePointer interface{}, baseUrl string, requestPath string) errors.EdgeX {
	req, err := createRequest(ctx, http.MethodDelete, baseUrl, requestPath, nil)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// DeleteRequestWithParams makes the delete request with URL query params and return the body
func DeleteRequestWithParams(ctx context.Context, returnValuePointer interface{}, baseUrl string, requestPath string, requestParams url.Values) errors.EdgeX {
	req, err := createRequest(ctx, http.MethodDelete, baseUrl, requestPath, requestParams)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}
