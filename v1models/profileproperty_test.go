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

package v1models

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
)

var TestMediaType = common.ContentTypeCBOR
var TestProfileProperty = ProfileProperty{Value: TestPropertyValue, Units: TestUnits}
var TestProfilePropertyEmpty = ProfileProperty{}

func TestProfileProperty_String(t *testing.T) {
	tests := []struct {
		name string
		pp   ProfileProperty
		want string
	}{
		{"profile property to string", TestProfileProperty,
			"{\"value\":" + TestPropertyValue.String() +
				",\"units\":" + TestUnits.String() +
				"}"},
		{"profile property to string", TestProfilePropertyEmpty, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pp.String(); got != tt.want {
				t.Errorf("ProfileProperty.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
