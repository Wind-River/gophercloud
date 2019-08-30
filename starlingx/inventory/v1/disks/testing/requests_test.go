/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/disks"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListDisks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDisksListSuccessfully(t)

	pages := 0
	err := disks.List(client.ServiceClient(),
		"f73dda8e-be3c-4704-ad1e-ed99e44b846e",
		disks.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := disks.ExtractDisks(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 disks, got %d", len(actual))
		}
		th.CheckDeepEquals(t, DiskHerp, actual[0])
		th.CheckDeepEquals(t, DiskDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllDisks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDisksListSuccessfully(t)

	allPages, err := disks.List(client.ServiceClient(), "f73dda8e-be3c-4704-ad1e-ed99e44b846e", disks.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := disks.ExtractDisks(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, DiskHerp, actual[0])
	th.CheckDeepEquals(t, DiskDerp, actual[1])
}

func TestGetDisk(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDiskGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := disks.Get(client, "c8e0a268-dd2e-4001-ba8c-21875325d01c").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, DiskDerp, *actual)
}
