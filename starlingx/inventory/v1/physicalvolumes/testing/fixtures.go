/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/physicalvolumes"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	UpdatedAtHerp      = "2019-11-07T21:23:55.371201+00:00"
	UpdatedAtDerp      = "2019-11-07T21:23:55.425455+00:00"
	PhysicalVolumeHerp = physicalvolumes.PhysicalVolume{
		ID:            "4c039907-54a4-4b8b-8241-ad2867565032",
		Type:          "partition",
		State:         "provisioned",
		HostID:        "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		DevicePath:    "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part4",
		DeviceNode:    "/dev/sda4",
		DeviceUUID:    "f15b94af-8912-4664-a63e-97e6e83d9d7d",
		VolumeGroupID: "40cc6c66-fc2d-45f4-80d8-f7d218a4cb62",
		CreatedAt:     "2019-11-07T20:52:14.328904+00:00",
		UpdatedAt:     &UpdatedAtHerp,
		LVMInfo: physicalvolumes.LVMInfo{
			Name:            "/dev/sda4",
			VolumeGroupName: "cgts-vg",
			UUID:            "0C8f58-qIub-FebU-vDku-EKzm-GXRT-a1qzEE",
			Size:            245853323264,
			TotalPE:         7327,
			AllocatedPE:     4954,
		},
	}
	PhysicalVolumeDerp = physicalvolumes.PhysicalVolume{
		ID:            "8b798534-4c87-40e7-9058-b2f9bbcb01bb",
		Type:          "partition",
		State:         "provisioned",
		HostID:        "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		DevicePath:    "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part5",
		DeviceNode:    "/dev/sda5",
		DeviceUUID:    "4c495b9d-a66f-4ffc-b109-2af359108f33",
		VolumeGroupID: "449aee64-342f-4255-9a23-b229b0589c1b",
		CreatedAt:     "2019-11-07T21:07:50.896648+00:00",
		UpdatedAt:     &UpdatedAtDerp,
		LVMInfo: physicalvolumes.LVMInfo{
			Name:            "/dev/sda5",
			VolumeGroupName: "nova-local",
			UUID:            "DlpMMX-XF5v-rUqn-xwR2-px5s-GhDc-j30Tam",
			Size:            36503027712,
			TotalPE:         8703,
			AllocatedPE:     8703,
		},
	}
)

const PhysicalVolumeListBody = `
{
    "ipvs": [
        {
            "lvm_pe_alloced": 4954,
            "lvm_pe_total": 7327,
            "ilvg_uuid": "40cc6c66-fc2d-45f4-80d8-f7d218a4cb62",
            "uuid": "4c039907-54a4-4b8b-8241-ad2867565032",
            "disk_or_part_device_path": "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part4",
            "lvm_pv_name": "/dev/sda4",
            "created_at": "2019-11-07T20:52:14.328904+00:00",
            "disk_or_part_device_node": "/dev/sda4",
            "forilvgid": 1,
            "disk_or_part_uuid": "f15b94af-8912-4664-a63e-97e6e83d9d7d",
            "updated_at": "2019-11-07T21:23:55.371201+00:00",
            "pv_state": "provisioned",
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "pv_type": "partition",
            "capabilities": {},
            "lvm_vg_name": "cgts-vg",
            "lvm_pv_uuid": "0C8f58-qIub-FebU-vDku-EKzm-GXRT-a1qzEE",
            "lvm_pv_size": 245853323264
        },
        {
            "lvm_pe_alloced": 8703,
            "lvm_pe_total": 8703,
            "ilvg_uuid": "449aee64-342f-4255-9a23-b229b0589c1b",
            "uuid": "8b798534-4c87-40e7-9058-b2f9bbcb01bb",
            "disk_or_part_device_path": "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part5",
            "lvm_pv_name": "/dev/sda5",
            "created_at": "2019-11-07T21:07:50.896648+00:00",
            "disk_or_part_device_node": "/dev/sda5",
            "forilvgid": 2,
            "disk_or_part_uuid": "4c495b9d-a66f-4ffc-b109-2af359108f33",
            "updated_at": "2019-11-07T21:23:55.425455+00:00",
            "pv_state": "provisioned",
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "pv_type": "partition",
            "capabilities": {},
            "lvm_vg_name": "nova-local",
            "lvm_pv_uuid": "DlpMMX-XF5v-rUqn-xwR2-px5s-GhDc-j30Tam",
            "lvm_pv_size": 36503027712
        }
    ]
}
`

const PhysicalVolumeSingleBody = `
{
	"lvm_pe_alloced": 8703,
	"lvm_pe_total": 8703,
	"ilvg_uuid": "449aee64-342f-4255-9a23-b229b0589c1b",
	"uuid": "8b798534-4c87-40e7-9058-b2f9bbcb01bb",
	"disk_or_part_device_path": "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part5",
	"lvm_pv_name": "/dev/sda5",
	"created_at": "2019-11-07T21:07:50.896648+00:00",
	"disk_or_part_device_node": "/dev/sda5",
	"forilvgid": 2,
	"disk_or_part_uuid": "4c495b9d-a66f-4ffc-b109-2af359108f33",
	"updated_at": "2019-11-07T21:23:55.425455+00:00",
	"pv_state": "provisioned",
	"ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
	"pv_type": "partition",
	"capabilities": {},
	"lvm_vg_name": "nova-local",
	"lvm_pv_uuid": "DlpMMX-XF5v-rUqn-xwR2-px5s-GhDc-j30Tam",
	"lvm_pv_size": 36503027712
}
`

func HandlePhysicalVolumeListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/ipvs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, PhysicalVolumeListBody)
	})
}

func HandlePhysicalVolumeGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ipvs/8b798534-4c87-40e7-9058-b2f9bbcb01bb", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, PhysicalVolumeSingleBody)
	})
}

func HandlePhysicalVolumeDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ipvs/8b798534-4c87-40e7-9058-b2f9bbcb01bb", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandlePhysicalVolumeCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/ipvs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
          "ilvg_uuid": "449aee64-342f-4255-9a23-b229b0589c1b",
		  "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		  "disk_or_part_uuid": "4c495b9d-a66f-4ffc-b109-2af359108f33",
		  "pv_type": ""
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}
