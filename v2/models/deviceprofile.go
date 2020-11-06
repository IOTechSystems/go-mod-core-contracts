//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// DeviceProfile and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceProfile
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type DeviceProfile struct {
	Timestamps
	Description     string
	Id              string
	Name            string
	Manufacturer    string
	Model           string
	Labels          []string
	DeviceResources []DeviceResource
	DeviceCommands  []ProfileResource
	CoreCommands    []Command
}
