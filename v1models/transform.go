// Copyright (C) 2021 IOTech Ltd

package v1models

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

const (
	primaryTable    = "primaryTable"
	startingAddress = "startingAddress"
)

// TransformProfileFromV2ToV1 transform v2 profile to v1
func TransformProfileFromV2ToV1(profile models.DeviceProfile) (DeviceProfile, errors.EdgeX) {
	v2dpDto := dtos.FromDeviceProfileModelToDTO(profile)
	err := v2dpDto.Validate()
	if err != nil {
		return DeviceProfile{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid v2 device profile", err)
	}

	v1dp := DeviceProfile{
		DescribedObject: DescribedObject{Description: profile.Description},
		Name:            profile.Name,
		Manufacturer:    profile.Manufacturer,
		Model:           profile.Model,
		Labels:          profile.Labels,
	}
	v1dp.DeviceResources = toV1DeviceResources(profile.DeviceResources)
	v1dp.DeviceCommands = toV1DeviceCommands(profile.DeviceCommands)
	v1dp.CoreCommands = toV1CoreCommand(profile.DeviceResources, profile.DeviceCommands)

	_, err = v1dp.Validate()
	if err != nil {
		return v1dp, errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid v1 device profile after transforming from v2 to v1", err)
	}
	return v1dp, nil
}

func toV1DeviceResources(deviceResources []models.DeviceResource) []DeviceResource {
	resources := make([]DeviceResource, len(deviceResources))
	for i, r := range deviceResources {
		resources[i] = DeviceResource{
			Name:        r.Name,
			Description: r.Description,
			Tags:        toV1Tags(r.Tags),
			Attributes:  toV1Attribute(r.Attributes),
			Properties: ProfileProperty{
				Value: PropertyValue{
					Type:         r.Properties.ValueType,
					ReadWrite:    r.Properties.ReadWrite,
					Minimum:      r.Properties.Minimum,
					Maximum:      r.Properties.Maximum,
					DefaultValue: r.Properties.DefaultValue,
					Mask:         r.Properties.Mask,
					Shift:        r.Properties.Shift,
					Scale:        r.Properties.Scale,
					Offset:       r.Properties.Offset,
					Base:         r.Properties.Base,
					Assertion:    r.Properties.Assertion,
					MediaType:    r.Properties.MediaType,
				},
				Units: Units{
					Type:         "String",
					ReadWrite:    common.ReadWrite_R,
					DefaultValue: r.Properties.Units,
				},
			},
		}
		if r.Properties.ValueType == common.ValueTypeFloat32 || r.Properties.ValueType == common.ValueTypeFloat64 ||
			r.Properties.ValueType == common.ValueTypeFloat32Array || r.Properties.ValueType == common.ValueTypeFloat64Array {
			resources[i].Properties.Value.FloatEncoding = "eNotation"
		}
	}

	return resources
}

func toV1Tags(tags map[string]interface{}) map[string]string {
	dto := make(map[string]string)
	for k, v := range tags {
		dto[k] = fmt.Sprintf("%v", v)
	}
	return dto
}

func toV1Attribute(attributes map[string]interface{}) map[string]string {
	dto := make(map[string]string)
	for k, v := range attributes {
		dto[k] = fmt.Sprintf("%v", v)
	}
	return dto
}

const (
	ResourceOperationGet = "get"
	ResourceOperationSet = "set"
)

func toV1DeviceCommands(deviceCommands []models.DeviceCommand) []ProfileResource {
	commands := make([]ProfileResource, len(deviceCommands))
	for i, c := range deviceCommands {
		commands[i] = ProfileResource{
			Name: c.Name,
		}
		if strings.Contains(c.ReadWrite, common.ReadWrite_R) {
			commands[i].Get = toV1Operation(ResourceOperationGet, c.ResourceOperations)
		}
		if strings.Contains(c.ReadWrite, common.ReadWrite_W) {
			commands[i].Set = toV1Operation(ResourceOperationSet, c.ResourceOperations)
		}
	}

	return commands
}

func toV1Operation(op string, resourceOperations []models.ResourceOperation) []ResourceOperation {
	operations := make([]ResourceOperation, len(resourceOperations))
	for i, ro := range resourceOperations {
		valueMappings := ro.Mappings
		if op == ResourceOperationSet && len(ro.Mappings) != 0 {
			valueMappings = reverseMapKeyValue(ro.Mappings)
		}
		operations[i] = ResourceOperation{
			Index:          strconv.Itoa(i),
			Operation:      op,
			DeviceResource: ro.DeviceResource,
			Mappings:       valueMappings,
		}
	}
	return operations
}

func reverseMapKeyValue(mappings map[string]string) map[string]string {
	valueMappings := make(map[string]string, len(mappings))
	for k, v := range mappings {
		valueMappings[v] = k
	}
	return valueMappings
}

func toV1CoreCommand(v2DeviceResources []models.DeviceResource, v2DeviceCommands []models.DeviceCommand) []Command {
	var commands []Command

	// Create v1 CoreCommands by v2DeviceCommands
	for _, cmd := range v2DeviceCommands {
		if cmd.IsHidden {
			continue
		}
		v1Command := Command{
			Name: cmd.Name,
		}
		if strings.Contains(cmd.ReadWrite, common.ReadWrite_R) {
			v1Command.Get = toV1GetAction(cmd.Name, cmd.ResourceOperations)
		}
		if strings.Contains(cmd.ReadWrite, common.ReadWrite_W) {
			v1Command.Put = toV1PutAction(cmd.Name, cmd.ResourceOperations)
		}
		commands = append(commands, v1Command)
	}

	// Create v1 CoreCommands by v2DeviceResources
	for _, resource := range v2DeviceResources {
		if resource.IsHidden {
			continue
		}
		// Skip if the resource exists in the v2 DeviceCommands
		if existFromV2DeviceCommands(resource.Name, v2DeviceCommands) {
			continue
		}

		v1Command := Command{
			Name: resource.Name,
		}
		if strings.Contains(resource.Properties.ReadWrite, common.ReadWrite_R) {
			v1Command.Get = toV1GetAction(resource.Name, []models.ResourceOperation{{DeviceResource: resource.Name}})
		}
		if strings.Contains(resource.Properties.ReadWrite, common.ReadWrite_W) {
			v1Command.Put = toV1PutAction(resource.Name, []models.ResourceOperation{{DeviceResource: resource.Name}})
		}
		commands = append(commands, v1Command)
	}

	return commands
}

func existFromV2DeviceCommands(resourceName string, v2DeviceCommands []models.DeviceCommand) bool {
	for _, cmd := range v2DeviceCommands {
		if resourceName == cmd.Name {
			return true
		}
	}
	return false
}

func toV1GetAction(cmdName string, resourceOperations []models.ResourceOperation) Get {
	expectedValues := make([]string, len(resourceOperations))
	for i, ro := range resourceOperations {
		expectedValues[i] = ro.DeviceResource
	}
	action := Action{
		Path: fmt.Sprintf("/api/v1/device/{deviceId}/%s", cmdName),
		Responses: []Response{
			{
				Code:           strconv.Itoa(http.StatusOK),
				Description:    fmt.Sprintf("Issue the Get command %s", cmdName),
				ExpectedValues: expectedValues,
			}, {
				Code:        strconv.Itoa(http.StatusInternalServerError),
				Description: "internal server error",
			},
		},
	}
	return Get{Action: action}
}

func toV1PutAction(cmdName string, resourceOperations []models.ResourceOperation) Put {
	parameterNames := make([]string, len(resourceOperations))
	for i, ro := range resourceOperations {
		parameterNames[i] = ro.DeviceResource
	}
	action := Action{
		Path: fmt.Sprintf("/api/v1/device/{deviceId}/%s", cmdName),
		Responses: []Response{
			{
				Code:        strconv.Itoa(http.StatusNoContent),
				Description: fmt.Sprintf("Issue the Put command %s", cmdName),
			}, {
				Code:        strconv.Itoa(http.StatusInternalServerError),
				Description: "internal server error",
			},
		},
	}
	return Put{
		Action:         action,
		ParameterNames: parameterNames,
	}
}

// TransformProfileFromV1ToV2 transform v1 profile to v2
func TransformProfileFromV1ToV2(profile DeviceProfile) (models.DeviceProfile, errors.EdgeX) {
	_, err := profile.Validate()
	if err != nil {
		return models.DeviceProfile{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid v1 device profile", err)
	}

	v2dp := models.DeviceProfile{
		Description:  profile.Description,
		Name:         profile.Name,
		Manufacturer: profile.Manufacturer,
		Model:        profile.Model,
		Labels:       profile.Labels,
	}
	v2dp.DeviceResources, err = toV2DeviceResources(profile)
	if err != nil {
		return models.DeviceProfile{}, errors.NewCommonEdgeXWrapper(err)
	}

	v2dp.DeviceCommands = toV2DeviceCommands(profile.DeviceCommands)
	for i, r := range v2dp.DeviceResources {
		v2dp.DeviceResources[i].IsHidden = isV2ResourceHidden(r.Name, profile.CoreCommands)
	}
	for i, cmd := range v2dp.DeviceCommands {
		v2dp.DeviceCommands[i].IsHidden = isV2DeviceCommandHidden(cmd.Name, profile.CoreCommands)
	}
	// Convert StartingAddress for Modbus protocol
	err = ConvertStartingAddressToZeroBased(&v2dp)
	if err != nil {
		return v2dp, errors.NewCommonEdgeX(errors.KindContractInvalid, "convert startingAddress from string to int for v2 failed", err)
	}
	v2dpDto := dtos.FromDeviceProfileModelToDTO(v2dp)
	err = v2dpDto.Validate()
	if err != nil {
		return v2dp, errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid v2 device profile after transforming from v1 to v2", err)
	}
	return v2dp, nil
}

func TransformResourceFromV1ToV2(r DeviceResource) (models.DeviceResource, errors.EdgeX) {
	v2dr := models.DeviceResource{
		Description: r.Description,
		Name:        r.Name,
		IsHidden:    false,
		Tags:        toV2Tags(r.Tags),
		Properties: models.ResourceProperties{
			ValueType:    strings.Title(strings.ToLower(r.Properties.Value.Type)),
			ReadWrite:    r.Properties.Value.ReadWrite,
			Units:        r.Properties.Units.DefaultValue,
			Minimum:      r.Properties.Value.Minimum,
			Maximum:      r.Properties.Value.Maximum,
			DefaultValue: r.Properties.Value.DefaultValue,
			Mask:         r.Properties.Value.Mask,
			Shift:        r.Properties.Value.Shift,
			Scale:        r.Properties.Value.Scale,
			Offset:       r.Properties.Value.Offset,
			Base:         r.Properties.Value.Base,
			Assertion:    r.Properties.Value.Assertion,
			MediaType:    r.Properties.Value.MediaType,
		},
		Attributes: toV2Attributes(r.Attributes),
	}

	v2drDTO := dtos.FromDeviceResourceModelToDTO(v2dr)
	err := v2drDTO.Validate()
	if err != nil {
		return v2dr, errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid v2 device resource after transforming from v1 to v2", err)
	}

	return v2dr, nil
}

func toV2DeviceResources(profile DeviceProfile) ([]models.DeviceResource, errors.EdgeX) {
	resources := make([]models.DeviceResource, len(profile.DeviceResources))
	for i, r := range profile.DeviceResources {
		v2, err := TransformResourceFromV1ToV2(r)
		if err != nil {
			return nil, errors.NewCommonEdgeXWrapper(err)
		}

		resources[i] = v2
	}

	return resources, nil
}

func isV2ResourceHidden(resourceName string, v1CoreCommands []Command) bool {
	// Check whether the resource exists in the v1 CoreCommands, if exists, the resource is not hidden.
	for _, v1CoreCommand := range v1CoreCommands {
		if v1CoreCommand.Name == resourceName {
			return false
		}
	}
	return true
}

func isV2DeviceCommandHidden(v2DeviceCommandName string, v1CoreCommands []Command) bool {
	for _, v1CoreCommand := range v1CoreCommands {
		if v1CoreCommand.Name == v2DeviceCommandName || v2SetCommandName(v1CoreCommand.Name) == v2DeviceCommandName {
			return false
		}
	}
	return true
}

func v2SetCommandName(cmdName string) string {
	return fmt.Sprintf("%s_Set", cmdName)
}

func toV2Tags(tags map[string]string) map[string]interface{} {
	dto := make(map[string]interface{})
	for k, v := range tags {
		dto[k] = v
	}
	return dto
}

func toV2Attributes(attributes map[string]string) map[string]interface{} {
	dto := make(map[string]interface{})
	for k, v := range attributes {
		dto[k] = v
	}
	return dto
}

func toV2DeviceCommands(deviceCommands []ProfileResource) []models.DeviceCommand {
	var commands []models.DeviceCommand
	for _, c := range deviceCommands {
		if len(c.Get) > 0 && len(c.Set) > 0 && len(c.Get) == len(c.Set) {
			command := models.DeviceCommand{
				Name:               c.Name,
				IsHidden:           true,
				ReadWrite:          common.ReadWrite_RW,
				ResourceOperations: toV2ResourceOperations(common.ReadWrite_RW, c.Get),
			}
			commands = append(commands, command)

		} else if len(c.Get) > 0 && len(c.Set) > 0 && len(c.Get) != len(c.Set) {
			readCommand := models.DeviceCommand{
				Name:               c.Name,
				IsHidden:           true,
				ReadWrite:          common.ReadWrite_R,
				ResourceOperations: toV2ResourceOperations(common.ReadWrite_R, c.Get),
			}
			commands = append(commands, readCommand)

			writeCommand := models.DeviceCommand{
				Name:               v2SetCommandName(c.Name),
				IsHidden:           true,
				ReadWrite:          common.ReadWrite_W,
				ResourceOperations: toV2ResourceOperations(common.ReadWrite_W, c.Set),
			}
			commands = append(commands, writeCommand)

		} else if len(c.Set) > 0 {
			command := models.DeviceCommand{
				Name:               c.Name,
				IsHidden:           true,
				ReadWrite:          common.ReadWrite_W,
				ResourceOperations: toV2ResourceOperations(common.ReadWrite_W, c.Set),
			}
			commands = append(commands, command)

		} else if len(c.Get) > 0 {
			command := models.DeviceCommand{
				Name:               c.Name,
				IsHidden:           true,
				ReadWrite:          common.ReadWrite_R,
				ResourceOperations: toV2ResourceOperations(common.ReadWrite_R, c.Get),
			}
			commands = append(commands, command)

		}
	}

	return commands
}

func toV2ResourceOperations(readWrite string, v1ros []ResourceOperation) []models.ResourceOperation {
	var v2ros []models.ResourceOperation
	for _, v1ro := range v1ros {
		v2ro := models.ResourceOperation{
			DeviceResource: v1ro.DeviceResource,
			DefaultValue:   v1ro.Parameter,
			Mappings:       v1ro.Mappings,
		}
		if readWrite == common.ReadWrite_W && len(v2ro.Mappings) != 0 {
			v2ro.Mappings = reverseMapKeyValue(v2ro.Mappings)
		}
		v2ros = append(v2ros, v2ro)
	}
	return v2ros
}

// ConvertStartingAddressToOneBased convert startingAddress attribute from zero-based to one-based when transforming from v2 to v1 profile
func ConvertStartingAddressToOneBased(profile *DeviceProfile) errors.EdgeX {
	if len(profile.DeviceResources) == 0 {
		return nil
	}
	if isModbusAttributes(profile.DeviceResources[0].Attributes) {
		for i, resource := range profile.DeviceResources {
			for k, v := range resource.Attributes {
				if k == startingAddress {
					val, err := strconv.Atoi(v)
					if err != nil {
						return errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid startingAddress", err)
					}
					profile.DeviceResources[i].Attributes[k] = fmt.Sprintf("%v", val+1)
				}
			}
		}
	}
	return nil
}

// ConvertStartingAddressToZeroBased convert startingAddress attribute from one-based to zero-based when transforming from v1 to v2 profile
func ConvertStartingAddressToZeroBased(profile *models.DeviceProfile) errors.EdgeX {
	if len(profile.DeviceResources) == 0 {
		return nil
	}
	if isModbusAttributes(toV1Attribute(profile.DeviceResources[0].Attributes)) {
		for i, resource := range profile.DeviceResources {
			for k, v := range resource.Attributes {
				if k == startingAddress {
					val, err := strconv.Atoi(fmt.Sprint(v))
					if err != nil {
						return errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid startingAddress", err)
					}
					profile.DeviceResources[i].Attributes[k] = val - 1
				}
			}
		}
	}
	return nil
}

func isModbusAttributes(attributes map[string]string) bool {
	_, isPrimaryTableExists := attributes[primaryTable]
	_, isStartingAddressExists := attributes[startingAddress]
	return isPrimaryTableExists && isStartingAddressExists
}
