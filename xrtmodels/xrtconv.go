// Copyright (C) 2022 IOTech Ltd

package xrtmodels

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

func toXrtProperties(protocol string, protocolProperties map[string]interface{}) errors.EdgeX {
	intProperties, floatProperties, boolProperties := propertyConversionList(protocol)

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

	for _, p := range boolProperties {
		propertyValue, ok := protocolProperties[p]
		if ok {
			// convert property value from interface{} to string, then to bool
			val, err := strconv.ParseBool(fmt.Sprintf("%v", propertyValue))
			if err != nil {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("fail to convert %v to bool", p), err)
			}
			protocolProperties[p] = val
		}
	}
	return nil
}

func toEdgeXProperties(protocol string, protocolProperties map[string]interface{}) map[string]string {
	intProperties, floatProperties, boolProperties := propertyConversionList(protocol)

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

	for _, p := range boolProperties {
		propertyValue, ok := protocolProperties[p]
		if ok {
			edgexProperties[p] = fmt.Sprintf("%v", propertyValue)
		}
	}
	return edgexProperties
}

func propertyConversionList(protocol string) ([]string, []string, []string) {
	var intProperties []string
	var floatProperties []string
	var boolProperties []string
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
	case common.EtherNetIPExplicitConnected:
		intProperties = []string{common.EtherNetIPRPI}
		boolProperties = []string{common.EtherNetIPSaveValue}
	case common.EtherNetIPO2T, common.EtherNetIPT2O:
		intProperties = []string{common.EtherNetIPRPI}
	case common.EtherNetIPKey:
		intProperties = []string{common.EtherNetIPVendorID, common.EtherNetIPDeviceType, common.EtherNetIPProductCode,
			common.EtherNetIPMajorRevision, common.EtherNetIPMinorRevision}
	}
	return intProperties, floatProperties, boolProperties
}

func ToEdgeXV2EventDTO(xrtEvent MultiResourcesResult) (dtos.Event, errors.EdgeX) {
	event := dtos.Event{
		DeviceName:  xrtEvent.Device,
		ProfileName: xrtEvent.Profile,
		SourceName:  xrtEvent.SourceName,
		Tags:        xrtEvent.Tags,
		Readings:    make([]dtos.BaseReading, len(xrtEvent.Readings)),
	}

	index := 0
	for resourceName, reading := range xrtEvent.Readings {
		valueType, err := common.NormalizeValueType(reading.Type)
		if err != nil {
			return event, errors.NewCommonEdgeXWrapper(err)
		}
		value, err := ParseXRTReadingValue(valueType, reading.Value)
		if err != nil {
			return event, errors.NewCommonEdgeXWrapper(err)
		}

		switch valueType {
		case common.ValueTypeBinary:
			if data, ok := value.([]byte); ok {
				event.Readings[index] = dtos.NewBinaryReading(xrtEvent.Profile, xrtEvent.Device, resourceName, data, "")
			} else {
				return event, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid binary value '%v'", value), nil)
			}
		case common.ValueTypeObject:
			event.Readings[index] = dtos.NewObjectReading(xrtEvent.Profile, xrtEvent.Device, resourceName, value)

		default:
			event.Readings[index], err = dtos.NewSimpleReading(xrtEvent.Profile, xrtEvent.Device, resourceName, valueType, value)
			if err != nil {
				return event, errors.NewCommonEdgeXWrapper(err)
			}
		}
		event.Readings[index].Origin = reading.Origin
		event.Readings[index].Tags = reading.Tags
		event.Origin = reading.Origin
		index++
	}

	return event, nil
}

// ParseXRTReadingValue parses the XRT reading value to EdgeX reading value
func ParseXRTReadingValue(valueType string, reading interface{}) (interface{}, errors.EdgeX) {
	// Since we receive the reading in JSON format, the JSON lib will unmarshal the reading to specified data type:
	// bool for JSON booleans,  float64 for JSON numbers, nil for JSON null
	// string for JSON strings, []interface{} for JSON arrays, map[string]interface{} for JSON objects
	var err error
	var val interface{}
	switch valueType {
	case common.ValueTypeString:
		strValue, ok := reading.(string)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid string value '%v'", reading), nil)
		}
		val = strValue
	case common.ValueTypeBool:
		boolValue, ok := reading.(bool)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid bool value '%v'", reading), nil)
		}
		val = boolValue
	case common.ValueTypeUint8:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid numbers '%v'", reading), nil)
		}
		val = uint8(float64Value)
	case common.ValueTypeUint16:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid numbers '%v'", reading), nil)
		}
		val = uint16(float64Value)
	case common.ValueTypeUint32:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", reading), nil)
		}
		val = uint32(float64Value)
	case common.ValueTypeUint64:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", reading), nil)
		}
		val = uint64(float64Value)
	case common.ValueTypeInt8:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", reading), nil)
		}
		val = int8(float64Value)
	case common.ValueTypeInt16:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", reading), nil)
		}
		val = int16(float64Value)
	case common.ValueTypeInt32:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", reading), nil)
		}
		val = int32(float64Value)
	case common.ValueTypeInt64:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", reading), nil)
		}
		val = int64(float64Value)
	case common.ValueTypeFloat32:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", reading), nil)
		}
		val = float32(float64Value)
	case common.ValueTypeFloat64:
		float64Value, ok := reading.(float64)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", reading), nil)
		}
		val = float64Value
	case common.ValueTypeBinary:
		// XRT transfer binary data in base64 encoded string
		strValue, ok := reading.(string)
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid string value '%v'", reading), nil)
		}
		val, err = base64.StdEncoding.DecodeString(strValue)
		if err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("fail to decode the base64 string '%v'", reading), err)
		}
	case common.ValueTypeBoolArray:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		boolArray := make([]bool, len(interfaceArray))
		for i, v := range interfaceArray {
			boolValue, ok := v.(bool)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid bool value '%v'", v), nil)
			}
			boolArray[i] = boolValue
		}
		val = boolArray
	case common.ValueTypeStringArray:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		stringArray := make([]string, len(interfaceArray))
		for i, v := range interfaceArray {
			strValue, ok := v.(string)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid string value '%v'", v), nil)
			}
			stringArray[i] = strValue
		}
		val = stringArray
	case common.ValueTypeUint8Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		uint8Array := make([]uint8, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid nunmber '%v'", v), nil)
			}
			uint8Array[i] = uint8(float64Value)
		}
		val = uint8Array
	case common.ValueTypeUint16Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		uint16Array := make([]uint16, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid nunmber '%v'", v), nil)
			}
			uint16Array[i] = uint16(float64Value)
		}
		val = uint16Array
	case common.ValueTypeUint32Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		uint32Array := make([]uint32, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid nunmber '%v'", v), nil)
			}
			uint32Array[i] = uint32(float64Value)
		}
		val = uint32Array
	case common.ValueTypeUint64Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		uint64Array := make([]uint64, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid nunmber '%v'", v), nil)
			}
			uint64Array[i] = uint64(float64Value)
		}
		val = uint64Array
	case common.ValueTypeInt8Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		int8Array := make([]int8, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid nunmber '%v'", v), nil)
			}
			int8Array[i] = int8(float64Value)
		}
		val = int8Array
	case common.ValueTypeInt16Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		int16Array := make([]int16, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid nunmber '%v'", v), nil)
			}
			int16Array[i] = int16(float64Value)
		}
		val = int16Array
	case common.ValueTypeInt32Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		int32Array := make([]int32, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", v), nil)
			}
			int32Array[i] = int32(float64Value)
		}
		val = int32Array
	case common.ValueTypeInt64Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		int64Array := make([]int64, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", v), nil)
			}
			int64Array[i] = int64(float64Value)
		}
		val = int64Array
	case common.ValueTypeFloat32Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		float32Array := make([]float32, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", v), nil)
			}
			float32Array[i] = float32(float64Value)
		}
		val = float32Array
	case common.ValueTypeFloat64Array:
		interfaceArray, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		float64Array := make([]float64, len(interfaceArray))
		for i, v := range interfaceArray {
			float64Value, ok := v.(float64)
			if !ok {
				return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid number '%v'", v), nil)
			}
			float64Array[i] = float64Value
		}
		val = float64Array
	case common.ValueTypeObject:
		val = reading
	case common.ValueTypeObjectArray:
		_, ok := reading.([]interface{})
		if !ok {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid array '%v'", reading), nil)
		}
		val = reading
	default:
		return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("none supported value type '%s'", valueType), nil)
	}

	return val, nil
}
