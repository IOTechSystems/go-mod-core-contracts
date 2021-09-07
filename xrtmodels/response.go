// Copyright (C) 2021 IOTech Ltd

package xrtmodels

import v1Model "github.com/edgexfoundry/go-mod-core-contracts/v2/v1/models"

type BaseResponse struct {
	Client    string `json:"client"`
	RequestId string `json:"request_id"`
}

type CommonResponse struct {
	BaseResponse `json:",inline"`
	Result       BaseResult `json:"result"`
}

type BaseResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type MultiResourcesResponse struct {
	BaseResponse `json:",inline"`
	Result       MultiResourcesResult `json:"result"`
}

type MultiResourcesResult struct {
	BaseResult `json:",inline"`
	Readings   map[string]Reading `json:"readings"`
}

type AsyncResourcesResult struct {
	DeviceName string             `json:"device"`
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

type MultiProfilesResponse struct {
	BaseResponse `json:",inline"`
	Result       MultiProfilesResult `json:"result"`
}

type MultiProfilesResult struct {
	BaseResult `json:",inline"`
	Profiles   []string `json:"profiles"`
}

type ProfileResponse struct {
	BaseResult `json:",inline"`
	Result     ProfileResult `json:"result"`
}

type ProfileResult struct {
	BaseResult `json:",inline"`
	Profile    v1Model.DeviceProfile `json:"profile"`
}
