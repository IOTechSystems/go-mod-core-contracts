//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package xlsx

// constants relates to the stylesheet names
const (
	devicesSheetName        = "Devices"
	mappingTableSheetName   = "MappingTable"
	autoEventsSheetName     = "AutoEvents"
	deviceInfoSheetName     = "DeviceInfo"
	deviceResourceSheetName = "DeviceResource"
	deviceCommandSheetName  = "DeviceCommand"
)

// constants relates to the header names
const (
	objectCol       = "object"
	pathCol         = "path"
	defaultValueCol = "default value"
)

// constants relates to the Device DTO field names
const (
	protocols    = "Protocols"
	protocolName = "ProtocolName"
	autoEvents   = "AutoEvents"
)

// constants relates to the DeviceResource/DeviceCommand DTO field names
const (
	attributes         = "Attributes"
	properties         = "Properties"
	resourceOperations = "ResourceOperations"
)

// constants relates to the Device protocol property keys
const (
	modbusRTUKey = "modbus-rtu"
	modbusTCPKey = "modbus-tcp"
)
