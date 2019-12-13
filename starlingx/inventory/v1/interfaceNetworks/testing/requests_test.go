/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/interfaceNetworks"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListInterfaceNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceNetworkListSuccessfully(t)

	pages := 0
	err := interfaceNetworks.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", interfaceNetworks.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := interfaceNetworks.ExtractInterfaceNetworks(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 InterfaceNetworks, got %d", len(actual))
		}
		th.CheckDeepEquals(t, InterfaceNetworkHerp, actual[0])
		th.CheckDeepEquals(t, InterfaceNetworkDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllInterfaceNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceNetworkListSuccessfully(t)

	allPages, err := interfaceNetworks.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		interfaceNetworks.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := interfaceNetworks.ExtractInterfaceNetworks(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, InterfaceNetworkHerp, actual[0])
	th.CheckDeepEquals(t, InterfaceNetworkDerp, actual[1])
}

func TestGetInterfaceNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceNetworkGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := interfaceNetworks.Get(client, "2afb1f71-b42e-403d-a0cd-f1138be8167e").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, InterfaceNetworkDerp, *actual)
}

func TestCreateInterfaceNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceNetworkCreationSuccessfully(t, InterfaceNetworkSingleBody)

	actual, err := interfaceNetworks.Create(client.ServiceClient(), interfaceNetworks.InterfaceNetworkOpts{
		InterfaceUUID: "fa7721c8-cb24-4f91-bb2e-736cabeff4e2",
		NetworkUUID:   "7ace62b9-907a-4f35-b88b-97f829f66330",
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, InterfaceNetworkDerp, *actual)
}

func TestDeleteInterfaceNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceNetworkDeletionSuccessfully(t)

	res := interfaceNetworks.Delete(client.ServiceClient(), "2afb1f71-b42e-403d-a0cd-f1138be8167e")
	th.AssertNoErr(t, res.Err)
}
