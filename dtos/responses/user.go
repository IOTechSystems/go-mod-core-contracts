//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

// UserResponse defines the Response Content for GET User DTOs.
type UserResponse struct {
	common.BaseResponse `json:",inline"`
	User                dtos.User `json:"user"`
}

func NewUserResponse(requestId string, message string, statusCode int, user dtos.User) UserResponse {
	return UserResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		User:         user,
	}
}

// MultiUsersResponse defines the Response Content for GET multiple User DTOs
type MultiUsersResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Users                             []dtos.User `json:"users"`
}

func NewMultiUsersResponse(requestId string, message string, statusCode int, totalCount uint32, users []dtos.User) MultiUsersResponse {
	return MultiUsersResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Users:                      users,
	}
}
