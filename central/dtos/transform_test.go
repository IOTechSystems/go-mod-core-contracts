//
// Copyright (C) 2024 IOTech Ltd
//

package dtos

import (
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/v2dtos"

	"github.com/stretchr/testify/require"
)

var (
	mockDeviceResourceName   = "mockRes1"
	v2MockResourceOperations = []v2dtos.ResourceOperation{{DeviceResource: mockDeviceResourceName}}
	v3MockResourceOperation  = []dtos.ResourceOperation{dtos.ResourceOperation(v2MockResourceOperations[0])}
	mockReadWrite            = "R"
	mockValueType            = "Float64"
	mockMin                  = 100.00
	mockMax                  = 10000.00
	mockV2ResProp            = v2dtos.ResourceProperties{
		ValueType: mockValueType,
		ReadWrite: mockReadWrite,
		Minimum:   fmt.Sprintf("%f", mockMin),
		Maximum:   fmt.Sprintf("%f", mockMax),
	}
	mockResProp = dtos.ResourceProperties{
		ValueType: mockValueType,
		ReadWrite: mockReadWrite,
		Minimum:   &mockMin,
		Maximum:   &mockMax,
	}
	mockAttributes  = map[string]any{"foo": "bar"}
	mockV2DeviceRes = v2dtos.DeviceResource{
		Name:       mockDeviceResourceName,
		IsHidden:   false,
		Properties: mockV2ResProp,
		Attributes: mockAttributes,
	}
	mockDeviceRes = dtos.DeviceResource{
		Name:       mockDeviceResourceName,
		IsHidden:   false,
		Properties: mockResProp,
		Attributes: mockAttributes,
	}
	mockDeviceCommandName = "mockDC1"
	mockV2DeviceCommand   = v2dtos.DeviceCommand{
		Name:               mockDeviceCommandName,
		IsHidden:           false,
		ReadWrite:          mockReadWrite,
		ResourceOperations: v2MockResourceOperations,
	}
	mockDeviceCommand = dtos.DeviceCommand{
		Name:               mockDeviceCommandName,
		IsHidden:           false,
		ReadWrite:          mockReadWrite,
		ResourceOperations: v3MockResourceOperation,
	}
)

func Test_TransformProfileFromV2ToV3(t *testing.T) {
	mockDeviceProfileName := "mockPro1"
	mockManufacturerName := "DNZ"
	mockV2Profile := v2dtos.DeviceProfile{
		ApiVersion: "v2",
		DeviceProfileBasicInfo: v2dtos.DeviceProfileBasicInfo{
			Name:         mockDeviceProfileName,
			Manufacturer: mockManufacturerName,
		},
		DeviceResources: []v2dtos.DeviceResource{mockV2DeviceRes},
		DeviceCommands:  []v2dtos.DeviceCommand{mockV2DeviceCommand},
	}
	result, err := TransformProfileFromV2ToV3(mockV2Profile)
	require.NoError(t, err)
	require.Equal(t, mockDeviceProfileName, result.Name)
	require.Equal(t, mockManufacturerName, result.Manufacturer)
	require.Equal(t, mockV2DeviceRes.Name, result.DeviceResources[0].Name)
	require.Equal(t, mockV2DeviceCommand.Name, result.DeviceCommands[0].Name)
}

func Test_transformResourceFromV2ToV3(t *testing.T) {
	result, err := transformResourceFromV2ToV3([]v2dtos.DeviceResource{mockV2DeviceRes})
	require.NoError(t, err)
	require.Equal(t, mockDeviceRes.Name, result[0].Name)
	require.Equal(t, mockDeviceRes.Attributes, result[0].Attributes)
}

func Test_transformResPropsFromV2ToV3(t *testing.T) {
	result, err := transformResPropsFromV2ToV3(mockV2ResProp)
	require.NoError(t, err)
	require.Equal(t, mockResProp.ValueType, result.ValueType)
	require.Equal(t, mockResProp.ReadWrite, result.ReadWrite)
	require.Equal(t, *mockResProp.Minimum, *result.Minimum)
	require.Equal(t, *mockResProp.Maximum, *result.Maximum)
}

func Test_transformCommandFromV2ToV3(t *testing.T) {
	expected := []dtos.DeviceCommand{mockDeviceCommand}
	results := transformCommandFromV2ToV3([]v2dtos.DeviceCommand{mockV2DeviceCommand})
	require.Equal(t, expected, results)
}

func Test_transformResourceOperationFromV2ToV3(t *testing.T) {
	result := transformResourceOperationFromV2ToV3(v2MockResourceOperations)
	require.Equal(t, v3MockResourceOperation, result)
}
