/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/networks"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkListSuccessfully(t)

	pages := 0
	err := networks.List(client.ServiceClient(), networks.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := networks.ExtractNetworks(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 networks, got %d", len(actual))
		}
		th.CheckDeepEquals(t, NetworkHerp, actual[0])
		th.CheckDeepEquals(t, NetworkDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkListSuccessfully(t)

	allPages, err := networks.List(client.ServiceClient(), networks.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := networks.ExtractNetworks(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, NetworkHerp, actual[0])
	th.CheckDeepEquals(t, NetworkDerp, actual[1])
}

func TestGetNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := networks.Get(client, "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, NetworkDerp, *actual)
}

func TestCreateNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkCreationSuccessfully(t, SingleNetworkBody)

	networkName := "Derp"
	networkType := "oam"
	networkDynamic := false
	poolUUID := "c7ac5a0c-606b-4fe0-9065-28a8c8fb78cc"
	actual, err := networks.Create(client.ServiceClient(), networks.NetworkOpts{
		Name:     &networkName,
		Type:     &networkType,
		Dynamic:  &networkDynamic,
		PoolUUID: &poolUUID,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, NetworkDerp, *actual)
}

func TestUpdateNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkUpdateSuccessfully(t)

	newName := "new-name"
	actual, err := networks.Update(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", networks.NetworkOpts{Name: &newName}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, NetworkDerp, *actual)
}

func TestDeleteNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkDeletionSuccessfully(t)

	res := networks.Delete(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e")
	th.AssertNoErr(t, res.Err)
}
