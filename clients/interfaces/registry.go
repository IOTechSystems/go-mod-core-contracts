//
// Copyright (C) 2023 IOTech Ltd
//

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

// RegistryClient defines the interface for interactions with the registry endpoint on the Edge Xpert core-keeper service.
type RegistryClient interface {
	Register(context.Context, requests.AddRegistrationRequest) errors.EdgeX
	UpdateRegister(context.Context, requests.AddRegistrationRequest) errors.EdgeX
	RegistrationByServiceId(context.Context, string) (responses.RegistrationResponse, errors.EdgeX)
	AllRegistry(context.Context) (responses.MultiRegistrationsResponse, errors.EdgeX)
	Deregister(context.Context, string) errors.EdgeX
}
