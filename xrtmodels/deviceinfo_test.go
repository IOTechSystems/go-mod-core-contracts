// Copyright (C) 2022 IOTech Ltd

package xrtmodels

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"

	"github.com/stretchr/testify/assert"
)

func TestProcessEtherNetIP(t *testing.T) {
	tests := []struct {
		name     string
		protocol map[string]map[string]interface{}
		expected map[string]map[string]interface{}
	}{
		{
			name: "process O2T and T2O properties",
			protocol: map[string]map[string]interface{}{
				common.EtherNetIP: {
					common.EtherNetIPAddress: "127.0.0.1",
				},
				common.EtherNetIPO2T: {
					common.EtherNetIPConnectionType: "p2p",
					common.EtherNetIPRPI:            10,
					common.EtherNetIPPriority:       "low",
					common.EtherNetIPOwnership:      "exclusive",
				},
				common.EtherNetIPT2O: {
					common.EtherNetIPConnectionType: "p2p",
					common.EtherNetIPRPI:            10,
					common.EtherNetIPPriority:       "low",
					common.EtherNetIPOwnership:      "exclusive",
				},
			},
			expected: map[string]map[string]interface{}{
				common.EtherNetIPXRT: {
					common.EtherNetIPAddress: "127.0.0.1",
					common.EtherNetIPO2T: map[string]interface{}{
						common.EtherNetIPConnectionType: "p2p",
						common.EtherNetIPRPI:            10,
						common.EtherNetIPPriority:       "low",
						common.EtherNetIPOwnership:      "exclusive",
					},
					common.EtherNetIPT2O: map[string]interface{}{
						common.EtherNetIPConnectionType: "p2p",
						common.EtherNetIPRPI:            10,
						common.EtherNetIPPriority:       "low",
						common.EtherNetIPOwnership:      "exclusive",
					},
				},
			},
		},
		{
			name: "process ExplicitConnected and Key properties",
			protocol: map[string]map[string]interface{}{
				common.EtherNetIP: {
					common.EtherNetIPAddress: "127.0.0.1",
				},
				common.EtherNetIPExplicitConnected: {
					common.EtherNetIPDeviceResource: "VendorID",
					common.EtherNetIPRPI:            10,
					common.EtherNetIPSaveValue:      true,
				},
				common.EtherNetIPKey: {
					common.EtherNetIPMethod:        "exact",
					common.EtherNetIPVendorID:      10,
					common.EtherNetIPDeviceType:    72,
					common.EtherNetIPProductCode:   50,
					common.EtherNetIPMajorRevision: 12,
					common.EtherNetIPMinorRevision: 2,
				},
			},
			expected: map[string]map[string]interface{}{
				common.EtherNetIPXRT: {
					common.EtherNetIPAddress: "127.0.0.1",
					common.EtherNetIPExplicitConnected: map[string]interface{}{
						common.EtherNetIPDeviceResource: "VendorID",
						common.EtherNetIPRPI:            10,
						common.EtherNetIPSaveValue:      true,
					},
					common.EtherNetIPKey: map[string]interface{}{
						common.EtherNetIPMethod:        "exact",
						common.EtherNetIPVendorID:      10,
						common.EtherNetIPDeviceType:    72,
						common.EtherNetIPProductCode:   50,
						common.EtherNetIPMajorRevision: 12,
						common.EtherNetIPMinorRevision: 2,
					},
				},
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			processEtherNetIP(testCase.protocol)
			assert.EqualValues(t, testCase.expected, testCase.protocol)
		})
	}
}
