//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xrtmodels

type DeviceStatus struct {
	Device      string `json:"device"`
	Operational bool   `json:"operational"`
	Type        string `json:"type"`
}
