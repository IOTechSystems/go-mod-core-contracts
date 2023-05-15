//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

// RuleResponse defines the Response Content for GET rule DTO.
type RuleResponse struct {
	common.BaseResponse `json:",inline"`
	Rule                dtos.Rule `json:"rule"`
}

// MultiRulesResponse defines the Response Content for GET multiple rule DTO.
type MultiRulesResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Rules                             []dtos.Rule `json:"rules"`
}

func NewMultiRulesResponse(requestId string, message string, statusCode int, totalCount uint32, rules []dtos.Rule) MultiRulesResponse {
	return MultiRulesResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Rules:                      rules,
	}
}
