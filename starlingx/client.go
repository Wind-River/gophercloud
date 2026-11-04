/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package starlingx

import "github.com/gophercloud/gophercloud"

func initClientOpts(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, clientType string) (*gophercloud.ServiceClient, error) {
	sc := new(gophercloud.ServiceClient)
	eo.ApplyDefaults(clientType)
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return sc, err
	}
	sc.ProviderClient = client
	sc.Endpoint = url
	sc.Type = clientType
	// Install FaultErrorContext on every StarlingX client built here, so error
	// responses carry the server's reason instead of a generic status message.
	// This covers both the sysinv (inventory) client and the nfv-vim
	// orchestration client under nfv/v1. parseFaultString handles the sysinv
	// {"error_message": "{...faultstring...}"} shape as well as the top-level
	// {"faultstring": ...} WSME body that nfv-vim returns, and falls back to the
	// default gophercloud message when the body carries no reason.
	sc.DefaultErrorContext = FaultErrorContext()
	return sc, nil
}

// NewInventoryV1 creates a ServiceClient that may be used with the v1 inventory
// package.
func NewInventoryV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "inventory")
}
