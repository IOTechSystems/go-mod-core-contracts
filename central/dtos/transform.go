//
// Copyright (C) 2024 IOTech Ltd
//

package dtos

import (
	"fmt"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/v2dtos"
)

// TransformProfileFromV2ToV3 converts the v2 DeviceProfile DTO to v3
func TransformProfileFromV2ToV3(v2Profile v2dtos.DeviceProfile) (dtos.DeviceProfile, errors.EdgeX) {
	resources, err := transformResourceFromV2ToV3(v2Profile.DeviceResources)
	if err != nil {
		return dtos.DeviceProfile{}, errors.NewCommonEdgeXWrapper(err)
	}

	profile := dtos.DeviceProfile{
		DeviceProfileBasicInfo: dtos.DeviceProfileBasicInfo{
			DBTimestamp:  dtos.DBTimestamp(v2Profile.DBTimestamp),
			Id:           v2Profile.DeviceProfileBasicInfo.Id,
			Name:         v2Profile.DeviceProfileBasicInfo.Name,
			Manufacturer: v2Profile.DeviceProfileBasicInfo.Manufacturer,
			Description:  v2Profile.DeviceProfileBasicInfo.Description,
			Model:        v2Profile.DeviceProfileBasicInfo.Model,
			Labels:       v2Profile.DeviceProfileBasicInfo.Labels,
		},
		DeviceResources: resources,
		DeviceCommands:  transformCommandFromV2ToV3(v2Profile.DeviceCommands),
		ApiVersion:      common.ApiVersion,
	}
	return profile, nil
}

// transformResourceFromV2ToV3 converts the v2 []DeviceResource DTO to v3
func transformResourceFromV2ToV3(v2Resources []v2dtos.DeviceResource) ([]dtos.DeviceResource, errors.EdgeX) {
	var deviceResources []dtos.DeviceResource
	for _, v2Res := range v2Resources {
		resProps, err := transformResPropsFromV2ToV3(v2Res.Properties)
		if err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindServerError, "failed to convert v2 ResourceProperties DTO to v3", err)
		}

		dr := dtos.DeviceResource{
			Description: v2Res.Description,
			Name:        v2Res.Name,
			IsHidden:    v2Res.IsHidden,
			Properties:  resProps,
			Attributes:  v2Res.Attributes,
			Tags:        v2Res.Tags,
			Tag:         v2Res.Tag,
		}
		deviceResources = append(deviceResources, dr)
	}
	return deviceResources, nil
}

// transformResPropsFromV2ToV3 converts the v2 ResourceProperties DTO to v3
func transformResPropsFromV2ToV3(v2ResProp v2dtos.ResourceProperties) (dtos.ResourceProperties, errors.EdgeX) {
	var err error
	var minimum, maximum, scale, offset, base float64
	var mask uint64
	var shift int64

	if v2ResProp.Minimum != "" {
		minimum, err = strconv.ParseFloat(v2ResProp.Minimum, 64)
		if err != nil {
			return dtos.ResourceProperties{}, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse minimum value '%s' to float64", v2ResProp.Minimum), err)
		}
	}
	if v2ResProp.Maximum != "" {
		maximum, err = strconv.ParseFloat(v2ResProp.Maximum, 64)
		if err != nil {
			return dtos.ResourceProperties{}, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse maximum value '%s' to float64", v2ResProp.Maximum), err)
		}
	}
	if v2ResProp.Mask != "" {
		mask, err = strconv.ParseUint(v2ResProp.Mask, 10, 64)
		if err != nil {
			return dtos.ResourceProperties{}, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse mask value '%s' to uint64", v2ResProp.Mask), err)
		}
	}
	if v2ResProp.Shift != "" {
		shift, err = strconv.ParseInt(v2ResProp.Shift, 10, 64)
		if err != nil {
			return dtos.ResourceProperties{}, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse shift value '%s' to int64", v2ResProp.Shift), err)
		}
	}
	if v2ResProp.Scale != "" {
		scale, err = strconv.ParseFloat(v2ResProp.Scale, 64)
		if err != nil {
			return dtos.ResourceProperties{}, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse scale value '%s' to float64", v2ResProp.Scale), err)
		}
	}
	if v2ResProp.Offset != "" {
		offset, err = strconv.ParseFloat(v2ResProp.Offset, 64)
		if err != nil {
			return dtos.ResourceProperties{}, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse offset value '%s' to float64", v2ResProp.Offset), err)
		}
	}
	if v2ResProp.Base != "" {
		base, err = strconv.ParseFloat(v2ResProp.Base, 64)
		if err != nil {
			return dtos.ResourceProperties{}, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse base value '%s' to float64", v2ResProp.Base), err)
		}
	}

	return dtos.ResourceProperties{
		ValueType:    v2ResProp.ValueType,
		ReadWrite:    v2ResProp.ReadWrite,
		Units:        v2ResProp.Units,
		Minimum:      &minimum,
		Maximum:      &maximum,
		DefaultValue: v2ResProp.DefaultValue,
		Mask:         &mask,
		Shift:        &shift,
		Scale:        &scale,
		Offset:       &offset,
		Base:         &base,
		Assertion:    v2ResProp.Assertion,
		MediaType:    v2ResProp.MediaType,
		Optional:     nil,
	}, nil
}

// transformCommandFromV2ToV3 converts the v2 []DeviceCommand DTO to v3
func transformCommandFromV2ToV3(v2Commands []v2dtos.DeviceCommand) []dtos.DeviceCommand {
	var deviceCommands []dtos.DeviceCommand
	for _, v2Command := range v2Commands {
		dc := dtos.DeviceCommand{
			Name:               v2Command.Name,
			IsHidden:           v2Command.IsHidden,
			ReadWrite:          v2Command.ReadWrite,
			ResourceOperations: transformResourceOperationFromV2ToV3(v2Command.ResourceOperations),
			Tags:               v2Command.Tags,
		}
		deviceCommands = append(deviceCommands, dc)
	}
	return deviceCommands
}

// transformResourceOperationFromV2ToV3 converts the v2 []ResourceOperation DTO to v3
func transformResourceOperationFromV2ToV3(v2ResOps []v2dtos.ResourceOperation) []dtos.ResourceOperation {
	var ros []dtos.ResourceOperation
	for _, v2ro := range v2ResOps {
		ro := dtos.ResourceOperation(v2ro)
		ros = append(ros, ro)
	}
	return ros
}
