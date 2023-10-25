//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"fmt"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"

	"github.com/xuri/excelize/v2"
)

// deviceXlsx stores the worksheets processed result and the converted Device DTOs
type deviceXlsx struct {
	fieldMappings  map[string]mappingField // fieldMappings defines all the device fields with default values defined in the xlsx
	Devices        []*dtos.Device          `json:"devices"`
	ValidateErrors []error                 `json:"validateErrors,omitempty"`
}

// convertToDTO parses the Devices sheet and convert the rows to Device DTOs
func (deviceXlsx *deviceXlsx) convertToDTO(xlsFile *excelize.File, protocol string) error {
	var header []string

	rows, err := xlsFile.GetRows(devicesSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s: %w", devicesSheetName, err)
	}

	for rowIndex, row := range rows {
		// parse the header row
		if rowIndex == 0 {
			// get the column count of the header row to see if any Object field from MappingTable sheet is not defined
			colCount := len(row)

			// assign the first row to header
			header = row

			if colCount != len(deviceXlsx.fieldMappings) {
				for objectField, mapping := range deviceXlsx.fieldMappings {
					if startsWithAutoEvents(mapping.path) {
						// if the mapping path starts with autoEvents, skip the check of the Devices sheet header column
						continue
					}

					// check if the mapping object is defined in the Devices sheet if the defaultValue is not empty
					// if not, insert the mapping object as a new column in the Devices sheet with defaultValue set in each data row
					if mapping.defaultValue != "" {
						err = checkMappingObject(xlsFile, devicesSheetName, &colCount, len(rows), mapping.defaultValue, objectField, &header)
						if err != nil {
							return fmt.Errorf("failed to check mapping object: %w", err)
						}
					}
				}
			}
		}
		break
	}

	rows, err = xlsFile.GetRows(devicesSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s after inserting misshing columns: %w", devicesSheetName, err)
	}

	// parse the device data rows
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}

		convertedDevice := dtos.Device{ProtocolName: protocol}
		_, err = readStruct(&convertedDevice, protocol, header, row)
		if err != nil {
			return fmt.Errorf("failed to unmarshal an excel row into Device DTO: %w", err)
		}

		// validate the device DTO
		err := common.Validate(convertedDevice)
		if err != nil {
			deviceErr := fmt.Errorf("device %s validation error: %v", convertedDevice.Name, err)
			deviceXlsx.ValidateErrors = append(deviceXlsx.ValidateErrors, deviceErr)
		} else {
			deviceXlsx.Devices = append(deviceXlsx.Devices, &convertedDevice)
		}
	}

	err = deviceXlsx.convertAutoEvents(xlsFile)
	if err != nil {
		return err
	}
	return nil
}

// convertAutoEvents parses the AutoEvents sheet and convert the rows to AutoEvent DTOs
func (deviceXlsx *deviceXlsx) convertAutoEvents(xlsFile *excelize.File) error {
	var header []string

	rows, err := xlsFile.GetRows(autoEventsSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s worksheet: %w", autoEventsSheetName, err)
	}

	for rowIndex, row := range rows {
		// parse the header row
		if rowIndex == 0 {
			// get the column count of the header row to see if any Object field from MappingTable sheet is not defined
			colCount := len(row)

			// assign the first row to header
			header = row

			// AutoEvents sheet should at least define 4 columns in the header row
			if colCount < 4 {
				for objectField, mapping := range deviceXlsx.fieldMappings {
					if !startsWithAutoEvents(mapping.path) {
						// if the mapping path doesn't start with autoEvents, skip the check of the AutoEvents sheet header column
						continue
					}

					// check if the mapping object is defined in the AutoEvents sheet if the defaultValue is not empty
					// if not, insert the mapping object as a new column in the Devices sheet with defaultValue set in each data row
					err = checkMappingObject(xlsFile, autoEventsSheetName, &colCount, len(rows), mapping.defaultValue, objectField, &header)
					if err != nil {
						return fmt.Errorf("failed to check mapping object: %w", err)
					}
				}
			}
		}
		break
	}

	rows, err = xlsFile.GetRows(autoEventsSheetName)
	if err != nil {
		return err
	}

	// parse the device data rows
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}

		autoEvent := dtos.AutoEvent{}
		deviceNames, err := readStruct(&autoEvent, "", header, row)
		if err != nil {
			return fmt.Errorf("failed to unmarshal an excel row into AutoEvent DTO: %w", err)
		}

		// validate the AutoEvent DTO
		err = common.Validate(autoEvent)
		if err != nil {
			autoEventErr := fmt.Errorf("autoEvent validation error: %v", err)
			deviceXlsx.ValidateErrors = append(deviceXlsx.ValidateErrors, autoEventErr)
		}

		for _, deviceName := range deviceNames {
			for _, device := range deviceXlsx.Devices {
				if deviceName == device.Name {
					device.AutoEvents = append(device.AutoEvents, autoEvent)
				}
			}
		}
	}

	return nil
}

// startsWithAutoEvents checks if the path name defined in MappingTable sheet starts with autoEvents
func startsWithAutoEvents(path string) bool {
	return strings.HasPrefix(strings.ToLower(path), strings.ToLower(autoEvents))
}
