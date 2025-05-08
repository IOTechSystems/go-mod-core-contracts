// Copyright (C) 2025 IOTech Ltd

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvisionWatcher_Clone(t *testing.T) {
	testProvisionWatcher := ProvisionWatcher{
		DBTimestamp:         DBTimestamp{},
		Id:                  "ca93c8fa-9919-4ec5-85d3-f81b2b6a7bc1",
		Name:                "TestProvisionWatcher",
		ServiceName:         "TestServiceName",
		Labels:              []string{"label1", "label2"},
		Identifiers:         map[string]string{"Address": "172.0.0.1", "Port": "8080"},
		BlockingIdentifiers: map[string][]string{"Address": {"127.0.0.1", "127.0.0.2"}, "Port": {"123", "456"}},
		AdminState:          Unlocked,
		DiscoveredDevice: DiscoveredDevice{
			ProfileName: "TestProfile",
			AdminState:  Locked,
			AutoEvents: []AutoEvent{
				{
					Interval: "10s", OnChange: false,
					SourceName: "TestDeviceResource",
				},
				{
					Interval: "15s", OnChange: true,
					SourceName: "TestDeviceResource2",
				},
			},
			Properties: map[string]any{
				"foo": "bar",
			},
		},
	}
	clone := testProvisionWatcher.Clone()
	assert.Equal(t, testProvisionWatcher, clone)
}
