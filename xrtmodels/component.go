//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xrtmodels

// constants of device connector settings
const (
	Name                               = "Name"
	EventTopic                         = "EventTopic"
	ReplyTopic                         = "ReplyTopic"
	RequestTopic                       = "RequestTopic"
	TelemetryTopic                     = "TelemetryTopic"
	OPCUAServerRequestTimeout          = "OPCUAServerRequestTimeout"
	OPCUAServerUseTelemetryValues      = "OPCUAServerUseTelemetryValues"
	OPCUAServerStaleTelemetryValueTime = "OPCUAServerStaleTelemetryValueTime"
	OPCUAServerTopicMiddlewarePrefix   = "OPCUAServerTopicMiddlewarePrefix"

	DeviceServiceCategory      = "XRT::DeviceService"
	DeviceServiceRunningStatus = "Running"
)

type Component struct {
	Category string         `json:"category"`
	Config   map[string]any `json:"config"`
	Name     string         `json:"name"`
	State    string         `json:"state"`
	Type     string         `json:"type"`
}
