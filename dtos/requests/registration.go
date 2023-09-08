//
// Copyright (C) 2023 IOTech Ltd
//

package requests

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

type AddRegistrationRequest struct {
	common.BaseRequest `json:",inline"`
	Registration       dtos.Registration `json:"registration"`
}

func (a *AddRegistrationRequest) Validate() error {
	err := a.Registration.Validate()
	if err != nil {
		return err
	}
	return nil
}
