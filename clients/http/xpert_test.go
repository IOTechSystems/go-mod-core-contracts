//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/require"
)

func Test_generalClient_XpertFetchConfiguration_JSON(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiConfigRoute, dtoCommon.ConfigResponse{})
	defer ts.Close()

	client := NewGeneralClient(ts.URL, NewNullAuthenticationInjector())
	res, err := client.XpertFetchConfiguration(context.Background())
	require.NoError(t, err)
	require.IsType(t, dtoCommon.ConfigResponse{}, res)
}

func Test_generalClient_XpertFetchConfiguration_TOML(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != common.ApiConfigRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		b, _ := toml.Marshal(dtoCommon.ConfigResponse{})
		_, _ = w.Write(b)
	}))
	defer ts.Close()

	client := NewGeneralClient(ts.URL, NewNullAuthenticationInjector())
	ctx := context.WithValue(context.Background(), common.ContextKeyContentType, common.ContentTypeTOML)
	res, err := client.XpertFetchConfiguration(ctx)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.ConfigResponse{}, res)
}
