//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xlsx

import (
	"fmt"
	"io"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"

	"github.com/xuri/excelize/v2"
	"golang.org/x/exp/slices"
)

const (
	validateErrResourcePrefix = "deviceResource_"
	validateErrCommandPrefix  = "deviceCommand_"
	validateErrProfilePrefix  = "deviceProfile_"
)

// requiredProfileSheets defines the required worksheet names in the xlsx file
var requiredProfileSheets = []string{deviceInfoSheetName, deviceResourceSheetName}

// deviceProfileXlsx stores the worksheets processed result and the converted DeviceProfile DTO
type deviceProfileXlsx struct {
	baseXlsx
	deviceProfile *dtos.DeviceProfile
}

func newDeviceProfileXlsx(file io.Reader) (Converter[*dtos.DeviceProfile], error) {
	// file io.Reader should be closed from the caller in another module
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}

	fieldMappings, err := convertMappingTable(f)
	if err != nil {
		return nil, err
	}
	return &deviceProfileXlsx{
		baseXlsx: baseXlsx{
			xlsFile:        f,
			fieldMappings:  fieldMappings,
			validateErrors: make(map[string]error),
		},
	}, nil
}

// ConvertToDTO parses the DeviceInfo/DeviceResource/DeviceCommand sheets and convert the rows to DeviceProfile DTO
func (dpXlsx *deviceProfileXlsx) ConvertToDTO() error {
	allSheetNames := dpXlsx.xlsFile.GetSheetList()

	err := checkRequiredSheets(allSheetNames, requiredProfileSheets)
	if err != nil {
		return err
	}

	convertedProfile := &dtos.DeviceProfile{}
	err = dpXlsx.convertDeviceInfo(convertedProfile)
	if err != nil {
		return err
	}

	// parse the DeviceResource sheet
	err = dpXlsx.convertDeviceResources(convertedProfile)
	if err != nil {
		return err
	}

	if slices.Contains(allSheetNames, deviceCommandSheetName) {
		// parse the DeviceCommand sheet
		err = dpXlsx.convertDeviceCommands(convertedProfile)
		if err != nil {
			return err
		}
	}

	// validate the device profile DTO
	err = convertedProfile.Validate()
	if err != nil {
		dpXlsx.validateErrors[validateErrProfilePrefix+convertedProfile.DeviceProfileBasicInfo.Name] = err
	} else if dpXlsx.validateErrors != nil {
		dpXlsx.deviceProfile = convertedProfile
	}
	return nil
}

// convertDeviceInfo parses the DeviceInfo sheet and convert the rows to DeviceProfile DTO
func (dpXlsx *deviceProfileXlsx) convertDeviceInfo(convertedProfile *dtos.DeviceProfile) error {
	var header []string
	cols, err := dpXlsx.xlsFile.GetCols(deviceInfoSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all columns from %s worksheet: %w", deviceInfoSheetName, err)
	}

	// checks at least 2 columns exists in the DeviceInfo sheet (1 header and 1 data column)
	// and parses the header column
	if len(cols) >= 2 {
		header = cols[0]
	} else {
		return fmt.Errorf("at least 2 columns need to be defined in %s worksheet", deviceInfoSheetName)
	}

	// parse the DeviceInfo data column
	_, err = readStruct(convertedProfile, header, cols[1], dpXlsx.fieldMappings)
	if err != nil {
		return fmt.Errorf("failed to unmarshal an xlsx column into DeviceProfile DTO: %w", err)
	}
	return nil
}

// convertDeviceResources parses the DeviceResource sheet and convert the rows to DeviceResource DTOs
func (dpXlsx *deviceProfileXlsx) convertDeviceResources(convertedProfile *dtos.DeviceProfile) error {
	var header []string
	rows, err := dpXlsx.xlsFile.GetRows(deviceResourceSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s worksheet: %w", deviceResourceSheetName, err)
	}

	// checks at least 2 rows exists in the DeviceResource sheet (1 header and 1 data row)
	// and parses the header row
	if len(rows) >= 2 {
		header = rows[0]
		err = dpXlsx.parseDeviceResourceHeader(&header, len(rows))
		if err != nil {
			return fmt.Errorf("failed to parse the header row from %s worksheet: %w", deviceResourceSheetName, err)
		}
	} else {
		return fmt.Errorf("at least 2 rows need to be defined in %s worksheet", deviceResourceSheetName)
	}

	// retrieve all rows again as new columns might be added while the Header row
	rows, err = dpXlsx.xlsFile.GetRows(deviceResourceSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s worksheet after inserting misshing columns: %w", deviceResourceSheetName, err)
	}

	// parse the device resource data rows
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}

		convertedDR := dtos.DeviceResource{}
		_, err = readStruct(&convertedDR, header, row, dpXlsx.fieldMappings)
		if err != nil {
			return fmt.Errorf("failed to unmarshal an xlsx row into DeviceResource DTO: %w", err)
		}

		// validate the DeviceResource DTO
		err = convertedDR.Validate()
		if err != nil {
			dpXlsx.validateErrors[validateErrResourcePrefix+convertedDR.Name] = err
		} else {
			convertedProfile.DeviceResources = append(convertedProfile.DeviceResources, convertedDR)
		}
	}
	return nil
}

func (dpXlsx *deviceProfileXlsx) parseDeviceResourceHeader(header *[]string, rowCount int) error {
	var err error
	// get the column count of the header row to see if any Object field from MappingTable sheet is not defined
	colCount := len(*header)

	for objectField, mapping := range dpXlsx.fieldMappings {
		// check if the mapping object is defined in the DeviceResource sheet if the defaultValue is not empty
		// if not, insert the mapping object as a new column in the DeviceResource sheet with defaultValue set in each data row
		if mapping.defaultValue != "" {
			err = checkMappingObject(dpXlsx.xlsFile, deviceResourceSheetName, &colCount, rowCount, mapping.defaultValue, objectField, header)
			if err != nil {
				return fmt.Errorf("failed to check mapping object: %w", err)
			}
		}
	}

	return nil
}

// convertDeviceCommands parses the DeviceCommand sheet and convert the rows to DeviceCommand DTOs
func (dpXlsx *deviceProfileXlsx) convertDeviceCommands(convertedProfile *dtos.DeviceProfile) error {
	var header []string
	rows, err := dpXlsx.xlsFile.GetRows(deviceCommandSheetName)
	if err != nil {
		return fmt.Errorf("failed to retrieve all rows from %s worksheet: %w", deviceCommandSheetName, err)
	}

	// checks at least 2 rows exists in the DeviceCommand sheet (1 header and 1 data row)
	if len(rows) >= 2 {
		header = rows[0]
	} else {
		return fmt.Errorf("at least 2 rows need to be defined in %s worksheet", deviceCommandSheetName)
	}

	// parse the Device Command data rows
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}

		convertedDC := dtos.DeviceCommand{}
		_, err = readStruct(&convertedDC, header, row, nil)
		if err != nil {
			return fmt.Errorf("failed to unmarshal an xlsx row into DeviceCommand DTO: %w", err)
		}

		// validate the DeviceCommand DTO
		err = common.Validate(convertedDC)
		if err != nil {
			dpXlsx.validateErrors[validateErrCommandPrefix+convertedDC.Name] = err
		} else {
			convertedProfile.DeviceCommands = append(convertedProfile.DeviceCommands, convertedDC)
		}
	}
	return nil
}

func (dpXlsx *deviceProfileXlsx) GetDTOs() *dtos.DeviceProfile {
	return dpXlsx.deviceProfile
}

func (dpXlsx *deviceProfileXlsx) GetValidateErrors() map[string]error {
	return dpXlsx.validateErrors
}