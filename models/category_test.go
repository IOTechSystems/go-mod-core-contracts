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

import "testing"

func TestNotificationsCategory_UnmarshalJSON(t *testing.T) {
	var swHealth = NotificationsCategory(Swhealth)
	var hwHealth = NotificationsCategory(Hwhealth)
	var security = NotificationsCategory(Security)
	var devicechanged = NotificationsCategory(DeviceChanged)

	tests := []struct {
		name    string
		as      *NotificationsCategory
		args    []byte
		wantErr bool
	}{
		{"Test marshal of sw health", &swHealth, []byte("\"SW_HEALTH\""), false},
		{"Test marshal of hw health", &hwHealth, []byte("\"HW_HEALTH\""), false},
		{"Test marshal of security", &security, []byte("\"SECURITY\""), false},
		{"Test marshal of devicechanged", &devicechanged, []byte("\"DEVICE_CHANGED\""), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.as.UnmarshalJSON(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("NotificationsCategory.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsNotificationsCategory(t *testing.T) {
	type args struct {
		as string
	}
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"test SW HEALTH", Swhealth, true},
		{"test HW HEALTH", Hwhealth, true},
		{"test SECURITY", Security, true},
		{"test DEVICE_CHANGED", DeviceChanged, true},
		{"test fail on non-notif cat", "foo", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotificationsCategory(tt.arg); got != tt.want {
				t.Errorf("IsNotificationsCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}
