//
// Copyright (C) 2023 IOTech Ltd
//

package http

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

var emptyResponse any

// RegistryClient is the REST client for invoking the registry APIs(/registry/*) from Core Keeper
type registryClient struct {
	baseUrl string
}

// NewRegistryClient creates an instance of RegistryClient
func NewRegistryClient(baseUrl string) interfaces.RegistryClient {
	return &registryClient{
		baseUrl: baseUrl,
	}
}

// Register registers a service instance
func (rc *registryClient) Register(ctx context.Context, req requests.AddRegistrationRequest) errors.EdgeX {
	err := utils.PostRequestWithRawData(ctx, &emptyResponse, rc.baseUrl, common.ApiRegisterRoute, nil, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	return nil
}

// UpdateRegister updates the registration data of the service
func (rc *registryClient) UpdateRegister(ctx context.Context, req requests.AddRegistrationRequest) errors.EdgeX {
	err := utils.PutRequest(ctx, &emptyResponse, rc.baseUrl, common.ApiRegisterRoute, nil, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	return nil
}

// RegistrationByServiceId returns the registration data by service id
func (rc *registryClient) RegistrationByServiceId(ctx context.Context, serviceId string) (responses.RegistrationResponse, errors.EdgeX) {
	requestPath := utils.EscapeAndJoinPath(common.ApiRegisterRoute, common.ServiceId, serviceId)
	res := responses.RegistrationResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllRegistry returns the registration data of all registered service
func (rc *registryClient) AllRegistry(ctx context.Context) (responses.MultiRegistrationsResponse, errors.EdgeX) {
	res := responses.MultiRegistrationsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, common.ApiAllRegistrationsRoute, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// Deregister deregisters a service by service id
func (rc *registryClient) Deregister(ctx context.Context, serviceId string) errors.EdgeX {
	requestPath := utils.EscapeAndJoinPath(common.ApiRegisterRoute, common.ServiceId, serviceId)
	err := utils.DeleteRequest(ctx, &emptyResponse, rc.baseUrl, requestPath)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	return nil
}