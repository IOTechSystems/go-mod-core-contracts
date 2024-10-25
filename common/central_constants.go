// Copyright (C) 2022-2024 IOTech Ltd

package common

const (
	IOTechPrefix = "IOTech_"

	ContextKeyContentType contextKey = ContentType
	ContentTypeCSV                   = "text/csv"

	BacnetIP             = "BACnet-IP"
	BacnetMSTP           = "BACnet-MSTP"
	BacnetDeviceInstance = "DeviceInstance"
	BacnetAddress        = "Address"
	BacnetPort           = "Port"
	BacnetCOVPropName    = "BACnet-COVs"
	BacnetCOV            = "COV"
	BacnetCOVConfirmed   = "Confirmed"
	BacnetCOVLifetime    = "Lifetime"

	Gps                   = "GPS"
	GpsGpsdPort           = "GpsdPort"
	GpsGpsdRetries        = "GpsdRetries"
	GpsGpsdConnTimeout    = "GpsdConnTimeout"
	GpsGpsdRequestTimeout = "GpsdRequestTimeout"

	ModbusTcp                       = "modbus-tcp"
	ModbusRtu                       = "modbus-rtu"
	ModbusUnitID                    = "UnitID"
	ModbusAddress                   = "Address"
	ModbusPort                      = "Port"
	ModbusBaudRate                  = "BaudRate"
	ModbusDataBits                  = "DataBits"
	ModbusParity                    = "Parity"
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

// constants relate to the remote edge node
const (
	BrokerName                            = "brokerName"
	EdgeInstName                          = "edgeinstName"
	DeviceServiceName                     = "deviceServiceName"
	TopicPatternFieldGroupName            = "GROUP_NAME"
	TopicPatternFieldInstName             = "INSTANCE_NAME"
	TopicPatternFieldCentralGroupName     = "EDGECENTRAL_GROUP_NAME"
	TopicPatternFieldCentralInstName      = "EDGECENTRAL_INSTANCE_NAME"
	TopicPatternFieldDeviceServiceName    = "DEVICE_SERVICE_NAME"
	TopicPatternFieldKeyCentralGroupName  = "${" + TopicPatternFieldCentralGroupName + "}"
	TopicPatternFieldKeyCentralInstName   = "${" + TopicPatternFieldCentralInstName + "}"
	TopicPatternFieldKeyDeviceServiceName = "${" + TopicPatternFieldDeviceServiceName + "}"
	CentralNodeRequestTopicKey            = "RequestTopic"
	CentralNodeReplyTopicKey              = "ReplyTopic"
)

// Constants relate to the service status error from sys-mgmt inspect operation
const (
	ServiceIsNotRunningButShouldBe = "service is not running but should be"
	ServiceIsRunningButShouldNotBe = "service is running but shouldn't be"
)

// Constants related to how services identify themselves in the Service Registry
const (
	SupportProvisionServiceKey          = "support-provision"
	SupportSparkplugServiceKey          = "support-sparkplug"
	SupportSparkplugHistorianServiceKey = "support-sparkplug-historian"
	SupportRulesEngineServiceKey        = "support-rulesengine"
)

// Constants related for Notification Category
const (
	DisconnectAlert      = "Disconnection"
	DeviceOperatingState = "DeviceOperatingState"
)

// Constants related for DeviceChangedNotification
const (
	DeviceCreateAction = "Device creation"
	DeviceUpdateAction = "Device update"
	DeviceRemoveAction = "Device removal"

	DeviceChangedNotificationCategory = "DEVICE_CHANGED"
)

// Constants related to System Events
const (
	DeviceSystemEventActionAdd    = "add"
	DeviceSystemEventActionUpdate = "update"
	DeviceSystemEventActionDelete = "delete"
)

// Constants related for Address
const (
	ZeroMQ = "ZeroMQ"
	HTTP   = "http"
	TCP    = "tcp"
	TCPS   = "tcps"
)

// Constants related for provisionWatcher discoveredDevice
const (
	ProtocolName       = IOTechPrefix + "ProtocolName"
	DeviceNamePattern  = IOTechPrefix + "DeviceNamePattern"
	DeviceDescription  = IOTechPrefix + "DeviceDescription"
	DeviceLabels       = IOTechPrefix + "DeviceLabels"
	ProfileNamePattern = IOTechPrefix + "ProfileNamePattern"
	ProfileDescription = IOTechPrefix + "ProfileDescription"
	ProfileLabels      = IOTechPrefix + "ProfileLabels"
	ProfileScanOptions = IOTechPrefix + "ProfileScanOptions"
)
