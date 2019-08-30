/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/ntp"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListNTPs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNTPListSuccessfully(t)

	pages := 0
	err := ntp.List(client.ServiceClient(), ntp.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := ntp.ExtractNTPs(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 NTPs, got %d", len(actual))
		}
		th.CheckDeepEquals(t, NTPHerp, actual[0])
		th.CheckDeepEquals(t, NTPDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllNTPs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNTPListSuccessfully(t)

	allPages, err := ntp.List(client.ServiceClient(), ntp.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ntp.ExtractNTPs(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, NTPHerp, actual[0])
	th.CheckDeepEquals(t, NTPDerp, actual[1])
}

func TestUpdateNTP(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNTPUpdateSuccessfully(t)

	newServer := "1.pool.ntp.org"
	actual, err := ntp.Update(client.ServiceClient(),
		"92939488-6f53-4913-aa15-ce89162751c6",
		ntp.NTPOpts{NTPServers: &newServer}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, NTPDerp, *actual)
}

func TestGetNTP(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNTPGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := ntp.Get(client, "92939488-6f53-4913-aa15-ce89162751c6").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, NTPDerp, *actual)
}
