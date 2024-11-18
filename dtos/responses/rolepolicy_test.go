//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"

	"github.com/stretchr/testify/require"
)

var (
	mockRole1 = "device-admin"
	mockRole2 = "cmd-admin"
)

func TestNewRolePolicyResponse(t *testing.T) {
	expectedRolePolicy := dtos.RolePolicy{Role: mockRole1}
	actual := NewRolePolicyResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedRolePolicy)
	require.Equal(t, expectedRequestId, actual.RequestId)
	require.Equal(t, expectedStatusCode, actual.StatusCode)
	require.Equal(t, expectedMessage, actual.Message)
	require.Equal(t, expectedRolePolicy, actual.RolePolicy)
}

func TestNewMultiRolePolicyResponse(t *testing.T) {
	expectedRolePolicies := []dtos.RolePolicy{
		{Role: mockRole1},
		{Role: mockRole2},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiRolePolicyResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedRolePolicies)

	require.Equal(t, expectedRequestId, actual.RequestId)
	require.Equal(t, expectedStatusCode, actual.StatusCode)
	require.Equal(t, expectedMessage, actual.Message)
	require.Equal(t, expectedTotalCount, actual.TotalCount)
	require.Equal(t, expectedRolePolicies, actual.RolePolicies)
}
