//
// Copyright (C) 2023 IOTech Ltd
//

package models

type Registration struct {
	DBTimestamp
	ServiceId     string
	Status        string
	Host          string
	Port          int
	HealthCheck   HealthCheck
	LastConnected int64
}

type HealthCheck struct {
	Interval string
	Path     string
	Type     string
}
