// Copyright (C) 2022 IOTech Ltd

package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	"github.com/pelletier/go-toml/v2"
)

// EscapeAndJoinPath escape and join the path variables
func EscapeAndJoinPath(apiRoutePath string, pathVariables ...string) string {
	elements := make([]string, len(pathVariables)+1)
	elements[0] = apiRoutePath // we don't need to escape the route path like /device, /reading, ...,etc.
	for i, e := range pathVariables {
		elements[i+1] = url.QueryEscape(e)
	}
	return path.Join(elements...)
}

// edgeXClientReqURI returns the non-encoded path?query that would be used in an HTTP request for u.
func edgeXClientReqURI(u *url.URL) string {
	result := u.Scheme + "://" + u.Host + u.Path
	if u.ForceQuery || u.RawQuery != "" {
		result += "?" + u.RawQuery
	}
	return result
}

// XpertGetRequest makes the get request and return the body
func XpertGetRequest(ctx context.Context, returnValuePointer interface{}, baseUrl string, requestPath string, requestParams url.Values) errors.EdgeX {
	req, edgexErr := createRequest(ctx, http.MethodGet, baseUrl, requestPath, requestParams)
	if edgexErr != nil {
		return errors.NewCommonEdgeXWrapper(edgexErr)
	}

	res, edgexErr := sendRequest(ctx, req)
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
