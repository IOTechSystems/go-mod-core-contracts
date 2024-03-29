//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2models

// AutoEvent and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/AutoEvent
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type AutoEvent struct {
	Interval   string
	OnChange   bool
	SourceName string
}
