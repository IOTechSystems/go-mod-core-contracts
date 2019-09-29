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
	"github.com/google/uuid"
	"testing"
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
