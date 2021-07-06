/*******************************************************************************
 * Copyright 2019 Dell Inc.
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
 *******************************************************************************/

package models

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/fxamacker/cbor/v2"
)

// Event represents a single measurable event read from a device
type Event struct {
	ID          string                 `json:"id,omitempty" codec:"id,omitempty"`             // ID uniquely identifies an event, for example a UUID
	Pushed      int64                  `json:"pushed,omitempty" codec:"pushed,omitempty"`     // Pushed is a timestamp indicating when the event was exported. If unexported, the value is zero.
	Device      string                 `json:"device,omitempty" codec:"device,omitempty"`     // Device identifies the source of the event, can be a device name or id. Usually the device name.
	Created     int64                  `json:"created,omitempty" codec:"created,omitempty"`   // Created is a timestamp indicating when the event was created.
	Modified    int64                  `json:"modified,omitempty" codec:"modified,omitempty"` // Modified is a timestamp indicating when the event was last modified.
	Origin      int64                  `json:"origin,omitempty" codec:"origin,omitempty"`     // Origin is a timestamp that can communicate the time of the original reading, prior to event creation
	Readings    []Reading              `json:"readings,omitempty" codec:"readings,omitempty"` // Readings will contain zero to many entries for the associated readings of a given event.
	Tags        map[string]interface{} `json:"tags,omitempty" codec:"tags,omitempty" xml:"-"` // Tags is an optional collection of key/value pairs that all the event to be tagged with custom information. Ignored for XML since maps not supported.
	isValidated bool                   // internal member used for validation check
	CommandName string                 `json:"commandName,omitempty" codec:"commandName,omitempty"`
}

func encodeAsCBOR(e Event) ([]byte, error) {
	bytes, err := cbor.Marshal(e)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

// UnmarshalJSON implements the Unmarshaler interface for the Event type
func (e *Event) UnmarshalJSON(data []byte) error {
	var err error
	type Alias struct {
		ID          *string                `json:"id"`
		Pushed      int64                  `json:"pushed"`
		Device      *string                `json:"device"`
		Created     int64                  `json:"created"`
		Modified    int64                  `json:"modified"`
		Origin      int64                  `json:"origin"`
		Readings    []Reading              `json:"readings"`
		Tags        map[string]interface{} `json:"tags"`
		CommandName *string                `json:"commandName"`
	}
	a := Alias{}

	// Error with unmarshaling
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}

	// Set the fields
	if a.ID != nil {
		e.ID = *a.ID
	}
	if a.Device != nil {
		e.Device = *a.Device
	}
	e.Pushed = a.Pushed
	e.Created = a.Created
	e.Modified = a.Modified
	e.Origin = a.Origin
	e.Readings = a.Readings
	if a.CommandName != nil {
		e.CommandName = *a.CommandName
	}
	e.Tags = a.Tags

	e.isValidated, err = e.Validate()
	return err
}

// Validate satisfies the Validator interface
func (e Event) Validate() (bool, error) {
	if !e.isValidated {
		if e.Device == "" {
			return false, NewErrContractInvalid("source device for event not specified")
		}
	}
	return true, nil
}

// String provides a JSON representation of the Event as a string
func (e Event) String() string {
	out, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}

	return string(out)
}

// CBOR provides a byte array CBOR-encoded representation of the Event
func (e Event) CBOR() []byte {
	cbor, err := encodeAsCBOR(e)
	if err != nil {
		return []byte{}
	}

	return cbor
}

// ToXML provides a XML representation of the Event as a string
func (e Event) ToXML() (string, error) {
	eventXml, err := xml.Marshal(e)
	if err != nil {
		return "", err
	}

	// The Tags field is being ignore from XML Marshaling since maps are not supported.
	// We have to provide our own marshaling of the Tags field if it is non-empty
	if len(e.Tags) > 0 {
		tagsXmlElements := []string{"<Tags>"}
		for key, value := range e.Tags {
			tag := fmt.Sprintf("<%s>%s</%s>", key, value, key)
			tagsXmlElements = append(tagsXmlElements, tag)
		}
		tagsXmlElements = append(tagsXmlElements, "</Tags>")
		tagsXml := strings.Join(tagsXmlElements, "")
		eventXml = []byte(strings.Replace(string(eventXml), "</Event>", tagsXml+"</Event>", 1))
	}

	return string(eventXml), nil
}
