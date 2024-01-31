// Copyright (C) 2023 IOTech Ltd

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseValueByDeviceResource(t *testing.T) {
	tests := []struct {
		name      string
		valueType string
		value     any
		expected  any
	}{
		{"string", ValueTypeString, "foo and bar", "foo and bar"},
		{"uint8", ValueTypeUint8, "127", uint8(127)},
		{"uint16", ValueTypeUint16, "127", uint16(127)},
		{"uint32", ValueTypeUint32, "127", uint32(127)},
		{"uint64", ValueTypeUint64, "127", uint64(127)},
		{"int8", ValueTypeInt8, "-127", int8(-127)},
		{"int16", ValueTypeInt16, "-127", int16(-127)},
		{"int32", ValueTypeInt32, "-127", int32(-127)},
		{"int64", ValueTypeInt64, "-127", int64(-127)},
		{"float32", ValueTypeFloat32, "0.123", float32(0.123)},
		{"float64", ValueTypeFloat64, "0.123", 0.123},
		{"string array - EdgeX readings", ValueTypeStringArray, "[foo, bar]", []string{"foo", "bar"}},
		{"string array - Set command payload ", ValueTypeStringArray, "[\"foo\",\"bar\"]", []string{"foo", "bar"}},
		{"bool array", ValueTypeBoolArray, "[true, false]", []bool{true, false}},
		{"uint8 array", ValueTypeUint8Array, "[100, 127]", []uint8{100, 127}},
		{"uint16 array", ValueTypeUint16Array, "[100, 127]", []uint16{100, 127}},
		{"uint32 array", ValueTypeUint32Array, "[100, 127]", []uint32{100, 127}},
		{"uint64 array", ValueTypeUint64Array, "[100, 127]", []uint64{100, 127}},
		{"int8 array", ValueTypeInt8Array, "[-127, 127]", []int8{-127, 127}},
		{"int16 array", ValueTypeInt16Array, "[-127, 127]", []int16{-127, 127}},
		{"int32 array", ValueTypeInt32Array, "[-127, 127]", []int32{-127, 127}},
		{"int64 array", ValueTypeInt64Array, "[-127, 127]", []int64{-127, 127}},
		{"float32 array", ValueTypeFloat32Array, "[-0.123, 0.123]", []float32{-0.123, 0.123}},
		{"float64 array", ValueTypeFloat64Array, "[-0.123, 0.123]", []float64{-0.123, 0.123}},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := ParseValueByDeviceResource(testCase.valueType, testCase.value)
			require.NoError(t, err)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
