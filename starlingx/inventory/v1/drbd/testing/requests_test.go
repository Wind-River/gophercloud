/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/drbd"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListDRBDs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDRBDListSuccessfully(t)

	pages := 0
	err := drbd.List(client.ServiceClient(), drbd.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := drbd.ExtractDRBDs(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 DRBDs, got %d", len(actual))
		}
		th.CheckDeepEquals(t, DRBDHerp, actual[0])
		th.CheckDeepEquals(t, DRBDDerp, actual[1])
		th.CheckDeepEquals(t, DRBDMerp, actual[2])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllDRBDs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDRBDListSuccessfully(t)

	allPages, err := drbd.List(client.ServiceClient(), drbd.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := drbd.ExtractDRBDs(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, DRBDHerp, actual[0])
	th.CheckDeepEquals(t, DRBDDerp, actual[1])
}

func TestGetDRBD(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDRBDGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := drbd.Get(client, "90c13f86-49db-4a5a-ba6e-22016fe96223").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, DRBDDerp, *actual)
}

func TestUpdateDRBD(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDRBDUpdateSuccessfully(t)

	newLinkUtil := 60
	actual, err := drbd.Update(client.ServiceClient(),
		"90c13f86-49db-4a5a-ba6e-22016fe96223",
		drbd.DRBDOpts{LinkUtilization: newLinkUtil}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, DRBDDerp, *actual)
}
