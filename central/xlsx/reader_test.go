//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/stretchr/testify/require"
)

func Test_readStruct(t *testing.T) {
	testStr := "testString"
	testInvalidDevice := dtos.Device{}
	testValidDevice := dtos.Device{}

	tests := []struct {
		name        string
		structPtr   *dtos.Device
		headerRow   []string
		dataRow     []string
		expectError bool
	}{
		{"readStruct with invalid ptr", nil, nil, nil, true},
		{"readStruct with invalid value type", &testInvalidDevice, []string{"LastConnected"}, []string{"test"}, true},
		{"readStruct with valid value type", &testValidDevice, []string{"Location"}, []string{"test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.structPtr != nil {
				_, err = readStruct(tt.structPtr, "", tt.headerRow, tt.dataRow)
			} else {
				_, err = readStruct(&testStr, "", tt.headerRow, tt.dataRow)
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

func Test_setProtocolPropMap(t *testing.T) {
	tests := []struct {
		name        string
		protocol    string
		prtProps    map[string]string
		expectError bool
	}{
		{"setProtocolPropMap with valid protocol", "modbus-rtu", map[string]string{"DataBits": "7"}, false},
		{"setProtocolPropMap with invalid protocol", "test", map[string]string{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := setProtocolPropMap(tt.protocol, tt.prtProps)
			if tt.expectError {
				require.Error(t, err, "Expected setProtocolPropMap error not occurred")
			} else {
				require.NoError(t, err, "Unexpected setProtocolPropMap error occurred")
				if prtProps, ok := result[tt.protocol]; ok {
					require.Equal(t, dtos.ProtocolProperties(tt.prtProps), prtProps)
				} else {
					require.Fail(t, "Unexpected setProtocolPropMap parse result")
				}
			}
		})
	}
}
