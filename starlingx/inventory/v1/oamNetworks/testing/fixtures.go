/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/oamNetworks"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var (
	OAMNetworkHerp = oamNetworks.OAMNetwork{
		UUID:           "fd5aaa82-b503-40e2-af45-9fc4411df7a0",
		OAMSubnet:      "10.10.20.0/24",
		OAMGatewayIP:   nil,
		OAMFloatingIP:  "10.10.20.2",
		OAMC0IP:        "10.10.20.3",
		OAMC1IP:        "10.10.20.4",
		OAMStartIP:     "10.10.20.1",
		OAMEndIP:       "10.10.20.254",
		RegionConfig:   false,
		ForISystemUUID: "607671a2-15a7-4f97-9297-c4e1804cde12",
		CreatedAt:      "2023-11-28T13:10:53.200531+00:00",
		Links: []gophercloud.Link{
			{
				Href: "http://192.168.204.2:6385/v1/iextoams/fd5aaa82-b503-40e2-af45-9fc4411df7a0",
				Rel:  "self",
			},
			{
				Href: "http://192.168.204.2:6385/iextoams/fd5aaa82-b503-40e2-af45-9fc4411df7a0",
				Rel:  "bookmark",
			},
		},
		UpdatedAt: nil,
	}
	OAMNetworkDerp = oamNetworks.OAMNetwork{
		UUID:          "727bd796-070f-40c2-8b9b-7ed674fd0fe7",
		OAMSubnet:     "10.10.20.0/24",
		OAMGatewayIP:  nil,
		OAMFloatingIP: "10.10.20.5",
		OAMC0IP:       "10.10.20.3",
		OAMC1IP:       "10.10.20.4",
	}
)

const OAMNetworkListBody = `
{
	"iextoams": [
		{
			"uuid": "fd5aaa82-b503-40e2-af45-9fc4411df7a0",
			"oam_subnet": "10.10.20.0/24",
			"oam_gateway_ip": null,
			"oam_floating_ip": "10.10.20.2",
			"oam_c0_ip": "10.10.20.3",
			"oam_c1_ip": "10.10.20.4",
			"oam_start_ip": "10.10.20.1",
			"oam_end_ip": "10.10.20.254",
			"region_config": false,
			"isystem_uuid": "607671a2-15a7-4f97-9297-c4e1804cde12",
			"links": [
				{
					"href": "http://192.168.204.2:6385/v1/iextoams/fd5aaa82-b503-40e2-af45-9fc4411df7a0",
					"rel": "self"
				}, {
					"href": "http://192.168.204.2:6385/iextoams/fd5aaa82-b503-40e2-af45-9fc4411df7a0",
					"rel": "bookmark"
				}
			],
			"created_at": "2023-11-28T13:10:53.200531+00:00",
			"updated_at": null
		}
	]
}
`

const SingleOAMNetworkBody = `
{
    "uuid": "727bd796-070f-40c2-8b9b-7ed674fd0fe7",
	"oam_subnet": "10.10.20.0/24",
	"oam_gateway_ip": null,
	"oam_floating_ip": "10.10.20.5",
	"oam_c0_ip": "10.10.20.3",
	"oam_c1_ip": "10.10.20.4"
}
`

func HandleOAMNetworkListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/iextoam", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, OAMNetworkListBody)
	})
}

func HandleOAMNetworkGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/iextoam/fd5aaa82-b503-40e2-af45-9fc4411df7a0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, SingleOAMNetworkBody)
	})
}

func HandleOAMNetworkUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/iextoam/727bd796-070f-40c2-8b9b-7ed674fd0fe7", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[ { "op": "replace", "path": "/oam_floating_ip", "value": "10.10.20.10" } ]`)
		fmt.Fprint(w, SingleOAMNetworkBody)
	})
}
