/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

/*
Package distro provides a client for integration with the export-distro service.
*/
package distro

import (
	"context"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// DistroClient defines the interface for interactions with the EdgeX Foundry export-distro service.
type DistroClient interface {
	// NotifyRegistrations facilitates several kinds of updates to registered export clients while the service is running.
	NotifyRegistrations(context.Context, models.NotifyUpdate) error
}

type distroRestClient struct {
	urlClient interfaces.URLClient
}

// NewDistroClient creates an instance of DistroClient
func NewDistroClient(urlClient interfaces.URLClient) DistroClient {
	return &distroRestClient{
		urlClient: urlClient,
	}
}

func (d *distroRestClient) NotifyRegistrations(ctx context.Context, update models.NotifyUpdate) error {
	return clients.UpdateRequest(ctx, "", update, d.urlClient)
}
