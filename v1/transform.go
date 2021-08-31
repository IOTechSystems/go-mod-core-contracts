// Copyright (C) 2021 IOTech Ltd

package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	v1Models "github.com/edgexfoundry/go-mod-core-contracts/v2/v1/models"
)

// TransformToV1DeviceProfile transform v2 profile to v1
func TransformToV1DeviceProfile(profile models.DeviceProfile) v1Models.DeviceProfile {
	dto := v1Models.DeviceProfile{
		DescribedObject: v1Models.DescribedObject{Description: profile.Description},
		Name:            profile.Name,
		Manufacturer:    profile.Manufacturer,
		Model:           profile.Model,
		Labels:          profile.Labels,
	}
	dto.DeviceResources = toV1DeviceResources(profile.DeviceResources)
	dto.DeviceCommands = toV1DeviceCommands(profile.DeviceCommands)
	dto.CoreCommands = toV1CoreCommand(profile.DeviceCommands)
	return dto
}

func toV1DeviceResources(deviceResources []models.DeviceResource) []v1Models.DeviceResource {
	resources := make([]v1Models.DeviceResource, len(deviceResources))
	for i, r := range deviceResources {
		resources[i] = v1Models.DeviceResource{
			Name:        r.Name,
			Description: r.Description,
			Tags:        toV1Tags(r.Tags),
			Attributes:  toV1Attribute(r.Attributes),
			Properties: v1Models.ProfileProperty{
				Value: v1Models.PropertyValue{
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
				Units: v1Models.Units{
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

func toV1DeviceCommands(deviceCommands []models.DeviceCommand) []v1Models.ProfileResource {
	commands := make([]v1Models.ProfileResource, len(deviceCommands))
	for i, c := range deviceCommands {
		commands[i] = v1Models.ProfileResource{
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

func toV1Operation(op string, resourceOperations []models.ResourceOperation) []v1Models.ResourceOperation {
	operations := make([]v1Models.ResourceOperation, len(resourceOperations))
	for i, ro := range resourceOperations {
		valueMappings := ro.Mappings
		if op == ResourceOperationSet && len(ro.Mappings) != 0 {
			valueMappings = reverseMapKeyValue(ro.Mappings)
		}
		operations[i] = v1Models.ResourceOperation{
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

func toV1CoreCommand(deviceCommands []models.DeviceCommand) []v1Models.Command {
	commands := make([]v1Models.Command, len(deviceCommands))
	for i, c := range deviceCommands {
		commands[i] = v1Models.Command{
			Name: c.Name,
		}
		if strings.Contains(c.ReadWrite, common.ReadWrite_R) {
			commands[i].Get = toV1GetAction(c.Name, c.ResourceOperations)
		}
		if strings.Contains(c.ReadWrite, common.ReadWrite_W) {
			commands[i].Put = toV1PutAction(c.Name, c.ResourceOperations)
		}
	}

	return commands
}

func toV1GetAction(cmdName string, resourceOperations []models.ResourceOperation) v1Models.Get {
	expectedValues := make([]string, len(resourceOperations))
	for i, ro := range resourceOperations {
		expectedValues[i] = ro.DeviceResource
	}
	action := v1Models.Action{
		Path: fmt.Sprintf("/api/v1/device/{deviceId}/%s", cmdName),
		Responses: []v1Models.Response{
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
	return v1Models.Get{Action: action}
}

func toV1PutAction(cmdName string, resourceOperations []models.ResourceOperation) v1Models.Put {
	parameterNames := make([]string, len(resourceOperations))
	for i, ro := range resourceOperations {
		parameterNames[i] = ro.DeviceResource
	}
	action := v1Models.Action{
		Path: fmt.Sprintf("/api/v1/device/{deviceId}/%s", cmdName),
		Responses: []v1Models.Response{
			{
				Code:        strconv.Itoa(http.StatusNoContent),
				Description: fmt.Sprintf("Issue the Put command %s", cmdName),
			}, {
				Code:        strconv.Itoa(http.StatusInternalServerError),
				Description: "internal server error",
			},
		},
	}
	return v1Models.Put{
		Action:         action,
		ParameterNames: parameterNames,
	}
}

// TransformToV2DeviceProfile transform v1 profile to v2
func TransformToV2DeviceProfile(profile v1Models.DeviceProfile) models.DeviceProfile {
	dto := models.DeviceProfile{
		Description:  profile.Description,
		Name:         profile.Name,
		Manufacturer: profile.Manufacturer,
		Model:        profile.Model,
		Labels:       profile.Labels,
	}
	dto.DeviceResources = toV2DeviceResources(profile)
	dto.DeviceCommands = toV2DeviceCommands(profile.DeviceCommands)
	return dto
}

func toV2DeviceResources(profile v1Models.DeviceProfile) []models.DeviceResource {
	resources := make([]models.DeviceResource, len(profile.DeviceResources))
	for i, r := range profile.DeviceResources {
		resources[i] = models.DeviceResource{
			Description: r.Description,
			Name:        r.Name,
			IsHidden:    isResourceHidden(r.Name, profile),
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
	}

	return resources
}

func isResourceHidden(resourceName string, profile v1Models.DeviceProfile) bool {
	// Check whether the resource exists in the CoreCommands, if exists, the resource is not hidden.
	for _, coreCommand := range profile.CoreCommands {
		for _, res := range coreCommand.Get.Responses {
			isContains := contains(res.ExpectedValues, resourceName)
			if isContains {
				return false
			}
		}
		isContains := contains(coreCommand.Put.ParameterNames, resourceName)
		if isContains {
			return false
		}
	}
	// Check whether the resource exists in the DeviceCommands, if exists, the resource is not hidden.
	for _, deviceCommand := range profile.DeviceCommands {
		for _, ro := range deviceCommand.Get {
			if ro.DeviceResource == resourceName {
				return false
			}
		}
		for _, ro := range deviceCommand.Set {
			if ro.DeviceResource == resourceName {
				return false
			}
		}
	}

	return true
}

func contains(list []string, e string) bool {
	for _, v := range list {
		if v == e {
			return true
		}
	}
	return false
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

func toV2DeviceCommands(deviceCommands []v1Models.ProfileResource) []models.DeviceCommand {
	commands := make([]models.DeviceCommand, len(deviceCommands))
	for i, c := range deviceCommands {
		commands[i] = models.DeviceCommand{
			Name:               c.Name,
			IsHidden:           false,
			ReadWrite:          "",
			ResourceOperations: nil,
		}
		var ros []models.ResourceOperation
		if len(c.Get) > 0 && len(c.Set) > 0 {
			commands[i].ReadWrite = common.ReadWrite_RW
			for _, getOp := range c.Get {
				for _, setOp := range c.Set {
					if getOp.DeviceResource == setOp.DeviceResource {
						ro := models.ResourceOperation{
							DeviceResource: getOp.DeviceResource,
							DefaultValue:   "",
							Mappings:       getOp.Mappings,
						}
						ros = append(ros, ro)
						break
					}
				}
			}
		} else if len(c.Set) > 0 {
			commands[i].ReadWrite = common.ReadWrite_W
			for _, op := range c.Set {
				ro := models.ResourceOperation{
					DeviceResource: op.DeviceResource,
					DefaultValue:   "",
					Mappings:       op.Mappings,
				}
				ros = append(ros, ro)
			}
		} else {
			commands[i].ReadWrite = common.ReadWrite_R
			for _, op := range c.Get {
				ro := models.ResourceOperation{
					DeviceResource: op.DeviceResource,
					DefaultValue:   "",
					Mappings:       op.Mappings,
				}
				ros = append(ros, ro)
			}
		}
		commands[i].ResourceOperations = ros
	}

	return commands
}
