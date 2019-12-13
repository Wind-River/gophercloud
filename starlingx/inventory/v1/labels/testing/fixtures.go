/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/labels"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	LabelHerp = labels.Label{
		ID:       "b38aa42e-89e7-4fa0-8e3e-7d656962cdf6",
		HostUUID: "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		Key:      "openstack-control-plane",
		Value:    "enabled",
	}
	LabelDerp = labels.Label{
		ID:       "c0940c5f-dedf-4010-9cca-4d9b7f5dacec",
		HostUUID: "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		Key:      "openstack-compute-node",
		Value:    "enabled",
	}
)

const LabelListBody = `
{
    "labels": [
        {
            "label_value": "enabled",
            "uuid": "b38aa42e-89e7-4fa0-8e3e-7d656962cdf6",
            "host_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "label_key": "openstack-control-plane"
        },
        {
            "label_value": "enabled",
            "uuid": "c0940c5f-dedf-4010-9cca-4d9b7f5dacec",
            "host_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "label_key": "openstack-compute-node"
        }
    ]
}
`

const LabelSingleBody = `
{
	"label_value": "enabled",
	"uuid": "c0940c5f-dedf-4010-9cca-4d9b7f5dacec",
	"host_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
	"label_key": "openstack-compute-node"
}
`

func HandleLabelListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/labels", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, LabelListBody)
	})
}

func HandleLabelGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/labels/c0940c5f-dedf-4010-9cca-4d9b7f5dacec", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, LabelSingleBody)
	})
}

func HandleLabelDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/labels/c0940c5f-dedf-4010-9cca-4d9b7f5dacec", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleLabelAssignSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/labels/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"openstack-compute-node": "enabled"
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}
