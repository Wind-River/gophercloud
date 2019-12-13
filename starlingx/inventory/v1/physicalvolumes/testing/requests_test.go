/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/physicalvolumes"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListPhysicalVolumes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePhysicalVolumeListSuccessfully(t)

	pages := 0
	err := physicalvolumes.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", physicalvolumes.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := physicalvolumes.ExtractPhysicalVolumes(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 PhysicalVolumes, got %d", len(actual))
		}
		th.CheckDeepEquals(t, PhysicalVolumeHerp, actual[0])
		th.CheckDeepEquals(t, PhysicalVolumeDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllPhysicalVolumes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePhysicalVolumeListSuccessfully(t)

	allPages, err := physicalvolumes.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		physicalvolumes.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := physicalvolumes.ExtractPhysicalVolumes(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PhysicalVolumeHerp, actual[0])
	th.CheckDeepEquals(t, PhysicalVolumeDerp, actual[1])
}

func TestGetPhysicalVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePhysicalVolumeGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := physicalvolumes.Get(client, "8b798534-4c87-40e7-9058-b2f9bbcb01bb").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, PhysicalVolumeDerp, *actual)
}

func TestCreatePhysicalVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePhysicalVolumeCreationSuccessfully(t, PhysicalVolumeSingleBody)

	actual, err := physicalvolumes.Create(client.ServiceClient(), physicalvolumes.PhysicalVolumeOpts{
		VolumeGroupID: "449aee64-342f-4255-9a23-b229b0589c1b",
		HostID:        "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		DeviceID:      "4c495b9d-a66f-4ffc-b109-2af359108f33",
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, PhysicalVolumeDerp, *actual)
}

func TestDeletePhysicalVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePhysicalVolumeDeletionSuccessfully(t)

	res := physicalvolumes.Delete(client.ServiceClient(), "8b798534-4c87-40e7-9058-b2f9bbcb01bb")
	th.AssertNoErr(t, res.Err)
}
