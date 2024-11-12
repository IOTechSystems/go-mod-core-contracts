//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTokenResponse(t *testing.T) {
	expectedToken := "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzEzODQxNDksImlzcyI6IklPVGVjaCIsImxjciI6MTczMTM4NTMxOSwidG9rZW5faWQiOiI3YWI4YTQwMS04ODEyLTRkMTgtOTgyMS05NTA3OGRiOGI2MDcifQ.arKj3pfSXRm2wH5chVaSMTBUA-cgSu_0CW2AbvXVEsiSIbB_KOt9p3pt2V1WWml2Tzvk7m_tLo-W_1HJVhuiCA" // nolint:gosec
	actual := NewTokenResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedToken)

	require.Equal(t, expectedRequestId, actual.RequestId)
	require.Equal(t, expectedStatusCode, actual.StatusCode)
	require.Equal(t, expectedMessage, actual.Message)
	require.Equal(t, expectedToken, actual.JWT)
}
