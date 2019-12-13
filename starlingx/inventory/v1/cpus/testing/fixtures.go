/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/cpus"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	CPUHerp = cpus.CPU{
		ID:           "e85f7367-16e6-44ab-a62f-25a081b49fe2",
		Processor:    0,
		LogicalCore:  8,
		PhysicalCore: 11,
		Thread:       0,
		Function:     "Application",
		CreatedAt:    "2019-11-11T17:03:26.929861+00:00",
	}
	CPUDerp = cpus.CPU{
		ID:           "0469d0bb-da3d-439d-bb6e-4720f1e74021",
		Processor:    0,
		LogicalCore:  28,
		PhysicalCore: 11,
		Thread:       1,
		Function:     "Application",
		CreatedAt:    "2019-11-11T17:03:26.938028+00:00",
	}
)

const CPUListBody = `
{
    "icpus": [
        {
            "allocated_function": "Application",
            "core": 11,
            "uuid": "e85f7367-16e6-44ab-a62f-25a081b49fe2",
            "thread": 0,
            "inode_uuid": "fa68b35e-6ce3-40f7-acfa-ed01e23a3a1c",
            "numa_node": 0,
            "created_at": "2019-11-11T17:03:26.929861+00:00",
            "cpu_model": "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
            "capabilities": {},
            "updated_at": null,
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "cpu_family": "6",
            "cpu": 8
        },
        {
            "allocated_function": "Application",
            "core": 11,
            "uuid": "0469d0bb-da3d-439d-bb6e-4720f1e74021",
            "thread": 1,
            "inode_uuid": "fa68b35e-6ce3-40f7-acfa-ed01e23a3a1c",
            "numa_node": 0,
            "created_at": "2019-11-11T17:03:26.938028+00:00",
            "cpu_model": "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
            "capabilities": {},
            "updated_at": null,
            "ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
            "cpu_family": "6",
            "cpu": 28
        }
    ]
}
`

const CPUSingleBody = `
{
	"allocated_function": "Application",
	"core": 11,
	"uuid": "0469d0bb-da3d-439d-bb6e-4720f1e74021",
	"thread": 1,
	"inode_uuid": "fa68b35e-6ce3-40f7-acfa-ed01e23a3a1c",
	"numa_node": 0,
	"created_at": "2019-11-11T17:03:26.938028+00:00",
	"cpu_model": "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
	"capabilities": {},
	"updated_at": null,
	"ihost_uuid": "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
	"cpu_family": "6",
	"cpu": 28
}
`

func HandleCPUListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/icpus", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CPUListBody)
	})
}

func HandleCPUGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/icpus/0469d0bb-da3d-439d-bb6e-4720f1e74021", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, CPUSingleBody)
	})
}

func HandleCPUUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/state/host_cpus_modify", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `[{"function": "platform", "sockets": [{"0": 3}]}]`)

		fmt.Fprintf(w, CPUSingleBody)
	})
}
