/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/partitions"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListDiskPartitions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDiskPartitionListSuccessfully(t)

	pages := 0
	err := partitions.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", partitions.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := partitions.ExtractDiskPartitions(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 DiskPartitions, got %d", len(actual))
		}
		th.CheckDeepEquals(t, DiskPartitionHerp, actual[0])
		th.CheckDeepEquals(t, DiskPartitionDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllDiskPartitions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDiskPartitionListSuccessfully(t)

	allPages, err := partitions.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		partitions.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := partitions.ExtractDiskPartitions(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, DiskPartitionHerp, actual[0])
	th.CheckDeepEquals(t, DiskPartitionDerp, actual[1])
}

func TestGetDiskPartition(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDiskPartitionGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := partitions.Get(client, "09aa5a82-aa89-4d86-a9d1-7410167d510b").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, DiskPartitionDerp, *actual)
}

func TestCreateDiskPartition(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDiskPartitionCreationSuccessfully(t, DiskPartitionSingleBody)

	TypeGUID := "ba5eba11-0000-1111-2222-000000000001"
	actual, err := partitions.Create(client.ServiceClient(), partitions.DiskPartitionOpts{
		Size:     250,
		HostID:   "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		TypeGUID: &TypeGUID,
		DiskID:   "d104a752-3186-4172-8a0b-e792321ebf37",
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, DiskPartitionDerp, *actual)
}

func TestDeleteDiskPartition(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDiskPartitionDeletionSuccessfully(t)

	res := partitions.Delete(client.ServiceClient(), "09aa5a82-aa89-4d86-a9d1-7410167d510b")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateDiskPartition(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDiskPartitionUpdateSuccessfully(t)

	size_mib := 400
	actual, err := partitions.Update(client.ServiceClient(),
		"09aa5a82-aa89-4d86-a9d1-7410167d510b",
		partitions.DiskPartitionOpts{Size: size_mib}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, DiskPartitionDerp, *actual)
}
