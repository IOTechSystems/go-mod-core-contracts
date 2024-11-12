//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"

type TokenResponse struct {
	common.BaseResponse `json:",inline"`
	JWT                 string `json:"jwt"`
}

// NewTokenResponse returns the JWT
func NewTokenResponse(requestId string, message string, statusCode int, token string) TokenResponse {
	return TokenResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		JWT:          token,
	}
}
