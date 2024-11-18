//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/require"
)

var (
	mockRole        = "mockAdmin"
	mockDescription = "A test admin role"
	mockPath        = "*/core-command/*"
	mockMethods     = []string{"GET", "PUT"}
	mockEffect      = "allow"

	mockAccessPolicyDTO = dtos.AccessPolicy{
		Path:        mockPath,
		HttpMethods: mockMethods,
		Effect:      mockEffect,
	}
	mockAccessPolicyModel = models.AccessPolicy{
		Path:        mockPath,
		HttpMethods: mockMethods,
		Effect:      mockEffect,
	}
	mockRolePolicyDTO = dtos.RolePolicy{
		Role:           mockRole,
		Description:    mockDescription,
		AccessPolicies: []dtos.AccessPolicy{mockAccessPolicyDTO},
	}
	testAddRolePolicyRequest = AddRolePolicyRequest{
		BaseRequest: dtoCommon.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: dtoCommon.NewVersionable(),
		},
		RolePolicy: mockRolePolicyDTO,
	}
)

func TestAddRolePolicyRequest_Validate(t *testing.T) {
	valid := testAddRolePolicyRequest

	emptyRole := valid
	emptyRole.RolePolicy.Role = ""

	emptyAccessPolicy := valid
	emptyAccessPolicy.RolePolicy.AccessPolicies = nil

	emptyPathAccessPolicy := mockAccessPolicyDTO
	emptyPathAccessPolicy.Path = ""
	invalidAccessPolicy := valid
	invalidAccessPolicy.RolePolicy.AccessPolicies = []dtos.AccessPolicy{emptyPathAccessPolicy}

	emptyMethodAccessPolicy := mockAccessPolicyDTO
	emptyMethodAccessPolicy.HttpMethods = nil
	emptyMethodPolicy := valid
	emptyMethodPolicy.RolePolicy.AccessPolicies = []dtos.AccessPolicy{emptyMethodAccessPolicy}

	invalidMethodAccessPolicy := mockAccessPolicyDTO
	invalidMethodAccessPolicy.HttpMethods = []string{"invalid"}
	invalidMethodPolicy := valid
	invalidMethodPolicy.RolePolicy.AccessPolicies = []dtos.AccessPolicy{invalidMethodAccessPolicy}
	tests := []struct {
		name        string
		RolePolicy  AddRolePolicyRequest
		expectError bool
	}{
		{"valid AddRolePolicyRequest", valid, false},
		{"invalid AddRolePolicyRequest, empty role", emptyRole, true},
		{"invalid AddRolePolicyRequest, empty AccessPolicy", emptyAccessPolicy, true},
		{"invalid AddRolePolicyRequest, invalid AccessPolicy", invalidAccessPolicy, true},
		{"invalid AddRolePolicyRequest, empty HttpMethod", emptyMethodPolicy, true},
		{"invalid AddRolePolicyRequest, invalid HttpMethod", invalidMethodPolicy, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.RolePolicy.Validate()
			if tt.expectError {
				require.Error(t, err, fmt.Sprintf("expect error but not : %s", tt.name))
			} else {
				require.NoError(t, err, fmt.Sprintf("unexpected error occurs : %s", tt.name))
			}
		})
	}
}

func TestAddRolePolicy_UnmarshalJSON(t *testing.T) {
	valid := testAddRolePolicyRequest
	resultTestBytes, _ := json.Marshal(testAddRolePolicyRequest)
	type args struct {
		data []byte
	}

	tests := []struct {
		name          string
		addRolePolicy AddRolePolicyRequest
		args          args
		wantErr       bool
	}{
		{"unmarshal AddRolePolicyRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddRolePolicyRequest, empty data", AddRolePolicyRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddRolePolicyRequest, string data", AddRolePolicyRequest{}, args{[]byte("Invalid AddUserRequest")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addRolePolicy
			err := tt.addRolePolicy.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, expected, tt.addRolePolicy, "Unmarshal did not result in expected AddRolePolicyRequest")
			}
		})
	}
}

func Test_AddRolePolicyReqToRolePolicyModels(t *testing.T) {
	requests := []AddRolePolicyRequest{testAddRolePolicyRequest}
	expectedRoleModels := []models.RolePolicy{
		{
			Role:           mockRole,
			Description:    mockDescription,
			AccessPolicies: []models.AccessPolicy{mockAccessPolicyModel},
		},
	}
	resultModels := AddRolePolicyReqToRolePolicyModels(requests)
	require.Equal(t, expectedRoleModels, resultModels, "AddRolePolicyReqToRolePolicyModels did not result in expected RolePolicy model")
}
