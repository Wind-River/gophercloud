/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/storagetiers"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	UpdatedAtHerp   = "2019-11-07T20:52:23.144016+00:00"
	UpdatedAtDerp   = "2019-11-07T20:52:23.144016+00:00"
	StorageTierHerp = storagetiers.StorageTier{
		ID:        "1249f869-1042-49de-9032-33f314bc46f7",
		Name:      "storage",
		Type:      "ceph",
		Status:    "in-use",
		Stors:     []int{0},
		ClusterID: "e7d497b9-085c-4a46-bfff-21d02ad9313a",
		BackendID: "b8d67a89-e604-4e32-a729-26b8ce565bfc",
		CreatedAt: "2019-11-07T20:47:54.669165+00:00",
		UpdatedAt: &UpdatedAtHerp,
	}
	StorageTierDerp = storagetiers.StorageTier{
		ID:        "4da7bca3-a7ea-4042-99ac-6f74f27c50b3",
		Name:      "storage",
		Type:      "ceph",
		Status:    "in-use",
		Stors:     []int{0},
		ClusterID: "e7d497b9-085c-4a46-bfff-21d02ad9313a",
		BackendID: "515380fa-5646-405d-bbdf-cce2d7761349",
		CreatedAt: "2019-11-07T20:47:54.669165+00:00",
		UpdatedAt: &UpdatedAtDerp,
	}
)

const StorageTierListBody = `
{
    "storage_tiers": [
        {
            "status": "in-use",
            "uuid": "1249f869-1042-49de-9032-33f314bc46f7",
            "stors": [
                0
            ],
            "created_at": "2019-11-07T20:47:54.669165+00:00",
            "updated_at": "2019-11-07T20:52:23.144016+00:00",
            "capabilities": {},
            "cluster_uuid": "e7d497b9-085c-4a46-bfff-21d02ad9313a",
            "backend_uuid": "b8d67a89-e604-4e32-a729-26b8ce565bfc",
            "type": "ceph",
            "name": "storage"
        },
        {
            "status": "in-use",
            "uuid": "4da7bca3-a7ea-4042-99ac-6f74f27c50b3",
            "stors": [
                0
            ],
            "created_at": "2019-11-07T20:47:54.669165+00:00",
            "updated_at": "2019-11-07T20:52:23.144016+00:00",
            "capabilities": {},
            "cluster_uuid": "e7d497b9-085c-4a46-bfff-21d02ad9313a",
            "backend_uuid": "515380fa-5646-405d-bbdf-cce2d7761349",
            "type": "ceph",
            "name": "storage"
        }
    ]
}
`

const StorageTierSingleBody = `
{
	"status": "in-use",
	"uuid": "4da7bca3-a7ea-4042-99ac-6f74f27c50b3",
	"stors": [
		0
	],
	"created_at": "2019-11-07T20:47:54.669165+00:00",
	"updated_at": "2019-11-07T20:52:23.144016+00:00",
	"capabilities": {},
	"cluster_uuid": "e7d497b9-085c-4a46-bfff-21d02ad9313a",
	"backend_uuid": "515380fa-5646-405d-bbdf-cce2d7761349",
	"type": "ceph",
	"name": "storage"
}
`

func HandleStorageTierListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/clusters/e7d497b9-085c-4a46-bfff-21d02ad9313a/storage_tiers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, StorageTierListBody)
	})
}

func HandleStorageTierGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/storage_tiers/4da7bca3-a7ea-4042-99ac-6f74f27c50b3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, StorageTierSingleBody)
	})
}

func HandleStorageTierDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/storage_tiers/4da7bca3-a7ea-4042-99ac-6f74f27c50b3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleStorageTierCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/storage_tiers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
		  "name": "test",
		  "cluster_uuid": "e7d497b9-085c-4a46-bfff-21d02ad9313a"
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}

func HandleStorageTierUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/storage_tiers/4da7bca3-a7ea-4042-99ac-6f74f27c50b3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[ {"op": "replace", "path": "/cluster_uuid", "value": "9db6c194-dc8b-4c9e-8a0e-a3d01fa1f5f1"},{ "op": "replace", "path": "/name", "value": "edited" } ]`)

		fmt.Fprintf(w, StorageTierSingleBody)
	})
}
