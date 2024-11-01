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
	testUsername  = "bob"
	testUsername2 = "alice"
)

func TestNewUserResponse(t *testing.T) {
	expectedUser := dtos.User{Name: testUsername}
	actual := NewUserResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedUser)

	require.Equal(t, expectedRequestId, actual.RequestId)
	require.Equal(t, expectedStatusCode, actual.StatusCode)
	require.Equal(t, expectedMessage, actual.Message)
	require.Equal(t, expectedUser, actual.User)
}

func TestNewMultiUsersResponse(t *testing.T) {
	expectedUsers := []dtos.User{
		{Name: testUsername},
		{Name: testUsername2},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiUsersResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedUsers)

	require.Equal(t, expectedRequestId, actual.RequestId)
	require.Equal(t, expectedStatusCode, actual.StatusCode)
	require.Equal(t, expectedMessage, actual.Message)
	require.Equal(t, expectedTotalCount, actual.TotalCount)
	require.Equal(t, expectedUsers, actual.Users)
}
