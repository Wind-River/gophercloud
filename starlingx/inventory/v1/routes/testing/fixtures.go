/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/routes"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	RouteHerp = routes.Route{
		ID:            "dbacbb57-b3cd-4b9c-b365-0ecf5dec4d60",
		Network:       "0.0.0.0",
		Prefix:        0,
		Gateway:       "192.168.59.1",
		Metric:        1,
		InterfaceName: "vlan11",
		InterfaceUUID: "9a01253c-17bd-453b-8dcf-038da4636100",
	}
	RouteDerp = routes.Route{
		ID:            "354968fc-6f18-46dc-93a1-6118280e3cee",
		Network:       "::",
		Prefix:        0,
		Gateway:       "fd00:0:0:b::1",
		Metric:        1,
		InterfaceName: "vlan11",
		InterfaceUUID: "c461e2de-74b8-4c88-a456-ae310e469afd",
	}
)

const RouteListBody = `
{
    "routes": [
		{
			"interface_uuid": "9a01253c-17bd-453b-8dcf-038da4636100",
			"uuid": "dbacbb57-b3cd-4b9c-b365-0ecf5dec4d60",
			"metric": 1,
			"prefix": 0,
			"ifname": "vlan11",
			"gateway": "192.168.59.1",
			"network": "0.0.0.0"
		},
		{
			"interface_uuid": "c461e2de-74b8-4c88-a456-ae310e469afd",
			"uuid": "354968fc-6f18-46dc-93a1-6118280e3cee",
			"metric": 1,
			"prefix": 0,
			"ifname": "vlan11",
			"gateway": "fd00:0:0:b::1",
			"network": "::"
		}
	]
}
`

const RouteSingleBody = `
{
	"interface_uuid": "c461e2de-74b8-4c88-a456-ae310e469afd",
	"uuid": "354968fc-6f18-46dc-93a1-6118280e3cee",
	"metric": 1,
	"prefix": 0,
	"ifname": "vlan11",
	"gateway": "fd00:0:0:b::1",
	"network": "::"
}
`

func HandleRouteListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/routes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RouteListBody)
	})
}

func HandleRouteGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/routes/354968fc-6f18-46dc-93a1-6118280e3cee", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, RouteSingleBody)
	})
}

func HandleRouteDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/routes/354968fc-6f18-46dc-93a1-6118280e3cee", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleRouteCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/routes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
          "network": "0.0.0.0",
		  "prefix": 0,
		  "gateway": "192.168.59.1",
		  "interface_uuid": "9a01253c-17bd-453b-8dcf-038da4636100",
		  "metric": 1
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}
