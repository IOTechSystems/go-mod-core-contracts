//
// Copyright (C) 2024 IOTech Ltd
//

package xlsx

import (
	"fmt"
	"io"
	"strconv"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

func initialXlsxFileReader() (*excelize.File, io.Reader, error) {
	f, err := createXlsxTemplateFile()
	if err != nil {
		return nil, nil, err
	}
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, nil, err
	}
	return f, buffer, nil
}

func Test_newDPXlsxWriter(t *testing.T) {
	f, buffer, err := initialXlsxFileReader()
	require.NoError(t, err)
	defer f.Close()

	xlsxWriter, err := newDPXlsxWriter(mockDeviceProfile, buffer)

	require.NoError(t, err)
	require.Equal(t, mockProfileName, xlsxWriter.deviceProfile.Name)
}

func Test_ConvertToXlsx(t *testing.T) {
	f, buffer, err := initialXlsxFileReader()
	require.NoError(t, err)
	defer f.Close()

	xlsxWriter, err := newDPXlsxWriter(mockDeviceProfile, buffer)
	require.NoError(t, err)

	err = xlsxWriter.ConvertToXlsx()
	require.NoError(t, err)
}

func Test_dpWriter_convertDeviceInfo(t *testing.T) {
	f, buffer, err := initialXlsxFileReader()
	require.NoError(t, err)
	defer f.Close()

	xlsxWriter, err := newDPXlsxWriter(mockDeviceProfile, buffer)
	require.NoError(t, err)

	err = xlsxWriter.convertDeviceInfo()
	require.NoError(t, err)

	value, err := xlsxWriter.xlsxFile.GetCellValue(deviceInfoSheetName, "B1")
	require.NoError(t, err)
	require.Equal(t, mockProfileName, value)
}

func Test_dpWriter_convertDeviceResources(t *testing.T) {
	f, buffer, err := initialXlsxFileReader()
	require.NoError(t, err)
	defer f.Close()

	mockAttrValue := "HOLDING_REGISTERS"
	minValue := float64(0)
	mockNestedIdAttrValue := 8
	mockResource := dtos.DeviceResource{
		Description: "this is the mockRes1 resource",
		Name:        "mockRes1",
		IsHidden:    false,
		Properties: dtos.ResourceProperties{
			ValueType: "Float32",
			ReadWrite: common.ReadWrite_R,
			Minimum:   &minValue,
		},
		Attributes: map[string]any{"primaryTable": mockAttrValue, "dataTypeId": map[string]any{"identifier": mockNestedIdAttrValue}},
	}
	mockDeviceProfile.DeviceResources = []dtos.DeviceResource{mockResource}
	xlsxWriter, err := newDPXlsxWriter(mockDeviceProfile, buffer)
	require.NoError(t, err)

	err = xlsxWriter.convertDeviceResources()
	require.NoError(t, err)

	value, err := xlsxWriter.xlsxFile.GetCellValue(deviceResourceSheetName, "A2")
	require.NoError(t, err)
	require.Equal(t, mockResource.Name, value)

	value, err = xlsxWriter.xlsxFile.GetCellValue(deviceResourceSheetName, "B2")
	require.NoError(t, err)
	require.Equal(t, strconv.FormatBool(mockResource.IsHidden), value)

	value, err = xlsxWriter.xlsxFile.GetCellValue(deviceResourceSheetName, "C2")
	require.NoError(t, err)
	require.Equal(t, mockResource.Description, value)

	value, err = xlsxWriter.xlsxFile.GetCellValue(deviceResourceSheetName, "D2")
	require.NoError(t, err)
	require.Equal(t, mockResource.Properties.ValueType, value)

	value, err = xlsxWriter.xlsxFile.GetCellValue(deviceResourceSheetName, "E2")
	require.NoError(t, err)
	require.Equal(t, mockResource.Properties.ReadWrite, value)

	value, err = xlsxWriter.xlsxFile.GetCellValue(deviceResourceSheetName, "F2")
	require.NoError(t, err)
	require.Equal(t, mockAttrValue, value)

	value, err = xlsxWriter.xlsxFile.GetCellValue(deviceResourceSheetName, "G2")
	require.NoError(t, err)
	require.Equal(t, strconv.FormatFloat(*mockResource.Properties.Minimum, 'g', -1, 64), value)

	value, err = xlsxWriter.xlsxFile.GetCellValue(deviceResourceSheetName, "H2")
	require.NoError(t, err)
	require.Equal(t, strconv.FormatInt(int64(mockNestedIdAttrValue), 10), value)
}

func Test_dpWriter_convertDeviceCommand(t *testing.T) {
	f, buffer, err := initialXlsxFileReader()
	require.NoError(t, err)
	defer f.Close()

	mockDeviceCommand := dtos.DeviceCommand{
		Name:     "mockCmd1",
		IsHidden: false,
	}
	mockDeviceProfile.DeviceCommands = []dtos.DeviceCommand{mockDeviceCommand}
	xlsxWriter, err := newDPXlsxWriter(mockDeviceProfile, buffer)
	require.NoError(t, err)

	err = xlsxWriter.convertDeviceCommand()
	require.NoError(t, err)

	value, err := xlsxWriter.xlsxFile.GetCellValue(deviceCommandSheetName, "B1")
	require.NoError(t, err)
	require.Equal(t, mockDeviceCommand.Name, value)
}

func Test_dpWriter_setResourceNameCells(t *testing.T) {
	f, buffer, err := initialXlsxFileReader()
	require.NoError(t, err)
	defer f.Close()

	mockResOp1 := dtos.ResourceOperation{
		DeviceResource: "res1",
	}
	mockResOp2 := dtos.ResourceOperation{
		DeviceResource: "res2",
	}
	mockResOp3 := dtos.ResourceOperation{
		DeviceResource: "res3",
	}
	mockResOPs := []dtos.ResourceOperation{mockResOp1, mockResOp2, mockResOp3}
	xlsxWriter, err := newDPXlsxWriter(mockDeviceProfile, buffer)
	require.NoError(t, err)

	startRow := 0
	colNumber := 0
	err = xlsxWriter.setResourceNameCells(startRow, colNumber, mockResOPs)
	require.NoError(t, err)

	value, err := xlsxWriter.xlsxFile.GetCellValue(deviceCommandSheetName, fmt.Sprintf("B%d", startRow+1))
	require.NoError(t, err)
	require.Equal(t, mockResOp1.DeviceResource, value)

	value, err = xlsxWriter.xlsxFile.GetCellValue(deviceCommandSheetName, fmt.Sprintf("B%d", startRow+2))
	require.NoError(t, err)
	require.Equal(t, mockResOp2.DeviceResource, value)

	value, err = xlsxWriter.xlsxFile.GetCellValue(deviceCommandSheetName, fmt.Sprintf("B%d", startRow+3))
	require.NoError(t, err)
	require.Equal(t, mockResOp3.DeviceResource, value)
}
