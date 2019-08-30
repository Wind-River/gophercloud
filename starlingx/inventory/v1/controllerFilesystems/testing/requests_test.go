/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/controllerFilesystems"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListFileSystems(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleFileSystemListSuccessfully(t)

	pages := 0
	err := controllerFilesystems.List(client.ServiceClient(), controllerFilesystems.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := controllerFilesystems.ExtractFileSystems(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 servers, got %d", len(actual))
		}
		th.CheckDeepEquals(t, FileSystemHerp, actual[0])
		th.CheckDeepEquals(t, FileSystemDerp, actual[1])
		th.CheckDeepEquals(t, FileSystemMerp, actual[2])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllFileSystems(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleFileSystemListSuccessfully(t)

	allPages, err := controllerFilesystems.List(client.ServiceClient(), controllerFilesystems.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := controllerFilesystems.ExtractFileSystems(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, FileSystemHerp, actual[0])
	th.CheckDeepEquals(t, FileSystemDerp, actual[1])
}

func TestGetFileSystem(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleFileSystemGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := controllerFilesystems.Get(client, "ff2e628d-23b2-4d73-b6b5-1c291ab6905a").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, FileSystemDerp, *actual)
}
