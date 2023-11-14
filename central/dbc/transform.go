//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dbc

import (
	"math"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"go.einride.tech/can/pkg/descriptor"
)

func valueType(s *descriptor.Signal) string {
	_, offsetFrac := math.Modf(s.Offset)
	_, scaleFrac := math.Modf(s.Scale)
	if offsetFrac != 0 || scaleFrac != 0 {
		return common.ValueTypeFloat64
	}
	if s.IsSigned {
		return common.ValueTypeInt64
	} else {
		return common.ValueTypeUint64
	}
}

func ConvertDBCtoProfile(data []byte) (profileDTOs []dtos.DeviceProfile, err error, validateErrors map[string]error) {
	compileResult, err := Compile("", data)
	if err != nil {
		return
	}

	validateErrors = make(map[string]error, len(compileResult.Database.Messages))

	for _, m := range compileResult.Database.Messages {
		var deviceResources []dtos.DeviceResource
		var deviceCommands []dtos.DeviceCommand
		var profileDto dtos.DeviceProfile
		for _, s := range m.Signals {
			deviceResource := dtos.DeviceResource{
				Name:        s.Name,
				Description: s.Description,
				Properties: dtos.ResourceProperties{
					ValueType:    valueType(s),
					ReadWrite:    common.ReadWrite_R,
					Units:        s.Unit,
					Minimum:      strconv.FormatFloat(s.Min, 'f', -1, 64),
					Maximum:      strconv.FormatFloat(s.Max, 'f', -1, 64),
					Scale:        strconv.FormatFloat(s.Scale, 'f', -1, 64),
					Offset:       strconv.FormatFloat(s.Offset, 'f', -1, 64),
					DefaultValue: strconv.FormatInt(int64(s.DefaultValue), 10),
				},
				Attributes: map[string]interface{}{
					BitStart:      s.Start,
					BitLen:        s.Length,
					LittleEndian:  !s.IsBigEndian,
					ReceiverNames: s.ReceiverNodes,
					MuxSignal:     s.IsMultiplexer,
					IsSigned:      s.IsSigned,
				},
			}
			if s.IsMultiplexed {
				deviceResource.Attributes[MuxNum] = s.MultiplexerValue
			}
			if len(s.ValueDescriptions) > 0 {
				var deviceCommand dtos.DeviceCommand
				deviceCommand.Name = s.Name
				deviceCommand.ReadWrite = common.ReadWrite_R
				mappings := make(map[string]string, len(s.ValueDescriptions))
				for _, valueDescription := range s.ValueDescriptions {
					mappings[strconv.FormatInt(valueDescription.Value, 10)] = valueDescription.Description
				}
				deviceCommand.ResourceOperations = []dtos.ResourceOperation{
					{
						DeviceResource: s.Name,
						DefaultValue:   strconv.FormatInt(int64(s.DefaultValue), 10),
						Mappings:       mappings,
					},
				}
				deviceCommands = append(deviceCommands, deviceCommand)
			}
			deviceResources = append(deviceResources, deviceResource)
		}
		profileDto.Name = m.Name
		profileDto.Description = m.Description
		profileDto.DeviceResources = deviceResources
		profileDto.DeviceCommands = deviceCommands

		if validateErr := common.Validate(profileDto); validateErr != nil {
			validateErrors[profileDto.Name] = validateErr
		} else {
			profileDTOs = append(profileDTOs, profileDto)
		}
	}
	return
}

func ConvertDBCtoDevice(data []byte, networkName, serviceName string) (deviceDTOs []dtos.Device, err error, validateErrors map[string]error) {
	compileResult, err := Compile("", data)
	if err != nil {
		return
	}

	validateErrors = make(map[string]error, len(compileResult.Database.Messages))

	for _, m := range compileResult.Database.Messages {
		deviceDTO := dtos.Device{
			Name:           m.Name,
			Description:    m.Description,
			AdminState:     models.Unlocked,
			OperatingState: models.Up,
			ProfileName:    m.Name,
			ServiceName:    serviceName,
			Protocols: map[string]dtos.ProtocolProperties{
				Canbus: {
					Network:  networkName,
					Standard: J1939,
					ID:       strconv.Itoa(int(m.ID)),
					DataSize: strconv.Itoa(int(m.Length)),
					Sender:   m.SenderNode,
				},
			},
		}
		validateErr := common.Validate(deviceDTO)
		if validateErr != nil {
			validateErrors[deviceDTO.Name] = validateErr
		} else {
			deviceDTOs = append(deviceDTOs, deviceDTO)
		}
	}
	return
}
