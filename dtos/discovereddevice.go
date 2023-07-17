//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v3/models"

type DiscoveredDevice struct {
	ProfileName string         `json:"profileName" yaml:"profileName" validate:"len=0|edgex-dto-no-reserved-chars"`
	AdminState  string         `json:"adminState" yaml:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	AutoEvents  []AutoEvent    `json:"autoEvents,omitempty" yaml:"autoEvents,omitempty" validate:"dive"`
	Properties  map[string]any `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Xpert
	DeviceNameTemplate string   `json:"deviceNameTemplate,omitempty" validate:"omitempty"`
	ProtocolName       string   `json:"protocolName" validate:"omitempty,len=0|edgex-dto-rfc3986-unreserved-chars"`
	DeviceDescription  string   `json:"deviceDescription"`
	DeviceLabels       []string `json:"deviceLabels"`

	ProfileNameTemplate string   `json:"profileNameTemplate" validate:"omitempty,len=0|edgex-dto-no-reserved-chars"`
	ProfileLabels       []string `json:"profileLabels"`
	ProfileDescription  string   `json:"profileDescription"`
}

type UpdateDiscoveredDevice struct {
	ProfileName *string        `json:"profileName" validate:"omitempty,len=0|edgex-dto-no-reserved-chars"`
	AdminState  *string        `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	AutoEvents  []AutoEvent    `json:"autoEvents" validate:"dive"`
	Properties  map[string]any `json:"properties"`

	// Xpert
	DeviceNameTemplate *string  `json:"deviceNameTemplate" validate:"omitempty"`
	ProtocolName       *string  `json:"protocolName" validate:"omitempty,len=0|edgex-dto-rfc3986-unreserved-chars"`
	DeviceDescription  *string  `json:"deviceDescription"`
	DeviceLabels       []string `json:"deviceLabels"`

	ProfileNameTemplate *string  `json:"profileNameTemplate" validate:"omitempty,len=0|edgex-dto-no-reserved-chars"`
	ProfileLabels       []string `json:"profileLabels"`
	ProfileDescription  *string  `json:"profileDescription"`
}

func ToDiscoveredDeviceModel(dto DiscoveredDevice) models.DiscoveredDevice {
	return models.DiscoveredDevice{
		ProfileName: dto.ProfileName,
		AdminState:  models.AdminState(dto.AdminState),
		AutoEvents:  ToAutoEventModels(dto.AutoEvents),
		Properties:  dto.Properties,

		// Xpert
		DeviceNameTemplate: dto.DeviceNameTemplate,
		ProtocolName:       dto.ProtocolName,
		DeviceDescription:  dto.DeviceDescription,
		DeviceLabels:       dto.DeviceLabels,

		ProfileNameTemplate: dto.ProfileNameTemplate,
		ProfileLabels:       dto.ProfileLabels,
		ProfileDescription:  dto.ProfileDescription,
	}
}

func FromDiscoveredDeviceModelToDTO(d models.DiscoveredDevice) DiscoveredDevice {
	return DiscoveredDevice{
		ProfileName: d.ProfileName,
		AdminState:  string(d.AdminState),
		AutoEvents:  FromAutoEventModelsToDTOs(d.AutoEvents),
		Properties:  d.Properties,

		// Xpert
		DeviceNameTemplate: d.DeviceNameTemplate,
		ProtocolName:       d.ProtocolName,
		DeviceDescription:  d.DeviceDescription,
		DeviceLabels:       d.DeviceLabels,

		ProfileNameTemplate: d.ProfileNameTemplate,
		ProfileLabels:       d.ProfileLabels,
		ProfileDescription:  d.ProfileDescription,
	}
}

func FromDiscoveredDeviceModelToUpdateDTO(d models.DiscoveredDevice) UpdateDiscoveredDevice {
	adminState := string(d.AdminState)
	return UpdateDiscoveredDevice{
		ProfileName: &d.ProfileName,
		AdminState:  &adminState,
		AutoEvents:  FromAutoEventModelsToDTOs(d.AutoEvents),
		Properties:  d.Properties,

		// Xpert
		DeviceNameTemplate: &d.DeviceNameTemplate,
		ProtocolName:       &d.ProtocolName,
		DeviceDescription:  &d.DeviceDescription,
		DeviceLabels:       d.DeviceLabels,

		ProfileNameTemplate: &d.ProfileNameTemplate,
		ProfileLabels:       d.ProfileLabels,
		ProfileDescription:  &d.ProfileDescription,
	}
}
