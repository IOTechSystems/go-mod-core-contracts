//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

// RolePolicyResponse defines the Response Content for GET RolePolicy DTOs
type RolePolicyResponse struct {
	common.BaseResponse `json:",inline"`
	RolePolicy          dtos.RolePolicy `json:"rolePolicy"`
}

func NewRolePolicyResponse(requestId string, message string, statusCode int, rolePolicy dtos.RolePolicy) RolePolicyResponse {
	return RolePolicyResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		RolePolicy:   rolePolicy,
	}
}

// MultiRolePolicyResponse defines the Response Content for GET multiple RolePolicy DTOs
type MultiRolePolicyResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	RolePolicies                      []dtos.RolePolicy `json:"rolePolicies"`
}

func NewMultiRolePolicyResponse(requestId string, message string, statusCode int, totalCount uint32, rolePolicies []dtos.RolePolicy) MultiRolePolicyResponse {
	return MultiRolePolicyResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		RolePolicies:               rolePolicies,
	}
}
