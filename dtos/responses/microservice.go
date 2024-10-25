//
// Copyright (C) 2022-2024 IOTech Ltd
//

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

// MultiMicroServicesResponse defines the Response Content for GET multiple MicroService DTOs.
type MultiMicroServicesResponse struct {
	common.BaseResponse `json:",inline"`
	MicroServices       []dtos.MicroService `json:"microservices"`
}

func NewMultiMicroServicesResponse(requestId string, message string, statusCode int, ms []dtos.MicroService) MultiMicroServicesResponse {
	return MultiMicroServicesResponse{
		BaseResponse:  common.NewBaseResponse(requestId, message, statusCode),
		MicroServices: ms,
	}
}
