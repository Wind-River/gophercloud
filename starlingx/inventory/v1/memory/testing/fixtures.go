/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/memory"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	//VM1GHugepagesPendingHerp = 0
	VM1GHugepagesPendingDerp = 0
	VM2MHugepagesPendingDerp = 1000
	UpdatedAtHerp            = "2019-11-12T16:30:21.594724+00:00"
	UpdatedAtDerp            = "2019-11-12T16:30:21.594724+00:00"

	MemoryHerp = memory.Memory{
		ID:                        "126dcf05-bfb6-4ba6-a405-9c0d0ba261ab",
		Processor:                 0,
		Total:                     11783,
		Available:                 11783,
		Platform:                  4600,
		PlatformMinimum:           6200,
		VM1GHugepagesCount:        0,
		VM1GHugepagesEnabled:      "False",
		VM1GHugepagesPossible:     11,
		VM2MHugepagesCount:        0,
		VM2MHugepagesPossible:     5891,
		VM2MHugepagesAvailable:    0,
		VSwitchHugepagesSize:      2,
		VSwitchHugepagesAvailable: 0,
		VSwitchHugepagesCount:     0,
		CreatedAt:                 "2019-11-07T20:52:09.564404+00:00",
		UpdatedAt:                 &UpdatedAtHerp,
	}
	MemoryDerp = memory.Memory{
		ID:                        "143965fc-695b-4c66-9db4-bb14d1fd41c0",
		Processor:                 0,
		Total:                     11783,
		Available:                 11783,
		Platform:                  4600,
		PlatformMinimum:           6200,
		VM1GHugepagesCount:        0,
		VM1GHugepagesEnabled:      "False",
		VM1GHugepagesPossible:     11,
		VM2MHugepagesCount:        200,
		VM2MHugepagesPossible:     5891,
		VM2MHugepagesPending:      &VM2MHugepagesPendingDerp,
		VM2MHugepagesAvailable:    0,
		VSwitchHugepagesSize:      2,
		VSwitchHugepagesAvailable: 0,
		VSwitchHugepagesCount:     0,
		CreatedAt:                 "2019-11-07T20:52:09.564404+00:00",
		UpdatedAt:                 &UpdatedAtDerp,
	}
)

const MemoryListBody = `
{
    "imemorys": [
        {
            "vm_hugepages_nr_1G_pending": null,
            "platform_reserved_mib": 4600,
            "memtotal_mib": 11783,
            "updated_at": "2019-11-12T16:30:21.594724+00:00",
            "vswitch_hugepages_reqd": null,
            "inode_uuid": "fa183823-c83a-4f90-befd-cdd6bb61468d",
            "capabilities": {},
            "memavail_mib": 11783,
            "minimum_platform_reserved_mib": 6200,
            "uuid": "126dcf05-bfb6-4ba6-a405-9c0d0ba261ab",
            "vm_hugepages_use_1G": "False",
            "vm_hugepages_possible_1G": 11,
            "vswitch_hugepages_avail": 0,
            "hugepages_configured": "True",
            "vm_hugepages_avail_1G": null,
            "vswitch_hugepages_size_mib": 2,
            "vm_hugepages_nr_1G": 0,
            "vm_hugepages_nr_4K": 3016448,
            "vm_hugepages_nr_2M": 0,
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "numa_node": 0,
            "vm_hugepages_nr_2M_pending": null,
            "vm_hugepages_possible_2M": 5891,
            "vm_hugepages_avail_2M": 0,
            "vswitch_hugepages_nr": 0,
            "created_at": "2019-11-07T20:52:09.564404+00:00"
		},
		{
            "vm_hugepages_nr_1G_pending": null,
            "platform_reserved_mib": 4600,
            "memtotal_mib": 11783,
            "updated_at": "2019-11-12T16:30:21.594724+00:00",
            "vswitch_hugepages_reqd": null,
            "inode_uuid": "d4b0c874-99d1-4975-bdd3-3d59bee44d0c",
            "capabilities": {},
            "memavail_mib": 11783,
            "minimum_platform_reserved_mib": 6200,
            "uuid": "143965fc-695b-4c66-9db4-bb14d1fd41c0",
            "vm_hugepages_use_1G": "False",
            "vm_hugepages_possible_1G": 11,
            "vswitch_hugepages_avail": 0,
            "hugepages_configured": "True",
            "vm_hugepages_avail_1G": null,
            "vswitch_hugepages_size_mib": 2,
            "vm_hugepages_nr_1G": 0,
            "vm_hugepages_nr_4K": 3016448,
            "vm_hugepages_nr_2M": 200,
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "numa_node": 0,
            "vm_hugepages_nr_2M_pending": 1000,
            "vm_hugepages_possible_2M": 5891,
            "vm_hugepages_avail_2M": 0,
            "vswitch_hugepages_nr": 0,
            "created_at": "2019-11-07T20:52:09.564404+00:00"
        }
    ]
}
`

const MemorySingleBody = `
{
	"vm_hugepages_nr_1G_pending": null,
	"platform_reserved_mib": 4600,
	"memtotal_mib": 11783,
	"updated_at": "2019-11-12T16:30:21.594724+00:00",
	"vswitch_hugepages_reqd": null,
	"inode_uuid": "d4b0c874-99d1-4975-bdd3-3d59bee44d0c",
	"capabilities": {},
	"memavail_mib": 11783,
	"minimum_platform_reserved_mib": 6200,
	"uuid": "143965fc-695b-4c66-9db4-bb14d1fd41c0",
	"vm_hugepages_use_1G": "False",
	"vm_hugepages_possible_1G": 11,
	"vswitch_hugepages_avail": 0,
	"hugepages_configured": "True",
	"vm_hugepages_avail_1G": null,
	"vswitch_hugepages_size_mib": 2,
	"vm_hugepages_nr_1G": 0,
	"vm_hugepages_nr_4K": 3016448,
	"vm_hugepages_nr_2M": 200,
	"ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
	"numa_node": 0,
	"vm_hugepages_nr_2M_pending": 1000,
	"vm_hugepages_possible_2M": 5891,
	"vm_hugepages_avail_2M": 0,
	"vswitch_hugepages_nr": 0,
	"created_at": "2019-11-07T20:52:09.564404+00:00"
}
`

func HandleMemoryListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/imemorys", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, MemoryListBody)
	})
}

func HandleMemoryGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/imemorys/143965fc-695b-4c66-9db4-bb14d1fd41c0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, MemorySingleBody)
	})
}

func HandleMemoryUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/imemorys/143965fc-695b-4c66-9db4-bb14d1fd41c0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[ { "op": "replace", "path": "/vm_hugepages_nr_2M_pending", "value": 400 } ]`)

		fmt.Fprintf(w, MemorySingleBody)
	})
}
