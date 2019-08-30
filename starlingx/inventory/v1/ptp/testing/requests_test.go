/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/ptp"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListPTPs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPListSuccessfully(t)

	pages := 0
	err := ptp.List(client.ServiceClient(), ptp.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := ptp.ExtractPTPs(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 PTPs, got %d", len(actual))
		}
		th.CheckDeepEquals(t, PTPHerp, actual[0])
		th.CheckDeepEquals(t, PTPDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllPTPs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPListSuccessfully(t)

	allPages, err := ptp.List(client.ServiceClient(), ptp.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ptp.ExtractPTPs(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PTPHerp, actual[0])
	th.CheckDeepEquals(t, PTPDerp, actual[1])
}

func TestUpdatePTP(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPUpdateSuccessfully(t)

	newTransport := "l2"
	actual, err := ptp.Update(client.ServiceClient(),
		"d87feed9-e351-40fc-8356-7bf6a59750ea",
		ptp.PTPOpts{Transport: &newTransport}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PTPDerp, *actual)
}

func TestGetPTP(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := ptp.Get(client, "d87feed9-e351-40fc-8356-7bf6a59750ea").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, PTPDerp, *actual)
}
