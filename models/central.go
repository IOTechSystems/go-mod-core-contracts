// Copyright (C) 2024 IOTech Ltd

package models

import (
	"fmt"
)

func (d DiscoveredDevice) StrValueFromProperties(key string) string {
	val := ""
	if propVal, ok := d.Properties[key]; ok {
		val = fmt.Sprintf("%v", propVal)
	}
	return val
}

func (d DiscoveredDevice) StrSliceFromProperties(key string) []string {
	if propVal, ok := d.Properties[key]; ok {
		switch slice := propVal.(type) {
		case []string:
			return slice
		case []any:
			strSlice := make([]string, len(slice))
			for i, v := range slice {
				strSlice[i] = fmt.Sprintf("%v", v)
			}
			return strSlice
		}
	}
	return []string{}
}
