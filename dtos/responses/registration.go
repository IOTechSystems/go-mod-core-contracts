//
// Copyright (C) 2023 IOTech Ltd
//

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
)

type RegistrationResponse struct {
	common.BaseResponse `json:",inline"`
	Registration        dtos.Registration `json:"registration"`
}

type MultiRegistrationsResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Registrations                     []dtos.Registration `json:"registrations"`
}

func NewRegistrationResponse(requestId string, message string, statusCode int, r dtos.Registration) RegistrationResponse {
	return RegistrationResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Registration: r,
	}
}

func NewMultiRegistrationsResponse(requestId string, message string, statusCode int, totalCount uint32, registrations []dtos.Registration) MultiRegistrationsResponse {
	return MultiRegistrationsResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Registrations:              registrations,
	}
}
