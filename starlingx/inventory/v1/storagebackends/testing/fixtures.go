/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/storagebackends"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	StorageBackendHerp = storagebackends.StorageBackend{
		ID:        "da0dd822-5b41-4402-a735-6b0f02a9d953",
		Backend:   "external",
		Name:      "shared_services",
		State:     "configured",
		Services:  "glance",
		CreatedAt: "2019-11-11T17:03:01.304313+00:00",
	}
	StorageBackendDerp = storagebackends.StorageBackend{
		ID:        "cebe7a5e-7b57-497b-a335-6e7cf93e98ee",
		Backend:   "ceph",
		Name:      "ceph-store",
		State:     "configured",
		Task:      "provision-storage",
		CreatedAt: "2019-11-11T17:03:36.193041+00:00",
	}
)

const StorageBackendListBody = `
{
    "storage_backends": [
        {
            "task": null,
            "uuid": "da0dd822-5b41-4402-a735-6b0f02a9d953",
            "created_at": "2019-11-11T17:03:01.304313+00:00",
            "updated_at": null,
            "capabilities": {},
            "services": "glance",
            "state": "configured",
            "isystem_uuid": null,
            "backend": "external",
            "name": "shared_services"
        },
        {
            "task": "provision-storage",
            "uuid": "cebe7a5e-7b57-497b-a335-6e7cf93e98ee",
            "created_at": "2019-11-11T17:03:36.193041+00:00",
            "updated_at": null,
            "capabilities": {
                "min_replication": "1",
                "replication": "2"
            },
            "services": null,
            "state": "configured",
            "isystem_uuid": "cc0149fb-b40d-4ff1-9dd9-2070a05aee74",
            "backend": "ceph",
            "name": "ceph-store"
        }
    ]
}
`

const StorageBackendSingleBody = `
{
	"task": "provision-storage",
	"uuid": "cebe7a5e-7b57-497b-a335-6e7cf93e98ee",
	"created_at": "2019-11-11T17:03:36.193041+00:00",
	"updated_at": null,
	"capabilities": {
		"min_replication": "1",
		"replication": "2"
	},
	"services": null,
	"state": "configured",
	"isystem_uuid": "cc0149fb-b40d-4ff1-9dd9-2070a05aee74",
	"backend": "ceph",
	"name": "ceph-store"
}
`

func HandleStorageBackendListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/storage_backend", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, StorageBackendListBody)
	})
}

func HandleStorageBackendGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/storage_backend/cebe7a5e-7b57-497b-a335-6e7cf93e98ee", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, StorageBackendSingleBody)
	})
}

func HandleStorageBackendDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/storage_backend/cebe7a5e-7b57-497b-a335-6e7cf93e98ee", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleStorageBackendCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/storage_backend", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
<<<<<<< HEAD
			"backend": "ceph",
			"name": "ceph-store",
			"capabilities": {},
			"confirmed": false
		}`)
=======
        	"backend": "ceph",
        	"name": "ceph-store",
        	"capabilities": {},
        	"confirmed": false
        }`)
>>>>>>> 08dba1af3efc54755bf0f357789e5b1b558e78ec

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}

func HandleStorageBackendUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/storage_backend/cebe7a5e-7b57-497b-a335-6e7cf93e98ee", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[
			{"path": "/capabilities", "value": {}, "op": "replace"},
			{"op": "replace", "path": "/confirmed", "value": false}
		]`)

		fmt.Fprintf(w, StorageBackendSingleBody)
	})
}
