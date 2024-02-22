// Copyright (C) 2024 IOTech Ltd

package models

import (
	"fmt"
)

func (d Device) StrValueFromProperties(key string) string {
	return strValueFromProperties(key, d.Properties)
}

func (d Device) StrSliceFromProperties(key string) []string {
	return strSliceFromProperties(key, d.Properties)
}

func (d DiscoveredDevice) StrValueFromProperties(key string) string {
	return strValueFromProperties(key, d.Properties)
}

func (d DiscoveredDevice) StrSliceFromProperties(key string) []string {
	return strSliceFromProperties(key, d.Properties)
}

func strValueFromProperties(key string, properties map[string]any) string {
	val := ""
	if propVal, ok := properties[key]; ok {
		val = fmt.Sprintf("%v", propVal)
	}
	return val
}

func strSliceFromProperties(key string, properties map[string]any) []string {
	if propVal, ok := properties[key]; ok {
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
