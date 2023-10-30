//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// convertMappingTable parses the MappingTable sheet and stores the default value for each device field
func convertMappingTable(xlsFile *excelize.File) (map[string]mappingField, error) {
	defaultValueMap := make(map[string]mappingField)

	rows, err := xlsFile.GetRows(mappingTableSheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all rows from %s: %w", mappingTableSheetName, err)
	}

	var objColIndex, pathColIndex, defaultColIndex, headerLength int
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			// read the header row and get the Object and DefaultValue column index
			for colIndex, colCell := range row {
				switch colCell {
				case objectCol:
					objColIndex = colIndex
				case pathCol:
					pathColIndex = colIndex
				case defaultValueCol:
					defaultColIndex = colIndex
				}
			}
			headerLength = len(row)
		} else {
			// Since GetRows method skips the continually blank cells in the tail of each row
			// Append empty string to the row to avoid invalid access to the objectColIndex or defaultValueColIndex element of each row slice
			// See Excelize doc: https://xuri.me/excelize/en/cell.html#GetRows
			if headerLength > len(row) {
				skippedCount := headerLength - len(row)
				for skippedCount > 0 {
					row = append(row, "")
					skippedCount--
				}
			}

			defaultValueMap[row[objColIndex]] = mappingField{path: row[pathColIndex], defaultValue: row[defaultColIndex]}
		}
	}

	return defaultValueMap, nil
}
