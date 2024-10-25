// Copyright (C) 2022-2024 IOTech Ltd

package xrtmodels

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

func TestProcessEtherNetIP(t *testing.T) {
	tests := []struct {
		name     string
		protocol map[string]dtos.ProtocolProperties
		expected map[string]dtos.ProtocolProperties
	}{
		{
			name: "process O2T and T2O properties",
			protocol: map[string]dtos.ProtocolProperties{
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
			expected: map[string]dtos.ProtocolProperties{
				common.EtherNetIPXRT: {
					common.EtherNetIPAddress: "127.0.0.1",
					common.EtherNetIPO2T: dtos.ProtocolProperties{
						common.EtherNetIPConnectionType: "p2p",
						common.EtherNetIPRPI:            10,
						common.EtherNetIPPriority:       "low",
						common.EtherNetIPOwnership:      "exclusive",
					},
					common.EtherNetIPT2O: dtos.ProtocolProperties{
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
			protocol: map[string]dtos.ProtocolProperties{
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
			expected: map[string]dtos.ProtocolProperties{
				common.EtherNetIPXRT: {
					common.EtherNetIPAddress: "127.0.0.1",
					common.EtherNetIPExplicitConnected: dtos.ProtocolProperties{
						common.EtherNetIPDeviceResource: "VendorID",
						common.EtherNetIPRPI:            10,
						common.EtherNetIPSaveValue:      true,
					},
					common.EtherNetIPKey: dtos.ProtocolProperties{
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

func TestToEdgeXV3Device(t *testing.T) {
	device := DeviceInfo{
		Device: dtos.Device{
			Name:           "test-ble-device",
			AdminState:     models.Unlocked,
			OperatingState: models.Up,
			ServiceName:    "device-ble",
			ProfileName:    "test-ble-profile",
			Protocols: map[string]dtos.ProtocolProperties{
				"BLE": {
					"MAC": "00:00:00:00:00:00",
				},
			},
		},
	}

	result := ToEdgeXV3Device(device, device.ServiceName)

	assert.Equal(t, device.Name, result.Name)
	assert.Equal(t, device.AdminState, result.AdminState)
	assert.Equal(t, device.OperatingState, result.OperatingState)
	assert.Equal(t, device.ServiceName, result.ServiceName)
	assert.Equal(t, device.Protocols, result.Protocols)
	assert.Equal(t, map[string]any{common.ProtocolName: "ble"}, result.Properties)
}
