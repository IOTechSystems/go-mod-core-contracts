// Copyright (C) 2021-2024 IOTech Ltd

package xrtmodels

import (
	"encoding/json"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

type DeviceInfo struct {
	dtos.Device
	Protocols map[string]map[string]interface{} `json:"protocols"`
}

// ToEdgeXV2Device converts the XRT model to EdgeX model
func ToEdgeXV2Device(device DeviceInfo, serviceName string) (models.Device, errors.EdgeX) {
	// Convert all properties to string for EdgeX
	protocols := make(map[string]models.ProtocolProperties)
	protocolName := ""
	for protocol, protocolProperties := range device.Protocols {
		protocols[protocol] = toEdgeXProperties(protocol, protocolProperties)
		protocolName = strings.ToLower(protocol)
	}

	autoEvents, err := autoEventsFromDeviceProperties(device.Properties)
	if err != nil {
		return models.Device{}, errors.NewCommonEdgeXWrapper(err)
	}

	return models.Device{
		Name:           device.Name,
		Description:    device.Description,
		AdminState:     models.Unlocked,
		OperatingState: models.Up,
		ProtocolName:   protocolName,
		Protocols:      protocols,
		LastConnected:  0,
		LastReported:   0,
		Labels:         device.Labels,
		Location:       device.Location,
		Tags:           device.Tags,
		ServiceName:    serviceName,
		ProfileName:    device.ProfileName,
		AutoEvents:     dtos.ToAutoEventModels(autoEvents),
		Notify:         false,
		Properties:     device.Properties,
	}, nil
}

// ToEdgeXV2DeviceDTO converts the deviceInfo to EdgeX DTO
// for the edgeNode API, the xpert-manager can convert the deviceInfo to the edgex format DTO
func ToEdgeXV2DeviceDTO(device DeviceInfo) (dtos.Device, errors.EdgeX) {
	deviceDTO := device.Device

	// Convert all properties to string for EdgeX
	deviceDTO.Protocols = make(map[string]dtos.ProtocolProperties)
	protocolName := ""
	for protocol, protocolProperties := range device.Protocols {
		deviceDTO.Protocols[protocol] = toEdgeXProperties(protocol, protocolProperties)
		protocolName = strings.ToLower(protocol)
	}
	deviceDTO.ProtocolName = protocolName

	// Retrieve the auto events from the device properties
	autoEvents, err := autoEventsFromDeviceProperties(device.Properties)
	if err != nil {
		return deviceDTO, errors.NewCommonEdgeXWrapper(err)
	}
	deviceDTO.AutoEvents = autoEvents
	return deviceDTO, nil
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

	// Process the specified protocol for XRT
	for protocol := range deviceInfo.Protocols {
		switch protocol {
		case common.EtherNetIP:
			processEtherNetIP(deviceInfo.Protocols)
		}
	}

	// Store the auto events data to device properties because the XRT not allow to update the autoEvents field in the state/device.json
	if deviceInfo.Properties == nil {
		deviceInfo.Properties = map[string]any{}
	}
	if deviceInfo.AutoEvents != nil {
		deviceInfo.Properties[common.IOTechPrefix+common.AutoEvents] = device.AutoEvents
		deviceInfo.AutoEvents = nil
	}

	return deviceInfo, nil
}

// autoEventsFromDeviceProperties extracts the auto events from the device properties
func autoEventsFromDeviceProperties(properties map[string]any) ([]dtos.AutoEvent, errors.EdgeX) {
	if autoEventsProperty, ok := properties[common.IOTechPrefix+common.AutoEvents]; ok {
		data, err := json.Marshal(autoEventsProperty)
		if err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindServerError, "failed to marshal the auto events from the device properties", err)
		}
		var autoEvents []dtos.AutoEvent
		err = json.Unmarshal(data, &autoEvents)
		if err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindServerError, "failed to unmarshal the auto events data", err)
		}
		delete(properties, common.IOTechPrefix+common.AutoEvents)
		return autoEvents, nil
	}
	return nil, nil
}

func processEtherNetIP(protocolProperties map[string]map[string]interface{}) {
	// Combine ExplicitConnected, O2T, T2O and Key into EtherNet-IP
	if v, ok := protocolProperties[common.EtherNetIP]; ok {
		protocolProperties[common.EtherNetIPXRT] = v
		delete(protocolProperties, common.EtherNetIP)
	}
	if v, ok := protocolProperties[common.EtherNetIPExplicitConnected]; ok {
		protocolProperties[common.EtherNetIPXRT][common.EtherNetIPExplicitConnected] = v
		delete(protocolProperties, common.EtherNetIPExplicitConnected)
	}
	if v, ok := protocolProperties[common.EtherNetIPO2T]; ok {
		protocolProperties[common.EtherNetIPXRT][common.EtherNetIPO2T] = v
		delete(protocolProperties, common.EtherNetIPO2T)
	}
	if v, ok := protocolProperties[common.EtherNetIPT2O]; ok {
		protocolProperties[common.EtherNetIPXRT][common.EtherNetIPT2O] = v
		delete(protocolProperties, common.EtherNetIPT2O)
	}
	if v, ok := protocolProperties[common.EtherNetIPKey]; ok {
		protocolProperties[common.EtherNetIPXRT][common.EtherNetIPKey] = v
		delete(protocolProperties, common.EtherNetIPKey)
	}
}
