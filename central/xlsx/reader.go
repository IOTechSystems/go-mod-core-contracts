//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xlsx

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
)

func readStruct(structPtr any, headerCol []string, row []string, mapppingTable map[string]mappingField) (any, error) {
	var extraReturnedCols any
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, errors.New("the structPtr argument should be a pointer of struct")
	}

	elementType := v.Elem().Type()
	rowElement := reflect.New(elementType).Elem()

	var err error
	switch elementType {
	case reflect.TypeOf(dtos.DeviceProfile{}):
		err = convertDTOStdTypeFields(&rowElement, row, headerCol, mapppingTable)
	case reflect.TypeOf(dtos.AutoEvent{}):
		extraReturnedCols, err = convertAutoEventFields(&rowElement, row, headerCol, mapppingTable)
	case reflect.TypeOf(dtos.Device{}):
		err = convertDeviceFields(&rowElement, row, headerCol, mapppingTable)
	case reflect.TypeOf(dtos.DeviceCommand{}):
		err = convertDeviceCommandFields(&rowElement, row, headerCol)
	case reflect.TypeOf(dtos.DeviceResource{}):
		err = convertResourcesFields(&rowElement, row, headerCol, mapppingTable)
	default:
		// skip the processing of the not found field name
		err = fmt.Errorf("unknown converted DTO type '%T'", elementType)
	}
	if err != nil {
		return nil, err
	}

	v.Elem().Set(rowElement)
	return extraReturnedCols, nil
}

// getStructFieldByHeader returns the passed structEle struct field by headerName
func getStructFieldByHeader(structEle *reflect.Value, colIndex int, headerCol []string) (string, reflect.Value) {
	var headerName string
	headerLastIndex := len(headerCol) - 1
	// check if row length is larger than the header
	if colIndex > headerLastIndex {
		headerName = headerCol[headerLastIndex]
	} else {
		headerName = headerCol[colIndex]
	}
	field := structEle.FieldByName(headerName)
	return headerName, field
}

// setStdStructFieldValue set the struct field with Go standard types to the xlsx cell value
func setStdStructFieldValue(originValue string, field reflect.Value) error {
	var fieldValue any
	switch field.Kind() {
	case reflect.String:
		fieldValue = originValue
	case reflect.Slice:
		values := strings.Split(originValue, common.CommaSeparator)
		fieldValue = values
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(originValue)
		if err != nil {
			return fmt.Errorf("failed to parse originValue '%v' to bool type: %w", originValue, err)
		}
		fieldValue = boolValue
	case reflect.Int64:
		int64Value, err := strconv.ParseInt(originValue, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse originValue '%v' to Int64 type: %w", originValue, err)
		}
		fieldValue = int64Value
	case reflect.Interface:
		fieldValue = originValue
	default:
		return fmt.Errorf("failed to parse originValue '%v' to %s type", originValue, field.Type())
	}

	field.Set(reflect.ValueOf(fieldValue))
	return nil
}

// convertDTOStdTypeFields unmarshalls the xlsx cells into the standard type fields of the DTO struct
func convertDTOStdTypeFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string, fieldMappings map[string]mappingField) error {
	for colIndex, cell := range xlsxRow {
		headerName, field := getStructFieldByHeader(rowElement, colIndex, headerCol)
		fieldValue := strings.TrimSpace(cell)

		if field.Kind() != reflect.Invalid {
			if fieldValue == "" {
				// set the struct field value to 'default value' defined in mapping Table if not empty
				if mapping, ok := fieldMappings[headerName]; ok && mapping.defaultValue != "" {
					fieldValue = mapping.defaultValue
				}
			}

			err := setStdStructFieldValue(fieldValue, field)
			if err != nil {
				return err
			}
		} else {
			// field not found in the DTO struct, skip this column
			continue
		}
	}
	return nil
}

// setProtocolPropMap sets the ProtocolProperties outer map key based on protocol and returns the Protocols map
func setProtocolPropMap(prtProps map[string]string, fieldMappings map[string]mappingField) (map[string]dtos.ProtocolProperties, error) {
	var protocol string
	prtPropMap := make(map[string]dtos.ProtocolProperties)

	if mapping, ok := fieldMappings[protocolName]; ok {
		protocol = mapping.defaultValue
	} else {
		return nil, errors.New("ProtocolName not defined in fieldMappings")
	}

	switch protocol {
	case modbusRTUKey:
		prtPropMap[modbusRTUKey] = prtProps
	case modbusTCPKey:
		prtPropMap[modbusTCPKey] = prtProps
	default:
		return nil, fmt.Errorf("unknown ProtocolProperties outer key for '%s' protocol", protocol)
	}
	return prtPropMap, nil
}

// convertDeviceFields convert the xlsx row to the Device DTO
func convertDeviceFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string, fieldMappings map[string]mappingField) error {
	if fieldMappings == nil {
		return errors.New("fieldMappings not defined while converting device fields")
	}
	protocolProperties := dtos.ProtocolProperties{}

	for colIndex, cell := range xlsxRow {
		headerName, field := getStructFieldByHeader(rowElement, colIndex, headerCol)
		fieldValue := strings.TrimSpace(cell)
		if fieldValue == "" {
			// set fieldValue to 'default value' defined in mapping Table if not empty
			if mapping, ok := fieldMappings[headerName]; ok && mapping.defaultValue != "" {
				fieldValue = mapping.defaultValue
			}
		}

		if field.Kind() != reflect.Invalid {
			// header matches the Device DTO field name (one of the Name, Description, AdminState, OperatingState, etc)
			err := setStdStructFieldValue(fieldValue, field)
			if err != nil {
				return err
			}
		} else {
			// set the cell to Protocols map if header not belongs to the above fields with standard types
			if fieldValue != "" {
				protocolProperties[headerName] = fieldValue
			}
		}
	}

	if len(protocolProperties) > 0 {
		prtField := rowElement.FieldByName(protocols)
		if prtField.Kind() == reflect.Invalid {
			return errors.New("failed to find Protocols field in Device DTO")
		}
		prtPropMap, err := setProtocolPropMap(protocolProperties, fieldMappings)
		if err != nil {
			return err
		}
		prtField.Set(reflect.ValueOf(prtPropMap))
	}
	return nil
}

// convertAutoEventFields convert the xlsx row to the AutoEvent DTO
func convertAutoEventFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string, fieldMappings map[string]mappingField) ([]string, error) {
	if fieldMappings == nil {
		return nil, errors.New("fieldMappings not defined while converting AutoEvent fields")
	}
	var deviceNames []string

	for colIndex, cell := range xlsxRow {
		headerName, field := getStructFieldByHeader(rowElement, colIndex, headerCol)
		fieldValue := strings.TrimSpace(cell)
		if fieldValue == "" {
			// set fieldValue to 'default value' defined in mapping Table if not empty
			if mapping, ok := fieldMappings[headerName]; ok && mapping.defaultValue != "" {
				fieldValue = mapping.defaultValue
			}
		}

		if field.Kind() != reflect.Invalid {
			// header matches the AutoEvent DTO field name (one of the Interval, OnChange, SourceName field)
			err := setStdStructFieldValue(fieldValue, field)
			if err != nil {
				return nil, err
			}
		} else {
			// the cell belongs to the "Reference Device Name" column, append it to deviceNames
			if fieldValue != "" {
				deviceNames = append(deviceNames, fieldValue)
			}
		}
	}

	return deviceNames, nil
}

// convertDeviceCommandFields convert the xlsx row to the DeviceCommand DTO
func convertDeviceCommandFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string) error {
	var resOpSlice []dtos.ResourceOperation
	for colIndex, cell := range xlsxRow {
		_, field := getStructFieldByHeader(rowElement, colIndex, headerCol)
		cell = strings.TrimSpace(cell)

		if field.Kind() != reflect.Invalid {
			// header matches the DeviceCommand field name (one of the Name, IsHidden or ReadWrite field name)
			err := setStdStructFieldValue(cell, field)
			if err != nil {
				return err
			}
		} else {
			// parse the rest ResourceName columns in the xlsx row and convert to the ResourceOperation DTO
			resOp := dtos.ResourceOperation{
				DeviceResource: cell,
			}
			resOpSlice = append(resOpSlice, resOp)
		}
	}

	if len(resOpSlice) > 0 {
		// set resOpSlice to the ResourceOperations field of DeviceCommand struct
		resOpField := rowElement.FieldByName(resourceOperations)
		if resOpField.Kind() == reflect.Invalid {
			return errors.New("failed to find ResourceOperations field in DeviceCommand DTO")
		}
		resOpField.Set(reflect.ValueOf(resOpSlice))
	}
	return nil
}

// convertResourcesFields convert the xlsx row to the DeviceResource DTO
func convertResourcesFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string, fieldMappings map[string]mappingField) error {
	if fieldMappings == nil {
		return errors.New("fieldMappings not defined while converting DeviceResource fields")
	}

	for colIndex, cell := range xlsxRow {
		headerName, field := getStructFieldByHeader(rowElement, colIndex, headerCol)
		fieldValue := strings.TrimSpace(cell)
		if fieldValue == "" {
			// set fieldValue to 'default value' defined in mapping Table if not empty
			if mapping, ok := fieldMappings[headerName]; ok && mapping.defaultValue != "" {
				fieldValue = mapping.defaultValue
			}
		}

		if field.Kind() != reflect.Invalid {
			// header matches the DeviceResource field name (one of the Name, Description or IsHidden field name)
			err := setStdStructFieldValue(fieldValue, field)
			if err != nil {
				return err
			}
		} else {
			resPropField := rowElement.FieldByName(properties).FieldByName(headerName)
			if resPropField.Kind() != reflect.Invalid {
				// header matches the ResourceProperties DTO field name (one of the ValueType, ReadWrite, Units, etc)
				err := setStdStructFieldValue(fieldValue, resPropField)
				if err != nil {
					return err
				}
			} else {
				// set the cell to Attributes map if header not belongs to Properties field
				if fieldValue != "" {
					attrMapField := rowElement.FieldByName(attributes)
					if attrMapField.Len() == 0 {
						// initialize the Attributes map
						attrMap := make(map[string]any)
						attrMapField.Set(reflect.MakeMap(reflect.TypeOf(attrMap)))
					}
					attrMapField.SetMapIndex(reflect.ValueOf(headerName), reflect.ValueOf(fieldValue))
				}
			}
		}
	}

	return nil
}
