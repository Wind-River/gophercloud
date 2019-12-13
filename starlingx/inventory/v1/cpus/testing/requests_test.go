/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/cpus"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListCPUs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCPUListSuccessfully(t)

	pages := 0
	err := cpus.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", cpus.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := cpus.ExtractCPUs(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 CPUs, got %d", len(actual))
		}
		th.CheckDeepEquals(t, CPUHerp, actual[0])
		th.CheckDeepEquals(t, CPUDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllCPUs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCPUListSuccessfully(t)

	allPages, err := cpus.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		cpus.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := cpus.ExtractCPUs(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, CPUHerp, actual[0])
	th.CheckDeepEquals(t, CPUDerp, actual[1])
}

func TestGetCPU(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCPUGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := cpus.Get(client, "0469d0bb-da3d-439d-bb6e-4720f1e74021").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, CPUDerp, *actual)
}

func TestUpdateCPU(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCPUUpdateSuccessfully(t)

	actual, err := cpus.Update(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		[]cpus.CPUOpts{cpus.CPUOpts{
			Function: "platform",
			Sockets:  []map[string]int{map[string]int{"0": 3}},
		}}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, CPUDerp, *actual)
}
