// Copyright (C) 2022-2023 IOTech Ltd

package xrtmodels

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToXrtDevice(t *testing.T) {
	deviceName := "test-device"
	profileName := "test-profile"
	serviceName := "device-bacnet-ip"
	device := models.Device{
		Name:         deviceName,
		ProtocolName: common.BacnetIP,
		Protocols: map[string]models.ProtocolProperties{
			common.BacnetIP: {
				common.BacnetDeviceInstance: "1234",
			},
		},
		ProfileName:    profileName,
		ServiceName:    serviceName,
		AdminState:     models.Unlocked,
		OperatingState: models.Up,
	}
	xrtDevice, err := ToXrtDevice(device)
	require.NoError(t, err)

	assert.Equal(t, deviceName, xrtDevice.Name)
	assert.Equal(t, common.BacnetIP, xrtDevice.ProtocolName)
	assert.Equal(t, 1234, xrtDevice.Protocols[common.BacnetIP][common.BacnetDeviceInstance])
	assert.Equal(t, profileName, xrtDevice.ProfileName)
	assert.Equal(t, serviceName, xrtDevice.ServiceName)
	assert.Equal(t, models.Unlocked, xrtDevice.AdminState)
	assert.Equal(t, models.Up, xrtDevice.OperatingState)
}

func TestToXrtProperties(t *testing.T) {
	tests := []struct {
		protocol   string
		properties map[string]interface{}
		expected   map[string]interface{}
	}{
		{
			protocol: common.BacnetIP,
			properties: map[string]interface{}{
				common.BacnetDeviceInstance: "1234",
			},
			expected: map[string]interface{}{
				common.BacnetDeviceInstance: 1234,
			},
		},
		{
			protocol: common.BacnetIP,
			properties: map[string]interface{}{
				common.BacnetDeviceInstance: "4194302",
			},
			expected: map[string]interface{}{
				common.BacnetDeviceInstance: 4194302,
			},
		},
		{
			protocol: common.BacnetIP,
			properties: map[string]interface{}{
				common.BacnetAddress: "192.168.60.123",
				common.BacnetPort:    "47809",
			},
			expected: map[string]interface{}{
				common.BacnetAddress: "192.168.60.123",
				common.BacnetPort:    47809,
			},
		},
		{
			protocol: common.Opcua,
			properties: map[string]interface{}{
				common.OpcuaRequestedSessionTimeout:    "1200000",
				common.OpcuaBrowseDepth:                "0",
				common.OpcuaBrowsePublishInterval:      "5.2",
				common.OpcuaConnectionReadingPostDelay: "0",
				common.OpcuaIDType:                     "1",
				common.OpcuaReadBatchSize:              "15",
				common.OpcuaWriteBatchSize:             "10",
				common.OpcuaNodesPerBrowse:             "5",
				common.OpcuaSessionKeepAliveInterval:   "1000",
			},
			expected: map[string]interface{}{
				common.OpcuaRequestedSessionTimeout:    1200000,
				common.OpcuaBrowseDepth:                0,
				common.OpcuaBrowsePublishInterval:      5.2,
				common.OpcuaConnectionReadingPostDelay: 0,
				common.OpcuaIDType:                     "1",
				common.OpcuaReadBatchSize:              15,
				common.OpcuaWriteBatchSize:             10,
				common.OpcuaNodesPerBrowse:             5,
				common.OpcuaSessionKeepAliveInterval:   1000.0,
			},
		},
		{
			protocol: common.ModbusTcp,
			properties: map[string]interface{}{
				common.ModbusUnitID:                    "1",
				common.ModbusPort:                      "1234",
				common.ModbusReadMaxHoldingRegisters:   "125",
				common.ModbusReadMaxInputRegisters:     "125",
				common.ModbusReadMaxBitsCoils:          "2000",
				common.ModbusReadMaxBitsDiscreteInputs: "2000",
				common.ModbusWriteMaxHoldingRegisters:  "123",
				common.ModbusWriteMaxBitsCoils:         "1968",
			},
			expected: map[string]interface{}{
				common.ModbusUnitID:                    1,
				common.ModbusPort:                      1234,
				common.ModbusReadMaxHoldingRegisters:   125,
				common.ModbusReadMaxInputRegisters:     125,
				common.ModbusReadMaxBitsCoils:          2000,
				common.ModbusReadMaxBitsDiscreteInputs: 2000,
				common.ModbusWriteMaxHoldingRegisters:  123,
				common.ModbusWriteMaxBitsCoils:         1968,
			},
		},
		{
			protocol: common.ModbusRtu,
			properties: map[string]interface{}{
				common.ModbusUnitID:                    "1",
				common.ModbusBaudRate:                  "0",
				common.ModbusDataBits:                  "5",
				common.ModbusStopBits:                  "1",
				common.ModbusReadMaxHoldingRegisters:   "125",
				common.ModbusReadMaxInputRegisters:     "125",
				common.ModbusReadMaxBitsCoils:          "2000",
				common.ModbusReadMaxBitsDiscreteInputs: "2000",
				common.ModbusWriteMaxHoldingRegisters:  "123",
				common.ModbusWriteMaxBitsCoils:         "1968",
			},
			expected: map[string]interface{}{
				common.ModbusUnitID:                    1,
				common.ModbusBaudRate:                  0,
				common.ModbusDataBits:                  5,
				common.ModbusStopBits:                  1,
				common.ModbusReadMaxHoldingRegisters:   125,
				common.ModbusReadMaxInputRegisters:     125,
				common.ModbusReadMaxBitsCoils:          2000,
				common.ModbusReadMaxBitsDiscreteInputs: 2000,
				common.ModbusWriteMaxHoldingRegisters:  123,
				common.ModbusWriteMaxBitsCoils:         1968,
			},
		},
		{
			protocol: common.EtherNetIPExplicitConnected,
			properties: map[string]interface{}{
				common.EtherNetIPDeviceResource: "VendorID",
				common.EtherNetIPRPI:            "3000",
				common.EtherNetIPSaveValue:      "true",
			},
			expected: map[string]interface{}{
				common.EtherNetIPDeviceResource: "VendorID",
				common.EtherNetIPRPI:            3000,
				common.EtherNetIPSaveValue:      true,
			},
		},
		{
			protocol: common.EtherNetIPO2T,
			properties: map[string]interface{}{
				common.EtherNetIPConnectionType: "p2p",
				common.EtherNetIPRPI:            "10",
				common.EtherNetIPPriority:       "low",
				common.EtherNetIPOwnership:      "exclusive",
			},
			expected: map[string]interface{}{
				common.EtherNetIPConnectionType: "p2p",
				common.EtherNetIPRPI:            10,
				common.EtherNetIPPriority:       "low",
				common.EtherNetIPOwnership:      "exclusive",
			},
		},
		{
			protocol: common.EtherNetIPT2O,
			properties: map[string]interface{}{
				common.EtherNetIPConnectionType: "p2p",
				common.EtherNetIPRPI:            "10",
				common.EtherNetIPPriority:       "low",
				common.EtherNetIPOwnership:      "exclusive",
			},
			expected: map[string]interface{}{
				common.EtherNetIPConnectionType: "p2p",
				common.EtherNetIPRPI:            10,
				common.EtherNetIPPriority:       "low",
				common.EtherNetIPOwnership:      "exclusive",
			},
		},
		{
			protocol: common.EtherNetIPKey,
			properties: map[string]interface{}{
				common.EtherNetIPMethod:        "exact",
				common.EtherNetIPVendorID:      "10",
				common.EtherNetIPDeviceType:    "72",
				common.EtherNetIPProductCode:   "50",
				common.EtherNetIPMajorRevision: "12",
				common.EtherNetIPMinorRevision: "2",
			},
			expected: map[string]interface{}{
				common.EtherNetIPMethod:        "exact",
				common.EtherNetIPVendorID:      10,
				common.EtherNetIPDeviceType:    72,
				common.EtherNetIPProductCode:   50,
				common.EtherNetIPMajorRevision: 12,
				common.EtherNetIPMinorRevision: 2,
			},
		},
		{
			protocol: common.EtherCAT,
			properties: map[string]interface{}{
				common.EtherCATNetworkIndex: "1",
				common.EtherCATSyncMode:     "3",
			},
			expected: map[string]interface{}{
				common.EtherCATNetworkIndex: 1,
				common.EtherCATSyncMode:     3,
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.protocol, func(t *testing.T) {
			err := toXrtProperties(testCase.protocol, testCase.properties)
			require.NoError(t, err)
			assert.EqualValues(t, testCase.expected, testCase.properties)
		})
	}
}

func TestToXrType_Invalid(t *testing.T) {
	tests := []struct {
		protocol   string
		properties map[string]interface{}
	}{
		{
			protocol: common.ModbusTcp,
			properties: map[string]interface{}{
				common.ModbusPort: "test",
			},
		},
		{
			protocol: common.Opcua,
			properties: map[string]interface{}{
				common.OpcuaBrowsePublishInterval: "test",
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.protocol, func(t *testing.T) {
			err := toXrtProperties(testCase.protocol, testCase.properties)
			require.Error(t, err)
		})
	}
}

func TestToEdgeXProperties(t *testing.T) {
	tests := []struct {
		protocol   string
		properties map[string]interface{}
		expected   map[string]string
	}{
		{
			protocol: common.BacnetIP,
			properties: map[string]interface{}{
				common.BacnetDeviceInstance: float64(1234),
			},
			expected: map[string]string{
				common.BacnetDeviceInstance: "1234",
			},
		},
		{
			protocol: common.BacnetIP,
			properties: map[string]interface{}{
				common.BacnetDeviceInstance: float64(4194302),
			},
			expected: map[string]string{
				common.BacnetDeviceInstance: "4194302",
			},
		},
		{
			protocol: common.Opcua,
			properties: map[string]interface{}{
				common.OpcuaRequestedSessionTimeout:    float64(1200000),
				common.OpcuaBrowseDepth:                float64(0),
				common.OpcuaBrowsePublishInterval:      5.2,
				common.OpcuaConnectionReadingPostDelay: float64(0),
				common.OpcuaIDType:                     float64(1),
			},
			expected: map[string]string{
				common.OpcuaRequestedSessionTimeout:    "1200000",
				common.OpcuaBrowseDepth:                "0",
				common.OpcuaBrowsePublishInterval:      "5.2",
				common.OpcuaConnectionReadingPostDelay: "0",
				common.OpcuaIDType:                     "1",
			},
		},
		{
			protocol: common.ModbusTcp,
			properties: map[string]interface{}{
				common.ModbusUnitID: float64(1),
				common.ModbusPort:   float64(1234),
			},
			expected: map[string]string{
				common.ModbusUnitID: "1",
				common.ModbusPort:   "1234",
			},
		},
		{
			protocol: common.ModbusRtu,
			properties: map[string]interface{}{
				common.ModbusUnitID:   float64(1),
				common.ModbusBaudRate: float64(0),
				common.ModbusDataBits: float64(5),
				common.ModbusStopBits: float64(1),
			},
			expected: map[string]string{
				common.ModbusUnitID:   "1",
				common.ModbusBaudRate: "0",
				common.ModbusDataBits: "5",
				common.ModbusStopBits: "1",
			},
		},
		{
			protocol: common.EtherCAT,
			properties: map[string]interface{}{
				common.EtherCATNetworkIndex: float64(1),
				common.EtherCATSyncMode:     float64(3),
			},
			expected: map[string]string{
				common.EtherCATNetworkIndex: "1",
				common.EtherCATSyncMode:     "3",
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.protocol, func(t *testing.T) {
			result := toEdgeXProperties(testCase.protocol, testCase.properties)
			assert.EqualValues(t, testCase.expected, result)
		})
	}
}
