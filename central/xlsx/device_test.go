//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xlsx

import (
	"errors"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

var (
	deviceHeaderStr   = []string{"Name", "Description", "ServiceName", "ProtocolName", "Labels", "AdminState", "Address", "BaudRate", "DataBits", "Parity", "StopBits", "UnitID", "ProfileName"}
	validDeviceHeader = []any{
		"Name", "Description", "ServiceName", "ProtocolName", "Labels", "AdminState", "Address", "BaudRate", "DataBits", "Parity", "StopBits", "UnitID", "ProfileName",
	}
	emptyValidateErr = map[string]error{}
)

func Test_newDeviceXlsx(t *testing.T) {
	f := excelize.NewFile()
	defer f.Close()

	_, err := f.NewSheet(mappingTableSheetName)
	require.NoError(t, err)
	err = createMappingTableSheet(f)
	require.NoError(t, err)
	buffer, err := f.WriteToBuffer()
	require.NoError(t, err)

	deviceXls, err := newDeviceXlsx(buffer)
	require.NoError(t, err)
	require.NotEmpty(t, deviceXls)
}

func mockExcelFile(sheetNames []string) (*excelize.File, error) {
	f := excelize.NewFile()

	for _, sheetName := range sheetNames {
		_, err := f.NewSheet(sheetName)
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}

func createMappingTableSheet(f *excelize.File) error {
	sw, err := f.NewStreamWriter(mappingTableSheetName)
	if err != nil {
		return err
	}

	err = sw.SetRow("A1",
		[]any{
			"Object", "Path", "Default Value",
		})
	if err != nil {
		return err
	}

	err = sw.SetRow("A2",
		[]any{
			"AdminState", "adminState", "UNLOCKED",
		})
	if err != nil {
		return err
	}

	err = sw.SetRow("A3",
		[]any{
			"OperatingState", "operatingState", "UP",
		})
	if err != nil {
		return err
	}

	err = sw.SetRow("A4",
		[]any{
			"ProtocolName", "protocolName", "modbus-rtu",
		})
	if err != nil {
		return err
	}

	err = sw.SetRow("A5",
		[]any{
			"Interval", "autoEvents[].interval", "1s",
		})
	if err != nil {
		return err
	}

	err = sw.Flush()
	if err != nil {
		return err
	}

	return nil
}

func createDeviceXlsxInst() (Converter[[]*dtos.Device], error) {
	f, err := mockExcelFile([]string{devicesSheetName, mappingTableSheetName})
	if err != nil {
		return nil, err
	}

	err = createMappingTableSheet(f)
	if err != nil {
		return nil, err
	}

	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	deviceXls, err := newDeviceXlsx(buffer)
	if err != nil {
		return nil, err
	}
	return deviceXls, err
}

func Test_convertToDTO(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.(*deviceXlsx).xlsFile.Close()

	sw, err := deviceX.(*deviceXlsx).xlsFile.NewStreamWriter(devicesSheetName)
	require.NoError(t, err)
	err = sw.SetRow("A1", validDeviceHeader)
	require.NoError(t, err)
	err = sw.SetRow("A2",
		[]any{
			"Sensor30001", "test-rtu-device 30001", "device-modbus", "modbus-rtu", "modbus-rtu-labels1,modbus-rtu-labels2", "LOCKED", "/dev/virtualport", "19200", "8", "O", "1", "247", "rtu-profile",
		})
	require.NoError(t, err)
	err = sw.Flush()
	require.NoError(t, err)
	require.NotEmpty(t, deviceX)

	err = deviceX.ConvertToDTO()
	require.NoError(t, err)

	devices := deviceX.GetDTOs()
	require.Equal(t, 1, len(devices))
	require.Equal(t, "Sensor30001", devices[0].Name)
}

func Test_parseDevicesHeader(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.(*deviceXlsx).xlsFile.Close()

	err = deviceX.(*deviceXlsx).xlsFile.SetSheetRow(devicesSheetName, "A1", &[]any{"Name"})
	require.NoError(t, err)

	err = deviceX.(*deviceXlsx).parseDevicesHeader(&deviceHeaderStr, 1)
	require.NoError(t, err)
}

func Test_convertAutoEvents_WithoutSheet(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.(*deviceXlsx).xlsFile.Close()

	err = deviceX.(*deviceXlsx).convertAutoEvents()
	require.Error(t, err, "AutoEvents sheet not exists error should be displayed")
}

func Test_convertAutoEvents_WithSheet(t *testing.T) {
	validAutoEventsHeader := []any{"Interval", "OnChange", "SourceName"}

	tests := []struct {
		name                string
		headerRow           []any
		dataRow             []any
		expectError         bool
		expectValidateError bool
	}{
		{"ConvertAutoEvents with row count less than 2", []any{"invalid"}, nil, false, false},
		{"ConvertAutoEvents with invalid data row", validAutoEventsHeader, []any{"xxx"}, false, true},
		{"ConvertAutoEvents with invalid Interval", validAutoEventsHeader, []any{"invalidInterval", "true", "temperature"}, false, true},
		{"ConvertAutoEvents with valid data row", validAutoEventsHeader, []any{"1s", "true", "temperature"}, false, false},
		{"ConvertAutoEvents with invalid OnChange", validAutoEventsHeader, []any{"1s", "notBool", "temperature"}, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deviceX, err := createDeviceXlsxInst()
			require.NoError(t, err)
			xlsFile := deviceX.(*deviceXlsx).xlsFile
			defer xlsFile.Close()

			_, err = xlsFile.NewSheet(autoEventsSheetName)
			require.NoError(t, err)

			headerRow := tt.headerRow
			err = xlsFile.SetSheetRow(autoEventsSheetName, "A1", &headerRow)
			require.NoError(t, err)
			if tt.dataRow != nil {
				dataRow := tt.dataRow
				err = xlsFile.SetSheetRow(autoEventsSheetName, "A2", &dataRow)
				require.NoError(t, err)
			}
			err = deviceX.(*deviceXlsx).convertAutoEvents()

			if tt.expectError {
				require.Error(t, err, "Expected convertAutoEvents error not generated")
			} else {
				require.NoError(t, err)
				if tt.expectValidateError {
					require.NotNil(t, deviceX.GetValidateErrors(), "Expected convertAutoEvents validation error not generated")
				} else {
					require.Equal(t, emptyValidateErr, deviceX.GetValidateErrors(), "Unexpected convertAutoEvents validation error")
				}
			}
		})
	}
}

func Test_parseAutoEventsHeader_Fail_WithoutAutoEventsSheet(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	xlsFile := deviceX.(*deviceXlsx).xlsFile
	defer xlsFile.Close()

	err = deviceX.(*deviceXlsx).parseAutoEventsHeader([]string{"Resource"}, 1)
	require.Error(t, err, "Expected parseAutoEventsHeader error not occurred")
}

func Test_parseAutoEventsHeader_Success_WithAutoEventsSheet(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	xlsFile := deviceX.(*deviceXlsx).xlsFile
	defer xlsFile.Close()

	_, err = xlsFile.NewSheet(autoEventsSheetName)
	require.NoError(t, err)

	err = deviceX.(*deviceXlsx).parseAutoEventsHeader([]string{"Resource"}, 1)
	require.NoError(t, err, "Unexpected parseAutoEventsHeader error occurred")
}

func Test_startsWithAutoEvents(t *testing.T) {
	result := startsWithAutoEvents("autoEvents[].interval")
	require.True(t, result, "Unexpected startsWithAutoEvents result")

	result = startsWithAutoEvents("name")
	require.False(t, result, "Unexpected startsWithAutoEvents result")
}

func Test_GetDTOs(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.(*deviceXlsx).xlsFile.Close()

	deviceDTOs := deviceX.GetDTOs()
	require.Nil(t, deviceDTOs)

	deviceName := "testDevice"
	mockDevice := dtos.Device{Name: deviceName}
	deviceX.(*deviceXlsx).devices = []*dtos.Device{&mockDevice}

	devices := deviceX.GetDTOs()
	require.Equal(t, 1, len(devices))
	require.Equal(t, deviceName, devices[0].Name)
}

func Test_GetValidateErrors(t *testing.T) {
	mockDeviceName := "mockDevice"
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.(*deviceXlsx).xlsFile.Close()

	validateErrs := deviceX.GetValidateErrors()
	require.Equal(t, validateErrs, emptyValidateErr)

	errMsg := "test error"
	mockError := errors.New(errMsg)
	deviceX.(*deviceXlsx).validateErrors[mockDeviceName] = mockError

	validateErrs = deviceX.GetValidateErrors()
	require.NotNil(t, validateErrs)
	if actualErr, ok := validateErrs[mockDeviceName]; ok {
		require.EqualError(t, actualErr, errMsg)
	} else {
		require.Fail(t, "Expected device validation error not found")
	}
}
