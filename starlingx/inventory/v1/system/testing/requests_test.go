/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/system"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListSystems(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSystemListSuccessfully(t)

	pages := 0
	err := system.List(client.ServiceClient(), system.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := system.ExtractSystems(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 servers, got %d", len(actual))
		}
		th.CheckDeepEquals(t, SystemHerp, actual[0])
		th.CheckDeepEquals(t, SystemDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllSystems(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSystemListSuccessfully(t)

	allPages, err := system.List(client.ServiceClient(), system.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := system.ExtractSystems(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SystemHerp, actual[0])
	th.CheckDeepEquals(t, SystemDerp, actual[1])
}

func TestUpdateSystem(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSystemUpdateSuccessfully(t)

	newName := "new-name"
	actual, err := system.Update(client.ServiceClient(), "6dc8ccb9-f687-40fa-9663-c0c286e65772", system.SystemOpts{Name: &newName}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, SystemDerp, *actual)
}

func TestGetSystem(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSystemGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := system.Get(client, "6dc8ccb9-f687-40fa-9663-c0c286e65772").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, SystemDerp, *actual)
}

/*
func TestToSystemListQuery(t *testing.T) {
	TestOps := system.ListOpts{

	}
}
*/

func TestGetDefaultSystem(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSystemGetDefaultSuccesfully(t)

	client := client.ServiceClient()
	actual, err := system.GetDefaultSystem(client)
	t.Log(err)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SystemHerp, *actual)
}
