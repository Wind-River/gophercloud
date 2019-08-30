/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/volumegroups"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	UpdatedAtHerp   = "2019-11-14T21:04:50.007663+00:00"
	UpdatedAtDerp   = "2019-11-14T21:04:50.029120+00:00"
	VolumeGroupHerp = volumegroups.VolumeGroup{
		ID:     "40cc6c66-fc2d-45f4-80d8-f7d218a4cb62",
		HostID: "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		State:  "provisioned",
		LVMInfo: volumegroups.LVMInfo{
			Name:                   "cgts-vg",
			GroupUUID:              "oMeU58-YxJj-yI2j-vbJV-lqi0-gZnn-7Ai1cx",
			Access:                 "wz--n-",
			Size:                   245853323264,
			AvailableSize:          79624667136,
			TotalPE:                7327,
			FreePE:                 2373,
			CurrentLogicalVolumes:  12,
			CurrentPhysicalVolumes: 1,
			MaximumPhysicalVolumes: 0,
		},
		CreatedAt: "2019-11-07T20:52:11.779459+00:00",
		UpdatedAt: &UpdatedAtHerp,
	}
	VolumeGroupDerp = volumegroups.VolumeGroup{
		ID:     "449aee64-342f-4255-9a23-b229b0589c1b",
		HostID: "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		State:  "provisioned",
		LVMInfo: volumegroups.LVMInfo{
			Name:                   "nova-local",
			GroupUUID:              "VI0rH6-sFJr-SaGV-QfmU-ogMu-UKMa-g45SjM",
			Access:                 "wz--n-",
			Size:                   36503027712,
			AvailableSize:          0,
			TotalPE:                8703,
			FreePE:                 0,
			CurrentLogicalVolumes:  1,
			CurrentPhysicalVolumes: 1,
			MaximumPhysicalVolumes: 0,
		},
		CreatedAt: "2019-11-07T21:07:46.112632+00:00",
		UpdatedAt: &UpdatedAtDerp,
	}
)

const VolumeGroupListBody = `
{
    "ilvgs": [
        {
            "lvm_vg_access": "wz--n-",
            "lvm_vg_size": 245853323264,
            "lvm_max_lv": 0,
            "lvm_vg_free_pe": 2373,
            "uuid": "40cc6c66-fc2d-45f4-80d8-f7d218a4cb62",
            "lvm_cur_lv": 12,
            "created_at": "2019-11-07T20:52:11.779459+00:00",
            "lvm_max_pv": 0,
            "updated_at": "2019-11-14T21:04:50.007663+00:00",
            "capabilities": {},
            "vg_state": "provisioned",
            "lvm_vg_avail_size": 79624667136,
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "lvm_cur_pv": 1,
            "lvm_vg_uuid": "oMeU58-YxJj-yI2j-vbJV-lqi0-gZnn-7Ai1cx",
            "lvm_vg_total_pe": 7327,
            "lvm_vg_name": "cgts-vg"
        },
        {
            "lvm_vg_access": "wz--n-",
            "lvm_vg_size": 36503027712,
            "lvm_max_lv": 0,
            "lvm_vg_free_pe": 0,
            "uuid": "449aee64-342f-4255-9a23-b229b0589c1b",
            "lvm_cur_lv": 1,
            "created_at": "2019-11-07T21:07:46.112632+00:00",
            "lvm_max_pv": 0,
            "updated_at": "2019-11-14T21:04:50.029120+00:00",
            "capabilities": {},
            "vg_state": "provisioned",
            "lvm_vg_avail_size": 0,
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "lvm_cur_pv": 1,
            "lvm_vg_uuid": "VI0rH6-sFJr-SaGV-QfmU-ogMu-UKMa-g45SjM",
            "lvm_vg_total_pe": 8703,
            "lvm_vg_name": "nova-local"
        }
    ]
}
`

const VolumeGroupSingleBody = `
{
	"lvm_vg_access": "wz--n-",
	"lvm_vg_size": 36503027712,
	"lvm_max_lv": 0,
	"lvm_vg_free_pe": 0,
	"uuid": "449aee64-342f-4255-9a23-b229b0589c1b",
	"lvm_cur_lv": 1,
	"created_at": "2019-11-07T21:07:46.112632+00:00",
	"lvm_max_pv": 0,
	"updated_at": "2019-11-14T21:04:50.029120+00:00",
	"capabilities": {},
	"vg_state": "provisioned",
	"lvm_vg_avail_size": 0,
	"ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
	"lvm_cur_pv": 1,
	"lvm_vg_uuid": "VI0rH6-sFJr-SaGV-QfmU-ogMu-UKMa-g45SjM",
	"lvm_vg_total_pe": 8703,
	"lvm_vg_name": "nova-local"
}
`

func HandleVolumeGroupListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/ilvgs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, VolumeGroupListBody)
	})
}

func HandleVolumeGroupGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ilvgs/449aee64-342f-4255-9a23-b229b0589c1b", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, VolumeGroupSingleBody)
	})
}

func HandleVolumeGroupDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ilvgs/449aee64-342f-4255-9a23-b229b0589c1b", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleVolumeGroupCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/ilvgs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
          "lvm_vg_name": "nova-local",
		  "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e"
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}

func HandleVolumeGroupUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ilvgs/449aee64-342f-4255-9a23-b229b0589c1b", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[ { "path": "/capabilities", "value": {}, "op": "replace"} ]`)

		fmt.Fprintf(w, VolumeGroupSingleBody)
	})
}
