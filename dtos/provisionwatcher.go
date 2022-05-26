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
	ProtocolName        string              `json:"protocolName" validate:"omitempty,edgex-dto-rfc3986-unreserved-chars"`
	ProfileName         string              `json:"profileName" validate:"omitempty,edgex-dto-rfc3986-unreserved-chars"`
	AutoEvents          []AutoEvent         `json:"autoEvents,omitempty" validate:"dive"`
	DeviceDescription   string              `json:"deviceDescription"`
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
	ProtocolName        *string             `json:"protocolName" validate:"omitempty,edgex-dto-rfc3986-unreserved-chars"`
	ProfileName         *string             `json:"profileName" validate:"omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	AutoEvents          []AutoEvent         `json:"autoEvents" validate:"dive"`
	DeviceDescription   *string             `json:"deviceDescription"`
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
		ProtocolName:        dto.ProtocolName,
		ProfileName:         dto.ProfileName,
		AutoEvents:          ToAutoEventModels(dto.AutoEvents),
		DeviceDescription:   dto.DeviceDescription,
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
		ProtocolName:        pw.ProtocolName,
		ProfileName:         pw.ProfileName,
		AutoEvents:          FromAutoEventModelsToDTOs(pw.AutoEvents),
		DeviceDescription:   pw.DeviceDescription,
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
		ProtocolName:        &pw.ProtocolName,
		ProfileName:         &pw.ProfileName,
		AutoEvents:          FromAutoEventModelsToDTOs(pw.AutoEvents),
		DeviceDescription:   &pw.DeviceDescription,
	}
	return dto
}
