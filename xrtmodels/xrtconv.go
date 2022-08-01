// Copyright (C) 2022 IOTech Ltd

package xrtmodels

import (
	"fmt"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

func toXrtProperties(protocol string, protocolProperties map[string]interface{}) errors.EdgeX {
	intProperties, floatProperties := propertyConversionList(protocol)

	for _, p := range intProperties {
		propertyValue, ok := protocolProperties[p]
		if ok {
			// convert property value from interface{} to string, then to int
			val, err := strconv.Atoi(fmt.Sprintf("%v", propertyValue))
			if err != nil {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("fail to convert %v to int", p), err)
			}
			protocolProperties[p] = int(val)
		}
	}

	for _, p := range floatProperties {
		propertyValue, ok := protocolProperties[p]
		if ok {
			// convert property value from interface{} to string, then to float
			val, err := strconv.ParseFloat(fmt.Sprintf("%v", propertyValue), 64)
			if err != nil {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("fail to convert %v to float", p), err)
			}
			protocolProperties[p] = val
		}
	}
	return nil
}

func toEdgeXProperties(protocol string, protocolProperties map[string]interface{}) map[string]string {
	intProperties, floatProperties := propertyConversionList(protocol)

	edgexProperties := make(map[string]string)
	for k, v := range protocolProperties {
		edgexProperties[k] = fmt.Sprintf("%v", v)
	}

	for _, p := range intProperties {
		propertyValue, ok := protocolProperties[p]
		if ok {
			// if we use fmt.fmt.Sprintf("%v", propertyValue) to convert the float to string,
			// the 4194148 become 4.194148e+06 and dot(.), plus(+) are invalid for metadata
			// so we can use %.0f to convert the float without the decimal point
			edgexProperties[p] = fmt.Sprintf("%.0f", propertyValue)
		}
	}

	for _, p := range floatProperties {
		propertyValue, ok := protocolProperties[p]
		if ok {
			switch val := propertyValue.(type) {
			case float64:
				// The -1 as the third parameter tells the function to print the fewest digits necessary to accurately represent the float
				// For example:
				//   strconv.FormatFloat(5.2, 'f', -1, 64) -> 5.2
				//   fmt.Sprintf("%f",5.2) -> 5.200000
				edgexProperties[p] = strconv.FormatFloat(val, 'f', -1, 64)
			}
		}
	}
	return edgexProperties
}

func propertyConversionList(protocol string) ([]string, []string) {
	var intProperties []string
	var floatProperties []string
	switch protocol {
	case common.BacnetIP, common.BacnetMSTP:
		intProperties = []string{common.BacnetDeviceInstance}
	case common.Gps:
		intProperties = []string{common.GpsGpsdPort, common.GpsGpsdRetries, common.GpsGpsdConnTimeout, common.GpsGpsdRequestTimeout}
	case common.ModbusTcp:
		intProperties = []string{common.ModbusUnitID, common.ModbusPort}
	case common.ModbusRtu:
		intProperties = []string{common.ModbusUnitID, common.ModbusBaudRate, common.ModbusDataBits, common.ModbusStopBits}
	case common.Opcua:
		intProperties = []string{common.OpcuaRequestedSessionTimeout, common.OpcuaBrowseDepth, common.OpcuaConnectionReadingPostDelay}
		floatProperties = []string{common.OpcuaBrowsePublishInterval}
	case common.S7:
		intProperties = []string{common.S7Rack, common.S7Slot}
	}
	return intProperties, floatProperties
}
