// Copyright (C) 2021 IOTech Ltd

package xrtmodels

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v1models"

	"github.com/google/uuid"
)

const (
	ProfileAddOperation    = "profile:add"
	ProfileUpdateOperation = "profile:update"
	ProfileListOperation   = "profile:list"
	ProfileGetOperation    = "profile:read"
	ProfileDeleteOperation = "profile:delete"

	DeviceAddOperation         = "device:add"
	DeviceUpdateOperation      = "device:update"
	DeviceResourceGetOperation = "device:get"
	DeviceGetOperation         = "device:read"
	DeviceResourceSetOperation = "device:put"
	DeviceDeleteOperation      = "device:delete"
	DeviceListOperation        = "device:list"

	ScheduleAddOperation    = "schedule:add"
	ScheduleListOperation   = "schedule:list"
	ScheduleDeleteOperation = "schedule:delete"

	DiscoveryTriggerOperation = "discovery:trigger"
)

type BaseRequest struct {
	Client    string `json:"client"`
	RequestId string `json:"request_id"`
	Op        string `json:"op"`
}

type AddProfileRequest struct {
	BaseRequest `json:",inline"`
	Profile     v1models.DeviceProfile `json:"profile"`
}

type UpdateProfileRequest struct {
	BaseRequest `json:",inline"`
	Profile     v1models.DeviceProfile `json:"profile"`
}

type ProfileRequest struct {
	BaseRequest `json:",inline"`
	Profile     string `json:"profile"`
}

type AddDeviceRequest struct {
	BaseRequest `json:",inline"`
	DeviceName  string     `json:"device"`
	DeviceInfo  DeviceInfo `json:"device_info"`
}

type UpdateDeviceRequest struct {
	BaseRequest `json:",inline"`
	DeviceName  string     `json:"device"`
	DeviceInfo  DeviceInfo `json:"device_info"`
}

type DeviceRequest struct {
	BaseRequest `json:",inline"`
	Device      string `json:"device"`
}

type GetResourcesRequest struct {
	BaseRequest `json:",inline"`
	DeviceName  string   `json:"device"`
	Resource    []string `json:"resource"`
}

type PutResourceRequest struct {
	BaseRequest `json:",inline"`
	DeviceName  string                 `json:"device"`
	Values      map[string]interface{} `json:"values"`
}

type ScheduleRequest struct {
	BaseRequest `json:",inline"`
	Schedule    string `json:"schedule"`
}

type AddScheduleRequest struct {
	BaseRequest `json:",inline"`
	Schedule    Schedule `json:"schedule"`
}

func NewBaseRequest(op string, clientName string) BaseRequest {
	return BaseRequest{
		Client:    clientName,
		RequestId: uuid.New().String(),
		Op:        op,
	}
}

// NewProfileAddRequest creates request with v1 device profile
func NewProfileAddRequest(profile v1models.DeviceProfile, clientName string) AddProfileRequest {
	req := AddProfileRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        ProfileAddOperation,
		},
		Profile: profile,
	}

	return req
}

// NewProfileUpdateRequest creates request with v1 device profile
func NewProfileUpdateRequest(profile v1models.DeviceProfile, clientName string) UpdateProfileRequest {
	req := UpdateProfileRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        ProfileUpdateOperation,
		},
		Profile: profile,
	}

	return req
}

func NewProfileGetRequest(profileName string, clientName string) ProfileRequest {
	return ProfileRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        ProfileGetOperation,
		},
		Profile: profileName,
	}
}

func NewProfileDeleteRequest(profileName string, clientName string) ProfileRequest {
	req := ProfileRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        ProfileDeleteOperation,
		},
		Profile: profileName,
	}
	return req
}

func NewDeviceAddRequest(device models.Device, clientName string) AddDeviceRequest {
	deviceRequest := AddDeviceRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        DeviceAddOperation,
		},
		DeviceName: device.Name,
		DeviceInfo: DeviceInfo{
			ProfileName: device.ProfileName,
			Protocols:   device.Protocols,
			Properties:  device.Properties,
		},
	}
	return deviceRequest
}

func NewDeviceUpdateRequest(device models.Device, clientName string) UpdateDeviceRequest {
	deviceRequest := UpdateDeviceRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        DeviceUpdateOperation,
		},
		DeviceName: device.Name,
		DeviceInfo: DeviceInfo{
			ProfileName: device.ProfileName,
			Protocols:   device.Protocols,
		},
	}
	return deviceRequest
}

func NewDeviceGetRequest(deviceName string, clientName string) DeviceRequest {
	req := DeviceRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        DeviceGetOperation,
		},
		Device: deviceName,
	}
	return req
}

func NewDeviceDeleteRequest(deviceName string, clientName string) DeviceRequest {
	req := DeviceRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        DeviceDeleteOperation,
		},
		Device: deviceName,
	}
	return req
}

func NewDeviceResourceGetRequest(deviceName string, clientName string, resources []string) GetResourcesRequest {
	req := GetResourcesRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        DeviceResourceGetOperation,
		},
		DeviceName: deviceName,
		Resource:   nil,
	}
	req.Resource = resources
	return req
}

func NewDeviceResourceSetRequest(deviceName string, clientName string, values map[string]interface{}) PutResourceRequest {
	req := PutResourceRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        DeviceResourceSetOperation,
		},
		DeviceName: deviceName,
		Values:     nil,
	}
	req.Values = values
	return req
}

func NewScheduleAddRequest(deviceName string, clientName string, schedule Schedule) AddScheduleRequest {
	req := AddScheduleRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        ScheduleAddOperation,
		},
		Schedule: schedule,
	}
	return req
}

func NewScheduleDeleteRequest(scheduleName string, clientName string) ScheduleRequest {
	req := ScheduleRequest{
		BaseRequest: BaseRequest{
			Client:    clientName,
			RequestId: uuid.New().String(),
			Op:        ScheduleDeleteOperation,
		},
		Schedule: scheduleName,
	}
	return req
}
