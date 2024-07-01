//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

// CronTransmissionRecordClient defines the interface for interactions with the CronTransmissionRecord endpoint on the EdgeX Foundry support-cron-scheduler service.
type CronTransmissionRecordClient interface {
	// AllCronTransmissionRecords query cron transmission records with start, end, offset, and limit
	AllCronTransmissionRecords(ctx context.Context, start, end int64, offset, limit int) (responses.MultiCronTransmissionRecordResponse, errors.EdgeX)
	// CronTransmissionRecordsByStatus queries cron transmission records with status, start, end, offset, and limit
	CronTransmissionRecordsByStatus(ctx context.Context, status string, start, end int64, offset, limit int) (responses.MultiCronTransmissionRecordResponse, errors.EdgeX)
	// CronTransmissionRecordsByJobName query cron transmission records with jobName, start, end, offset, and limit
	CronTransmissionRecordsByJobName(ctx context.Context, jobName string, start, end int64, offset, limit int) (responses.MultiCronTransmissionRecordResponse, errors.EdgeX)
	// CronTransmissionRecordsByJobNameAndStatus query cron transmission records with jobName, status, start, end, offset, and limit
	CronTransmissionRecordsByJobNameAndStatus(ctx context.Context, jobName, status string, start, end int64, offset, limit int) (responses.MultiCronTransmissionRecordResponse, errors.EdgeX)
}
