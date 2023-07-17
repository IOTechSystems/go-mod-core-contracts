//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

func (g *generalClient) XpertFetchConfiguration(ctx context.Context) (res dtoCommon.ConfigResponse, err errors.EdgeX) {
	err = utils.XpertGetRequest(ctx, &res, g.baseUrl, common.ApiConfigRoute, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return res, nil
}
