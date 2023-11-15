//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"fmt"
	"io"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
)

type mappingField struct {
	defaultValue string // the default value defined in the MappingTable sheet
	path         string // the path value defined in the MappingTable sheet
}

func ConvertDeviceXlsx(file io.Reader) (Converter[[]*dtos.Device], error) {
	deviceX, err := newDeviceXlsx(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create deviceXlsx instance: %w", err)
	}

	err = deviceX.convertToDTO()
	if err != nil {
		return nil, err
	}

	return deviceX, nil
}

func ConvertDeviceProfileXlsx(file io.Reader) (Converter[*dtos.DeviceProfile], error) {
	deviceProfileX, err := newDeviceProfileXlsx(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create deviceProfileXlsx instance: %w", err)
	}

	err = deviceProfileX.convertToDTO()
	if err != nil {
		return nil, err
	}

	return deviceProfileX, nil
}
