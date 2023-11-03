//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

// Converter interface provides an abstraction for parsing the xlsx file content
type Converter interface {
	// ConvertToDTO parses the xlsx file content to DTOs
	convertToDTO() error
	// GetDTOs returns the coverted DTOs
	GetDTOs() any
	// GetValidateErrors returns the deviceName-validationError key-value map while parsing the excel data rows to DTOs
	GetValidateErrors() map[string]error
}
