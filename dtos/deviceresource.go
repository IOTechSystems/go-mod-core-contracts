//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// DeviceResource and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.1.0#/DeviceResource
type DeviceResource struct {
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Name        string                 `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	IsHidden    bool                   `json:"isHidden" yaml:"isHidden"`
	Tag         string                 `json:"tag,omitempty" yaml:"tag,omitempty"`
	Tags        map[string]interface{} `json:"tags,omitempty" yaml:"tags,omitempty"`
	Properties  ResourceProperties     `json:"properties" yaml:"properties"`
	Attributes  map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// Validate satisfies the Validator interface
func (dr *DeviceResource) Validate() error {
	err := common.Validate(dr)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid DeviceResource.", err)
	}

	return nil
}

// ToDeviceResourceModel transforms the DeviceResource DTO to the DeviceResource model
func ToDeviceResourceModel(d DeviceResource) models.DeviceResource {
	return models.DeviceResource{
		Description: d.Description,
		Name:        d.Name,
		IsHidden:    d.IsHidden,
		Tag:         d.Tag,
		Tags:        d.Tags,
		Properties:  ToResourcePropertiesModel(d.Properties),
		Attributes:  d.Attributes,
	}
}

// ToDeviceResourceModels transforms the DeviceResource DTOs to the DeviceResource models
func ToDeviceResourceModels(deviceResourceDTOs []DeviceResource) []models.DeviceResource {
	deviceResourceModels := make([]models.DeviceResource, len(deviceResourceDTOs))
	for i, d := range deviceResourceDTOs {
		deviceResourceModels[i] = ToDeviceResourceModel(d)
	}
	return deviceResourceModels
}

// FromDeviceResourceModelToDTO transforms the DeviceResource model to the DeviceResource DTO
func FromDeviceResourceModelToDTO(d models.DeviceResource) DeviceResource {
	return DeviceResource{
		Description: d.Description,
		Name:        d.Name,
		IsHidden:    d.IsHidden,
		Tag:         d.Tag,
		Tags:        d.Tags,
		Properties:  FromResourcePropertiesModelToDTO(d.Properties),
		Attributes:  d.Attributes,
	}
}

// FromDeviceResourceModelsToDTOs transforms the DeviceResource models to the DeviceResource DTOs
func FromDeviceResourceModelsToDTOs(deviceResourceModels []models.DeviceResource) []DeviceResource {
	deviceResourceDTOs := make([]DeviceResource, len(deviceResourceModels))
	for i, d := range deviceResourceModels {
		deviceResourceDTOs[i] = FromDeviceResourceModelToDTO(d)
	}
	return deviceResourceDTOs
}
