//
// Copyright (C) 2023-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMultiRulesResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedRules := []dtos.Rule{
		{Name: "rule1", Rule: []byte("rule1")},
		{Name: "rule2", Rule: []byte("rule2")},
	}
	expectedTotalCount := uint32(len(expectedRules))
	expectedMessage := "message"
	actual := NewMultiRulesResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedRules)
	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedRules, actual.Rules)
}
