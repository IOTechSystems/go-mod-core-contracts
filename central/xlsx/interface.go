//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xlsx

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type AllowedDTOTypes interface {
	*dtos.DeviceProfile | []*dtos.Device
}

type Converter[T AllowedDTOTypes] interface {
	// ConvertToDTO parses the xlsx file content to DTOs
	ConvertToDTO() errors.EdgeX
	// GetDTOs returns the coverted DTOs
	GetDTOs() T
	// GetValidateErrors returns the deviceName-validationError key-value map while parsing the excel data rows to DTOs
	GetValidateErrors() map[string]error
}
