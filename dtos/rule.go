//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

type Rule struct {
	Name string `json:"name" validate:"required"`
	Rule []byte `json:"rule" validate:"required"`
}

func NewRule(name string, rule []byte) Rule {
	return Rule{
		Name: name,
		Rule: rule,
	}
}
