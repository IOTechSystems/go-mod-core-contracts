//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type ScheduleJob struct {
	DBTimestamp
	Id         string
	Name       string
	Definition ScheduleDef
	Actions    []ScheduleAction
	AdminState AdminState
	Labels     []string
	Properties map[string]any
}

type ScheduleDef interface {
	GetBaseScheduleDef() BaseScheduleDef
}

type BaseScheduleDef struct {
	Type ScheduleDefType
}

type DurationScheduleDef struct {
	BaseScheduleDef
	// Duration is the time interval between two consecutive executions
	Duration int64
}

func (d DurationScheduleDef) GetBaseScheduleDef() BaseScheduleDef {
	return d.BaseScheduleDef
}

type CronScheduleDef struct {
	BaseScheduleDef
	// Crontab is the cron expression
	Crontab string
}

func (c CronScheduleDef) GetBaseScheduleDef() BaseScheduleDef {
	return c.BaseScheduleDef
}

type ScheduleAction interface {
	GetBaseScheduleAction() BaseScheduleAction
}

type BaseScheduleAction struct {
	Type        ScheduleActionType
	ContentType string
	Payload     []byte
}

type MessageBusAction struct {
	BaseScheduleAction
	Topic string
}

func (m MessageBusAction) GetBaseScheduleAction() BaseScheduleAction {
	return m.BaseScheduleAction
}

type RESTAction struct {
	BaseScheduleAction
	Address         string
	InjectEdgeXAuth bool
}

func (r RESTAction) GetBaseScheduleAction() BaseScheduleAction {
	return r.BaseScheduleAction
}

type DeviceControlAction struct {
	BaseScheduleAction
	DeviceName string
	SourceName string
}

func (d DeviceControlAction) GetBaseScheduleAction() BaseScheduleAction {
	return d.BaseScheduleAction
}

// ScheduleDefType is used to identify the schedule definition type, i.e., Duration or Cron
type ScheduleDefType string

// ScheduleActionType is used to identify the schedule action type, i.e., MessageBus, REST, or DeviceControl
type ScheduleActionType string
