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
	"fmt"
	"testing"

	"github.com/google/uuid"
)

var testRegistration = Registration{ID: uuid.New().String(), Name: "Test Registration", Compression: CompNone,
	Format: FormatJSON, Destination: DestZMQ, Addressable: TestAddressable}

func TestRegistrationValidation(t *testing.T) {
	valid := testRegistration

	invalidName := testRegistration
	invalidName.Name = ""

	invalidCompression := testRegistration
	invalidCompression.Compression = "blah"

	invalidFormat := testRegistration
	invalidFormat.Format = "blah"

	invalidDestination := testRegistration
	invalidDestination.Destination = "blah"

	invalidEncryption := testRegistration
	invalidEncryption.Encryption.Algo = "blah"

	invalidProcessFrequency := testRegistration
	invalidProcessFrequency.ProcessFrequency = "blah"

	invalidProcessFrequency1ms := testRegistration
	invalidProcessFrequency1ms.ProcessFrequency = "1ms"

	validProcessFrequency1s := testRegistration
	validProcessFrequency1s.ProcessFrequency = "1s"

	validProcessFrequency1m := testRegistration
	validProcessFrequency1m.ProcessFrequency = "1m"

	validProcessFrequency1h := testRegistration
	validProcessFrequency1h.ProcessFrequency = "1h"

	validProcessFrequency1h1m1s := testRegistration
	validProcessFrequency1h1m1s.ProcessFrequency = "1h1m1s"

	tests := []struct {
		name        string
		r           Registration
		expectError bool
	}{
		{"valid registration", valid, false},
		{"invalid registration name", invalidName, true},
		{"invalid registration compression", invalidCompression, true},
		{"invalid registration format", invalidFormat, true},
		{"invalid registration destination", invalidDestination, true},
		{"invalid registration encryption", invalidEncryption, true},
		{"invalid registration processFrequency format", invalidProcessFrequency, true},
		{"invalid registration processFrequency 1ms", invalidProcessFrequency1ms, true},
		{"valid registration processFrequency 1s", validProcessFrequency1s, false},
		{"valid registration processFrequency 1m", validProcessFrequency1m, false},
		{"valid registration processFrequency 1h", validProcessFrequency1h, false},
		{"valid registration processFrequency 1h1m1s", validProcessFrequency1h1m1s, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

func TestMaxRetentionEvents(t *testing.T) {
	r := &Registration{}
	r.UnmarshalJSON([]byte("{\"name\":\"n\",\"addressable\":{\"name\":\"n\"},\"format\":\"JSON\",\"destination\":\"MQTT_TOPIC\"}"))

	tests := []struct {
		name        string
		data        string
		expectValue uint
		expectError bool
	}{
		{"invalid MaxRetentionEvents: string", "{\"maxRetentionEvents\":\"s\"}",
			DefaultMaxRetentionEvents, true},
		{"invalid MaxRetentionEvents: negative number", "{\"maxRetentionEvents\":-1}",
			DefaultMaxRetentionEvents, true},
		{"invalid MaxRetentionEvents: float", "{\"maxRetentionEvents\":3.1415}",
			DefaultMaxRetentionEvents, true},
		{fmt.Sprintf("default MaxRetentionEvents is %d", DefaultMaxRetentionEvents),
			"{\"maxRetentionEvents\":0}", DefaultMaxRetentionEvents, false},
		{"normal case", "{\"maxRetentionEvents\":1000}", 1000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.UnmarshalJSON([]byte(tt.data))
			if err != nil {
				if !tt.expectError {
					t.Errorf("unexpected error: %v", err)
				}
			} else {
				if tt.expectError {
					t.Errorf("did not receive expected error: %s", tt.name)
				}
				if err == nil && r.MaxRetentionEvents != tt.expectValue {
					t.Errorf("expected value: %d, but got: %d", tt.expectValue, r.MaxRetentionEvents)
				}
			}
		})
	}
}

func TestMaxBatchEvents(t *testing.T) {
	r := &Registration{}
	r.UnmarshalJSON([]byte("{\"name\":\"n\",\"addressable\":{\"name\":\"n\"},\"format\":\"JSON\",\"destination\":\"MQTT_TOPIC\"}"))

	tests := []struct {
		name        string
		data        string
		expectError bool
	}{
		{"invalid MaxBatchEvents: string", "{\"maxRetentionEvents\":\"s\"}", true},
		{"invalid MaxBatchEvents: negative number", "{\"maxRetentionEvents\":-1}", true},
		{"invalid MaxBatchEvents: float", "{\"maxRetentionEvents\":3.1415}", true},
		{"normal case", "{\"maxBatchEvents\":100}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.UnmarshalJSON([]byte(tt.data))
			if err != nil {
				if !tt.expectError {
					t.Errorf("unexpected error: %v", err)
				}
			} else {
				if tt.expectError {
					t.Errorf("did not receive expected error: %s", tt.name)
				}
			}
		})
	}
}
