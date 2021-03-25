/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/osds"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	UpdatedAtHerp       = "2019-11-07T21:23:45.122436+00:00"
	JournalPathHerp     = "/dev/disk/by-path/pci-0000:00:01.1-ata-1.1-part2"
	JournalSizeHerp     = 1024
	JournalNodeHerp     = "/dev/sdb2"
	JournalLocationHerp = "0738a917-ef8d-4222-91f2-e9802b167261"

	UpdatedAtDerp       = "2019-11-07T21:23:45.122436+00:00"
	JournalPathDerp     = "/dev/disk/by-path/pci-0000:00:01.1-ata-1.1-part2"
	JournalSizeDerp     = 1024
	JournalNodeDerp     = "/dev/sdb2"
	JournalLocationDerp = "982acece-40ac-4df0-b87d-eda51425e7bc"

	OSDHerp = osds.OSD{
		ID:       "0738a917-ef8d-4222-91f2-e9802b167261",
		Function: "osd",
		HostID:   "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		DiskID:   "ec945089-f367-4a3c-b2cf-87d81a99320e",
		State:    "configured",
		TierName: "storage",
		TierUUID: "1249f869-1042-49de-9032-33f314bc46f7",
		JournalInfo: osds.JournalInfo{
			Path:     &JournalPathHerp,
			Size:     &JournalSizeHerp,
			Node:     &JournalNodeHerp,
			Location: &JournalLocationHerp,
		},
		CreatedAt: "2019-11-07T21:07:12.588595+00:00",
		UpdatedAt: &UpdatedAtHerp,
	}
	OSDDerp = osds.OSD{
		ID:       "dfa582bd-8f1d-4e8c-9c8a-1167243d592a",
		Function: "osd",
		HostID:   "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		DiskID:   "b4465948-b5b1-4ea2-b81f-cda62162e4be",
		State:    "configured",
		TierName: "storage",
		TierUUID: "313c44bb-959d-46cc-82f7-6df7c2d7b057",
		JournalInfo: osds.JournalInfo{
			Path:     &JournalPathDerp,
			Size:     &JournalSizeDerp,
			Node:     &JournalNodeDerp,
			Location: &JournalLocationDerp,
		},
		CreatedAt: "2019-11-07T21:07:12.588595+00:00",
		UpdatedAt: &UpdatedAtHerp,
	}
)

const OSDListBody = `
{
    "istors": [
        {
            "function": "osd",
            "uuid": "0738a917-ef8d-4222-91f2-e9802b167261",
            "journal_size_mib": 1024,
            "journal_path": "/dev/disk/by-path/pci-0000:00:01.1-ata-1.1-part2",
            "created_at": "2019-11-07T21:07:12.588595+00:00",
            "updated_at": "2019-11-07T21:23:45.122436+00:00",
            "capabilities": {},
            "journal_node": "/dev/sdb2",
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "state": "configured",
            "idisk_uuid": "ec945089-f367-4a3c-b2cf-87d81a99320e",
            "tier_name": "storage",
            "osdid": 0,
            "tier_uuid": "1249f869-1042-49de-9032-33f314bc46f7",
            "journal_location": "0738a917-ef8d-4222-91f2-e9802b167261"
        },
        {
            "function": "osd",
            "uuid": "dfa582bd-8f1d-4e8c-9c8a-1167243d592a",
            "journal_size_mib": 1024,
            "journal_path": "/dev/disk/by-path/pci-0000:00:01.1-ata-1.1-part2",
            "created_at": "2019-11-07T21:07:12.588595+00:00",
            "updated_at": "2019-11-07T21:23:45.122436+00:00",
            "capabilities": {},
            "journal_node": "/dev/sdb2",
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "state": "configured",
            "idisk_uuid": "b4465948-b5b1-4ea2-b81f-cda62162e4be",
            "tier_name": "storage",
            "osdid": 0,
            "tier_uuid": "313c44bb-959d-46cc-82f7-6df7c2d7b057",
            "journal_location": "982acece-40ac-4df0-b87d-eda51425e7bc"
        }
    ]
}
`

const OSDSingleBody = `
{
	"function": "osd",
	"uuid": "dfa582bd-8f1d-4e8c-9c8a-1167243d592a",
	"journal_size_mib": 1024,
	"journal_path": "/dev/disk/by-path/pci-0000:00:01.1-ata-1.1-part2",
	"created_at": "2019-11-07T21:07:12.588595+00:00",
	"updated_at": "2019-11-07T21:23:45.122436+00:00",
	"capabilities": {},
	"journal_node": "/dev/sdb2",
	"ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
	"state": "configured",
	"idisk_uuid": "b4465948-b5b1-4ea2-b81f-cda62162e4be",
	"tier_name": "storage",
	"osdid": 0,
	"tier_uuid": "313c44bb-959d-46cc-82f7-6df7c2d7b057",
	"journal_location": "982acece-40ac-4df0-b87d-eda51425e7bc"
}
`

func HandleOSDListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/istors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, OSDListBody)
	})
}

func HandleOSDGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/istors/dfa582bd-8f1d-4e8c-9c8a-1167243d592a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, OSDSingleBody)
	})
}

func HandleOSDDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/istors/dfa582bd-8f1d-4e8c-9c8a-1167243d592a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleOSDCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/istors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"function": "osd",
			"idisk_uuid": "b4465948-b5b1-4ea2-b81f-cda62162e4be",
			"ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
			"journal_location": "982acece-40ac-4df0-b87d-eda51425e7bc",
			"journal_size_mib": 1048576,
			"tier_uuid": "313c44bb-959d-46cc-82f7-6df7c2d7b057"
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}

func HandleOSDUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/istors/dfa582bd-8f1d-4e8c-9c8a-1167243d592a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[
			{"op": "replace", "path": "/journal_location", "value": "bb5c7623-0eb2-4581-9942-7abac4a4be86"},
          	{"op": "replace", "path": "/journal_size_mib", "value": 2097152}]`)

		fmt.Fprintf(w, OSDSingleBody)
	})
}
