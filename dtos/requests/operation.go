//
// Copyright (C) 2021-2023 IOTech Ltd
//

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

// OperationRequest defines the Request Content for SMA POST Operation.
// This object and its properties correspond to the OperationRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/system-agent/2.1.0#/OperationRequest
type OperationRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ServiceName           string   `json:"serviceName" validate:"required"`
	Action                string   `json:"action" validate:"oneof='start' 'stop' 'restart' 'rm' 'inspect'"`
	Flags                 []string `json:"flags"` // to identify whether the app service container will be started with additional flags
}

// Validate satisfies the Validator interface
func (o *OperationRequest) Validate() error {
	err := common.Validate(o)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the OperationRequest type
func (o *OperationRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ServiceName string
		Action      string
		Flags       []string
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*o = OperationRequest(alias)

	if err := o.Validate(); err != nil {
		return err
	}

	return nil
}
