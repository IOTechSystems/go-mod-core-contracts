//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

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

func (scheduleJob *ScheduleJob) UnmarshalJSON(b []byte) error {
	var alias struct {
		DBTimestamp
		Id         string
		Name       string
		Definition any
		Actions    []any
		AdminState AdminState
		Labels     []string
		Properties map[string]any
	}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal ScheduleJob.", err)
	}

	def, err := instantiateScheduleDef(alias.Definition)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	actions := make([]ScheduleAction, len(alias.Actions))
	for i, a := range alias.Actions {
		action, err := instantiateScheduleAction(a)
		if err != nil {
			return errors.NewCommonEdgeXWrapper(err)
		}
		actions[i] = action
	}

	*scheduleJob = ScheduleJob{
		DBTimestamp: alias.DBTimestamp,
		Id:          alias.Id,
		Name:        alias.Name,
		Definition:  def,
		Actions:     actions,
		AdminState:  alias.AdminState,
		Labels:      alias.Labels,
		Properties:  alias.Properties,
	}
	return nil
}

type ScheduleDef interface {
	GetBaseScheduleDef() BaseScheduleDef
}

// instantiateScheduleDef instantiate the interface to the corresponding schedule definition type
func instantiateScheduleDef(i any) (def ScheduleDef, err error) {
	d, err := json.Marshal(i)
	if err != nil {
		return def, errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to marshal ScheduleDef.", err)
	}
	return unmarshalScheduleDef(d)
}

func unmarshalScheduleDef(b []byte) (def ScheduleDef, err error) {
	var alias struct {
		Type string
	}
	if err = json.Unmarshal(b, &alias); err != nil {
		return def, errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal ScheduleDef.", err)
	}
	switch alias.Type {
	case common.DefInterval:
		var intervalDef IntervalScheduleDef
		if err = json.Unmarshal(b, &intervalDef); err != nil {
			return def, errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal INTERVAL ScheduleDef.", err)
		}
		def = intervalDef
	case common.DefCron:
		var cronDef CronScheduleDef
		if err = json.Unmarshal(b, &cronDef); err != nil {
			return def, errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal CRON ScheduleDef.", err)
		}
		def = cronDef
	default:
		return def, errors.NewCommonEdgeX(errors.KindContractInvalid, "Unsupported schedule definition type", err)
	}
	return def, nil
}

type BaseScheduleDef struct {
	Type ScheduleDefType
}

type IntervalScheduleDef struct {
	BaseScheduleDef
	// Interval specifies the time interval between two consecutive executions
	Interval string
}

func (d IntervalScheduleDef) GetBaseScheduleDef() BaseScheduleDef {
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
	// WithEmptyPayload returns a copy of the ScheduleAction with empty payload, which is used by ScheduleActionRecord to remove the payload before storing the record into database
	WithEmptyPayload() ScheduleAction
}

// instantiateScheduleAction instantiate the interface to the corresponding schedule action type
func instantiateScheduleAction(i any) (action ScheduleAction, err error) {
	a, err := json.Marshal(i)
	if err != nil {
		return action, errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to marshal ScheduleAction.", err)
	}
	return unmarshalScheduleAction(a)
}

func unmarshalScheduleAction(b []byte) (action ScheduleAction, err error) {
	var alias struct {
		Type string
	}
	if err = json.Unmarshal(b, &alias); err != nil {
		return action, errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal ScheduleAction.", err)
	}
	switch alias.Type {
	case common.ActionEdgeXMessageBus:
		var edgeXMessageBusAction EdgeXMessageBusAction
		if err = json.Unmarshal(b, &edgeXMessageBusAction); err != nil {
			return action, errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal EDGEXMESSAGEBUS ScheduleAction.", err)
		}
		action = edgeXMessageBusAction
	case common.ActionREST:
		var restAction RESTAction
		if err = json.Unmarshal(b, &restAction); err != nil {
			return action, errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal REST ScheduleAction.", err)
		}
		action = restAction
	case common.ActionDeviceControl:
		var deviceControlAction DeviceControlAction
		if err = json.Unmarshal(b, &deviceControlAction); err != nil {
			return action, errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal DEVICECONTROL ScheduleAction.", err)
		}
		action = deviceControlAction
	default:
		return action, errors.NewCommonEdgeX(errors.KindContractInvalid, "Unsupported schedule action type", err)
	}
	return action, nil
}

type BaseScheduleAction struct {
	Type        ScheduleActionType
	ContentType string
	Payload     []byte
}

type EdgeXMessageBusAction struct {
	BaseScheduleAction
	Topic string
}

func (m EdgeXMessageBusAction) GetBaseScheduleAction() BaseScheduleAction {
	return m.BaseScheduleAction
}
func (m EdgeXMessageBusAction) WithEmptyPayload() ScheduleAction {
	m.Payload = nil
	return m
}

type RESTAction struct {
	BaseScheduleAction
	Address         string
	Method          string
	InjectEdgeXAuth bool
}

func (r RESTAction) GetBaseScheduleAction() BaseScheduleAction {
	return r.BaseScheduleAction
}
func (r RESTAction) WithEmptyPayload() ScheduleAction {
	r.Payload = nil
	return r
}

type DeviceControlAction struct {
	BaseScheduleAction
	DeviceName string
	SourceName string
}

func (d DeviceControlAction) GetBaseScheduleAction() BaseScheduleAction {
	return d.BaseScheduleAction
}
func (d DeviceControlAction) WithEmptyPayload() ScheduleAction {
	d.Payload = nil
	return d
}

// ScheduleDefType is used to identify the schedule definition type, i.e., INTERVAL or CRON
type ScheduleDefType string

// ScheduleActionType is used to identify the schedule action type, i.e., EDGEXMESSAGEBUS, REST, or DEVICECONTROL
type ScheduleActionType string
