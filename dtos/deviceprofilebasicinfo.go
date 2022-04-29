//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

type DeviceProfileBasicInfo struct {
	Id           string   `json:"id" validate:"omitempty,uuid" yaml:"id,omitempty"`
	Name         string   `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Manufacturer string   `json:"manufacturer" yaml:"manufacturer,omitempty"`
	Description  string   `json:"description" yaml:"description,omitempty"`
	Model        string   `json:"model" yaml:"model,omitempty"`
	Labels       []string `json:"labels" yaml:"labels,flow,omitempty"`
}

type UpdateDeviceProfileBasicInfo struct {
	Id           *string  `json:"id" validate:"required_without=Name,edgex-dto-uuid"`
	Name         *string  `json:"name" validate:"required_without=Id,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Manufacturer *string  `json:"manufacturer"`
	Description  *string  `json:"description"`
	Model        *string  `json:"model"`
	Labels       []string `json:"labels"`
}
