/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/snmpCommunity"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	SNMPCommunityHerp = snmpCommunity.SNMPCommunity{
		ID:        "638112d6-49b5-43ed-8f9a-af4c239bfaaf",
		Community: "test",
		Access:    "ro",
		View:      ".1",
	}
	SNMPCommunityDerp = snmpCommunity.SNMPCommunity{
		ID:        "5144712c-3133-4a91-b3b7-8767da6f0680",
		Community: "test2",
		Access:    "ro",
		View:      ".1",
	}
)

const SNMPCommunityListBody = `
{
    "icommunity": [
        {
            "access": "ro",
            "uuid": "638112d6-49b5-43ed-8f9a-af4c239bfaaf",
            "community": "test",
            "view": ".1"
        },
        {
            "access": "ro",
            "uuid": "5144712c-3133-4a91-b3b7-8767da6f0680",
            "community": "test2",
            "view": ".1"
        }
    ]
}
`

const SNMPCommunitySingleBody = `
{
	"access": "ro",
	"uuid": "5144712c-3133-4a91-b3b7-8767da6f0680",
	"community": "test2",
	"view": ".1"
}
`

func HandleSNMPCommunityListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/icommunity", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, SNMPCommunityListBody)
	})
}

func HandleSNMPCommunityGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/icommunity/5144712c-3133-4a91-b3b7-8767da6f0680", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, SNMPCommunitySingleBody)
	})
}

func HandleSNMPCommunityDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/icommunity/5144712c-3133-4a91-b3b7-8767da6f0680", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleSNMPCommunityCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/icommunity", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"community": "test2"
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}
