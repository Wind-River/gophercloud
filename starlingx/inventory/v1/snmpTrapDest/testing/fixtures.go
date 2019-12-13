/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/snmpTrapDest"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	SNMPTrapDestHerp = snmpTrapDest.SNMPTrapDest{
		ID:        "e9fb08f0-0795-4db2-8c22-b45b57fb9e39",
		Community: "638112d6-49b5-43ed-8f9a-af4c239bfaaf",
		IPAddress: "192.168.2.1",
		Type:      "snmpv2c_trap",
		Port:      162,
		Transport: "udp",
	}
	SNMPTrapDestDerp = snmpTrapDest.SNMPTrapDest{
		ID:        "e30da84e-4e9f-4b6e-969f-2a4c41345e47",
		Community: "5144712c-3133-4a91-b3b7-8767da6f0680",
		IPAddress: "192.168.2.100",
		Type:      "snmpv2c_trap",
		Port:      162,
		Transport: "udp",
	}
)

const SNMPTrapDestListBody = `
{
    "itrapdest": [
        {
            "uuid": "e9fb08f0-0795-4db2-8c22-b45b57fb9e39",
            "ip_address": "192.168.2.1",
            "community": "638112d6-49b5-43ed-8f9a-af4c239bfaaf",
            "type": "snmpv2c_trap",
            "port": 162,
            "transport": "udp"
        },
        {
            "uuid": "e30da84e-4e9f-4b6e-969f-2a4c41345e47",
            "ip_address": "192.168.2.100",
            "community": "5144712c-3133-4a91-b3b7-8767da6f0680",
            "type": "snmpv2c_trap",
            "port": 162,
            "transport": "udp"
        }
    ]
}
`

const SNMPTrapDestSingleBody = `
{
	"uuid": "e30da84e-4e9f-4b6e-969f-2a4c41345e47",
	"ip_address": "192.168.2.100",
	"community": "5144712c-3133-4a91-b3b7-8767da6f0680",
	"type": "snmpv2c_trap",
	"port": 162,
	"transport": "udp"
}
`

func HandleSNMPTrapDestListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/itrapdest", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, SNMPTrapDestListBody)
	})
}

func HandleSNMPTrapDestGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/itrapdest/e30da84e-4e9f-4b6e-969f-2a4c41345e47", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, SNMPTrapDestSingleBody)
	})
}

func HandleSNMPTrapDestDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/itrapdest/e30da84e-4e9f-4b6e-969f-2a4c41345e47", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleSNMPTrapDestCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/itrapdest", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
          "ip_address": "192.168.2.100",
		  "community": "5144712c-3133-4a91-b3b7-8767da6f0680"
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}
