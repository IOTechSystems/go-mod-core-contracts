package v1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	v1Models "github.com/edgexfoundry/go-mod-core-contracts/v2/v1/models"
)

// TransformToV1DeviceProfile transform v2 profile to v1
func TransformToV1DeviceProfile(profile models.DeviceProfile) v1Models.DeviceProfile {
	dto := v1Models.DeviceProfile{
		DescribedObject: v1Models.DescribedObject{},
		Name:            profile.Name,
		Manufacturer:    profile.Manufacturer,
		Model:           profile.Model,
		Labels:          profile.Labels,
		DeviceResources: nil,
		DeviceCommands:  nil,
	}
	dto.DeviceResources = toV1DeviceResources(profile.DeviceResources)
	dto.DeviceCommands = toV1DeviceCommands(profile.DeviceCommands)
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
					Type:          r.Properties.ValueType,
					ReadWrite:     r.Properties.ReadWrite,
					Minimum:       r.Properties.Minimum,
					Maximum:       r.Properties.Maximum,
					DefaultValue:  r.Properties.DefaultValue,
					Mask:          r.Properties.Mask,
					Shift:         r.Properties.Shift,
					Scale:         r.Properties.Scale,
					Offset:        r.Properties.Offset,
					Base:          r.Properties.Base,
					Assertion:     r.Properties.Assertion,
					FloatEncoding: "eNotation",
					MediaType:     r.Properties.MediaType,
				},
				Units: v1Models.Units{
					Type:         "String",
					ReadWrite:    common.ReadWrite_RW,
					DefaultValue: r.Properties.Units,
				},
			},
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
		operations[i] = v1Models.ResourceOperation{
			Index:          strconv.Itoa(i),
			Operation:      op,
			DeviceResource: ro.DeviceResource,
			Mappings:       ro.Mappings,
		}
	}
	return operations
}

// TransformToV2DeviceProfile transform v1 profile to v2
func TransformToV2DeviceProfile(profile v1Models.DeviceProfile) models.DeviceProfile {
	dto := models.DeviceProfile{
		Description:     profile.Description,
		Name:            profile.Name,
		Manufacturer:    profile.Manufacturer,
		Model:           profile.Model,
		Labels:          profile.Labels,
		DeviceResources: nil,
		DeviceCommands:  nil,
	}
	dto.DeviceResources = toV2DeviceResources(profile.DeviceResources)
	dto.DeviceCommands = toV2DeviceCommands(profile.DeviceCommands)
	return dto
}

func toV2DeviceResources(deviceResources []v1Models.DeviceResource) []models.DeviceResource {
	resources := make([]models.DeviceResource, len(deviceResources))
	for i, r := range deviceResources {
		resources[i] = models.DeviceResource{
			Description: r.Description,
			Name:        r.Name,
			IsHidden:    false,
			Tags:        toV2Tags(r.Tags),
			Properties: models.ResourceProperties{
				ValueType:    r.Properties.Value.Type,
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
		if len(c.Get) > 0 && len(c.Set) > 0 {
			commands[i].ReadWrite = common.ReadWrite_RW
		} else if len(c.Set) > 0 {
			commands[i].ReadWrite = common.ReadWrite_W
		} else {
			commands[i].ReadWrite = common.ReadWrite_R
		}
		var ros []models.ResourceOperation
		for _, op := range c.Get {
			ro := models.ResourceOperation{
				DeviceResource: op.DeviceResource,
				DefaultValue:   "",
				Mappings:       op.Mappings,
			}
			ros = append(ros, ro)
		}
		for _, op := range c.Set {
			exists := existFromResourceDTO(op.DeviceResource, ros)
			if exists {
				continue
			}
			ro := models.ResourceOperation{
				DeviceResource: op.DeviceResource,
				DefaultValue:   "",
				Mappings:       op.Mappings,
			}
			ros = append(ros, ro)
		}
		commands[i].ResourceOperations = ros
	}

	return commands
}

func existFromResourceDTO(resourceName string, ros []models.ResourceOperation) bool {
	for _, ro := range ros {
		if ro.DeviceResource == resourceName {
			return true
		}
	}
	return false
}
