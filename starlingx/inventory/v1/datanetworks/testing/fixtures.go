/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/datanetworks"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	DataNetworkHerp = datanetworks.DataNetwork{
		ID:        "05331897-b9ab-40a3-8029-b87099a0190f",
		Name:      "physnet1",
		Type:      "vlan",
		MTU:       1500,
		CreatedAt: "2019-11-07T21:06:40.787345+00:00",
	}
	DataNetworkDerp = datanetworks.DataNetwork{
		ID:        "c71a1cfa-4bf9-48fc-a8ed-70f8797fc442",
		Name:      "physnet0",
		Type:      "vlan",
		MTU:       1500,
		CreatedAt: "2019-11-07T21:06:37.231386+00:00",
	}
)

const DataNetworkListBody = `
{
	"datanetworks":[
		{
			"description": null,
			"updated_at": null,
			"created_at": "2019-11-07T21:06:40.787345+00:00",
			"port_num": null,
			"uuid": "05331897-b9ab-40a3-8029-b87099a0190f",
			"mtu": 1500,
			"multicast_group": null,
			"mode": null,
			"ttl": null,
			"id": 2,
			"network_type": "vlan",
			"name": "physnet1"
		},
		{
			"description": null,
			"updated_at": null,
			"created_at": "2019-11-07T21:06:37.231386+00:00",
			"port_num": null,
			"uuid": "c71a1cfa-4bf9-48fc-a8ed-70f8797fc442",
			"mtu": 1500,
			"multicast_group": null,
			"mode": null,
			"ttl": null,
			"id": 1,
			"network_type": "vlan",
			"name": "physnet0"
		}
	]
}
`

const DataNetworkSingleBody = `
{
	"description": null,
	"updated_at": null,
	"created_at": "2019-11-07T21:06:37.231386+00:00",
	"port_num": null,
	"uuid": "c71a1cfa-4bf9-48fc-a8ed-70f8797fc442",
	"mtu": 1500,
	"multicast_group": null,
	"mode": null,
	"ttl": null,
	"id": 1,
	"network_type": "vlan",
	"name": "physnet0"
}
`

func HandleDataNetworkListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/datanetworks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DataNetworkListBody)
	})
}

func HandleDataNetworkGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/datanetworks/c71a1cfa-4bf9-48fc-a8ed-70f8797fc442", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, DataNetworkSingleBody)
	})
}

func HandleDataNetworkDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/datanetworks/c71a1cfa-4bf9-48fc-a8ed-70f8797fc442", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleDataNetworkCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/datanetworks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
          "network_type": "vlan",
		  "name": "physnet0",
		  "mtu": 1500
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}

func HandleDataNetworkUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/datanetworks/c71a1cfa-4bf9-48fc-a8ed-70f8797fc442", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[ { "op": "replace", "path": "/mtu", "value": 2000 } ]`)

		fmt.Fprintf(w, DataNetworkSingleBody)
	})
}
