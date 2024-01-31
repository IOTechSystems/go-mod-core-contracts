//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type DeviceService struct {
	DBTimestamp
	Id          string
	Name        string
	Description string
	Labels      []string
	BaseAddress string
	AdminState  AdminState

	// Central
	TransformScript string
}
