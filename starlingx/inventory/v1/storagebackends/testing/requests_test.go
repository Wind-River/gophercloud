/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/storagebackends"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListStorageBackends(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageBackendListSuccessfully(t)

	pages := 0
	err := storagebackends.List(client.ServiceClient(), storagebackends.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := storagebackends.ExtractStorageBackends(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 StorageBackends, got %d", len(actual))
		}
		th.CheckDeepEquals(t, StorageBackendHerp, actual[0])
		th.CheckDeepEquals(t, StorageBackendDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllStorageBackends(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageBackendListSuccessfully(t)

	allPages, err := storagebackends.List(client.ServiceClient(),
		storagebackends.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := storagebackends.ExtractStorageBackends(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, StorageBackendHerp, actual[0])
	th.CheckDeepEquals(t, StorageBackendDerp, actual[1])
}

func TestGetStorageBackend(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageBackendGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := storagebackends.Get(client, "cebe7a5e-7b57-497b-a335-6e7cf93e98ee").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, StorageBackendDerp, *actual)
}

func TestCreateStorageBackend(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageBackendCreationSuccessfully(t, StorageBackendSingleBody)

	Backend := "ceph"
	Name := "ceph-store"
	Capabilities := map[string]interface{}{}
	actual, err := storagebackends.Create(client.ServiceClient(), storagebackends.StorageBackendOpts{
		Backend:      &Backend,
		Name:         &Name,
		Capabilities: &Capabilities,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, StorageBackendDerp, *actual)
}

func TestDeleteStorageBackend(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageBackendDeletionSuccessfully(t)

	res := storagebackends.Delete(client.ServiceClient(), "cebe7a5e-7b57-497b-a335-6e7cf93e98ee")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateStorageBackend(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStorageBackendUpdateSuccessfully(t)

	capabilities := map[string]interface{}{}
	actual, err := storagebackends.Update(client.ServiceClient(),
		"cebe7a5e-7b57-497b-a335-6e7cf93e98ee",
		storagebackends.StorageBackendOpts{Capabilities: &capabilities}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, StorageBackendDerp, *actual)
}
