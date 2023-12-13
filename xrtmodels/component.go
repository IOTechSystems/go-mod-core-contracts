//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xrtmodels

type Component struct {
	Category string         `json:"category"`
	Config   map[string]any `json:"config"`
	Name     string         `json:"name"`
	State    string         `json:"state"`
	Type     string         `json:"type"`
}
