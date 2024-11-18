//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v4/models"

type RolePolicy struct {
	DBTimestamp
	Id             string         `json:"id"`
	Role           string         `json:"role" validate:"required,edgex-dto-none-empty-string"`
	Description    string         `json:"description,omitempty"`
	AccessPolicies []AccessPolicy `json:"accessPolicies" validate:"gt=0,dive,required"`
}

type AccessPolicy struct {
	Path        string   `json:"path" validate:"required,edgex-dto-none-empty-string"`
	HttpMethods []string `json:"httpMethods" validate:"gt=0,dive,oneof=GET PUT POST PATCH DELETE,required"`
	Effect      string   `json:"effect" validate:"required,oneof=allow deny"`
}

// ToRolePolicyModel transforms the RolePolicy DTO to the RolePolicy Model
func ToRolePolicyModel(rolePolicy RolePolicy) models.RolePolicy {
	return models.RolePolicy{
		Id:             rolePolicy.Id,
		Role:           rolePolicy.Role,
		Description:    rolePolicy.Description,
		AccessPolicies: ToAccessPolicyModels(rolePolicy.AccessPolicies),
	}
}

// FromRolePolicyModelToDTO transforms the RolePolicy model to the RolePolicy DTO
func FromRolePolicyModelToDTO(r models.RolePolicy) RolePolicy {
	return RolePolicy{
		DBTimestamp:    DBTimestamp(r.DBTimestamp),
		Id:             r.Id,
		Role:           r.Role,
		Description:    r.Description,
		AccessPolicies: FromAccessPolicyModelsToDTOs(r.AccessPolicies),
	}
}

// FromRolePolicyModelsToDTOs transforms the RolePolicy model array to the RolePolicy DTO array
func FromRolePolicyModelsToDTOs(rolePolicies []models.RolePolicy) []RolePolicy {
	dtos := make([]RolePolicy, len(rolePolicies))
	for i, r := range rolePolicies {
		dtos[i] = FromRolePolicyModelToDTO(r)
	}
	return dtos
}

// ToAccessPolicyModel transforms the AccessPolicy DTO to the AccessPolicy model
func ToAccessPolicyModel(accessPolicyDTO AccessPolicy) models.AccessPolicy {
	return models.AccessPolicy{
		Path:        accessPolicyDTO.Path,
		HttpMethods: accessPolicyDTO.HttpMethods,
		Effect:      accessPolicyDTO.Effect,
	}
}

// ToAccessPolicyModels transforms the AccessPolicy DTO array to the AccessPolicy model array
func ToAccessPolicyModels(accessPolicyDTOs []AccessPolicy) []models.AccessPolicy {
	accessPolicyModels := make([]models.AccessPolicy, len(accessPolicyDTOs))
	for i, a := range accessPolicyDTOs {
		accessPolicyModels[i] = ToAccessPolicyModel(a)
	}
	return accessPolicyModels
}

// FromAccessPolicyModelToDTO transforms the AccessPolicy Model to the AccessPolicy DTO
func FromAccessPolicyModelToDTO(d models.AccessPolicy) AccessPolicy {
	return AccessPolicy{
		Path:        d.Path,
		HttpMethods: d.HttpMethods,
		Effect:      d.Effect,
	}
}

// FromAccessPolicyModelsToDTOs transforms the AccessPolicy model array to the AccessPolicy DTO array
func FromAccessPolicyModelsToDTOs(accessPolicies []models.AccessPolicy) []AccessPolicy {
	dtos := make([]AccessPolicy, len(accessPolicies))
	for i, a := range accessPolicies {
		dtos[i] = FromAccessPolicyModelToDTO(a)
	}
	return dtos
}
