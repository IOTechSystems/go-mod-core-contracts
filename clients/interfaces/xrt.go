// Copyright (C) 2023 IOTech Ltd

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/xrtmodels"
)

// XrtClient defines the interface for interactions with the XRT MQTT Management API.
type XrtClient interface {
	DeviceByName(ctx context.Context, name string) (xrtmodels.DeviceInfo, errors.EdgeX)
	DeviceProfileByName(ctx context.Context, name string) (models.DeviceProfile, errors.EdgeX)
}
