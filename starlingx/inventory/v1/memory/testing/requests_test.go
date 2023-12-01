/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/memory"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListMemory(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleMemoryListSuccessfully(t)

	pages := 0
	err := memory.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", memory.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := memory.ExtractMemorys(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 Memories, got %d", len(actual))
		}
		th.CheckDeepEquals(t, MemoryHerp, actual[0])
		th.CheckDeepEquals(t, MemoryDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllMemories(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleMemoryListSuccessfully(t)

	allPages, err := memory.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		memory.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := memory.ExtractMemorys(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, MemoryHerp, actual[0])
	th.CheckDeepEquals(t, MemoryDerp, actual[1])
}

func TestGetMemory(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleMemoryGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := memory.Get(client, "143965fc-695b-4c66-9db4-bb14d1fd41c0").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, MemoryDerp, *actual)
}

func TestUpdateMemory(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleMemoryUpdateSuccessfully(t)

	vm_hugepages_nr_2M := 400
	actual, err := memory.Update(client.ServiceClient(),
		"143965fc-695b-4c66-9db4-bb14d1fd41c0",
		memory.MemoryOpts{VMHugepages2M: &vm_hugepages_nr_2M}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, MemoryDerp, *actual)
}
