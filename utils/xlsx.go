//
// Copyright (C) 2023 IOTech Ltd
//

package utils

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// addMissingColumn adds the column defined in MappingTable with defaultValue to the provided sheetName
func addMissingColumn(xlsFile *excelize.File, sheetName string, totalColCount, totalRowCount int, defaultValue, headerName string) error {
	// if DefaultValue column of the missing field in the MappingTable is not empty
	// Append a new column before colName for this missing field with DefaultValue set on each row
	colName, err := excelize.ColumnNumberToName(totalColCount + 1)
	if err != nil {
		return fmt.Errorf("failed to covert column number to name: %w", err)
	}

	err = xlsFile.InsertCols(sheetName, colName, 1)
	if err != nil {
		return fmt.Errorf("failed to insert empty column to %s: %w", sheetName, err)
	}

	// create the new column values to append to the worksheet for the missing column
	newColValues := make([]any, totalRowCount)
	for i := range newColValues {
		if i == 0 {
			newColValues[i] = headerName
		} else {
			newColValues[i] = defaultValue
		}
	}

	err = xlsFile.SetSheetCol(sheetName, colName+"1", &newColValues)
	if err != nil {
		return fmt.Errorf("failed to set new column to %s in %s: %w", colName, sheetName, err)
	}

	return nil
}
