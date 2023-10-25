//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"io"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"

	"github.com/xuri/excelize/v2"
)

type mappingField struct {
	defaultValue string // the default value defined in the MappingTable sheet
	path         string // the path value defined in the MappingTable sheet
}

// ConvertXlsx transforms the xlsx file to the Converter interface
func ConvertXlsx(file io.Reader, dto any) (Converter, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fieldMappings, err := convertMappingTable(f)
	if err != nil {
		return nil, err
	}

	converter, err := processXlsxByType(f, fieldMappings, dto)
	if err != nil {
		return nil, err
	}

	return converter, nil
}

// processXlsxByType processes the xlsx file based on the dto type and the mapping table definition
func processXlsxByType(f *excelize.File, fieldMappings map[string]mappingField, dto any) (Converter, error) {
	var err error
	var converter Converter

	switch dto.(type) {
	case dtos.Device:
		deviceX := new(deviceXlsx)
		deviceX.fieldMappings = fieldMappings
		err = deviceX.convertToDTO(f, deviceX.fieldMappings[protocolName].defaultValue)
		if err != nil {
			return nil, err
		}
		converter = deviceX
	case dtos.DeviceProfile:
	}

	return converter, nil
}
