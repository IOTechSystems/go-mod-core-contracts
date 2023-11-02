//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"reflect"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

func initialXlsxFile(sheetNames []string) (*excelize.File, error) {
	f, err := mockExcelFile(sheetNames)
	if err != nil {
		return nil, err
	}

	err = createMappingTableSheet(f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func Test_ConvertXlsx_InvalidDTOType(t *testing.T) {
	f, err := initialXlsxFile([]string{mappingTableSheetName})
	require.NoError(t, err)
	defer f.Close()

	buffer, err := f.WriteToBuffer()

	_, err = ConvertXlsx(buffer, reflect.TypeOf("test"))
	require.Error(t, err, "Expected invalid DTO Type error not occurred")
}

func Test_ConvertXlsx_WithoutDeviceSheet(t *testing.T) {
	f, err := initialXlsxFile([]string{mappingTableSheetName})
	require.NoError(t, err)
	defer f.Close()

	buffer, err := f.WriteToBuffer()

	_, err = ConvertXlsx(buffer, reflect.TypeOf(dtos.Device{}))
	require.Error(t, err, "Expected no Devices sheet error not occurred")
}

func Test_ConvertXlsx_WithDeviceSheet(t *testing.T) {
	f, err := initialXlsxFile([]string{mappingTableSheetName, devicesSheetName})
	require.NoError(t, err)
	defer f.Close()

	sw, err := f.NewStreamWriter(devicesSheetName)
	err = sw.SetRow("A1", validDeviceHeader)
	require.NoError(t, err)
	err = sw.SetRow("A2",
		[]any{
			"Sensor30001", "test-rtu-device 30001", "device-modbus", "modbus-rtu", "modbus-rtu-labels1,modbus-rtu-labels2", "LOCKED", "/dev/virtualport", "19200", "8", "O", "1", "247", "rtu-profile",
		})
	require.NoError(t, err)
	err = sw.Flush()
	require.NoError(t, err)

	buffer, err := f.WriteToBuffer()

	_, err = ConvertXlsx(buffer, reflect.TypeOf(dtos.Device{}))
	require.NoError(t, err, "Unexpected ConvertXlsx error occurred")
}
