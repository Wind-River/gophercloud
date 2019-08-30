/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/interfaceDataNetworks"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	InterfaceDataNetworkHerp = interfaceDataNetworks.InterfaceDataNetwork{
		UUID:            "b1579be2-474b-4c44-b034-f7bfb7fd84b7",
		ID:              1,
		DataNetworkUUID: "c71a1cfa-4bf9-48fc-a8ed-70f8797fc442",
		DataNetworkID:   1,
		DataNetworkName: "physnet0",
		InterfaceName:   "data0",
		InterfaceUUID:   "2b9fba21-7b7c-42c1-bcab-50841370f049",
	}
	InterfaceDataNetworkDerp = interfaceDataNetworks.InterfaceDataNetwork{
		UUID:            "0038c026-ef30-4bca-8931-5d7255ebd34e",
		ID:              2,
		DataNetworkUUID: "05331897-b9ab-40a3-8029-b87099a0190f",
		DataNetworkID:   2,
		DataNetworkName: "physnet1",
		InterfaceName:   "data1",
		InterfaceUUID:   "953daf34-af6f-4bd9-a37f-1f1fd2a9cce5",
	}
)

const InterfaceDataNetworkListBody = `
{
    "interface_datanetworks": [
        {
            "datanetwork_uuid": "c71a1cfa-4bf9-48fc-a8ed-70f8797fc442",
            "forihostid": 1,
            "datanetwork_id": 1,
            "uuid": "b1579be2-474b-4c44-b034-f7bfb7fd84b7",
            "datanetwork_name": "physnet0",
            "ifname": "data0",
            "interface_uuid": "2b9fba21-7b7c-42c1-bcab-50841370f049",
            "id": 1
        },
        {
            "datanetwork_uuid": "05331897-b9ab-40a3-8029-b87099a0190f",
            "forihostid": 1,
            "datanetwork_id": 2,
            "uuid": "0038c026-ef30-4bca-8931-5d7255ebd34e",
            "datanetwork_name": "physnet1",
            "ifname": "data1",
            "interface_uuid": "953daf34-af6f-4bd9-a37f-1f1fd2a9cce5",
            "id": 2
        }
    ]
}
`

const InterfaceDataNetworkSingleBody = `
{
	"datanetwork_uuid": "05331897-b9ab-40a3-8029-b87099a0190f",
	"forihostid": 1,
	"datanetwork_id": 2,
	"uuid": "0038c026-ef30-4bca-8931-5d7255ebd34e",
	"datanetwork_name": "physnet1",
	"ifname": "data1",
	"interface_uuid": "953daf34-af6f-4bd9-a37f-1f1fd2a9cce5",
	"id": 2
}
`

func HandleInterfaceDataNetworkListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/interface_datanetworks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, InterfaceDataNetworkListBody)
	})
}

func HandleInterfaceDataNetworkGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/interface_datanetworks/0038c026-ef30-4bca-8931-5d7255ebd34e", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, InterfaceDataNetworkSingleBody)
	})
}

func HandleInterfaceDataNetworkDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/interface_datanetworks/0038c026-ef30-4bca-8931-5d7255ebd34e", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleInterfaceDataNetworkCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/interface_datanetworks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
          "interface_uuid": "953daf34-af6f-4bd9-a37f-1f1fd2a9cce5",
		  "datanetwork_uuid": "05331897-b9ab-40a3-8029-b87099a0190f"
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}
