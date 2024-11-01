// Copyright (C) 2021-2024 IOTech Ltd

package v1models

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/v2models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	TestDescription = "TestDescription"

	TestSourceSwitchName                 = "Switch"
	TestSourceSwitchDescription          = "On/Off , 0-OFF 1-ON"
	TestSourceOperationModeName          = "OperationMode"
	TestSourceOperationModeDescription   = "0-Cool 1-Heat 2-Auto 3-Dry 4-HAUX 5-Fan 6-HH 8-VAM Auto 9-VAM Bypass 10-VAM Heat Exc 11-VAM Normal"
	TestSourceRoomTemperatureName        = "RoomTemperature"
	TestSourceRoomTemperatureDescription = "Room Temperature x10 °C (Read Only)"
	TestSourceTemperatureName            = "Temperature"
	TestSourceTemperatureDescription     = "Temperature x10 °C"

	TestDeviceCommandValuesName = "Values"
)

var (
	TestSourceSwitchTags          = map[string]string{"source": "switch"}
	TestSourceOperationModeTags   = map[string]string{"source": "operation mode"}
	TestSourceRoomTemperatureTags = map[string]string{"source": "room temperature"}
	TestSourceTemperatureTags     = map[string]string{"source": "temperature"}
)

var testLabels = []string{"HVAC", "Air conditioner"}

var testV1DeviceResources = []DeviceResource{
	{
		Name:        TestSourceSwitchName,
		Description: TestSourceSwitchDescription,
		Tags:        TestSourceSwitchTags,
		Attributes: map[string]string{
			"primaryTable": "COILS", "startingAddress": "1",
		},
		Properties: ProfileProperty{
			Value: PropertyValue{
				Type:         "Bool",
				ReadWrite:    common.ReadWrite_RW,
				DefaultValue: "true",
			},
			Units: Units{
				Type:         "String",
				ReadWrite:    common.ReadWrite_R,
				DefaultValue: "On/Off",
			},
		},
	}, {
		Name:        TestSourceOperationModeName,
		Description: TestSourceOperationModeDescription,
		Tags:        TestSourceOperationModeTags,
		Attributes: map[string]string{
			"primaryTable": "HOLDING_REGISTERS", "startingAddress": "2",
		},
		Properties: ProfileProperty{
			Value: PropertyValue{
				Type:      "Int16",
				ReadWrite: common.ReadWrite_RW,
			},
			Units: Units{
				Type:         "String",
				ReadWrite:    common.ReadWrite_R,
				DefaultValue: "Operation Mode",
			},
		},
	}, {
		Name:        TestSourceRoomTemperatureName,
		Description: TestSourceRoomTemperatureDescription,
		Tags:        TestSourceRoomTemperatureTags,
		Attributes: map[string]string{
			"primaryTable": "INPUT_REGISTERS", "startingAddress": "4",
		},
		Properties: ProfileProperty{
			Value: PropertyValue{
				Type:          "Float32",
				ReadWrite:     common.ReadWrite_R,
				Scale:         "0.1",
				FloatEncoding: "eNotation",
			},
			Units: Units{
				Type:         "String",
				ReadWrite:    common.ReadWrite_R,
				DefaultValue: "degrees Celsius",
			},
		},
	}, {
		Name:        TestSourceTemperatureName,
		Description: TestSourceTemperatureDescription,
		Tags:        TestSourceTemperatureTags,
		Attributes: map[string]string{
			"primaryTable": "HOLDING_REGISTERS", "startingAddress": "5",
		},
		Properties: ProfileProperty{
			Value: PropertyValue{
				Type:          "Float64",
				ReadWrite:     common.ReadWrite_RW,
				Scale:         "0.1",
				FloatEncoding: "eNotation",
			},
			Units: Units{
				Type:         "String",
				ReadWrite:    common.ReadWrite_R,
				DefaultValue: "degrees Celsius",
			},
		},
	}}

func v1ProfileData() DeviceProfile {

	var testDeviceCommands = []ProfileResource{{
		Name: TestDeviceCommandValuesName,
		Get: []ResourceOperation{{
			Index: "0", Operation: "get", DeviceResource: TestSourceSwitchName,
			Mappings: map[string]string{"true": "ON", "false": "OFF"},
		}, {
			Index: "1", Operation: "get", DeviceResource: TestSourceOperationModeName,
			Mappings: map[string]string{
				"0": "Cool", "1": "Heat", "2": "Auto", "3": "Dry", "4": "HAUX", "5": "Fan", "6": "HH", "8": "VAM Auto", "9": "VAM Bypass", "10": "VAM Heat", "11": "VAM Normal",
			},
		}, {
			Index: "2", Operation: "get", DeviceResource: TestSourceTemperatureName,
		}},
		Set: []ResourceOperation{{
			Index: "0", Operation: "set", DeviceResource: TestSourceSwitchName,
			Mappings: map[string]string{"ON": "true", "OFF": "false"},
		}, {
			Index: "1", Operation: "set", DeviceResource: TestSourceOperationModeName,
			Mappings: map[string]string{
				"Cool": "0", "Heat": "1", "Auto": "2", "Dry": "3", "HAUX": "4", "Fan": "5", "HH": "6", "VAM Auto": "8", "VAM Bypass": "9", "VAM Heat": "10", "VAM Normal": "11",
			},
		}, {
			Index: "2", Operation: "set", DeviceResource: TestSourceTemperatureName,
		}},
	}}

	var testCoreCommands = []Command{
		{
			Name: TestDeviceCommandValuesName,
			Get: Get{
				Action: Action{
					Path: "/api/v1/device/{deviceId}/Values",
					Responses: []Response{
						{
							Code: "200", Description: "Issue the Get command Values",
							ExpectedValues: []string{TestSourceSwitchName, TestSourceOperationModeName, TestSourceTemperatureName},
						}, {
							Code: "500", Description: "internal server error",
						},
					},
				},
			},
			Put: Put{
				Action: Action{
					Path: "/api/v1/device/{deviceId}/Values",
					Responses: []Response{
						{
							Code: "204", Description: "Issue the Put command Values",
						}, {
							Code: "500", Description: "internal server error",
						},
					},
				},
				ParameterNames: []string{TestSourceSwitchName, TestSourceOperationModeName, TestSourceTemperatureName},
			},
		},
	}

	return DeviceProfile{
		Name:         TestProfileName,
		Manufacturer: TestManufacturer,
		Model:        TestModel,
		Labels:       testLabels,
		DescribedObject: DescribedObject{
			Description: TestDescription,
		},

		DeviceResources: testV1DeviceResources,
		DeviceCommands:  testDeviceCommands,
		CoreCommands:    testCoreCommands,
	}
}

func v2ProfileData() v2models.DeviceProfile {
	var testDeviceResources = []v2models.DeviceResource{{
		Name:        TestSourceSwitchName,
		IsHidden:    true,
		Description: TestSourceSwitchDescription,
		Tags:        toV2Tags(TestSourceSwitchTags),
		Attributes: map[string]interface{}{
			"primaryTable": "COILS", "startingAddress": 0,
		},
		Properties: v2models.ResourceProperties{
			ValueType:    "Bool",
			ReadWrite:    common.ReadWrite_RW,
			DefaultValue: "true",
			Units:        "On/Off",
		},
	}, {
		Name:        TestSourceOperationModeName,
		IsHidden:    true,
		Description: TestSourceOperationModeDescription,
		Tags:        toV2Tags(TestSourceOperationModeTags),
		Attributes: map[string]interface{}{
			"primaryTable": "HOLDING_REGISTERS", "startingAddress": 1,
		},
		Properties: v2models.ResourceProperties{
			ValueType: "Int16",
			ReadWrite: common.ReadWrite_RW,
			Units:     "Operation Mode",
		},
	}, {
		Name:        TestSourceRoomTemperatureName,
		IsHidden:    true,
		Description: TestSourceRoomTemperatureDescription,
		Tags:        toV2Tags(TestSourceRoomTemperatureTags),
		Attributes: map[string]interface{}{
			"primaryTable": "INPUT_REGISTERS", "startingAddress": 3,
		},
		Properties: v2models.ResourceProperties{
			ValueType: "Float32",
			ReadWrite: common.ReadWrite_R,
			Scale:     "0.1",
			Units:     "degrees Celsius",
		},
	}, {
		Name:        TestSourceTemperatureName,
		IsHidden:    true,
		Description: TestSourceTemperatureDescription,
		Tags:        toV2Tags(TestSourceTemperatureTags),
		Attributes: map[string]interface{}{
			"primaryTable": "HOLDING_REGISTERS", "startingAddress": 4,
		},
		Properties: v2models.ResourceProperties{
			ValueType: "Float64",
			ReadWrite: common.ReadWrite_RW,
			Scale:     "0.1",
			Units:     "degrees Celsius",
		},
	}}

	var testDeviceCommands = []v2models.DeviceCommand{{
		Name:      TestDeviceCommandValuesName,
		ReadWrite: common.ReadWrite_RW,
		ResourceOperations: []v2models.ResourceOperation{
			{
				DeviceResource: TestSourceSwitchName,
				Mappings:       map[string]string{"true": "ON", "false": "OFF"},
			}, {
				DeviceResource: TestSourceOperationModeName,
				Mappings: map[string]string{
					"0": "Cool", "1": "Heat", "2": "Auto", "3": "Dry", "4": "HAUX", "5": "Fan", "6": "HH", "8": "VAM Auto", "9": "VAM Bypass", "10": "VAM Heat", "11": "VAM Normal",
				},
			}, {
				DeviceResource: TestSourceTemperatureName,
			},
		},
	}}
	return v2models.DeviceProfile{
		ApiVersion:   common.ApiVersion,
		Name:         TestProfileName,
		Manufacturer: TestManufacturer,
		Model:        TestModel,
		Labels:       testLabels,
		Description:  TestDescription,

		DeviceResources: testDeviceResources,
		DeviceCommands:  testDeviceCommands,
	}
}

func TestTransformProfileFromV1ToV2(t *testing.T) {
	data := v1ProfileData()
	data.DeviceCommands = []ProfileResource{{
		Name: TestDeviceCommandValuesName,
		Get: []ResourceOperation{{
			Index: "0", Operation: "get", DeviceResource: TestSourceSwitchName,
			Mappings: map[string]string{"true": "ON", "false": "OFF"},
		}},
		Set: []ResourceOperation{{
			Index: "0", Operation: "set", DeviceResource: TestSourceSwitchName,
			Mappings: map[string]string{"ON": "true", "OFF": "false"},
		}, {
			Index: "1", Operation: "set", DeviceResource: TestSourceOperationModeName,
			Mappings: map[string]string{
				"Cool": "0", "Heat": "1", "Auto": "2", "Dry": "3", "HAUX": "4", "Fan": "5", "HH": "6", "VAM Auto": "8", "VAM Bypass": "9", "VAM Heat": "10", "VAM Normal": "11",
			},
		}},
	}}
	data.CoreCommands = []Command{
		{
			Name: TestDeviceCommandValuesName,
			Get: Get{
				Action: Action{
					Path: "/api/v1/device/{deviceId}/Values",
					Responses: []Response{
						{
							Code: "200", Description: "Issue the Get command Values",
							ExpectedValues: []string{TestSourceSwitchName},
						}, {
							Code: "500", Description: "internal server error",
						},
					},
				},
			},
			Put: Put{
				Action: Action{
					Path: "/api/v1/device/{deviceId}/Values",
					Responses: []Response{
						{
							Code: "204", Description: "Issue the Put command Values",
						}, {
							Code: "500", Description: "internal server error",
						},
					},
				},
				ParameterNames: []string{TestSourceSwitchName, TestSourceOperationModeName},
			},
		},
	}
	expected := v2ProfileData()
	expected.DeviceCommands = []v2models.DeviceCommand{{
		Name:      TestDeviceCommandValuesName,
		IsHidden:  false,
		ReadWrite: common.ReadWrite_R,
		ResourceOperations: []v2models.ResourceOperation{
			{
				DeviceResource: TestSourceSwitchName,
				Mappings:       map[string]string{"true": "ON", "false": "OFF"},
			},
		},
	}, {
		Name:      v2SetCommandName(TestDeviceCommandValuesName),
		IsHidden:  false,
		ReadWrite: common.ReadWrite_W,
		ResourceOperations: []v2models.ResourceOperation{
			{
				DeviceResource: TestSourceSwitchName,
				Mappings:       map[string]string{"true": "ON", "false": "OFF"},
			}, {
				DeviceResource: TestSourceOperationModeName,
				Mappings: map[string]string{
					"0": "Cool", "1": "Heat", "2": "Auto", "3": "Dry", "4": "HAUX", "5": "Fan", "6": "HH", "8": "VAM Auto", "9": "VAM Bypass", "10": "VAM Heat", "11": "VAM Normal",
				},
			},
		},
	}}

	var tests = []struct {
		name     string
		data     DeviceProfile
		expected v2models.DeviceProfile
	}{
		{"transform profile from v1 to v2", v1ProfileData(), v2ProfileData()},
		{"get operation size is different to set operation", data, expected},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := TransformProfileFromV1ToV2(testCase.data)
			require.NoError(t, err)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestTransformProfileFromV2ToV1(t *testing.T) {
	expected := v1ProfileData()
	data := v2ProfileData()

	actual, err := TransformProfileFromV2ToV1(data)
	require.NoError(t, err)
	err = ConvertStartingAddressToOneBased(&actual)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestTransformResourceFromV1ToV2(t *testing.T) {
	data := v1ProfileData()
	expected := v2ProfileData()

	for i, dr := range data.DeviceResources {
		actual, err := TransformResourceFromV1ToV2(dr)
		require.NoError(t, err)
		assert.Equal(t, actual.Tags, expected.DeviceResources[i].Tags)
		assert.Equal(t, actual.Properties, expected.DeviceResources[i].Properties)
	}
}
