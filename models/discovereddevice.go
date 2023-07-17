//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type DiscoveredDevice struct {
	ProfileName string
	AdminState  AdminState
	AutoEvents  []AutoEvent
	Properties  map[string]any

	// Xpert
	DeviceNameTemplate string
	ProtocolName       string
	DeviceDescription  string
	DeviceLabels       []string

	ProfileNameTemplate string
	ProfileLabels       []string
	ProfileDescription  string
}
