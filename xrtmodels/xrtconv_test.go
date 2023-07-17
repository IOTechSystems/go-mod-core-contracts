// Copyright (C) 2022-2023 IOTech Ltd

package xrtmodels

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"

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
	assert.Equal(t, "1234", xrtDevice.Protocols[common.BacnetIP][common.BacnetDeviceInstance])
	assert.Equal(t, profileName, xrtDevice.ProfileName)
	assert.Equal(t, serviceName, xrtDevice.ServiceName)
	assert.Equal(t, models.Unlocked, xrtDevice.AdminState)
	assert.Equal(t, models.Up, xrtDevice.OperatingState)
}
