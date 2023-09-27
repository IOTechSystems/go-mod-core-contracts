//
// Copyright (C) 2023 IOTech Ltd
//

package dtos

import (
	"errors"
	"fmt"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

type Registration struct {
	DBTimestamp   `json:",inline"`
	ServiceId     string      `json:"serviceId"`
	Status        string      `json:"status"`
	Host          string      `json:"host"`
	Port          int         `json:"port" `
	HealthCheck   HealthCheck `json:",inline"`
	LastConnected int64       `json:"lastConnected"`
}

type HealthCheck struct {
	Interval string `json:"interval"`
	Path     string `json:"path"`
	Type     string `json:"type"`
}

func ToRegistrationModel(dto Registration) models.Registration {
	var r models.Registration
	r.ServiceId = dto.ServiceId
	r.Status = dto.Status
	r.Host = dto.Host
	r.Port = dto.Port
	r.LastConnected = dto.LastConnected
	r.HealthCheck.Type = dto.HealthCheck.Type
	r.HealthCheck.Path = dto.HealthCheck.Path
	r.HealthCheck.Interval = dto.HealthCheck.Interval

	return r
}

func FromRegistrationModelToDTO(r models.Registration) Registration {
	var dto Registration
	dto.DBTimestamp = DBTimestamp(r.DBTimestamp)
	dto.ServiceId = r.ServiceId
	dto.Status = r.Status
	dto.Host = r.Host
	dto.Port = r.Port
	dto.LastConnected = r.LastConnected
	dto.HealthCheck.Type = r.HealthCheck.Type
	dto.HealthCheck.Path = r.HealthCheck.Path
	dto.HealthCheck.Interval = r.HealthCheck.Interval

	return dto
}

func (r *Registration) Validate() error {
	// check if either the ServiceId, Port or HealthCheck.Type field is empty
	if r.ServiceId == "" {
		return errors.New(" the ServiceId field is empty")
	}
	if r.Port == 0 {
		return errors.New(" the Port field is empty")
	}
	if r.HealthCheck.Type == "" {
		return errors.New(" the HealthCheck Type field is empty")
	}
	// check if the Interval field is a valid duration string
	_, err := time.ParseDuration(r.HealthCheck.Interval)
	if err != nil {
		return fmt.Errorf("health check interval is not in Go duration string format: %s", err.Error())
	}
	// check if the health status value is UP, DOWN, UNKNOWN, or HALT
	// if the value is invalid or empty, assign UNKNOWN to the status value
	switch r.Status {
	case models.Up, models.Down, models.Unknown, models.Halt:
		break
	default:
		r.Status = models.Unknown
	}
	return nil
}
