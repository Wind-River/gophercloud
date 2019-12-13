/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/storagetiers"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListStorageTiers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageTierListSuccessfully(t)

	pages := 0
	err := storagetiers.List(client.ServiceClient(), storagetiers.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := storagetiers.ExtractStorageTiers(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 StorageTiers, got %d", len(actual))
		}
		th.CheckDeepEquals(t, StorageTierHerp, actual[0])
		th.CheckDeepEquals(t, StorageTierDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllStorageTiers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageTierListSuccessfully(t)

	allPages, err := storagetiers.List(client.ServiceClient(),
		storagetiers.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := storagetiers.ExtractStorageTiers(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, StorageTierHerp, actual[0])
	th.CheckDeepEquals(t, StorageTierDerp, actual[1])
}

func TestGetStorageTier(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageTierGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := storagetiers.Get(client, "4da7bca3-a7ea-4042-99ac-6f74f27c50b3").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, StorageTierDerp, *actual)
}

func TestCreateStorageTier(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageTierCreationSuccessfully(t, StorageTierSingleBody)

	actual, err := storagetiers.Create(client.ServiceClient(), storagetiers.StorageTierOpts{
		Name:      "test",
		ClusterID: "e7d497b9-085c-4a46-bfff-21d02ad9313a",
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, StorageTierDerp, *actual)
}

func TestDeleteStorageTier(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageTierDeletionSuccessfully(t)

	res := storagetiers.Delete(client.ServiceClient(), "4da7bca3-a7ea-4042-99ac-6f74f27c50b3")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateStorageTier(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageTierUpdateSuccessfully(t)

	actual, err := storagetiers.Update(client.ServiceClient(),
		"4da7bca3-a7ea-4042-99ac-6f74f27c50b3",
		storagetiers.StorageTierOpts{ClusterID: "9db6c194-dc8b-4c9e-8a0e-a3d01fa1f5f1", Name: "edited"}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, StorageTierDerp, *actual)
}
