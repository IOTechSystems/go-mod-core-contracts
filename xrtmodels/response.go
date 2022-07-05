// Copyright (C) 2021 IOTech Ltd

package xrtmodels

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

const (
	// https://github.com/IOTechSystems/xrt/blob/master/include/devsdk/devsdk.h#L27
	XrtSdkStatusOk               = 0
	XrtSdkStatusNotFound         = 1
	XrtSdkStatusNotSupported     = 2
	XrtSdkStatusInvalidOperation = 3
	XrtSdkStatusAlreadyExists    = 7
)

type BaseResponse struct {
	Client    string `json:"client"`
	RequestId string `json:"request_id"`
	Type      string `json:"type"`
}

type CommonResponse struct {
	BaseResponse `json:",inline"`
	Result       BaseResult `json:"result"`
}

type BaseResult struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error"`
}

func (result BaseResult) Error() errors.EdgeX {
	switch result.Status {
	case XrtSdkStatusOk:
		return nil
	case XrtSdkStatusNotFound:
		return errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, result.ErrorMessage, nil)
	case XrtSdkStatusNotSupported:
		return errors.NewCommonEdgeX(errors.KindNotImplemented, result.ErrorMessage, nil)
	case XrtSdkStatusInvalidOperation:
		return errors.NewCommonEdgeX(errors.KindInvalidId, result.ErrorMessage, nil)
	case XrtSdkStatusAlreadyExists:
		return errors.NewCommonEdgeX(errors.KindDuplicateName, result.ErrorMessage, nil)
	default:
		return errors.NewCommonEdgeX(errors.KindServerError, result.ErrorMessage, nil)
	}
}

type MultiResourcesResponse struct {
	BaseResponse `json:",inline"`
	Result       MultiResourcesResult `json:"result"`
}

type MultiResourcesResult struct {
	BaseResult `json:",inline"`
	Device     string             `json:"device"`
	Profile    string             `json:"profile"`
	Readings   map[string]Reading `json:"readings"`
}

type Reading struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

type MultiDevicesResponse struct {
	BaseResponse `json:",inline"`
	Result       MultiDevicesResult `json:"result"`
}

type MultiDevicesResult struct {
	BaseResult `json:",inline"`
	Devices    []string `json:"devices"`
}

type DeviceResponse struct {
	BaseResponse `json:",inline"`
	Result       DeviceResult `json:"result"`
}

type DeviceResult struct {
	BaseResult `json:",inline"`
	Device     DeviceInfo `json:"device"`
}

type DiscoveredDevicesResult struct {
	BaseResult `json:",inline"`
	Devices    map[string]DeviceInfo `json:"devices"`
}

type MultiProfilesResponse struct {
	BaseResponse `json:",inline"`
	Result       MultiProfilesResult `json:"result"`
}

type MultiProfilesResult struct {
	BaseResult `json:",inline"`
	Profiles   []string `json:"profiles"`
}

type ProfileResponse struct {
	BaseResponse `json:",inline"`
	Result       ProfileResult `json:"result"`
}

type ProfileResult struct {
	BaseResult `json:",inline"`
	Profile    models.DeviceProfile `json:"profile"`
}

type MultiSchedulesResponse struct {
	BaseResponse `json:",inline"`
	Result       MultiSchedulesResult `json:"result"`
}

type MultiSchedulesResult struct {
	BaseResult `json:",inline"`
	Schedules  []string `json:"schedules"`
}
