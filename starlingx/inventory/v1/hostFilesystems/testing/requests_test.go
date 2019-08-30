/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/hostFilesystems"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListFileSystems(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleFileSystemListSuccessfully(t)

	pages := 0
	err := hostFilesystems.List(client.ServiceClient(),
		"d99637e9-5451-45c6-98f4-f18968e43e91",
		hostFilesystems.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := hostFilesystems.ExtractFileSystems(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 file systems, got %d", len(actual))
		}
		th.CheckDeepEquals(t, FileSystemHerp, actual[0])
		th.CheckDeepEquals(t, FileSystemDerp, actual[1])

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

	allPages, err := hostFilesystems.List(client.ServiceClient(),
		"d99637e9-5451-45c6-98f4-f18968e43e91",
		hostFilesystems.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := hostFilesystems.ExtractFileSystems(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, FileSystemHerp, actual[0])
	th.CheckDeepEquals(t, FileSystemDerp, actual[1])
}

func TestGetFileSystem(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleFileSystemGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := hostFilesystems.Get(client, "1a43b9e1-6360-46c1-adbe-81987a732e94").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, FileSystemDerp, *actual)
}

func TestUpdateFileSystem(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleFileSystemUpdateSuccessfully(t)

	newOpt := hostFilesystems.FileSystemOpts{
		Name: "Derp",
		Size: 50, // new size
	}
	newOpts := []hostFilesystems.FileSystemOpts{newOpt}
	err := hostFilesystems.Update(client.ServiceClient(),
		"d99637e9-5451-45c6-98f4-f18968e43e91",
		newOpts).ExtractErr()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
}
