//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type RolePolicy struct {
	DBTimestamp
	Id             string
	Role           string
	Description    string
	AccessPolicies []AccessPolicy
}

type AccessPolicy struct {
	Path        string
	HttpMethods []string
	Effect      string
}
