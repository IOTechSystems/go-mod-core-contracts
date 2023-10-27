//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"fmt"
	"io"
	"reflect"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"

	"github.com/xuri/excelize/v2"
)

type mappingField struct {
	defaultValue string // the default value defined in the MappingTable sheet
	path         string // the path value defined in the MappingTable sheet
}

// ConvertXlsx transforms the xlsx file to the Converter interface
func ConvertXlsx(file io.Reader, dtoType reflect.Type) (Converter, error) {
	// check if the dto is a valid DTO type before read the xlsx file
	switch dtoType {
	case reflect.TypeOf(dtos.Device{}), reflect.TypeOf(dtos.DeviceProfile{}):
	default:
		return nil, fmt.Errorf("unable to parse the xlsx file to invalid DTO type '%T'", dtoType)
	}

	var converter Converter

	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fieldMappings, err := convertMappingTable(f)
	if err != nil {
		return nil, err
	}

	switch dtoType {
	case reflect.TypeOf(dtos.Device{}):
		deviceX := newDeviceXlsx(f, fieldMappings)
		converter = deviceX
	case reflect.TypeOf(dtos.DeviceProfile{}):
	}

	err = converter.convertToDTO()
	if err != nil {
		return nil, err
	}

	return converter, nil
}
