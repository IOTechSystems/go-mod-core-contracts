//
// Copyright (C) 2022-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type User struct {
	DBTimestamp
	Id          string
	Name        string
	DisplayName string
	Password    string
	Description string
	Roles       []string
}
