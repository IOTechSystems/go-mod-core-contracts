//
// Copyright (C) 2020-2021 IOTech Ltd
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package common

// ConfigResponse defines the configuration for the targeted service.
type ConfigResponse struct {
	Versionable `json:",inline" toml:",inline"`
	Config      interface{} `json:"config" toml:"config"`
	ServiceName string      `json:"serviceName" toml:"serviceName"`
}

// NewConfigResponse creates new ConfigResponse with all fields set appropriately
func NewConfigResponse(serviceConfig interface{}, serviceName string) ConfigResponse {
	return ConfigResponse{
		Versionable: NewVersionable(),
		Config:      serviceConfig,
		ServiceName: serviceName,
	}
}
