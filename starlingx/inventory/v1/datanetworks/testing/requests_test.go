/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/datanetworks"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListDataNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDataNetworkListSuccessfully(t)

	pages := 0
	err := datanetworks.List(client.ServiceClient(), datanetworks.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := datanetworks.ExtractDataNetworks(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 DataNetworks, got %d", len(actual))
		}
		th.CheckDeepEquals(t, DataNetworkHerp, actual[0])
		th.CheckDeepEquals(t, DataNetworkDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllDataNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDataNetworkListSuccessfully(t)

	allPages, err := datanetworks.List(client.ServiceClient(),
		datanetworks.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := datanetworks.ExtractDataNetworks(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, DataNetworkHerp, actual[0])
	th.CheckDeepEquals(t, DataNetworkDerp, actual[1])
}

func TestGetDataNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDataNetworkGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := datanetworks.Get(client, "c71a1cfa-4bf9-48fc-a8ed-70f8797fc442").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, DataNetworkDerp, *actual)
}

func TestCreateDataNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDataNetworkCreationSuccessfully(t, DataNetworkSingleBody)

	network_type := "vlan"
	name := "physnet0"
	mtu := 1500
	actual, err := datanetworks.Create(client.ServiceClient(), datanetworks.DataNetworkOpts{
		Type: &network_type,
		Name: &name,
		MTU:  &mtu,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, DataNetworkDerp, *actual)
}

func TestDeleteDataNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDataNetworkDeletionSuccessfully(t)

	res := datanetworks.Delete(client.ServiceClient(), "c71a1cfa-4bf9-48fc-a8ed-70f8797fc442")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateDataNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDataNetworkUpdateSuccessfully(t)

	mtu := 2000
	actual, err := datanetworks.Update(client.ServiceClient(),
		"c71a1cfa-4bf9-48fc-a8ed-70f8797fc442",
		datanetworks.DataNetworkOpts{MTU: &mtu}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, DataNetworkDerp, *actual)
}
