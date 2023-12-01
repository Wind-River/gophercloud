/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/interfaceDataNetworks"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListInterfaceDataNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceDataNetworkListSuccessfully(t)

	pages := 0
	err := interfaceDataNetworks.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		interfaceDataNetworks.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {

		pages++
		actual, err := interfaceDataNetworks.ExtractInterfaceDataNetworks(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 InterfaceDataNetworks, got %d", len(actual))
		}
		th.CheckDeepEquals(t, InterfaceDataNetworkHerp, actual[0])
		th.CheckDeepEquals(t, InterfaceDataNetworkDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllInterfaceDataNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceDataNetworkListSuccessfully(t)

	allPages, err := interfaceDataNetworks.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		interfaceDataNetworks.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := interfaceDataNetworks.ExtractInterfaceDataNetworks(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, InterfaceDataNetworkHerp, actual[0])
	th.CheckDeepEquals(t, InterfaceDataNetworkDerp, actual[1])
}

func TestGetInterfaceDataNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceDataNetworkGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := interfaceDataNetworks.Get(client, "0038c026-ef30-4bca-8931-5d7255ebd34e").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, InterfaceDataNetworkDerp, *actual)
}

func TestCreateDataNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceDataNetworkCreationSuccessfully(t, InterfaceDataNetworkSingleBody)

	interface_uuid := "953daf34-af6f-4bd9-a37f-1f1fd2a9cce5"
	datanetwork_uuid := "05331897-b9ab-40a3-8029-b87099a0190f"
	actual, err := interfaceDataNetworks.Create(client.ServiceClient(), interfaceDataNetworks.InterfaceDataNetworkOpts{
		InterfaceUUID:   interface_uuid,
		DataNetworkUUID: datanetwork_uuid,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, InterfaceDataNetworkDerp, *actual)
}

func TestDeleteInterfaceDataNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceDataNetworkDeletionSuccessfully(t)

	res := interfaceDataNetworks.Delete(client.ServiceClient(), "0038c026-ef30-4bca-8931-5d7255ebd34e")
	th.AssertNoErr(t, res.Err)
}
