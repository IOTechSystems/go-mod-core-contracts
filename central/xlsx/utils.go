//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
	"golang.org/x/exp/slices"
)

// checkMappingObject checks if the object field from MappingTable is defined in the provided workseet
// if not found, adds the new object column with defaultValue to the provided sheetName
func checkMappingObject(xlsFile *excelize.File, sheetName string, totalColCount *int, totalRowCount int,
	defaultValue, objectField string, header *[]string) error {
	found := false
	for _, headerCell := range *header {
		if headerCell == objectField {
			// check if the mapping object field is defined in the Devices sheet header column
			found = true
			break
		}
	}

	if found || defaultValue == "" {
		return nil
	}

	// if DefaultValue column of the missing field in the MappingTable is not empty
	// Append a new column before colName for this missing field with DefaultValue set on each row
	colName, err := excelize.ColumnNumberToName(*totalColCount + 1)
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
			newColValues[i] = objectField
		} else {
			newColValues[i] = defaultValue
		}
	}

	err = xlsFile.SetSheetCol(sheetName, colName+"1", &newColValues)
	if err != nil {
		return fmt.Errorf("failed to set new column to %s in %s: %w", colName, sheetName, err)
	}

	*header = append(*header, objectField)
	*totalColCount++
	return nil
}

// checkRequiredSheets examines if all the required sheets are defined in the xlsx
func checkRequiredSheets(allSheetNames, requiredSheets []string) error {
	for _, requiredSheet := range requiredSheets {
		if !slices.Contains(allSheetNames, requiredSheet) {
			return fmt.Errorf("%s worksheet not found in the file", requiredSheet)
		}
	}
	return nil
}
