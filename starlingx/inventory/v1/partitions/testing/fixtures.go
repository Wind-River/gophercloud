/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/partitions"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	PhysicalVolumeIDHerp = "8b798534-4c87-40e7-9058-b2f9bbcb01bb"
	PhysicalVolumeIDDerp = "266b16a8-73da-46b7-967b-dbc1ae96a9c9"
	UpdatedAtHerp        = "2019-11-07T21:23:55.447690+00:00"
	UpdatedAtDerp        = "2019-11-07T21:23:55.447690+00:00"
	DiskPartitionHerp    = partitions.DiskPartition{
		ID:               "4c495b9d-a66f-4ffc-b109-2af359108f33",
		HostID:           "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		DiskID:           "e8ad3e30-e00c-4b2a-9845-ca4e8079b5b7",
		DevicePath:       "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part5",
		DeviceNode:       "/dev/sda5",
		TypeName:         "LVM Physical Volume",
		TypeGUID:         "ba5eba11-0000-1111-2222-000000000001",
		Size:             34816,
		Start:            254998,
		End:              289814,
		PhysicalVolumeID: &PhysicalVolumeIDHerp,
		Status:           1,
		CreatedAt:        "2019-11-07T21:07:42.097201+00:00",
		UpdatedAt:        &UpdatedAtHerp,
	}
	DiskPartitionDerp = partitions.DiskPartition{
		ID:               "09aa5a82-aa89-4d86-a9d1-7410167d510b",
		HostID:           "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		DiskID:           "d104a752-3186-4172-8a0b-e792321ebf37",
		DevicePath:       "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part5",
		DeviceNode:       "/dev/sda5",
		TypeName:         "LVM Physical Volume",
		TypeGUID:         "ba5eba11-0000-1111-2222-000000000001",
		Size:             254997,
		Start:            0,
		End:              254997,
		PhysicalVolumeID: &PhysicalVolumeIDDerp,
		Status:           1,
		CreatedAt:        "2019-11-07T21:07:42.097201+00:00",
		UpdatedAt:        &UpdatedAtDerp,
	}
)

const DiskPartitionListBody = `
{
    "partitions": [
        {
            "status": 1,
            "device_path": "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part5",
            "start_mib": 254998,
            "uuid": "4c495b9d-a66f-4ffc-b109-2af359108f33",
            "capabilities": {},
            "created_at": "2019-11-07T21:07:42.097201+00:00",
            "type_name": "LVM Physical Volume",
            "updated_at": "2019-11-07T21:23:55.447690+00:00",
            "device_node": "/dev/sda5",
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "ipv_uuid": "8b798534-4c87-40e7-9058-b2f9bbcb01bb",
            "end_mib": 289814,
            "idisk_uuid": "e8ad3e30-e00c-4b2a-9845-ca4e8079b5b7",
            "type_guid": "ba5eba11-0000-1111-2222-000000000001",
            "size_mib": 34816
        },
        {
            "status": 1,
            "device_path": "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part5",
            "start_mib": 0,
            "uuid": "09aa5a82-aa89-4d86-a9d1-7410167d510b",
            "capabilities": {},
            "created_at": "2019-11-07T21:07:42.097201+00:00",
            "type_name": "LVM Physical Volume",
            "updated_at": "2019-11-07T21:23:55.447690+00:00",
            "device_node": "/dev/sda5",
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "ipv_uuid": "266b16a8-73da-46b7-967b-dbc1ae96a9c9",
            "end_mib": 254997,
            "idisk_uuid": "d104a752-3186-4172-8a0b-e792321ebf37",
            "type_guid": "ba5eba11-0000-1111-2222-000000000001",
            "size_mib": 254997
        }
    ]
}
`

const DiskPartitionSingleBody = `
{
	"status": 1,
	"device_path": "/dev/disk/by-path/pci-0000:00:01.1-ata-1.0-part5",
	"start_mib": 0,
	"uuid": "09aa5a82-aa89-4d86-a9d1-7410167d510b",
	"capabilities": {},
	"created_at": "2019-11-07T21:07:42.097201+00:00",
	"type_name": "LVM Physical Volume",
	"updated_at": "2019-11-07T21:23:55.447690+00:00",
	"device_node": "/dev/sda5",
	"ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
	"ipv_uuid": "266b16a8-73da-46b7-967b-dbc1ae96a9c9",
	"end_mib": 254997,
	"idisk_uuid": "d104a752-3186-4172-8a0b-e792321ebf37",
	"type_guid": "ba5eba11-0000-1111-2222-000000000001",
	"size_mib": 254997
}
`

func HandleDiskPartitionListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/partitions", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DiskPartitionListBody)
	})
}

func HandleDiskPartitionGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/partitions/09aa5a82-aa89-4d86-a9d1-7410167d510b", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, DiskPartitionSingleBody)
	})
}

func HandleDiskPartitionCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/partitions", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "idisk_uuid": "d104a752-3186-4172-8a0b-e792321ebf37",
            "type_guid": "ba5eba11-0000-1111-2222-000000000001",
            "size_mib": 256000
        }`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}

func HandleDiskPartitionUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/partitions/09aa5a82-aa89-4d86-a9d1-7410167d510b", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[
          {
            "op": "replace",
            "path": "/ihost_uuid",
            "value": ""
          },
          {
            "op": "replace",
            "path": "/idisk_uuid",
            "value": ""
		  },
		  {
            "op": "replace",
            "path": "/size_mib",
            "value": 409600
          } ]`)

		fmt.Fprintf(w, DiskPartitionSingleBody)
	})
}

func HandleDiskPartitionDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/partitions/09aa5a82-aa89-4d86-a9d1-7410167d510b", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}
