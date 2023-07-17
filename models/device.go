//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type Device struct {
	DBTimestamp
	Id             string
	Name           string
	Description    string
	AdminState     AdminState
	OperatingState OperatingState
	Protocols      map[string]ProtocolProperties
	Labels         []string
	Location       interface{}
	ServiceName    string
	ProfileName    string
	AutoEvents     []AutoEvent
	Tags           map[string]any
	Properties     map[string]any

	// Xpert
	ProtocolName string
}

// ProtocolProperties contains the device connection information in key/value pair
type ProtocolProperties map[string]any

// AdminState controls the range of values which constitute valid administrative states for a device
type AdminState string

// OperatingState is an indication of the operations of the device.
type OperatingState string
