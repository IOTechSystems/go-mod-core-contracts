//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import "github.com/xuri/excelize/v2"

// Converter interface provides an abstraction for parsing the xlsx file content
type Converter interface {
	// ConvertToDTO parses the xlsx file content to DTOs
	convertToDTO(*excelize.File, string) error
}
