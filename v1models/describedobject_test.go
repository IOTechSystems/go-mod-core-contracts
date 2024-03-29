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
	"strconv"
	"testing"
)

var TestDescribedObject = DescribedObject{Timestamps: testTimestamps, Description: testDescription}
var TestEmptyDescribedObject = DescribedObject{}

func TestDescribedObject_String(t *testing.T) {
	tests := []struct {
		name string
		o    DescribedObject
		want string
	}{
		{"described object to string", TestDescribedObject,
			"{\"created\":" + strconv.FormatInt(TestDescribedObject.Created, 10) +
				",\"modified\":" + strconv.FormatInt(TestDescribedObject.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestDescribedObject.Origin, 10) +
				",\"description\":\"" + testDescription + "\"}"},
		{"empty described object to string", TestEmptyDescribedObject, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("DescribedObject.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
