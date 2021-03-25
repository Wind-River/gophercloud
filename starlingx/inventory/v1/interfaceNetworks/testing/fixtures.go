/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/interfaceNetworks"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	InterfaceNetworkHerp = interfaceNetworks.InterfaceNetwork{
		UUID:          "170cff91-d763-4361-8417-287b2aa5c71e",
		ID:            2,
		NetworkUUID:   "7bbd6f31-bcd1-4c65-859d-162a6c118543",
		NetworkID:     5,
		NetworkType:   "cluster-host",
		NetworkName:   "cluster-host",
		InterfaceName: "lo",
		InterfaceUUID: "927087ba-8f6f-4d1f-8ca0-0cd53b01e5c6",
	}
	InterfaceNetworkDerp = interfaceNetworks.InterfaceNetwork{
		UUID:          "2afb1f71-b42e-403d-a0cd-f1138be8167e",
		ID:            3,
		NetworkUUID:   "7ace62b9-907a-4f35-b88b-97f829f66330",
		NetworkID:     3,
		NetworkType:   "oam",
		NetworkName:   "oam",
		InterfaceName: "enp0s3",
		InterfaceUUID: "fa7721c8-cb24-4f91-bb2e-736cabeff4e2",
	}
)

const InterfaceNetworkListBody = `
{
    "interface_networks": [
        {
            "forihostid": 1,
            "network_uuid": "7bbd6f31-bcd1-4c65-859d-162a6c118543",
            "uuid": "170cff91-d763-4361-8417-287b2aa5c71e",
            "network_id": 5,
            "network_type": "cluster-host",
            "ifname": "lo",
            "interface_uuid": "927087ba-8f6f-4d1f-8ca0-0cd53b01e5c6",
            "network_name": "cluster-host",
            "id": 2
        },
        {
            "forihostid": 1,
            "network_uuid": "7ace62b9-907a-4f35-b88b-97f829f66330",
            "uuid": "2afb1f71-b42e-403d-a0cd-f1138be8167e",
            "network_id": 3,
            "network_type": "oam",
            "ifname": "enp0s3",
            "interface_uuid": "fa7721c8-cb24-4f91-bb2e-736cabeff4e2",
            "network_name": "oam",
            "id": 3
        }
    ]
}
`

const InterfaceNetworkSingleBody = `
{
	"forihostid": 1,
	"network_uuid": "7ace62b9-907a-4f35-b88b-97f829f66330",
	"uuid": "2afb1f71-b42e-403d-a0cd-f1138be8167e",
	"network_id": 3,
	"network_type": "oam",
	"ifname": "enp0s3",
	"interface_uuid": "fa7721c8-cb24-4f91-bb2e-736cabeff4e2",
	"network_name": "oam",
	"id": 3
}
`

func HandleInterfaceNetworkListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/interface_networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, InterfaceNetworkListBody)
	})
}

func HandleInterfaceNetworkGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/interface_networks/2afb1f71-b42e-403d-a0cd-f1138be8167e", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, InterfaceNetworkSingleBody)
	})
}

func HandleInterfaceNetworkDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/interface_networks/2afb1f71-b42e-403d-a0cd-f1138be8167e", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleInterfaceNetworkCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/interface_networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
          "interface_uuid": "fa7721c8-cb24-4f91-bb2e-736cabeff4e2",
		  "network_uuid": "7ace62b9-907a-4f35-b88b-97f829f66330"
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}
