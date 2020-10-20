/*******************************************************************************
 * Copyright 2019 Dell Technologies Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 *******************************************************************************/

package models

import (
	"encoding/json"
)

// Channel supports transmissions and notifications with fields for delivery via email, REST, ZeroMQ, or MQTT
type Channel struct {
	Type               ChannelType `json:"type,omitempty"`               // Type indicates whether the channel facilitates email, REST, ZeroMQ, or MQTT
	MailAddresses      []string    `json:"mailAddresses,omitempty"`      // MailAddresses contains email addresses
	Url                string      `json:"url,omitempty"`                // URL contains a REST API destination or message bus endpoint
	Port               int         `json:"port,omitempty"`               // Port specifies the port to which the channel is bound
	Topic              string      `json:"topic,omitempty"`              // Topic for filtering messages, usually be used by MQTT or ZeroMQ
	SecretSource       string      `json:"secretSource,omitempty"`       // SecretSource indicates where to get secrets
	SecretPath         string      `json:"secretPath,omitempty"`         // SecretPath specifies the path in the secret provider from which to retrieve secrets
	ClientKeyFilePath  string      `json:"clientKeyFilePath,omitempty"`  // ClientKeyFilePath specifies the path of the client private key file
	ClientCertFilePath string      `json:"clientCertFilePath,omitempty"` // ClientCertFilePath specifies the path of the client certificate file
	ServerCaFilePath   string      `json:"serverCaFilePath,omitempty"`   // ServerCaFilePath specifies the path of the server CA certificate file
	QoS                int         `json:"qos,omitempty"`                // QoS specifies the Quality of Service levels
}

func (c Channel) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
