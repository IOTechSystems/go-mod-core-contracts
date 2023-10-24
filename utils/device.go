//
// Copyright (C) 2023 IOTech Ltd
//

package utils

import (
	"fmt"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"

	"github.com/xuri/excelize/v2"
)

type DeviceXlsx struct {
	fieldMappings map[string]mappingField // fieldMappings defines all the device fields with default values defined in the xlsx
	devices       []*dtos.Device
	validateErr   []error
}

// ConvertDevice parses the Devices sheet and convert the rows to Device DTOs
func (deviceXlsx *DeviceXlsx) ConvertDevice(xlsFile *excelize.File, protocol string) error {
	var header []string

	rows, err := xlsFile.GetRows(devicesSheetName)
	if err != nil {
		return err
	}

	for rowIndex, row := range rows {
		// parse the header row and store the corresponding device field of each column
		// parse the Devices sheet header row and store the missing Object fields from MappingTable sheet
		if rowIndex == 0 {
			// get the column count of the header row to see if any Object field from MappingTable sheet is not defined
			colCount := len(row)

			for _, colCell := range row {
				header = append(header, colCell)
			}

			if colCount != len(deviceXlsx.fieldMappings) {
				for objectField, mapping := range deviceXlsx.fieldMappings {
					if startsWithAutoEvents(mapping.path) {
						// if the mapping path starts with autoEvents, skip the check of the Devices sheet header column
						continue
					}

					found := false
					for _, colCell := range row {
						if colCell == objectField {
							found = true
							break
						}
					}
					if !found && mapping.defaultValue != "" {
						err = addMissingColumn(xlsFile, devicesSheetName, colCount, len(rows), mapping.defaultValue, objectField)
						if err != nil {
							return fmt.Errorf("failed to add missing column: %w", err)
						}

						// add the new added column header to the header slice
						header = append(header, objectField)
						colCount++
					}
				}
			}
		}
		break
	}

	rows, err = xlsFile.GetRows(devicesSheetName)
	if err != nil {
		return err
	}

	// parse the device data rows
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}

		convertedDevice := dtos.Device{ProtocolName: protocol}
		_, err = ReadStruct(&convertedDevice, protocol, header, row)
		if err != nil {
			return err
		}

		// validate the device DTO
		err := common.Validate(convertedDevice)
		if err != nil {
			deviceErr := fmt.Errorf("device %s validation error: %v", convertedDevice.Name, err)
			deviceXlsx.validateErr = append(deviceXlsx.validateErr, deviceErr)
		} else {
			deviceXlsx.devices = append(deviceXlsx.devices, &convertedDevice)
		}
	}

	err = deviceXlsx.convertAutoEvents(xlsFile)
	if err != nil {
		return err
	}
	return nil
}

// convertAutoEvents parses the AutoEvents sheet and convert the rows to AutoEvent DTOs
func (deviceXlsx *DeviceXlsx) convertAutoEvents(xlsFile *excelize.File) error {
	var header []string

	rows, err := xlsFile.GetRows(autoEventsSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s worksheet: %w", autoEventsSheetName, err)
	}

	for rowIndex, row := range rows {
		// parse the header row and store the corresponding device field of each column and
		// parse the AutoEvents sheet header row and store the missing Object fields from MappingTable sheet
		if rowIndex == 0 {
			// get the column count of the header row to see if any Object field from MappingTable sheet is not defined
			colCount := len(row)

			for _, colCell := range row {
				header = append(header, colCell)
			}

			// AutoEvents sheet should define 4 columns in the header row
			if colCount != 4 {
				for objectField, mapping := range deviceXlsx.fieldMappings {
					if !startsWithAutoEvents(mapping.path) {
						// if the mapping path doesn't start with autoEvents, skip the check of the AutoEvents sheet header column
						continue
					}

					found := false
					for _, colCell := range row {
						if colCell == objectField {
							found = true
							break
						}
					}
					if !found && mapping.defaultValue != "" {
						err = addMissingColumn(xlsFile, autoEventsSheetName, colCount, len(rows), mapping.defaultValue, objectField)
						if err != nil {
							return fmt.Errorf("failed to add missing column: %w", err)
						}

						// add the new added column header to the header slice
						header = append(header, objectField)
						colCount++
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
		deviceNames, err := ReadStruct(&autoEvent, "", header, row)
		if err != nil {
			return fmt.Errorf("failed to unmarshal an excel row into AutoEvent DTO: %w", err)
		}

		// validate the AutoEvent DTO
		err = common.Validate(autoEvent)
		if err == nil {
			autoEventErr := fmt.Errorf("autoEvent validation error: %v", err)
			deviceXlsx.validateErr = append(deviceXlsx.validateErr, autoEventErr)
		}

		for _, deviceName := range deviceNames {
			for _, device := range deviceXlsx.devices {
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
