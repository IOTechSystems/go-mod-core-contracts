//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xlsx

import (
	"fmt"
	"io"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type mappingField struct {
	defaultValue string // the default value defined in the MappingTable sheet
	path         string // the path value defined in the MappingTable sheet
}

func ConvertDeviceXlsx(file io.Reader) (Converter[[]*dtos.Device], errors.EdgeX) {
	deviceX, err := newDeviceXlsx(file)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "failed to create deviceXlsx instance", err)
	}

	err = deviceX.ConvertToDTO()
	if err != nil {
		return nil, errors.NewCommonEdgeXWrapper(err)
	}

	return deviceX, nil
}

func ConvertDeviceProfileXlsx(file io.Reader) (Converter[*dtos.DeviceProfile], error) {
	deviceProfileX, err := newDeviceProfileXlsx(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create deviceProfileXlsx instance: %w", err)
	}

	err = deviceProfileX.ConvertToDTO()
	if err != nil {
		return nil, err
	}

	return deviceProfileX, nil
}
