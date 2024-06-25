/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/addresspools"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListAddressPool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressPoolListSuccessfully(t)

	pages := 0
	err := addresspools.List(client.ServiceClient(), addresspools.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := addresspools.ExtractAddressPools(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 AddressPools, got %d", len(actual))
		}

		th.CheckDeepEquals(t, AddressPoolHerp, actual[0])
		th.CheckDeepEquals(t, AddressPoolDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllAddressPools(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressPoolListSuccessfully(t)

	allPages, err := addresspools.List(client.ServiceClient(), addresspools.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := addresspools.ExtractAddressPools(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, AddressPoolHerp, actual[0])
	th.CheckDeepEquals(t, AddressPoolDerp, actual[1])
}

func TestGetAddressPool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressPoolGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := addresspools.Get(client, "123914e3-36e4-41a8-a702-d9f6e54d7140").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, AddressPoolDerp, *actual)
}

func TestCreateAddressPool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressPoolCreationSuccessfully(t, AddressPoolSingleBody)

	name := "pxeboot"
	prefix := 24
	network := "169.254.202.0"
	order := "random"
	ranges := [][]string{{"169.254.202.1", "169.254.202.254"}}
	floating_addr := "169.254.202.2"
	c0_addr := "169.254.202.3"
	c1_addr := "169.254.202.4"
	
	actual, err := addresspools.Create(client.ServiceClient(), addresspools.AddressPoolOpts{
		Name:               &name,
		Prefix:             &prefix,
		Network:            &network,
		Order:              &order,
		Ranges:             &ranges,
		FloatingAddress:    &floating_addr,
		Controller0Address: &c0_addr,
		Controller1Address: &c1_addr,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, AddressPoolDerp, *actual)
}

func TestDeleteAddressPool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressPoolDeletionSuccessfully(t)

	res := addresspools.Delete(client.ServiceClient(), "123914e3-36e4-41a8-a702-d9f6e54d7140")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateAddressPool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddressPoolUpdateSuccessfully(t)

	name := "Changed"
	ranges := [][]string{{"169.254.202.1", "169.254.202.200"}}
	actual, err := addresspools.Update(client.ServiceClient(),
		"123914e3-36e4-41a8-a702-d9f6e54d7140",
		addresspools.AddressPoolOpts{Name: &name, Ranges: &ranges}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, AddressPoolDerp, *actual)
}
