//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dbc

import (
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/require"
)

func TestConvertDBCtoDevice(t *testing.T) {
	networkName := "vcan0"
	serviceName := "device-can"

	ioReader, err := os.Open("dbc_sample.dbc")
	defer func() {
		err := ioReader.Close()
		require.NoError(t, err)
	}()
	require.NoError(t, err)

	data, err := io.ReadAll(ioReader)
	require.NoError(t, err)

	deviceDTOs, err, _ := ConvertDBCtoDevice(data, networkName, serviceName)
	require.NoError(t, err)
	require.NotEmpty(t, deviceDTOs)

	expectedDeviceDTO := dtos.Device{
		Name:           "EEC2",
		Description:    "Electronic Engine Controller 2",
		AdminState:     models.Unlocked,
		OperatingState: models.Up,
		ProfileName:    "EEC2",
		ServiceName:    serviceName,
		Protocols: map[string]dtos.ProtocolProperties{
			Canbus: {
				Network:  networkName,
				Standard: J1939,
				ID:       "217056254",
				DataSize: "8",
				Sender:   "Vector__XXX",
			},
		},
	}
	require.EqualValues(t, expectedDeviceDTO, deviceDTOs[0], "Generated Device DTO doesn't match the expected value.")
}

func TestConvertDBCtoProfile(t *testing.T) {
	ioReader, err := os.Open("dbc_sample.dbc")
	defer func() {
		err := ioReader.Close()
		require.NoError(t, err)
	}()
	require.NoError(t, err)

	data, err := io.ReadAll(ioReader)
	require.NoError(t, err)

	profileDTOs, err, _ := ConvertDBCtoProfile(data)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(profileDTOs) == 0 {
		t.Errorf("Expected 1 DeviceProfile, but got 0")
	}

	expectedProfileDTO := dtos.DeviceProfile{
		DeviceProfileBasicInfo: dtos.DeviceProfileBasicInfo{
			Name:        "EEC2",
			Description: "Electronic Engine Controller 2",
		},
		DeviceResources: []dtos.DeviceResource{
			{
				Name:        "Accelerator_Pedal_1_Low_Idle_Swi",
				Description: "Switch signal which indicates the state of the accelerator pedal 1 low idle switch.  The low idle switch is defined in SAE Recommended Practice J1843.",
				Properties: dtos.ResourceProperties{
					ValueType:    common.ValueTypeUint64,
					ReadWrite:    common.ReadWrite_R,
					Units:        "bit",
					Minimum:      "0",
					Maximum:      "3",
					Scale:        "1",
					Offset:       "0",
					DefaultValue: "0",
				},
				Attributes: map[string]interface{}{
					BitStart:      uint8(0),
					BitLen:        uint8(2),
					LittleEndian:  true,
					ReceiverNames: []string{"Vector__XXX"},
					MuxSignal:     false,
					IsSigned:      false,
				},
			},
		},
		DeviceCommands: []dtos.DeviceCommand{
			{
				Name:      "Accelerator_Pedal_1_Low_Idle_Swi",
				ReadWrite: common.ReadWrite_R,
				ResourceOperations: []dtos.ResourceOperation{
					{
						DeviceResource: "Accelerator_Pedal_1_Low_Idle_Swi",
						DefaultValue:   "0",
						Mappings: map[string]string{
							"0": "Accelerator pedal 1 not in low idle condition",
							"1": "Accelerator pedal 1 in low idle condition",
							"2": "Error",
							"3": "Not available",
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(profileDTOs[0], expectedProfileDTO) {
		t.Errorf("Generated DeviceProfile DTO doesn't match the expected value.")
	}
}
