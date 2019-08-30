/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/addresses"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListAddresses(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressListSuccessfully(t)

	pages := 0
	err := addresses.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", addresses.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := addresses.ExtractAddresses(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 Addresses, got %d", len(actual))
		}
		th.CheckDeepEquals(t, AddressHerp, actual[0])
		th.CheckDeepEquals(t, AddressDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllAddresses(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressListSuccessfully(t)

	allPages, err := addresses.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		addresses.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := addresses.ExtractAddresses(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, AddressHerp, actual[0])
	th.CheckDeepEquals(t, AddressDerp, actual[1])
}

func TestGetAddress(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := addresses.Get(client, "37fe7dd1-51dd-4579-8904-1fdc71f4dbd1").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, AddressDerp, *actual)
}

func TestCreateAddress(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressCreationSuccessfully(t, AddressSingleBody)

	interfaceUUID := "62b00dc7-0549-4418-84a4-117c1f74b8d4"
	address := "4.3.2.1"
	prefix := 24
	actual, err := addresses.Create(client.ServiceClient(), addresses.AddressOpts{
		InterfaceUUID: &interfaceUUID,
		Address:       &address,
		Prefix:        &prefix,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, AddressDerp, *actual)
}

func TestDeleteAddress(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressDeletionSuccessfully(t)

	res := addresses.Delete(client.ServiceClient(), "37fe7dd1-51dd-4579-8904-1fdc71f4dbd1")
	th.AssertNoErr(t, res.Err)
}
