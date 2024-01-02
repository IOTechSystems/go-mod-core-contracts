//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2models

type DBTimestamp struct {
	Created  int64 // Created is a timestamp indicating when the entity was created.
	Modified int64 // Modified is a timestamp indicating when the entity was last modified.
}
