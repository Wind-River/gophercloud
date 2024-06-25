/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/addresspools"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	AddressPoolHerp = addresspools.AddressPool{
		ID:                 "727bd796-070f-40c2-8b9b-7ed674fd0fe7",
		Name:               "management",
		Prefix:             24,
		Network:            "192.168.204.0",
		Gateway:            nil,
		Order:              "random",
		Ranges:             [][]string{{"192.268.206.100", "192.168.204.254"}},
		FloatingAddress:    "192.168.204.2",
		Controller0Address: "192.168.204.3",
		Controller1Address: "192.168.204.4",
	}
	AddressPoolDerp = addresspools.AddressPool{
		ID:                 "123914e3-36e4-41a8-a702-d9f6e54d7140",
		Name:               "pxeboot",
		Prefix:             24,
		Network:            "169.254.202.0",
		Gateway:            nil,
		Order:              "random",
		Ranges:             [][]string{{"169.254.202.1", "169.254.202.254"}},
		FloatingAddress:    "169.254.202.2",
		Controller0Address: "169.254.202.3",
		Controller1Address: "169.254.202.4",
	}
)

const AddressPoolListBody = `
{
    "addrpools": [
        {
            "gateway_address": null,
            "network": "192.168.204.0",
            "name": "management",
            "ranges": [
                [
                    "192.268.206.100",
                    "192.168.204.254"
                ]
            ],
			"floating_address": "192.168.204.2",
			"controller0_address": "192.168.204.3",
			"controller1_address": "192.168.204.4",
            "prefix": 24,
            "order": "random",
            "uuid": "727bd796-070f-40c2-8b9b-7ed674fd0fe7"
        },
        {
            "gateway_address": null,
            "network": "169.254.202.0",
            "name": "pxeboot",
            "ranges": [
                [
                    "169.254.202.1",
                    "169.254.202.254"
                ]
            ],
			"floating_address": "169.254.202.2",
			"controller0_address": "169.254.202.3",
			"controller1_address": "169.254.202.4",
            "prefix": 24,
            "order": "random",
            "uuid": "123914e3-36e4-41a8-a702-d9f6e54d7140"
        }
    ]
}
`

const AddressPoolSingleBody = `
{
	"gateway_address": null,
	"network": "169.254.202.0",
	"name": "pxeboot",
	"ranges": [
		[
			"169.254.202.1",
			"169.254.202.254"
		]
	],
	"floating_address": "169.254.202.2",
	"controller0_address": "169.254.202.3",
	"controller1_address": "169.254.202.4",
	"prefix": 24,
	"order": "random",
	"uuid": "123914e3-36e4-41a8-a702-d9f6e54d7140"
}
`

func HandleAddressPoolListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/addrpools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, AddressPoolListBody)
	})
}

func HandleAddressPoolGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/addrpools/123914e3-36e4-41a8-a702-d9f6e54d7140", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, AddressPoolSingleBody)
	})
}

func HandleAddressPoolDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/addrpools/123914e3-36e4-41a8-a702-d9f6e54d7140", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleAddressPoolCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/addrpools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"network": "169.254.202.0",
			"name": "pxeboot",
			"ranges": [
				[
					"169.254.202.1",
					"169.254.202.254"
				]
			],
			"floating_address": "169.254.202.2",
			"controller0_address": "169.254.202.3",
			"controller1_address": "169.254.202.4",		
			"prefix": 24,
			"order": "random"
		}`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}

func HandleAddressPoolUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/addrpools/123914e3-36e4-41a8-a702-d9f6e54d7140", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[ { "op": "replace", "path": "/name", "value": "Changed" }, {"op": "replace", "path": "/ranges", "value": [["169.254.202.1", "169.254.202.200"]]} ]`)

		fmt.Fprintf(w, AddressPoolSingleBody)
	})
}
