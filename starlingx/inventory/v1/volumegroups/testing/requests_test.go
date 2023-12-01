/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/volumegroups"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListVolumeGroups(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleVolumeGroupListSuccessfully(t)

	pages := 0
	err := volumegroups.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", volumegroups.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := volumegroups.ExtractVolumeGroups(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 VolumeGroups, got %d", len(actual))
		}
		th.CheckDeepEquals(t, VolumeGroupHerp, actual[0])
		th.CheckDeepEquals(t, VolumeGroupDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllVolumeGroups(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleVolumeGroupListSuccessfully(t)

	allPages, err := volumegroups.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		volumegroups.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := volumegroups.ExtractVolumeGroups(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, VolumeGroupHerp, actual[0])
	th.CheckDeepEquals(t, VolumeGroupDerp, actual[1])
}

func TestGetVolumeGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleVolumeGroupGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := volumegroups.Get(client, "449aee64-342f-4255-9a23-b229b0589c1b").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, VolumeGroupDerp, *actual)
}

func TestCreateVolumeGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleVolumeGroupCreationSuccessfully(t, VolumeGroupSingleBody)

	HostID := "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e"
	Name := "nova-local"
	actual, err := volumegroups.Create(client.ServiceClient(), volumegroups.VolumeGroupOpts{
		HostID: &HostID,
		Name:   &Name,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, VolumeGroupDerp, *actual)
}

func TestDeleteVolumeGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleVolumeGroupDeletionSuccessfully(t)

	res := volumegroups.Delete(client.ServiceClient(), "449aee64-342f-4255-9a23-b229b0589c1b")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateVolumeGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleVolumeGroupUpdateSuccessfully(t)

	Capabilities := volumegroups.CapabilitiesOpts{}
	actual, err := volumegroups.Update(client.ServiceClient(),
		"449aee64-342f-4255-9a23-b229b0589c1b",
		volumegroups.VolumeGroupOpts{
			Capabilities: &Capabilities}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, VolumeGroupDerp, *actual)
}
