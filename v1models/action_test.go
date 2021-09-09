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

import "testing"

var TestExpectedvalues = []string{testExpectedvalue1, testExpectedvalue2}
var TestAction = Action{testActionPath, []Response{{testCode, testDescription, TestExpectedvalues}}, ""}
var EmptyAction = Action{}

func TestAction_String(t *testing.T) {
	tests := []struct {
		name   string
		action Action
		want   string
	}{
		{"full action", TestAction, "{\"path\":\"" + testActionPath + "\",\"responses\":[{\"code\":\"" + testCode + "\",\"description\":\"" + testDescription + "\",\"expectedValues\":[\"" + testExpectedvalue1 + "\",\"" + testExpectedvalue2 + "\"]}]}"},
		{"empty action", EmptyAction, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.action.String(); got != tt.want {
				t.Errorf("Action.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
