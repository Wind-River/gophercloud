/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/snmpTrapDest"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListSNMPTrapDests(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPTrapDestListSuccessfully(t)

	pages := 0
	err := snmpTrapDest.List(client.ServiceClient(), snmpTrapDest.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := snmpTrapDest.ExtractSNMPTrapDests(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 SNMPTrapDests, got %d", len(actual))
		}
		th.CheckDeepEquals(t, SNMPTrapDestHerp, actual[0])
		th.CheckDeepEquals(t, SNMPTrapDestDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllSNMPTrapDests(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPTrapDestListSuccessfully(t)

	allPages, err := snmpTrapDest.List(client.ServiceClient(),
		snmpTrapDest.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := snmpTrapDest.ExtractSNMPTrapDests(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SNMPTrapDestHerp, actual[0])
	th.CheckDeepEquals(t, SNMPTrapDestDerp, actual[1])
}

func TestGetSNMPTrapDest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPTrapDestGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := snmpTrapDest.Get(client, "e30da84e-4e9f-4b6e-969f-2a4c41345e47").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, SNMPTrapDestDerp, *actual)
}

func TestCreateSNMPTrapDest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPTrapDestCreationSuccessfully(t, SNMPTrapDestSingleBody)

	community := "5144712c-3133-4a91-b3b7-8767da6f0680"
	ipaddress := "192.168.2.100"
	actual, err := snmpTrapDest.Create(client.ServiceClient(), snmpTrapDest.SNMPTrapDestOpts{
		Community: &community,
		IPAddress: &ipaddress,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, SNMPTrapDestDerp, *actual)
}

func TestDeleteSNMPTrapDest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPTrapDestDeletionSuccessfully(t)

	res := snmpTrapDest.Delete(client.ServiceClient(), "e30da84e-4e9f-4b6e-969f-2a4c41345e47")
	th.AssertNoErr(t, res.Err)
}
