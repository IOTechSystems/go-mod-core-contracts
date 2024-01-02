// Copyright (C) 2022-2023 IOTech Ltd

package common

const (
	ContextKeyContentType contextKey = ContentType
	ContentTypeCSV                   = "text/csv"

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

// constants relate to the remote edge node
const (
	BrokerName                            = "brokerName"
	EdgeNodeName                          = "edgenodeName"
	DeviceServiceName                     = "deviceServiceName"
	TopicPatternFieldGroupName            = "GROUP_NAME"
	TopicPatternFieldNodeName             = "NODE_NAME"
	TopicPatternFieldCentralGroupName     = "EDGECENTRAL_GROUP_NAME"
	TopicPatternFieldCentralNodeName      = "EDGECENTRAL_NODE_NAME"
	TopicPatternFieldDeviceServiceName    = "DEVICE_SERVICE_NAME"
	TopicPatternFieldKeyCentralGroupName  = "${" + TopicPatternFieldCentralGroupName + "}"
	TopicPatternFieldKeyCentralNodeName   = "${" + TopicPatternFieldCentralNodeName + "}"
	TopicPatternFieldKeyDeviceServiceName = "${" + TopicPatternFieldDeviceServiceName + "}"
	CentralNodeRequestTopicKey            = "RequestTopic"
	CentralNodeReplyTopicKey              = "ReplyTopic"
)

// constants relate to the service status error from sys-mgmt inspect operation
const (
	ServiceIsNotRunningButShouldBe = "service is not running but should be"
	ServiceIsRunningButShouldNotBe = "service is running but shouldn't be"
)
