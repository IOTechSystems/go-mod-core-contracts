// Copyright (C) 2021 IOTech Ltd

package xrtmodels

import (
	"encoding/json"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

type DeviceInfo struct {
	dtos.Device
	Protocols map[string]map[string]interface{} `json:"protocols"`
}

// ToEdgeXV2Device converts the XRT model to EdgeX model
func ToEdgeXV2Device(device DeviceInfo, serviceName string) models.Device {
	// Convert all properties to string for EdgeX
	protocols := make(map[string]models.ProtocolProperties)
	protocolName := ""
	for protocol, protocolProperties := range device.Protocols {
		protocols[protocol] = toEdgeXProperties(protocol, protocolProperties)
		protocolName = strings.ToLower(protocol)
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

// ToXrtDevice converts the EdgeX model to XRT model
func ToXrtDevice(device models.Device) (deviceInfo DeviceInfo, edgexErr errors.EdgeX) {
	deviceData, err := json.Marshal(device)
	if err != nil {
		return deviceInfo, errors.NewCommonEdgeXWrapper(err)
	}
	err = json.Unmarshal(deviceData, &deviceInfo)
	if err != nil {
		return deviceInfo, errors.NewCommonEdgeXWrapper(err)
	}

	// Convert the EdgeX protocol properties to xrt protocol properties
	for protocol, v := range deviceInfo.Protocols {
		err = toXrtProperties(protocol, v)
		if err != nil {
			return deviceInfo, errors.NewCommonEdgeXWrapper(err)
		}
	}

	return deviceInfo, nil
}
