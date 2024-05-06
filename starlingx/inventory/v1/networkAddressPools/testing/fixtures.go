/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2024 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/networkAddressPools"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var (
	NetworkAddressPoolHerp = networkAddressPools.NetworkAddressPool{
		UUID:            "11111111-a6e5-425e-9317-995da88d6694",
		NetworkUUID:     "11111111-0000-425e-9317-995da88d6694",
		AddressPoolUUID: "11111111-1111-425e-9317-995da88d6694",
		NetworkName:     "oam",
		AddressPoolName: "oam-ipv4",
	}

	NetworkAddressPoolDerp = networkAddressPools.NetworkAddressPool{
		UUID:            "22222222-a6e5-425e-9317-995da88d6694",
		NetworkUUID:     "22222222-0000-425e-9317-995da88d6694",
		AddressPoolUUID: "22222222-1111-425e-9317-995da88d6694",
		NetworkName:     "oam",
		AddressPoolName: "oam-ipv6",
	}
)

const NetworkAddressPoolListBody = `
{
    "network_addresspools": [
        {
			"uuid": "11111111-a6e5-425e-9317-995da88d6694",
			"network_uuid": "11111111-0000-425e-9317-995da88d6694",
			"address_pool_uuid": "11111111-1111-425e-9317-995da88d6694",
			"network_name": "oam",
			"address_pool_name": "oam-ipv4"
		},
		{
			"uuid": "22222222-a6e5-425e-9317-995da88d6694",
			"network_uuid": "22222222-0000-425e-9317-995da88d6694",
			"address_pool_uuid": "22222222-1111-425e-9317-995da88d6694",
			"network_name": "oam",
			"address_pool_name": "oam-ipv6"
		}
    ]
}
`

const SingleNetworkAddressPoolBody = `
{
	"uuid": "11111111-a6e5-425e-9317-995da88d6694",
	"network_uuid": "11111111-0000-425e-9317-995da88d6694",
	"address_pool_uuid": "11111111-1111-425e-9317-995da88d6694",
	"network_name": "oam",
	"address_pool_name": "oam-ipv4"
}
`

func HandleNetworkAddressPoolListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/network_addresspools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, NetworkAddressPoolListBody)
	})
}

func HandleNetworkAddressPoolGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/network_addresspools/11111111-a6e5-425e-9317-995da88d6694", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, SingleNetworkAddressPoolBody)
	})
}

func HandleNetworkAddressPoolDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/network_addresspools/11111111-a6e5-425e-9317-995da88d6694", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleNetworkAddressPoolCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/network_addresspools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
          "network_uuid": "11111111-0000-425e-9317-995da88d6694",
          "network_name": "oam",
          "address_pool_uuid": "11111111-1111-425e-9317-995da88d6694",
          "address_pool_name": "oam-ipv4"
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}
