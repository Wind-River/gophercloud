/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/cephmonitors"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListCephMonitors(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCephMonListSuccessfully(t)

	pages := 0
	err := cephmonitors.List(client.ServiceClient(), cephmonitors.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := cephmonitors.ExtractCephMonitors(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 ceph monitors, got %d", len(actual))
		}
		th.CheckDeepEquals(t, CephMonitorHerp, actual[0])
		th.CheckDeepEquals(t, CephMonitorDerp, actual[1])
		th.CheckDeepEquals(t, CephMonitorMerp, actual[2])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllCephMonitors(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCephMonListSuccessfully(t)

	allPages, err := cephmonitors.List(client.ServiceClient(), cephmonitors.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := cephmonitors.ExtractCephMonitors(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, CephMonitorHerp, actual[0])
	th.CheckDeepEquals(t, CephMonitorDerp, actual[1])
}

func TestCreateCephMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCephMonCreationSuccessfully(t, CephMonSingleBody)

	hostUUID := "d99637e9-5451-45c6-98f4-f18968e43e91"
	size := 30
	actual, err := cephmonitors.Create(client.ServiceClient(), cephmonitors.CephMonitorOpts{
		HostUUID: &hostUUID,
		Size:     &size,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, CephMonitorDerp, *actual)
}

func TestUpdateCephMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCephMonUpdateSuccessfully(t)

	newSize := 35
	actual, err := cephmonitors.Update(client.ServiceClient(),
		"79c25f45-9e3d-42bf-ab94-7c6b193934b7",
		cephmonitors.CephMonitorOpts{Size: &newSize}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, CephMonitorDerp, *actual)
}

func TestDeleteCephMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCephMonDeletionSuccessfully(t)

	res := cephmonitors.Delete(client.ServiceClient(), "79c25f45-9e3d-42bf-ab94-7c6b193934b7")
	th.AssertNoErr(t, res.Err)
}
