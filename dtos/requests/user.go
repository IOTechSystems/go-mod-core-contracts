//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

// AddUserRequest defines the Request Content for POST User DTO.
type AddUserRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	User                  dtos.User `json:"user"`
}

// Validate satisfies the Validator interface
func (a *AddUserRequest) Validate() error {
	err := common.Validate(a)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddUserRequest type
func (a *AddUserRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		User dtos.User
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*a = AddUserRequest(alias)

	// validate AddUserRequest DTO
	if err := a.Validate(); err != nil {
		return err
	}
	return nil
}

// AddUserReqToUserModels transforms the AddUserRequest DTO array to the User model array
func AddUserReqToUserModels(addRequests []AddUserRequest) (users []models.User) {
	for _, req := range addRequests {
		d := dtos.ToUserModel(req.User)
		users = append(users, d)
	}
	return users
}

// UpdateUserRequest defines the Request Content for PATCH User DTO.
type UpdateUserRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	User                  dtos.UpdateUser `json:"user"`
}

// Validate satisfies the Validator interface
func (u UpdateUserRequest) Validate() error {
	err := common.Validate(u)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateUserRequest type
func (u *UpdateUserRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		User dtos.UpdateUser
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*u = UpdateUserRequest(alias)

	// validate UpdateUserRequest DTO
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}
