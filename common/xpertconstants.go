// Copyright (C) 2022-2023 IOTech Ltd

package common

// Service name
const (
	SupportRulesEngineServiceKey        = "support-rulesengine"
	CoreKeeperServiceKey                = "core-keeper"
	SupportProvisionServiceKey          = "support-provision"
	SupportSparkplugServiceKey          = "support-sparkplug"
	SupportSparkplugHistorianServiceKey = "support-sparkplug-historian"
)

// Content types supported by the APIs
const (
	ContextKeyContentType contextKey = ContentType
	ContentTypeForm                  = "application/x-www-form-urlencoded"
)

// API route
const (
	ApiNotificationByIdsRoute              = ApiNotificationRoute + "/" + Ids + "/{" + Ids + "}"
	ApiNotificationAcknowledgeByIdsRoute   = ApiNotificationRoute + "/" + Acknowledge + "/" + Ids + "/{" + Ids + "}"
	ApiNotificationUnacknowledgeByIdsRoute = ApiNotificationRoute + "/" + Unacknowledge + "/" + Ids + "/{" + Ids + "}"

	ApiRuleRoute       = ApiBase + "/rule"
	ApiAllRulesRoute   = ApiRuleRoute + "/" + All
	ApiRuleByNameRoute = ApiRuleRoute + "/" + Name + "/{" + Name + "}"

	ApiKVSRoute      = ApiBase + "/kvs"
	ApiKVSByKeyRoute = ApiKVSRoute + "/" + Key + "/{" + Key + "}"
)

// API route path parameter
const (
	Ids           = "ids"
	User          = "user"
	Group         = "group"
	PublicKey     = "rsa_public_key"
	Ack           = "ack"
	Acknowledge   = "acknowledge"
	Unacknowledge = "unacknowledge"
	Key           = "key"
	Flatten       = "flatten" //query string to specify if the request json payload should be flattened to update multiple keys with the same prefix
	KeyOnly       = "keyOnly" //query string to specify if the response will only return the keys of the specified query key prefix, without values and metadata
	Plaintext     = "plaintext"
)

// Constants for Address
const (
	ZeroMQ = "ZeroMQ"
	HTTP   = "http"
	TCP    = "tcp"
	TCPS   = "tcps"
)

// Value type
const (
	ValueTypeObjectArray = "ObjectArray"
)

// Protocol properties
const (
	BacnetIP             = "BACnet-IP"
	BacnetMSTP           = "BACnet-MSTP"
	BacnetDeviceInstance = "DeviceInstance"
	BacnetAddress        = "Address"
	BacnetPort           = "Port"

	Gps                   = "GPS"
	GpsGpsdPort           = "GpsdPort"
	GpsGpsdRetries        = "GpsdRetries"
	GpsGpsdConnTimeout    = "GpsdConnTimeout"
	GpsGpsdRequestTimeout = "GpsdRequestTimeout"

	ModbusTcp                       = "modbus-tcp"
	ModbusRtu                       = "modbus-rtu"
	ModbusUnitID                    = "UnitID"
	ModbusPort                      = "Port"
	ModbusBaudRate                  = "BaudRate"
	ModbusDataBits                  = "DataBits"
	ModbusStopBits                  = "StopBits"
	ModbusReadMaxHoldingRegisters   = "ReadMaxHoldingRegisters"
	ModbusReadMaxInputRegisters     = "ReadMaxInputRegisters"
	ModbusReadMaxBitsCoils          = "ReadMaxBitsCoils"
	ModbusReadMaxBitsDiscreteInputs = "ReadMaxBitsDiscreteInputs"
	ModbusWriteMaxHoldingRegisters  = "WriteMaxHoldingRegisters"
	ModbusWriteMaxBitsCoils         = "WriteMaxBitsCoils"

	Opcua                           = "OPC-UA"
	OpcuaRequestedSessionTimeout    = "RequestedSessionTimeout"
	OpcuaBrowseDepth                = "BrowseDepth"
	OpcuaBrowsePublishInterval      = "BrowsePublishInterval"
	OpcuaConnectionReadingPostDelay = "ConnectionReadingPostDelay"
	OpcuaIDType                     = "IDType"
	OpcuaReadBatchSize              = "ReadBatchSize"
	OpcuaWriteBatchSize             = "WriteBatchSize"
	OpcuaNodesPerBrowse             = "NodesPerBrowse"
	OpcuaSessionKeepAliveInterval   = "SessionKeepAliveInterval"

	S7     = "S7"
	S7Rack = "Rack"
	S7Slot = "Slot"

	EtherNetIP                  = "ethernet-ip"
	EtherNetIPXRT               = "EtherNet-IP" // XRT only accept EtherNet-IP as protocol name
	EtherNetIPO2T               = "O2T"
	EtherNetIPT2O               = "T2O"
	EtherNetIPExplicitConnected = "ExplicitConnected"
	EtherNetIPDeviceResource    = "DeviceResource"
	EtherNetIPSaveValue         = "SaveValue"
	EtherNetIPConnectionType    = "ConnectionType"
	EtherNetIPRPI               = "RPI"
	EtherNetIPPriority          = "Priority"
	EtherNetIPOwnership         = "Ownership"
	EtherNetIPKey               = "Key"
	EtherNetIPMethod            = "Method"
	EtherNetIPVendorID          = "VendorID"
	EtherNetIPDeviceType        = "DeviceType"
	EtherNetIPProductCode       = "ProductCode"
	EtherNetIPMajorRevision     = "MajorRevision"
	EtherNetIPMinorRevision     = "MinorRevision"
	EtherNetIPAddress           = "Address"
)

// Constants for Notification Category
const (
	DisconnectAlert      = "Disconnection"
	DeviceOperatingState = "DeviceOperatingState"
)

// Constants for DeviceChangedNotification
const (
	DeviceCreateAction = "Device creation"
	DeviceUpdateAction = "Device update"
	DeviceRemoveAction = "Device removal"

	DeviceChangedNotificationCategory = "DEVICE_CHANGED"
)
