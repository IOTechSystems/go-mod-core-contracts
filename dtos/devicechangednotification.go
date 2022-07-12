// Copyright (C) 2022 IOTech Ltd

package dtos

import (
	"encoding/json"
)

type DeviceNotificationContent struct {
	DeviceName        string `json:"deviceName"`
	DeviceServiceName string `json:"deviceServiceName"`
	ActionType        string `json:"actionType"`
	OperatingState    string `json:"operatingState"`
	AdminState        string `json:"adminState"`
}

func NewDeviceNotificationContent(device Device, action string) DeviceNotificationContent {
	return DeviceNotificationContent{
		DeviceName:        device.Name,
		DeviceServiceName: device.ServiceName,
		ActionType:        action,
		OperatingState:    device.OperatingState,
		AdminState:        device.AdminState,
	}
}

func (d DeviceNotificationContent) String() (string, error) {
	if b, err := json.Marshal(d); err == nil {
		return string(b), nil
	} else {
		return "", err
	}
}
