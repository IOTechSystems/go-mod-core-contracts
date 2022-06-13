//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type User struct {
	DBTimestamp
	Id        string
	Name      string
	Group     string
	PublicKey string
}
