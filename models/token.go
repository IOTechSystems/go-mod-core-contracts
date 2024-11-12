//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// TokenDetails stores the token id, client ip, user agent and other token statuses information to the database for a session
type TokenDetails struct {
	Id        string
	ClientIP  string
	UserAgent string
	UserId    string
	Revoked   bool
	ExpTime   int64
}
