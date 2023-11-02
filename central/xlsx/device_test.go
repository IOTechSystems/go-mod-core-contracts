//
// Copyright (C) 2023 IOTech Ltd
//

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
)

func Test_newDeviceXlsx(t *testing.T) {
	f := excelize.NewFile()
	defer f.Close()

	_, err := f.NewSheet(mappingTableSheetName)
	require.NoError(t, err)

	buffer, err := f.WriteToBuffer()
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

func createDeviceXlsxInst() (*deviceXlsx, error) {
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
	defer deviceX.xlsFile.Close()

	sw, err := deviceX.xlsFile.NewStreamWriter(devicesSheetName)
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

	err = deviceX.convertToDTO()
	require.NoError(t, err)

	require.Equal(t, 1, len(deviceX.devices))
	require.Equal(t, "Sensor30001", deviceX.devices[0].Name)
}

func Test_parseDevicesHeader(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.xlsFile.Close()

	err = deviceX.xlsFile.SetSheetRow(devicesSheetName, "A1", &[]any{"Name"})
	require.NoError(t, err)

	err = deviceX.parseDevicesHeader(&deviceHeaderStr, 1)
	require.NoError(t, err)
}

func Test_convertAutoEvents_WithoutSheet(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.xlsFile.Close()

	err = deviceX.convertAutoEvents()
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
			defer deviceX.xlsFile.Close()

			_, err = deviceX.xlsFile.NewSheet(autoEventsSheetName)
			require.NoError(t, err)

			err = deviceX.xlsFile.SetSheetRow(autoEventsSheetName, "A1", &tt.headerRow)
			require.NoError(t, err)
			if tt.dataRow != nil {
				err = deviceX.xlsFile.SetSheetRow(autoEventsSheetName, "A2", &tt.dataRow)
				require.NoError(t, err)
			}
			err = deviceX.convertAutoEvents()

			if tt.expectError {
				require.Error(t, err, "Expected convertAutoEvents error not generated")
			} else {
				require.NoError(t, err)
				if tt.expectValidateError {
					require.NotNil(t, deviceX.ValidateErrors, "Expected convertAutoEvents validation error not generated")
				} else {
					require.Nil(t, deviceX.ValidateErrors, "Unexpected convertAutoEvents validation error")
				}
			}
		})
	}
}

func Test_parseAutoEventsHeader_Fail_WithoutAutoEventsSheet(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.xlsFile.Close()

	err = deviceX.parseAutoEventsHeader([]string{"Resource"}, 1)
	require.Error(t, err, "Expected parseAutoEventsHeader error not occurred")
}

func Test_parseAutoEventsHeader_Success_WithAutoEventsSheet(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.xlsFile.Close()

	_, err = deviceX.xlsFile.NewSheet(autoEventsSheetName)
	require.NoError(t, err)

	err = deviceX.parseAutoEventsHeader([]string{"Resource"}, 1)
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
	defer deviceX.xlsFile.Close()

	deviceDTOs := deviceX.GetDTOs()
	require.Nil(t, deviceDTOs)

	deviceName := "testDevice"
	mockDevice := dtos.Device{Name: deviceName}
	deviceX.devices = []*dtos.Device{&mockDevice}

	deviceDTOs = deviceX.GetDTOs()
	require.NotNil(t, deviceDTOs)
	if devices, ok := deviceDTOs.([]*dtos.Device); ok {
		require.Equal(t, 1, len(devices))
		require.Equal(t, deviceName, devices[0].Name)
	} else {
		require.Fail(t, "Unexpected GetDTOs data type")
	}
}

func Test_GetValidateErrors(t *testing.T) {
	deviceX, err := createDeviceXlsxInst()
	require.NoError(t, err)
	defer deviceX.xlsFile.Close()

	validateErrs := deviceX.GetValidateErrors()
	require.Nil(t, validateErrs)

	errMsg := "test error"
	mockError := errors.New(errMsg)
	deviceX.ValidateErrors = []error{mockError}

	validateErrs = deviceX.GetValidateErrors()
	require.NotNil(t, validateErrs)
	require.EqualError(t, validateErrs[0], errMsg)
}
