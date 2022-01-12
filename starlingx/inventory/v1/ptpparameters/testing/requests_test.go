/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2022 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/ptpparameters"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"testing"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestListPTPParameters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPParameterListSuccessfully(t)

	pages := 0
	err := ptpparameters.List(client.ServiceClient(), ptpparameters.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := ptpparameters.ExtractPTPParameters(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 ptp parameters, got %d", len(actual))
		}
		th.CheckDeepEquals(t, PTPParameterHerp, actual[0])
		th.CheckDeepEquals(t, PTPParameterDerp, actual[1])

		return true, nil
	})
	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllPTPParameters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPParameterListSuccessfully(t)

	allPages, err := ptpparameters.List(client.ServiceClient(), ptpparameters.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ptpparameters.ExtractPTPParameters(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PTPParameterHerp, actual[0])
	th.CheckDeepEquals(t, PTPParameterDerp, actual[1])
}

func TestGetPTPParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPParameterGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := ptpparameters.Get(client, "868e0ab8-2bc3-4d92-b736-de06bb8feb12").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, PTPParameterDerp, *actual)
}

func TestDeletePTPParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPParameterDeletionSuccessfully(t)

	res := ptpparameters.Delete(client.ServiceClient(), "868e0ab8-2bc3-4d92-b736-de06bb8feb12")
	th.AssertNoErr(t, res.Err)
}
