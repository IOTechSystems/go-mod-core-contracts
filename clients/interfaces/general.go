//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

type GeneralClient interface {
	// FetchConfiguration obtains configuration information from the target service.
	FetchConfiguration(ctx context.Context) (common.ConfigResponse, errors.EdgeX)

	// Central
	// FetchMetrics obtains metrics information from the target service.
	FetchMetrics(ctx context.Context) (common.MetricsResponse, errors.EdgeX)
	// XpertFetchConfiguration obtains configuration information from the target service.
	// In comparison with FetchConfiguration, this function supports both JSON and TOML formats.
	XpertFetchConfiguration(ctx context.Context) (common.ConfigResponse, errors.EdgeX)
}
