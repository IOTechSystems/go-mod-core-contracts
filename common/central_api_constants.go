// Copyright (C) 2024 IOTech Ltd

package common

const (
	ContentTypeForm = "application/x-www-form-urlencoded"

	Ids           = "ids"
	User          = "user"
	Group         = "group"
	PublicKey     = "rsa_public_key"
	Ack           = "ack"
	Acknowledge   = "acknowledge"
	Unacknowledge = "unacknowledge"
	NoCallback    = "nocallback" //query string to ask core-metadata not to invoke DS callback
)

const (
	ApiMetricsRoute      = ApiBase + "/metrics"
	ApiMultiMetricsRoute = ApiSystemRoute + "/metrics"

	ApiNotificationByIdsRoute              = ApiNotificationRoute + "/" + Ids + "/:" + Ids
	ApiNotificationAcknowledgeByIdsRoute   = ApiNotificationRoute + "/" + Acknowledge + "/" + Ids + "/:" + Ids
	ApiNotificationUnacknowledgeByIdsRoute = ApiNotificationRoute + "/" + Unacknowledge + "/" + Ids + "/:" + Ids

	ApiRuleRoute       = ApiBase + "/rule"
	ApiAllRulesRoute   = ApiRuleRoute + "/" + All
	ApiRuleByNameRoute = ApiRuleRoute + "/" + Name + "/:" + Name

	ApiCoreCommandsByDeviceNameRoute = ApiBase + "/command/device" + "/" + Name + "/:" + Name
)
