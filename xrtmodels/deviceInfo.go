// Copyright (C) 2021 IOTech Ltd

package xrtmodels

import (
	"fmt"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

type DeviceInfo struct {
	dtos.Device
	Protocols map[string]map[string]interface{} `json:"protocols"`
}

func ToEdgeXV2Device(device DeviceInfo, serviceName string) models.Device {
	// Convert all properties to string for EdgeX
	protocols := make(map[string]models.ProtocolProperties)
	protocolName := ""
	for protocol, protocolProperties := range device.Protocols {
		protocols[protocol] = make(map[string]string)
		protocolName = strings.ToLower(protocol)
		for k, v := range protocolProperties {
			protocols[protocol][k] = fmt.Sprintf("%v", v)
		}
	}
	return models.Device{
		Name:           device.Name,
		Description:    "",
		AdminState:     models.Unlocked,
		OperatingState: models.Up,
		ProtocolName:   protocolName,
		Protocols:      protocols,
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
