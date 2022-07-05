//
// Copyright (C) 2021-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"
)

type ProvisionWatcher struct {
	DBTimestamp         `json:",inline"`
	Id                  string              `json:"id,omitempty" yaml:"id,omitempty" validate:"omitempty,uuid"`
	Name                string              `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	ServiceName         string              `json:"serviceName" yaml:"serviceName" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Labels              []string            `json:"labels,omitempty" yaml:"labels,omitempty"`
	Identifiers         map[string]string   `json:"identifiers" yaml:"identifiers" validate:"gt=0,dive,keys,required,endkeys,required"`
	BlockingIdentifiers map[string][]string `json:"blockingIdentifiers,omitempty" yaml:"blockingIdentifiers,omitempty"`
	AdminState          string              `json:"adminState" yaml:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	DiscoveredDevice    DiscoveredDevice    `json:"discoveredDevice" yaml:"discoveredDevice" validate:"dive"`

	// Xpert
	DeviceNameTemplate string      `json:"deviceNameTemplate"`
	ProfileName        string      `json:"profileName" validate:"omitempty,edgex-dto-rfc3986-unreserved-chars"`
	AdminState         string      `json:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	AutoEvents         []AutoEvent `json:"autoEvents,omitempty" validate:"dive"`
	ProtocolName       string      `json:"protocolName" validate:"omitempty,edgex-dto-rfc3986-unreserved-chars"`
	DeviceDescription  string      `json:"deviceDescription"`
	DeviceLabels       []string    `json:"deviceLabels"`

	ProfileNameTemplate string   `json:"profileNameTemplate"`
	ProfileLabels       []string `json:"profileLabels"`
	ProfileDescription  string   `json:"profileDescription"`
}

type UpdateProvisionWatcher struct {
	Id                  *string                `json:"id" validate:"required_without=Name,edgex-dto-uuid"`
	Name                *string                `json:"name" validate:"required_without=Id,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	ServiceName         *string                `json:"serviceName" validate:"omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Labels              []string               `json:"labels"`
	Identifiers         map[string]string      `json:"identifiers" validate:"omitempty,gt=0,dive,keys,required,endkeys,required"`
	BlockingIdentifiers map[string][]string    `json:"blockingIdentifiers"`
	AdminState          *string                `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	DiscoveredDevice    UpdateDiscoveredDevice `json:"discoveredDevice"`

	// Xpert
	DeviceNameTemplate *string     `json:"deviceNameTemplate"`
	ProfileName        *string     `json:"profileName" validate:"omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	AdminState         *string     `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	AutoEvents         []AutoEvent `json:"autoEvents" validate:"dive"`
	ProtocolName       *string     `json:"protocolName" validate:"omitempty,edgex-dto-rfc3986-unreserved-chars"`
	DeviceDescription  *string     `json:"deviceDescription"`
	DeviceLabels       []string    `json:"deviceLabels"`

	ProfileNameTemplate *string  `json:"profileNameTemplate"`
	ProfileLabels       []string `json:"profileLabels"`
	ProfileDescription  *string  `json:"profileDescription"`
}

// ToProvisionWatcherModel transforms the ProvisionWatcher DTO to the ProvisionWatcher model
func ToProvisionWatcherModel(dto ProvisionWatcher) models.ProvisionWatcher {
	return models.ProvisionWatcher{
		DBTimestamp:         models.DBTimestamp(dto.DBTimestamp),
		Id:                  dto.Id,
		Name:                dto.Name,
		ServiceName:         dto.ServiceName,
		Labels:              dto.Labels,
		Identifiers:         dto.Identifiers,
		BlockingIdentifiers: dto.BlockingIdentifiers,
		AdminState:          models.AdminState(dto.AdminState),
		DiscoveredDevice:    ToDiscoveredDeviceModel(dto.DiscoveredDevice),

		// Xpert
		DeviceNameTemplate: dto.DeviceNameTemplate,
		ProfileName:        dto.ProfileName,
		AdminState:         models.AdminState(dto.AdminState),
		AutoEvents:         ToAutoEventModels(dto.AutoEvents),
		ProtocolName:       dto.ProtocolName,
		DeviceDescription:  dto.DeviceDescription,
		DeviceLabels:       dto.DeviceLabels,

		ProfileNameTemplate: dto.ProfileNameTemplate,
		ProfileLabels:       dto.ProfileLabels,
		ProfileDescription:  dto.ProfileDescription,
	}
}

// FromProvisionWatcherModelToDTO transforms the ProvisionWatcher Model to the ProvisionWatcher DTO
func FromProvisionWatcherModelToDTO(pw models.ProvisionWatcher) ProvisionWatcher {
	return ProvisionWatcher{
		DBTimestamp:         DBTimestamp(pw.DBTimestamp),
		Id:                  pw.Id,
		Name:                pw.Name,
		ServiceName:         pw.ServiceName,
		Labels:              pw.Labels,
		Identifiers:         pw.Identifiers,
		BlockingIdentifiers: pw.BlockingIdentifiers,
		AdminState:          string(pw.AdminState),
		DiscoveredDevice:    FromDiscoveredDeviceModelToDTO(pw.DiscoveredDevice),

		// Xpert
		DeviceNameTemplate: pw.DeviceNameTemplate,
		ProfileName:        pw.ProfileName,
		AdminState:         string(pw.AdminState),
		AutoEvents:         FromAutoEventModelsToDTOs(pw.AutoEvents),
		ProtocolName:       pw.ProtocolName,
		DeviceDescription:  pw.DeviceDescription,
		DeviceLabels:       pw.DeviceLabels,

		ProfileNameTemplate: pw.ProfileNameTemplate,
		ProfileLabels:       pw.ProfileLabels,
		ProfileDescription:  pw.ProfileDescription,
	}
}

// FromProvisionWatcherModelToUpdateDTO transforms the ProvisionWatcher Model to the UpdateProvisionWatcher DTO
func FromProvisionWatcherModelToUpdateDTO(pw models.ProvisionWatcher) UpdateProvisionWatcher {
	adminState := string(pw.AdminState)
	dto := UpdateProvisionWatcher{
		Id:                  &pw.Id,
		Name:                &pw.Name,
		ServiceName:         &pw.ServiceName,
		Labels:              pw.Labels,
		Identifiers:         pw.Identifiers,
		BlockingIdentifiers: pw.BlockingIdentifiers,
		AdminState:          &adminState,
		DiscoveredDevice:    FromDiscoveredDeviceModelToUpdateDTO(pw.DiscoveredDevice),

		// Xpert
		DeviceNameTemplate: &pw.DeviceNameTemplate,
		ProfileName:        &pw.ProfileName,
		AdminState:         &adminState,
		AutoEvents:         FromAutoEventModelsToDTOs(pw.AutoEvents),
		ProtocolName:       &pw.ProtocolName,
		DeviceDescription:  &pw.DeviceDescription,
		DeviceLabels:       pw.DeviceLabels,

		ProfileNameTemplate: &pw.ProfileNameTemplate,
		ProfileLabels:       pw.ProfileLabels,
		ProfileDescription:  &pw.ProfileDescription,
	}
	return dto
}
