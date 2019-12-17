/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/osds"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListOSDs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOSDListSuccessfully(t)

	pages := 0
	err := osds.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", osds.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := osds.ExtractOSDs(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 OSDs, got %d", len(actual))
		}
		th.CheckDeepEquals(t, OSDHerp, actual[0])
		th.CheckDeepEquals(t, OSDDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllOSDs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOSDListSuccessfully(t)

	allPages, err := osds.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		osds.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := osds.ExtractOSDs(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, OSDHerp, actual[0])
	th.CheckDeepEquals(t, OSDDerp, actual[1])
}

func TestGetOSD(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOSDGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := osds.Get(client, "dfa582bd-8f1d-4e8c-9c8a-1167243d592a").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, OSDDerp, *actual)
}

func TestCreateOSD(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOSDCreationSuccessfully(t, OSDSingleBody)

	HostID := "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e"
	Function := "osd"
	DiskID := "b4465948-b5b1-4ea2-b81f-cda62162e4be"
	TierUUID := "313c44bb-959d-46cc-82f7-6df7c2d7b057"
	JournalLocation := "982acece-40ac-4df0-b87d-eda51425e7bc"
	JournalSize := 1024
	actual, err := osds.Create(client.ServiceClient(), osds.OSDOpts{
		HostID:          &HostID,
		Function:        &Function,
		DiskID:          &DiskID,
		TierUUID:        &TierUUID,
		JournalLocation: &JournalLocation,
		JournalSize:     &JournalSize,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, OSDDerp, *actual)
}

func TestDeleteOSD(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOSDDeletionSuccessfully(t)

	res := osds.Delete(client.ServiceClient(), "dfa582bd-8f1d-4e8c-9c8a-1167243d592a")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateOSD(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOSDUpdateSuccessfully(t)

	JournalLocation := "bb5c7623-0eb2-4581-9942-7abac4a4be86"
	JournalSize := 2048
	actual, err := osds.Update(client.ServiceClient(),
		"dfa582bd-8f1d-4e8c-9c8a-1167243d592a",
		osds.OSDOpts{JournalLocation: &JournalLocation, JournalSize: &JournalSize}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, OSDDerp, *actual)
}
