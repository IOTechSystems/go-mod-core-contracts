//
// Copyright (C) 2023 IOTech Ltd
//

package xlsx

import (
	"fmt"
	"io"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"

	"github.com/xuri/excelize/v2"
	"golang.org/x/exp/slices"
)

// requiredSheets defines the required worksheet names in the xlsx file
var requiredSheets = []string{devicesSheetName}

// deviceXlsx stores the worksheets processed result and the converted Device DTOs
type deviceXlsx struct {
	xlsFile        *excelize.File
	fieldMappings  map[string]mappingField // fieldMappings defines all the device fields with default values defined in the xlsx
	devices        []*dtos.Device
	ValidateErrors []error
}

func newDeviceXlsx(file io.Reader) (*deviceXlsx, error) {
	// file io.Reader should be closed from the caller in another module
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}

	fieldMappings, err := convertMappingTable(f)
	if err != nil {
		return nil, err
	}
	return &deviceXlsx{xlsFile: f, fieldMappings: fieldMappings}, nil
}

// convertToDTO parses the Devices sheet and convert the rows to Device DTOs
func (deviceXlsx *deviceXlsx) convertToDTO() error {
	allSheetNames := deviceXlsx.xlsFile.GetSheetList()

	err := checkRequiredSheets(allSheetNames, requiredSheets)
	if err != nil {
		return err
	}

	var header []string
	xlsFile := deviceXlsx.xlsFile
	protocol := deviceXlsx.fieldMappings[protocolName].defaultValue

	rows, err := xlsFile.GetRows(devicesSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s worksheet: %w", devicesSheetName, err)
	}

	// checks at least 2 rows exists in the Devices sheet (1 header and 1 data row)
	// and parses the header row
	if len(rows) >= 2 {
		header = rows[0]
		err = deviceXlsx.parseDevicesHeader(&header, len(rows))
		if err != nil {
			return fmt.Errorf("failed to parse the header row from %s worksheet: %w", devicesSheetName, err)
		}
	} else {
		return fmt.Errorf("at least 2 rows need to be defined in %s worksheet", devicesSheetName)
	}

	// retrieve all rows again as new columns might be added while the Header row
	rows, err = xlsFile.GetRows(devicesSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s worksheet after inserting misshing columns: %w", devicesSheetName, err)
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
			deviceXlsx.devices = append(deviceXlsx.devices, &convertedDevice)
		}
	}

	if slices.Contains(allSheetNames, autoEventsSheetName) {
		err = deviceXlsx.convertAutoEvents()
		if err != nil {
			return fmt.Errorf("failed to convert AutoEvents worksheet: %w", err)
		}
	}

	return nil
}

func (deviceXlsx *deviceXlsx) parseDevicesHeader(header *[]string, rowCount int) error {
	var err error
	// get the column count of the header row to see if any Object field from MappingTable sheet is not defined
	colCount := len(*header)

	for objectField, mapping := range deviceXlsx.fieldMappings {
		if startsWithAutoEvents(mapping.path) {
			// if the mapping path starts with autoEvents, skip the check of the Devices sheet header column
			continue
		}

		// check if the mapping object is defined in the Devices sheet if the defaultValue is not empty
		// if not, insert the mapping object as a new column in the Devices sheet with defaultValue set in each data row
		if mapping.defaultValue != "" {
			err = checkMappingObject(deviceXlsx.xlsFile, devicesSheetName, &colCount, rowCount, mapping.defaultValue, objectField, header)
			if err != nil {
				return fmt.Errorf("failed to check mapping object: %w", err)
			}
		}
	}

	return nil
}

// convertAutoEvents parses the AutoEvents sheet and convert the rows to AutoEvent DTOs
func (deviceXlsx *deviceXlsx) convertAutoEvents() error {
	var header []string
	xlsFile := deviceXlsx.xlsFile

	rows, err := xlsFile.GetRows(autoEventsSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s worksheet: %w", autoEventsSheetName, err)
	}

	// checks at least 2 rows exists in the AutoEvents sheet (1 header and 1 data row)
	// and parses the header row
	if len(rows) >= 2 {
		header = rows[0]
		// parse the header row
		// get the column count of the header row to see if any Object field from MappingTable sheet is not defined
		colCount := len(header)

		// AutoEvents sheet should at least define 2 columns in the header row (SourceName and Reference Device Name)
		if colCount < 2 {
			err = deviceXlsx.parseAutoEventsHeader(header, len(rows))
			if err != nil {
				return fmt.Errorf("failed to parse the header row from %s worksheet: %w", autoEventsSheetName, err)
			}
		}
	} else {
		return nil
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
			for _, device := range deviceXlsx.devices {
				if deviceName == device.Name {
					device.AutoEvents = append(device.AutoEvents, autoEvent)
				}
			}
		}
	}

	return nil
}

func (deviceXlsx *deviceXlsx) parseAutoEventsHeader(header []string, rowCount int) error {
	var err error
	colCount := len(header)
	newColCount := &colCount

	for objectField, mapping := range deviceXlsx.fieldMappings {
		if !startsWithAutoEvents(mapping.path) {
			// if the mapping path doesn't start with autoEvents, skip the check of the AutoEvents sheet header column
			continue
		}

		// check if the mapping object is defined in the AutoEvents sheet if the defaultValue is not empty
		// if not, insert the mapping object as a new column in the Devices sheet with defaultValue set in each data row
		err = checkMappingObject(deviceXlsx.xlsFile, autoEventsSheetName, newColCount, rowCount, mapping.defaultValue, objectField, &header)
		if err != nil {
			return fmt.Errorf("failed to check mapping object: %w", err)
		}
	}

	return nil
}

// startsWithAutoEvents checks if the path name defined in MappingTable sheet starts with autoEvents
func startsWithAutoEvents(path string) bool {
	return strings.HasPrefix(strings.ToLower(path), strings.ToLower(autoEvents))
}

func (deviceXlsx *deviceXlsx) GetDTOs() any {
	return deviceXlsx.devices
}

func (deviceXlsx *deviceXlsx) GetValidateErrors() []error {
	return deviceXlsx.ValidateErrors
}
