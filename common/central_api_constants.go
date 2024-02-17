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
	Key           = "key"
	ServiceId     = "serviceId"
	Flatten       = "flatten"      //query string to specify if the request json payload should be flattened to update multiple keys with the same prefix
	KeyOnly       = "keyOnly"      //query string to specify if the response will only return the keys of the specified query key prefix, without values and metadata
	Plaintext     = "plaintext"    //query string to specify if the response will return the stored plain text value of the key(s) without any encoding
	Deregistered  = "deregistered" //query string to specify if the response will return the registries of deregistered services
	NoCallback    = "nocallback"   //query string to ask core-metadata not to invoke DS callback
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

	ApiKVSRoute                     = ApiBase + "/kvs"
	ApiKVSByKeyRoute                = ApiKVSRoute + "/" + Key + "/:" + Key
	ApiRegisterRoute                = ApiBase + "/registry"
	ApiAllRegistrationsRoute        = ApiRegisterRoute + "/" + All
	ApiRegistrationByServiceIdRoute = ApiRegisterRoute + "/" + ServiceId + "/:" + ServiceId

	ApiCoreCommandsByDeviceNameRoute = ApiBase + "/command/device" + "/" + Name + "/:" + Name
)
