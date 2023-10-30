//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"fmt"
	"io"
	"reflect"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
)

type mappingField struct {
	defaultValue string // the default value defined in the MappingTable sheet
	path         string // the path value defined in the MappingTable sheet
}

// ConvertXlsx transforms the xlsx file to the Converter interface
func ConvertXlsx(file io.Reader, dtoType reflect.Type) (Converter, error) {
	var converter Converter
	var err error

	switch dtoType {
	case reflect.TypeOf(dtos.Device{}):
		deviceX, err := newDeviceXlsx(file)
		converter = deviceX
		if err != nil {
			return nil, fmt.Errorf("failed to create deviceXlsx instance: %w", err)
		}
	case reflect.TypeOf(dtos.DeviceProfile{}):
	default:
		return nil, fmt.Errorf("unable to parse the xlsx file to invalid DTO type '%T'", dtoType)
	}

	err = converter.convertToDTO()
	if err != nil {
		return nil, err
	}

	return converter, nil
}
