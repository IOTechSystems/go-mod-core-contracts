//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
)

func TestNewCronTransmissionRecordResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	expectedTransmissionRecord := dtos.CronTransmissionRecord{JobName: "testJob"}
	actual := NewCronTransmissionRecordResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTransmissionRecord)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTransmissionRecord, actual.TransmissionRecord)
}

func TestNewMultiCronTransmissionRecordResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	expectedCronTransmissionRecords := []dtos.CronTransmissionRecord{
		{
			JobName: "testJob1",
		},
		{
			JobName: "testJob2",
		},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiCronTransmissionRecordResponse(expectedRequestId, expectedMessage, expectedStatusCode, uint32(len(expectedCronTransmissionRecords)), expectedCronTransmissionRecords)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedCronTransmissionRecords, actual.TransmissionRecords)
}
