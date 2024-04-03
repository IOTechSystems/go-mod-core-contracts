// Copyright (C) 2022-2024 IOTech Ltd

package xrtmodels

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testDeviceInfo() DeviceInfo {
	return DeviceInfo{
		Device: dtos.Device{
			Name:           "test-name",
			Description:    "test description",
			AdminState:     "UNLOCKED",
			OperatingState: "UP",
			Labels:         []string{"test label"},
			Location:       []string{"test location"},
			Tags:           map[string]any{"test": "tag"},
			ServiceName:    "device-modbus",
			ProfileName:    "test-profile",
			AutoEvents:     nil,
			ProtocolName:   "modbus",
			Protocols:      nil,
			Properties: map[string]any{
				common.IOTechPrefix + common.AutoEvents: []map[string]any{
					{"interval": "1h", "onChange": false, "sourceName": "source1"},
					{"interval": "1m", "onChange": true, "sourceName": "source2"},
				},
			},
		},
		Protocols: map[string]map[string]interface{}{
			"modbus": {
				"Address": "127.0.0.1",
				"Port":    1502,
				"UnitID":  1,
			},
		},
	}
}

func TestProcessEtherNetIP(t *testing.T) {
	tests := []struct {
		name     string
		protocol map[string]map[string]interface{}
		expected map[string]map[string]interface{}
	}{
		{
			name: "process O2T and T2O properties",
			protocol: map[string]map[string]interface{}{
				common.EtherNetIP: {
					common.EtherNetIPAddress: "127.0.0.1",
				},
				common.EtherNetIPO2T: {
					common.EtherNetIPConnectionType: "p2p",
					common.EtherNetIPRPI:            10,
					common.EtherNetIPPriority:       "low",
					common.EtherNetIPOwnership:      "exclusive",
				},
				common.EtherNetIPT2O: {
					common.EtherNetIPConnectionType: "p2p",
					common.EtherNetIPRPI:            10,
					common.EtherNetIPPriority:       "low",
					common.EtherNetIPOwnership:      "exclusive",
				},
			},
			expected: map[string]map[string]interface{}{
				common.EtherNetIPXRT: {
					common.EtherNetIPAddress: "127.0.0.1",
					common.EtherNetIPO2T: map[string]interface{}{
						common.EtherNetIPConnectionType: "p2p",
						common.EtherNetIPRPI:            10,
						common.EtherNetIPPriority:       "low",
						common.EtherNetIPOwnership:      "exclusive",
					},
					common.EtherNetIPT2O: map[string]interface{}{
						common.EtherNetIPConnectionType: "p2p",
						common.EtherNetIPRPI:            10,
						common.EtherNetIPPriority:       "low",
						common.EtherNetIPOwnership:      "exclusive",
					},
				},
			},
		},
		{
			name: "process ExplicitConnected and Key properties",
			protocol: map[string]map[string]interface{}{
				common.EtherNetIP: {
					common.EtherNetIPAddress: "127.0.0.1",
				},
				common.EtherNetIPExplicitConnected: {
					common.EtherNetIPDeviceResource: "VendorID",
					common.EtherNetIPRPI:            10,
					common.EtherNetIPSaveValue:      true,
				},
				common.EtherNetIPKey: {
					common.EtherNetIPMethod:        "exact",
					common.EtherNetIPVendorID:      10,
					common.EtherNetIPDeviceType:    72,
					common.EtherNetIPProductCode:   50,
					common.EtherNetIPMajorRevision: 12,
					common.EtherNetIPMinorRevision: 2,
				},
			},
			expected: map[string]map[string]interface{}{
				common.EtherNetIPXRT: {
					common.EtherNetIPAddress: "127.0.0.1",
					common.EtherNetIPExplicitConnected: map[string]interface{}{
						common.EtherNetIPDeviceResource: "VendorID",
						common.EtherNetIPRPI:            10,
						common.EtherNetIPSaveValue:      true,
					},
					common.EtherNetIPKey: map[string]interface{}{
						common.EtherNetIPMethod:        "exact",
						common.EtherNetIPVendorID:      10,
						common.EtherNetIPDeviceType:    72,
						common.EtherNetIPProductCode:   50,
						common.EtherNetIPMajorRevision: 12,
						common.EtherNetIPMinorRevision: 2,
					},
				},
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			processEtherNetIP(testCase.protocol)
			assert.EqualValues(t, testCase.expected, testCase.protocol)
		})
	}
}

func TestToEdgeXV2Device(t *testing.T) {
	deviceInfo := testDeviceInfo()
	res, err := ToEdgeXV2Device(deviceInfo, deviceInfo.ServiceName)
	require.NoError(t, err)
	assert.EqualValues(t, deviceInfo.Name, res.Name)
	assert.EqualValues(t, deviceInfo.Description, res.Description)
	assert.EqualValues(t, deviceInfo.AdminState, res.AdminState)
	assert.EqualValues(t, deviceInfo.OperatingState, res.OperatingState)
	assert.EqualValues(t, deviceInfo.Labels, res.Labels)
	assert.EqualValues(t, deviceInfo.Location, res.Location)
	assert.EqualValues(t, deviceInfo.ServiceName, res.ServiceName)
	assert.EqualValues(t, deviceInfo.ProfileName, res.ProfileName)
	assert.EqualValues(t, deviceInfo.Tags, res.Tags)
	assert.EqualValues(t, deviceInfo.ProtocolName, res.ProtocolName)
	assert.EqualValues(t, map[string]models.ProtocolProperties{"modbus": {"Address": "127.0.0.1", "Port": "1502", "UnitID": "1"}}, res.Protocols)
	assert.EqualValues(t, []models.AutoEvent{{Interval: "1h", OnChange: false, SourceName: "source1"}, {Interval: "1m", OnChange: true, SourceName: "source2"}}, res.AutoEvents)
}

func TestToEdgeXV2DeviceDTO(t *testing.T) {
	deviceInfo := testDeviceInfo()
	res, err := ToEdgeXV2DeviceDTO(deviceInfo)
	require.NoError(t, err)
	assert.EqualValues(t, deviceInfo.Name, res.Name)
	assert.EqualValues(t, deviceInfo.Description, res.Description)
	assert.EqualValues(t, deviceInfo.AdminState, res.AdminState)
	assert.EqualValues(t, deviceInfo.OperatingState, res.OperatingState)
	assert.EqualValues(t, deviceInfo.Labels, res.Labels)
	assert.EqualValues(t, deviceInfo.Location, res.Location)
	assert.EqualValues(t, deviceInfo.ServiceName, res.ServiceName)
	assert.EqualValues(t, deviceInfo.ProfileName, res.ProfileName)
	assert.EqualValues(t, deviceInfo.Tags, res.Tags)
	assert.EqualValues(t, deviceInfo.ProtocolName, res.ProtocolName)
	assert.EqualValues(t, map[string]dtos.ProtocolProperties{"modbus": {"Address": "127.0.0.1", "Port": "1502", "UnitID": "1"}}, res.Protocols)
	assert.EqualValues(t, []dtos.AutoEvent{{Interval: "1h", OnChange: false, SourceName: "source1"}, {Interval: "1m", OnChange: true, SourceName: "source2"}}, res.AutoEvents)
}
