//
// Copyright (C) 2023 IOTech Ltd
//

package utils

import (
	"io"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"

	"github.com/xuri/excelize/v2"
)

type mappingField struct {
	defaultValue string // the default value defined in the MappingTable sheet
	path         string // the path value defined in the MappingTable sheet
}

// FromSpreadsheetToDTO transforms the Device spreadsheets to the UpdateDevice DTO slice
func FromSpreadsheetToDTO(file io.Reader) ([]*dtos.Device, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	deviceXlsx := new(DeviceXlsx)
	fieldMappings, err := ConvertMappingTable(f)
	if err != nil {
		return nil, err
	}
	deviceXlsx.fieldMappings = fieldMappings

	err = deviceXlsx.ConvertDevice(f, deviceXlsx.fieldMappings[protocolName].defaultValue)
	if err != nil {
		return nil, err
	}

	return deviceXlsx.devices, nil
}
