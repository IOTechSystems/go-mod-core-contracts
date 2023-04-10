// Copyright (C) 2023 IOTech Ltd

package messaging

import (
	"context"
	"fmt"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/xrtmodels"
)

func (c *xrtClient) AllDeviceProfiles(ctx context.Context) ([]string, errors.EdgeX) {
	request := xrtmodels.NewAllProfilesRequest(clientName)
	var response xrtmodels.MultiProfilesResponse

	err := c.sendXrtRequest(ctx, request.RequestId, request, &response)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.Kind(err), "failed to query profile list", err)
	}
	return response.Result.Profiles, nil
}

func (c *xrtClient) DeviceProfileByName(ctx context.Context, name string) (models.DeviceProfile, errors.EdgeX) {
	request := xrtmodels.NewProfileGetRequest(name, clientName)
	var response xrtmodels.ProfileResponse

	err := c.sendXrtRequest(ctx, request.RequestId, request, &response)
	if err != nil {
		return models.DeviceProfile{}, errors.NewCommonEdgeX(errors.Kind(err), "failed to query profile", err)
	}
	return response.Result.Profile, nil
}

func (c *xrtClient) AddDeviceProfile(ctx context.Context, device models.DeviceProfile) errors.EdgeX {
	request := xrtmodels.NewProfileAddRequest(device, clientName)
	var response xrtmodels.CommonResponse

	err := c.sendXrtRequest(ctx, request.RequestId, request, &response)
	if err != nil {
		return errors.NewCommonEdgeX(errors.Kind(err), "failed to add profile", err)
	}
	return nil
}

func (c *xrtClient) UpdateDeviceProfile(ctx context.Context, device models.DeviceProfile) errors.EdgeX {
	request := xrtmodels.NewProfileUpdateRequest(device, clientName)
	var response xrtmodels.CommonResponse

	err := c.sendXrtRequest(ctx, request.RequestId, request, &response)
	if err != nil {
		return errors.NewCommonEdgeX(errors.Kind(err), "failed to update profile", err)
	}
	return nil
}

func (c *xrtClient) DeleteDeviceProfileByName(ctx context.Context, name string) errors.EdgeX {
	request := xrtmodels.NewProfileDeleteRequest(name, clientName)
	var response xrtmodels.CommonResponse

	err := c.sendXrtRequest(ctx, request.RequestId, request, &response)
	if err != nil {
		return errors.NewCommonEdgeX(errors.Kind(err), fmt.Sprintf("failed to delete profile %s", name), err)
	}
	return nil
}
