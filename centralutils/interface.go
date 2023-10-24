//
// Copyright (C) 2023 IOTech Ltd
//

package centralutils

import "github.com/xuri/excelize/v2"

// Xlsx interface provides an abstraction for parsing the xlsx file content
type Xlsx interface {
	// ConvertToDTO parses the xlsx file content to DTOs
	ConvertToDTO(*excelize.File, string) error
}
