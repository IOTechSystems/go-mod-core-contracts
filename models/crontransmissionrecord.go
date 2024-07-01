//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type CronTransmissionRecord struct {
	Id          string
	JobName     string
	Action      ScheduleAction
	Status      CronTransmissionStatus
	ScheduledAt int64
	Created     int64
}

// CronTransmissionStatus indicates the most recent success/failure of a given transmission attempt or a missed transmission.
type CronTransmissionStatus string
