/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/snmpCommunity"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListSNMPCommunitys(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPCommunityListSuccessfully(t)

	pages := 0
	err := snmpCommunity.List(client.ServiceClient(), snmpCommunity.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := snmpCommunity.ExtractSNMPCommunities(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 SNMPCommunitys, got %d", len(actual))
		}
		th.CheckDeepEquals(t, SNMPCommunityHerp, actual[0])
		th.CheckDeepEquals(t, SNMPCommunityDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllSNMPCommunitys(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPCommunityListSuccessfully(t)

	allPages, err := snmpCommunity.List(client.ServiceClient(),
		snmpCommunity.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := snmpCommunity.ExtractSNMPCommunities(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SNMPCommunityHerp, actual[0])
	th.CheckDeepEquals(t, SNMPCommunityDerp, actual[1])
}

func TestGetSNMPCommunity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPCommunityGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := snmpCommunity.Get(client, "5144712c-3133-4a91-b3b7-8767da6f0680").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, SNMPCommunityDerp, *actual)
}

func TestCreateSNMPCommunity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPCommunityCreationSuccessfully(t, SNMPCommunitySingleBody)

	Community := "test2"
	actual, err := snmpCommunity.Create(client.ServiceClient(), snmpCommunity.SNMPCommunityOpts{
		Community: &Community,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, SNMPCommunityDerp, *actual)
}

func TestDeleteSNMPCommunity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSNMPCommunityDeletionSuccessfully(t)

	res := snmpCommunity.Delete(client.ServiceClient(), "5144712c-3133-4a91-b3b7-8767da6f0680")
	th.AssertNoErr(t, res.Err)
}
