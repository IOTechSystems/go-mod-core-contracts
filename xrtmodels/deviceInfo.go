// Copyright (C) 2021 IOTech Ltd

package xrtmodels

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

type DeviceInfo struct {
	Name        string                               `json:"name"`
	ProfileName string                               `json:"profile"`
	Protocols   map[string]models.ProtocolProperties `json:"protocols"`
	Properties  map[string]interface{}               `json:"properties"`
}

func ToV2Device(device DeviceInfo, serviceName string) models.Device {
	return models.Device{
		Name:           device.Name,
		Description:    "",
		AdminState:     models.Unlocked,
		OperatingState: models.Up,
		Protocols:      device.Protocols,
		LastConnected:  0,
		LastReported:   0,
		Labels:         nil,
		Location:       nil,
		ServiceName:    serviceName,
		ProfileName:    device.ProfileName,
		AutoEvents:     nil,
		Notify:         false,
		Properties:     device.Properties,
	}
}
