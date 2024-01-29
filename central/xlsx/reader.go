//
// Copyright (C) 2023-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xlsx

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

func readStruct(structPtr any, headerCol []string, row []string, mapppingTable map[string]mappingField) (any, errors.EdgeX) {
	var extraReturnedCols any
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "the structPtr argument should be a pointer of struct", nil)
	}

	elementType := v.Elem().Type()
	rowElement := reflect.New(elementType).Elem()

	var err errors.EdgeX
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
		err = errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("unknown converted DTO type '%T'", elementType), nil)
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
func setStdStructFieldValue(originValue string, field reflect.Value) errors.EdgeX {
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
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("failed to parse originValue '%v' to bool type", originValue), err)
		}
		fieldValue = boolValue
	case reflect.Int64:
		int64Value, err := strconv.ParseInt(originValue, 10, 64)
		if err != nil {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("failed to parse originValue '%v' to Int64 type", originValue), err)
		}
		fieldValue = int64Value
	case reflect.Interface:
		fieldValue = originValue
	default:
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("failed to parse originValue '%v' to %s type", originValue, field.Type()), nil)
	}

	field.Set(reflect.ValueOf(fieldValue))
	return nil
}

// convertDTOStdTypeFields unmarshalls the xlsx cells into the standard type fields of the DTO struct
func convertDTOStdTypeFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string, fieldMappings map[string]mappingField) errors.EdgeX {
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
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("error occurred on '%s' column", headerName), err)
			}
		} else {
			// field not found in the DTO struct, skip this column
			continue
		}
	}
	return nil
}

// setProtocolPropMap sets the ProtocolProperties outer map key based on protocol and returns the Protocols map
func setProtocolPropMap(prtProps map[string]string, fieldMappings map[string]mappingField) (map[string]dtos.ProtocolProperties, errors.EdgeX) {
	var protocol string
	prtPropMap := make(map[string]dtos.ProtocolProperties)

	if mapping, ok := fieldMappings[protocolName]; ok {
		protocol = mapping.defaultValue
	} else {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "ProtocolName not defined in fieldMappings", nil)
	}

	switch protocol {
	case modbusRTUKey:
		prtPropMap[modbusRTUKey] = prtProps
	case modbusTCPKey:
		prtPropMap[modbusTCPKey] = prtProps
	default:
		return nil, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("unknown ProtocolProperties outer key for '%s' protocol", protocol), nil)
	}
	return prtPropMap, nil
}

// convertDeviceFields convert the xlsx row to the Device DTO
func convertDeviceFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string, fieldMappings map[string]mappingField) errors.EdgeX {
	if fieldMappings == nil {
		return errors.NewCommonEdgeX(errors.KindServerError, "fieldMappings not defined while converting device fields", nil)
	}
	protocolProperties := dtos.ProtocolProperties{}
	tagsMap := make(map[string]any)

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
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("error occurred on '%s' column", headerName), err)
			}
		} else {
			// header not belongs to the above fields with standard types
			// map the cell to the Protocols or Tags field
			if fieldValue != "" {
				// get the Path defined in the MappingTable
				if mapping, ok := fieldMappings[headerName]; ok && mapping.path != "" {
					path := mapping.path
					fieldPrefix := strings.SplitN(path, mappingPathSeparator, 2)[0]
					switch fieldPrefix {
					case strings.ToLower(protocols):
						// set the cell to Protocols map
						protocolProperties[headerName] = fieldValue
					case strings.ToLower(tags):
						// set the cell to Tags map
						tagsMap[headerName] = fieldValue
					default:
						// unknown column header
						continue
					}
				}
			}
		}
	}

	// set Protocols field to the Device DTO struct
	if len(protocolProperties) > 0 {
		prtField := rowElement.FieldByName(protocols)
		if prtField.Kind() == reflect.Invalid {
			return errors.NewCommonEdgeX(errors.KindServerError, "failed to find Protocols field in Device DTO", nil)
		}
		prtPropMap, err := setProtocolPropMap(protocolProperties, fieldMappings)
		if err != nil {
			return err
		}
		prtField.Set(reflect.ValueOf(prtPropMap))
	}
	// set Tags field to the Device DTO struct
	if len(tagsMap) > 0 {
		tagsField := rowElement.FieldByName(tags)
		if tagsField.Kind() == reflect.Invalid {
			return errors.NewCommonEdgeX(errors.KindServerError, "failed to find Tags field in Device DTO", nil)
		}
		tagsField.Set(reflect.ValueOf(tagsMap))
	}

	return nil
}

// convertAutoEventFields convert the xlsx row to the AutoEvent DTO
func convertAutoEventFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string, fieldMappings map[string]mappingField) ([]string, errors.EdgeX) {
	if fieldMappings == nil {
		return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "fieldMappings not defined while converting AutoEvent fields", nil)
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
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("error occurred on '%s' column", headerName), err)
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
func convertDeviceCommandFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string) errors.EdgeX {
	var resOpSlice []dtos.ResourceOperation
	for colIndex, cell := range xlsxRow {
		headerName, field := getStructFieldByHeader(rowElement, colIndex, headerCol)
		cell = strings.TrimSpace(cell)

		if field.Kind() != reflect.Invalid {
			// header matches the DeviceCommand field name (one of the Name, IsHidden or ReadWrite field name)
			err := setStdStructFieldValue(cell, field)
			if err != nil {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("error occurred on '%s' column", headerName), err)
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
			return errors.NewCommonEdgeX(errors.KindServerError, "failed to find ResourceOperations field in DeviceCommand DTO", nil)
		}
		resOpField.Set(reflect.ValueOf(resOpSlice))
	}
	return nil
}

// convertResourcesFields convert the xlsx row to the DeviceResource DTO
func convertResourcesFields(rowElement *reflect.Value, xlsxRow []string, headerCol []string, fieldMappings map[string]mappingField) errors.EdgeX {
	if fieldMappings == nil {
		return errors.NewCommonEdgeX(errors.KindServerError, "fieldMappings not defined while converting DeviceResource fields", nil)
	}

	for colIndex, cell := range xlsxRow {
		headerName, field := getStructFieldByHeader(rowElement, colIndex, headerCol)
		fieldValue := strings.TrimSpace(cell)

		var fieldMapping mappingField
		// set fieldValue to 'default value' defined in mapping Table if not empty
		if mapping, ok := fieldMappings[headerName]; ok {
			fieldMapping = mapping
			if fieldValue == "" && mapping.defaultValue != "" {
				fieldValue = mapping.defaultValue
			}
		}

		if field.Kind() != reflect.Invalid {
			// header matches the DeviceResource field name (one of the Name, Description or IsHidden field name)
			err := setStdStructFieldValue(fieldValue, field)
			if err != nil {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("error occurred on '%s' column", headerName), err)
			}
		} else {
			if fieldValue != "" && strings.Contains(strings.ToLower(fieldMapping.path), strings.ToLower(tags)) {
				// set the cell to Device Resource tags map
				tagsMapField := rowElement.FieldByName(tags)
				if tagsMapField.Len() == 0 {
					// initialize the Attributes map
					tagsMap := make(map[string]any)
					tagsMapField.Set(reflect.MakeMap(reflect.TypeOf(tagsMap)))
				}

				tagValue := convertCellToAny(fieldValue)
				tagsMapField.SetMapIndex(reflect.ValueOf(headerName), reflect.ValueOf(tagValue))
			}

			resPropField := rowElement.FieldByName(properties).FieldByName(headerName)
			if resPropField.Kind() != reflect.Invalid {
				// header matches the ResourceProperties DTO field name (one of the ValueType, ReadWrite, Units, etc)
				err := setStdStructFieldValue(fieldValue, resPropField)
				if err != nil {
					return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("error occurred on '%s' column", headerName), err)
				}
			} else {
				// set the cell to Attributes map if header not belongs to Properties field and mapping path contains 'attributes'
				if fieldValue != "" && strings.Contains(strings.ToLower(fieldMapping.path), strings.ToLower(attributes)) {
					attrMapField := rowElement.FieldByName(attributes)
					if attrMapField.Len() == 0 {
						// initialize the Attributes map
						attrMap := make(map[string]any)
						attrMapField.Set(reflect.MakeMap(reflect.TypeOf(attrMap)))
					}

					attrValue := convertCellToAny(fieldValue)
					attrMapField.SetMapIndex(reflect.ValueOf(headerName), reflect.ValueOf(attrValue))
				}
			}
		}
	}

	return nil
}

// to check if the cell value can be converted to int, float, bool data types
// if not, keep the cell value as string
func convertCellToAny(fieldValue string) any {
	var convertedValue any
	if intValue, err := strconv.ParseInt(fieldValue, 10, 16); err == nil {
		convertedValue = intValue
	} else if floatValue, err := strconv.ParseFloat(fieldValue, 64); err == nil {
		convertedValue = floatValue
	} else if boolValue, err := strconv.ParseBool(fieldValue); err == nil {
		convertedValue = boolValue
	} else {
		convertedValue = fieldValue
	}
	return convertedValue
}
