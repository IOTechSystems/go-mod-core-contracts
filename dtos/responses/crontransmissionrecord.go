//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
)

// CronTransmissionRecordResponse defines the Response Content for GET CronTransmissionRecord DTO.
type CronTransmissionRecordResponse struct {
	common.BaseResponse `json:",inline"`
	TransmissionRecord  dtos.CronTransmissionRecord `json:"transmissionRecord"`
}

func NewCronTransmissionRecordResponse(requestId string, message string, statusCode int, transmissionRecord dtos.CronTransmissionRecord) CronTransmissionRecordResponse {
	return CronTransmissionRecordResponse{
		BaseResponse:       common.NewBaseResponse(requestId, message, statusCode),
		TransmissionRecord: transmissionRecord,
	}
}

// MultiCronTransmissionRecordResponse defines the Response Content for GET multiple CronTransmissionRecord DTOs.
type MultiCronTransmissionRecordResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	TransmissionRecords               []dtos.CronTransmissionRecord `json:"transmissionRecords"`
}

func NewMultiCronTransmissionRecordResponse(requestId string, message string, statusCode int, totalCount uint32, transmissionRecords []dtos.CronTransmissionRecord) MultiCronTransmissionRecordResponse {
	return MultiCronTransmissionRecordResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		TransmissionRecords:        transmissionRecords,
	}
}
