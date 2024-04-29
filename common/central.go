// Copyright (C) 2023-2024 IOTech Ltd

package common

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

// ParseValueByDeviceResource parses value for the specified value type
func ParseValueByDeviceResource(valueType string, value any) (any, errors.EdgeX) {
	var err error

	// Support writing the null value for specific protocol like BACnet.
	// For example, the user send a put request with JSON body {"test-resource": null}, then the device service will receive the nil value and marshal to null before sending to the XRT.
	if value == nil {
		return nil, nil
	}

	v := fmt.Sprint(value)

	if valueType != ValueTypeString && strings.TrimSpace(v) == "" {
		return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("empty string is invalid for %v value type", valueType), nil)
	}

	switch valueType {
	case ValueTypeString:
		return value, nil
	case ValueTypeStringArray:
		var arr []string
		err = json.Unmarshal([]byte(v), &arr)
		if err != nil {
			// try to handle the nonstandard json format, for example, [foo, bar]
			strArr := strings.Split(strings.Trim(v, "[]"), ",")
			for _, u := range strArr {
				arr = append(arr, strings.TrimSpace(u))
			}
			return arr, nil
		}
		return arr, nil
	case ValueTypeBool:
		boolVal, err := strconv.ParseBool(v)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return boolVal, nil
	case ValueTypeBoolArray:
		var arr []bool
		err = json.Unmarshal([]byte(v), &arr)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return arr, nil
	case ValueTypeUint8:
		var n uint64
		n, err = strconv.ParseUint(v, 10, 8)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return uint8(n), nil
	case ValueTypeUint8Array:
		var arr []uint8
		strArr := strings.Split(strings.Trim(v, "[]"), ",")
		for _, u := range strArr {
			n, err := strconv.ParseUint(strings.TrimSpace(u), 10, 8)
			if err != nil {
				errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
				return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
			}
			arr = append(arr, uint8(n))
		}
		return arr, nil
	case ValueTypeUint16:
		var n uint64
		n, err = strconv.ParseUint(v, 10, 16)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return uint16(n), nil
	case ValueTypeUint16Array:
		var arr []uint16
		strArr := strings.Split(strings.Trim(v, "[]"), ",")
		for _, u := range strArr {
			n, err := strconv.ParseUint(strings.TrimSpace(u), 10, 16)
			if err != nil {
				errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
				return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
			}
			arr = append(arr, uint16(n))
		}
		return arr, nil
	case ValueTypeUint32:
		var n uint64
		n, err = strconv.ParseUint(v, 10, 32)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return uint32(n), nil
	case ValueTypeUint32Array:
		var arr []uint32
		strArr := strings.Split(strings.Trim(v, "[]"), ",")
		for _, u := range strArr {
			n, err := strconv.ParseUint(strings.TrimSpace(u), 10, 32)
			if err != nil {
				errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
				return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
			}
			arr = append(arr, uint32(n))
		}
		return arr, nil
	case ValueTypeUint64:
		var n uint64
		n, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return n, nil
	case ValueTypeUint64Array:
		var arr []uint64
		strArr := strings.Split(strings.Trim(v, "[]"), ",")
		for _, u := range strArr {
			n, err := strconv.ParseUint(strings.TrimSpace(u), 10, 64)
			if err != nil {
				errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
				return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
			}
			arr = append(arr, n)
		}
		return arr, nil
	case ValueTypeInt8:
		var n int64
		n, err = strconv.ParseInt(v, 10, 8)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return int8(n), nil
	case ValueTypeInt8Array:
		var arr []int8
		err = json.Unmarshal([]byte(v), &arr)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return arr, nil
	case ValueTypeInt16:
		var n int64
		n, err = strconv.ParseInt(v, 10, 16)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return int16(n), nil
	case ValueTypeInt16Array:
		var arr []int16
		err = json.Unmarshal([]byte(v), &arr)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return arr, nil
	case ValueTypeInt32:
		var n int64
		n, err = strconv.ParseInt(v, 10, 32)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return int32(n), nil
	case ValueTypeInt32Array:
		var arr []int32
		err = json.Unmarshal([]byte(v), &arr)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return arr, nil
	case ValueTypeInt64:
		var n int64
		n, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return n, nil
	case ValueTypeInt64Array:
		var arr []int64
		err = json.Unmarshal([]byte(v), &arr)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return arr, nil
	case ValueTypeFloat32:
		var val float64
		val, err = strconv.ParseFloat(v, 32)
		if err == nil {
			return float32(val), nil
		}
		if numError, ok := err.(*strconv.NumError); ok {
			if numError.Err == strconv.ErrRange {
				return nil, errors.NewCommonEdgeX(errors.KindServerError, "NumError", err)
			}
		}
		var decodedToBytes []byte
		decodedToBytes, err = base64.StdEncoding.DecodeString(v)
		if err == nil {
			var val float32
			val, err = float32FromBytes(decodedToBytes)
			if err != nil {
				return nil, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("fail to parse %v to float32", v), err)
			} else if math.IsNaN(float64(val)) {
				return nil, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("fail to parse %v to float32, unexpected result %v", v, val), nil)
			} else {
				return val, nil
			}
		}
	case ValueTypeFloat32Array:
		var arr []float32
		err = json.Unmarshal([]byte(v), &arr)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return arr, nil
	case ValueTypeFloat64:
		var val float64
		val, err = strconv.ParseFloat(v, 64)
		if err == nil {
			return val, nil
		}
		if numError, ok := err.(*strconv.NumError); ok {
			if numError.Err == strconv.ErrRange {
				return nil, errors.NewCommonEdgeX(errors.KindServerError, "NumError", err)
			}
		}
		var decodedToBytes []byte
		decodedToBytes, err = base64.StdEncoding.DecodeString(v)
		if err == nil {
			val, err = float64FromBytes(decodedToBytes)
			if err != nil {
				return nil, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("fail to parse %v to float64", v), err)
			} else if math.IsNaN(val) {
				return nil, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("fail to parse %v to float64, unexpected result %v", v, val), nil)
			} else {
				return val, nil
			}
		}
	case ValueTypeFloat64Array:
		var arr []float64
		err = json.Unmarshal([]byte(v), &arr)
		if err != nil {
			errMsg := fmt.Sprintf("failed to convert set parameter %s to ValueType %s", v, valueType)
			return nil, errors.NewCommonEdgeX(errors.KindServerError, errMsg, err)
		}
		return arr, nil
	case ValueTypeObject:
		return value, nil
	case ValueTypeObjectArray:
		return value, nil
	default:
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "unrecognized value type", nil)
	}
	return value, nil
}

func float32FromBytes(numericValue []byte) (res float32, err error) {
	reader := bytes.NewReader(numericValue)
	err = binary.Read(reader, binary.BigEndian, &res)
	return
}

func float64FromBytes(numericValue []byte) (res float64, err error) {
	reader := bytes.NewReader(numericValue)
	err = binary.Read(reader, binary.BigEndian, &res)
	return
}
