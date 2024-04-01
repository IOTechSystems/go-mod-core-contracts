//
// Copyright (C) 2023-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xlsx

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/stretchr/testify/require"
)

func Test_readStruct(t *testing.T) {
	testStr := "testString"
	testValidDevice := dtos.Device{}
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)

	validMappings := deviceX.(*deviceXlsx).fieldMappings
	tests := []struct {
		name        string
		structPtr   *dtos.Device
		headerRow   []string
		dataRow     []string
		expectError bool
	}{
		{"readStruct with invalid ptr", nil, nil, nil, true},
		{"readStruct with valid value type", &testValidDevice, []string{"Location"}, []string{"test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.structPtr != nil {
				_, err = readStruct(tt.structPtr, tt.headerRow, tt.dataRow, validMappings)
			} else {
				_, err = readStruct(&testStr, tt.headerRow, tt.dataRow, validMappings)
			}
			if tt.expectError {
				require.Error(t, err, "Expected readStruct parse error not occurred")
			} else {
				require.NoError(t, err, "Unexpected readStruct parse error occurred")
				require.Equal(t, "test", testValidDevice.Location)
			}
		})
	}
}

func Test_getStructFieldByHeader(t *testing.T) {
	rowElement := reflect.New(reflect.TypeOf(dtos.DeviceProfile{})).Elem()
	colNameWithoutSpace := "Manufacturer"
	colNameWithSpace := " " + colNameWithoutSpace
	headerCol := []string{"Name", colNameWithSpace}
	headerName, field := getStructFieldByHeader(&rowElement, 0, headerCol)
	require.Equal(t, "Name", headerName)
	require.Equal(t, reflect.String, field.Kind())

	headerName2, field2 := getStructFieldByHeader(&rowElement, 1, headerCol)
	require.Equal(t, colNameWithoutSpace, headerName2)
	require.Equal(t, reflect.String, field2.Kind())
}

func Test_setStdStructFieldValue(t *testing.T) {
	rowElement := reflect.New(reflect.TypeOf(dtos.Device{})).Elem()
	labels := rowElement.FieldByName("Labels")
	tests := []struct {
		name        string
		cellValue   string
		field       reflect.Value
		expectError bool
	}{
		{"setStdStructFieldValue - fail to parse cell to bool field", "invalid", reflect.ValueOf(true), true},
		{"setStdStructFieldValue - success to parse cell to slice field", "a,b,c", labels, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := setStdStructFieldValue(tt.cellValue, tt.field)
			if tt.expectError {
				require.Error(t, err, "Expected cell conversion error not occurred")
			} else {
				require.NoError(t, err, "Unexpected error occurred")

			}
		})
	}
}

func Test_setProtocolPropMap_WithoutMappingTableSheet(t *testing.T) {
	_, err := setProtocolPropMap(map[string]any{"DataBits": "7"}, nil)
	require.Error(t, err, "Expected fieldMapping not defined error not occurred")
}

func Test_setProtocolPropMap_WithMappingTableSheet(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)

	validMappings := deviceX.(*deviceXlsx).fieldMappings
	invalidMappings := make(map[string]mappingField)
	invalidMappings["ProtocolName"] = mappingField{defaultValue: "invalidPrt"}
	tests := []struct {
		name          string
		prtProps      map[string]any
		fieldMappings map[string]mappingField
		expectError   bool
	}{
		{"setProtocolPropMap with valid protocol", map[string]any{"DataBits": "7"}, validMappings, false},
		{"setProtocolPropMap with invalid protocol", map[string]any{}, invalidMappings, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := setProtocolPropMap(tt.prtProps, tt.fieldMappings)
			if tt.expectError {
				require.Error(t, err, "Expected setProtocolPropMap error not occurred")
			} else {
				require.NoError(t, err, "Unexpected setProtocolPropMap error occurred")
				if protocol, ok := tt.fieldMappings["ProtocolName"]; ok {
					if prtProps, ok := result[protocol.defaultValue]; ok {
						require.Equal(t, dtos.ProtocolProperties(tt.prtProps), prtProps)
					} else {
						require.Fail(t, "Unexpected setProtocolPropMap parse result")
					}
				} else {
					require.Fail(t, "ProtocolName not found in tt.fieldMapping")
				}
			}
		})
	}
}

func Test_convertDeviceFields(t *testing.T) {
	rowElement := reflect.New(reflect.TypeOf(dtos.Device{})).Elem()
	headerCol := []string{"Name", "LastConnected"}
	validDataRow := []string{"TestDevice", "0"}
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)

	validMappings := deviceX.(*deviceXlsx).fieldMappings

	tests := []struct {
		name          string
		rowElement    *reflect.Value
		dataRow       []string
		headerCol     []string
		fieldMappings map[string]mappingField
		expectError   bool
	}{
		{"Invalid convertDeviceFields - no fieldMappings", &rowElement, validDataRow, headerCol, nil, true},
		{"Valid convertDeviceFields", &rowElement, validDataRow, headerCol, validMappings, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := convertDeviceFields(tt.rowElement, tt.dataRow, tt.headerCol, tt.fieldMappings)
			if tt.expectError {
				require.Error(t, err, "Expected convertDeviceFields error not occurred")
			} else {
				require.NoError(t, err, "Unexpected convertDeviceFields error occurred")
			}
		})
	}
}

func Test_convertAutoEventFields(t *testing.T) {
	rowElement := reflect.New(reflect.TypeOf(dtos.AutoEvent{})).Elem()
	headerCol := []string{"Interval", "OnChange", "Reference Device Name"}
	invalidDataRow := []string{"3s", "invalid"}
	expectedDevice := "testDevice1"
	validDataRow := []string{"3s", "true", expectedDevice}
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)

	validMappings := deviceX.(*deviceXlsx).fieldMappings

	tests := []struct {
		name          string
		rowElement    *reflect.Value
		dataRow       []string
		headerCol     []string
		fieldMappings map[string]mappingField
		expectError   bool
	}{
		{"Invalid convertAutoEventFields - no fieldMappings", &rowElement, validDataRow, headerCol, nil, true},
		{"Invalid convertAutoEventFields - invalid OnChange cell", &rowElement, invalidDataRow, headerCol, validMappings, true},
		{"Valid convertAutoEventFields", &rowElement, validDataRow, headerCol, validMappings, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deviceNames, err := convertAutoEventFields(tt.rowElement, tt.dataRow, tt.headerCol, tt.fieldMappings)
			if tt.expectError {
				require.Error(t, err, "Expected convertAutoEventFields error not occurred")
			} else {
				require.NoError(t, err, "Unexpected convertAutoEventFields error occurred")
				require.Equal(t, 1, len(deviceNames))
				require.Equal(t, expectedDevice, deviceNames[0])
			}
		})
	}
}

func Test_convertDeviceCommandFields(t *testing.T) {
	rowElement := reflect.New(reflect.TypeOf(dtos.DeviceCommand{})).Elem()
	headerCol := []string{"Name", "IsHidden"}
	invalidDataRow := []string{"testCommand", "invalid"}
	validDataRow := []string{"testCommand", "true"}

	tests := []struct {
		name        string
		rowElement  *reflect.Value
		dataRow     []string
		headerCol   []string
		expectError bool
	}{
		{"Invalid convertDeviceCommandFields - invalid IsHidden cell", &rowElement, invalidDataRow, headerCol, true},
		{"Valid convertDeviceCommandFields", &rowElement, validDataRow, headerCol, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := convertDeviceCommandFields(tt.rowElement, tt.dataRow, tt.headerCol)
			if tt.expectError {
				require.Error(t, err, "Expected convertAutoEventFields error not occurred")
			} else {
				require.NoError(t, err, "Unexpected convertAutoEventFields error occurred")
			}
		})
	}
}

func Test_convertResourcesFields(t *testing.T) {
	headerCol := []string{"Name", "IsHidden", "ValueType", "nodeAttribute"}
	invalidIsHiddenRow := []string{"testCommand", "invalid", "Int64"}
	validDataRow := []string{"testCommand", "true", "Int64", "value"}
	deviceX, err := createDeviceProfileXlsxInst()
	require.NoError(t, err)

	validMappings := deviceX.(*deviceProfileXlsx).fieldMappings
	tests := []struct {
		name          string
		dataRow       []string
		headerCol     []string
		fieldMappings map[string]mappingField
		expectError   bool
	}{
		{"Invalid convertResourcesFields - no fieldMappings", validDataRow, headerCol, nil, true},
		{"Invalid convertResourcesFields - invalid IsHidden cell", invalidIsHiddenRow, headerCol, validMappings, true},
		{"Valid convertResourcesFields", validDataRow, headerCol, validMappings, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structPtr := dtos.DeviceResource{}
			v := reflect.ValueOf(&structPtr)
			elementType := v.Elem().Type()
			element := reflect.New(elementType).Elem()

			err := convertResourcesFields(&element, tt.dataRow, tt.headerCol, tt.fieldMappings)
			v.Elem().Set(element)
			if tt.expectError {
				require.Error(t, err, "Expected convertResourcesFields error not occurred")
			} else {
				require.NoError(t, err, "Unexpected convertResourcesFields error occurred")
				require.Equal(t, validDataRow[0], structPtr.Name)
				require.Equal(t, validDataRow[1], strconv.FormatBool(structPtr.IsHidden))
				require.Equal(t, validDataRow[2], structPtr.Properties.ValueType)
				require.Equal(t, validDataRow[3], structPtr.Attributes[headerCol[3]])
			}
		})
	}
}

func Test_convertResourcesFields_Nested_Attributes(t *testing.T) {
	nestedAttrName1 := "dataTypeId.identifier"
	nestedAttrName2 := "dataTypeId.identifierType"

	headerCol := []string{"Name", nestedAttrName1, nestedAttrName2}
	dataRow := []string{"testCommand", "8", "NUMERIC"}
	deviceProfileX, err := createDeviceProfileXlsxInst()
	require.NoError(t, err)

	fieldMappings := deviceProfileX.(*deviceProfileXlsx).fieldMappings

	structPtr := dtos.DeviceResource{}
	v := reflect.ValueOf(&structPtr)
	elementType := v.Elem().Type()
	element := reflect.New(elementType).Elem()

	err = convertResourcesFields(&element, dataRow, headerCol, fieldMappings)
	require.NoError(t, err)
	v.Elem().Set(element)

	require.Equal(t, dataRow[0], structPtr.Name)

	// check the converted nested attributes int64 value
	splitAttrNames := strings.Split(nestedAttrName1, mappingPathSeparator)
	if innerAttr, ok := structPtr.Attributes[splitAttrNames[0]].(map[string]any); ok {
		if attrVal, innerOk := innerAttr[splitAttrNames[1]].(int64); innerOk {
			require.Equal(t, dataRow[1], strconv.FormatInt(attrVal, 10))
		}
	}

	// check the converted nested attributes string value
	splitAttrNames = strings.Split(nestedAttrName2, mappingPathSeparator)
	if innerAttr, ok := structPtr.Attributes[splitAttrNames[0]].(map[string]any); ok {
		if attrVal, innerOk := innerAttr[splitAttrNames[1]]; innerOk {
			require.Equal(t, dataRow[2], attrVal)
		}
	}
}
