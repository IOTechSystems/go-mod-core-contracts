//
// Copyright (C) 2023 IOTech Ltd
//

package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
)

// ReadStruct parses the xlsx data row to the struct type of the structPtr argument
func ReadStruct(structPtr interface{}, protocol string, headerCol []string, row []string) ([]string, error) {
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, errors.New("the structPtr argument should be a pointer of struct")
	}

	var returnedColumns []string
	elementType := v.Elem().Type()
	rowElement := reflect.New(elementType).Elem()

	headerLastIndex := len(headerCol) - 1
	// define the protocol property map for Device DTO
	prtPropertyMap := make(map[string]string)
	for colIndex, cell := range row {
		// check if row length is larger than the header
		if colIndex > headerLastIndex {
			break
		}

		headerName := headerCol[colIndex]
		field := rowElement.FieldByName(headerName)

		if field.Kind() != reflect.Invalid {
			var fieldValue interface{}
			switch field.Kind() {
			case reflect.Slice:
				values := strings.Split(cell, common.CommaSeparator)
				fieldValue = values
			case reflect.Bool:
				boolValue, err := strconv.ParseBool(strings.TrimSpace(cell))
				if err != nil {
					return nil, fmt.Errorf("failed to parse cell '%v' to bool type: %w", cell, err)
				}
				fieldValue = boolValue
			default:
				fieldValue = cell
			}
			// if the column header matches the struct field, set the cell value to the struct
			field.Set(reflect.ValueOf(fieldValue))
		} else {
			// headerName belongs to the Protocols fields of Device DTO
			//if elementType == reflect.TypeOf(dtos.Device{}) {
			//	prtPropertyMap[headerName] = cell
			//}
			switch elementType {
			case reflect.TypeOf(dtos.Device{}):
				prtPropertyMap[headerName] = cell
			case reflect.TypeOf(dtos.AutoEvent{}):
				returnedColumns = append(returnedColumns, cell)
			}
		}
	}

	if len(prtPropertyMap) > 0 {
		// set ProtocolProperties map to the Protocols field of Device DTO
		prtProp := setProtocolPropMap(protocol, prtPropertyMap)
		prtPropertyField := rowElement.FieldByName(protocols)
		prtPropertyField.Set(reflect.ValueOf(prtProp))
	}

	v.Elem().Set(rowElement)

	return returnedColumns, nil
}

// setProtocolPropMap sets the ProtocolProperties outer map key based on protocol and returns the map
func setProtocolPropMap(protocol string, prtProps map[string]string) map[string]dtos.ProtocolProperties {
	prtPropMap := make(map[string]dtos.ProtocolProperties)
	switch protocol {
	case modbusRTUKey:
		prtPropMap = map[string]dtos.ProtocolProperties{modbusRTUKey: prtProps}
	}
	return prtPropMap
}
