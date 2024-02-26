// Copyright (C) 2022 IOTech Ltd

package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pelletier/go-toml/v2"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

// CentralGetRequest makes the get request and return the body
func CentralGetRequest(ctx context.Context, returnValuePointer interface{}, baseUrl string, requestPath string, requestParams url.Values, authInjector interfaces.AuthenticationInjector) errors.EdgeX {
	req, edgexErr := createRequest(ctx, http.MethodGet, baseUrl, requestPath, requestParams)
	if edgexErr != nil {
		return errors.NewCommonEdgeXWrapper(edgexErr)
	}

	res, edgexErr := sendRequest(ctx, req, authInjector)
	if edgexErr != nil {
		return errors.NewCommonEdgeXWrapper(edgexErr)
	}
	// Check the response content length to avoid json unmarshal error
	if len(res) == 0 {
		return nil
	}

	var err error
	if FromContext(ctx, common.ContextKeyContentType) == common.ContentTypeTOML {
		err = toml.Unmarshal(res, returnValuePointer)
	} else {
		err = json.Unmarshal(res, returnValuePointer)
	}

	if err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}
