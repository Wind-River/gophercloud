/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2022 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/ptpinstances"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"testing"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestListPTPInstances(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInstanceListSuccessfully(t)

	pages := 0
	err := ptpinstances.List(client.ServiceClient(), ptpinstances.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := ptpinstances.ExtractPTPInstances(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 ptp instances, got %d", len(actual))
		}
		th.CheckDeepEquals(t, PTPInstanceHerp, actual[0])
		th.CheckDeepEquals(t, PTPInstanceDerp, actual[1])

		return true, nil
	})
	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllPTPInstances(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInstanceListSuccessfully(t)

	allPages, err := ptpinstances.List(client.ServiceClient(), ptpinstances.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ptpinstances.ExtractPTPInstances(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PTPInstanceHerp, actual[0])
	th.CheckDeepEquals(t, PTPInstanceDerp, actual[1])
}

func TestCreatePTPInstance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInstanceCreationSuccessfully(t, PTPInstanceSingleBody)

	ptpInstanceName := "phc2sys1"
	ptpInstanceService := "phc2sys"
	actual, err := ptpinstances.Create(client.ServiceClient(), ptpinstances.PTPInstanceOpts{
		Name:    &ptpInstanceName,
		Service: &ptpInstanceService,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, PTPInstanceHerp, *actual)
}

func TestGetPTPInstance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInstanceGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := ptpinstances.Get(client, "fa5defce-2546-4786-ae58-7bb08e2105fc").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, PTPInstanceHerp, *actual)
}

func TestDeletePTPInstance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInstanceDeletionSuccessfully(t)

	res := ptpinstances.Delete(client.ServiceClient(), "fa5defce-2546-4786-ae58-7bb08e2105fc")
	th.AssertNoErr(t, res.Err)
}

func TestAddPTPParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddPTPParameterSuccessfully(t)

	newValue := "domainNumber=24"

	actual, err := ptpinstances.AddPTPParamToPTPInst(client.ServiceClient(),
		herpUUID,
		ptpinstances.PTPParamToPTPInstOpts{
			Parameter: &newValue,
		}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PTPInstanceHerpUpdated, *actual)
}

func TestRemovePTPParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRemovePTPParameterSuccessfully(t)

	newValue := "domainNumber=24"

	actual, err := ptpinstances.RemovePTPParamFromPTPInst(client.ServiceClient(),
		herpUUID,
		ptpinstances.PTPParamToPTPInstOpts{
			Parameter: &newValue,
		}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PTPInstanceHerp, *actual)
}

func TestAddToHost(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddInstanceToHostSuccessfully(t)

	herbID := 2

	actual, err := ptpinstances.AddPTPInstanceToHost(client.ServiceClient(),
	    controllerHostID,
		ptpinstances.PTPInstToHostOpts{
			PTPInstanceID: &herbID,
		}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PTPInstanceHerpAddToHost, *actual)
}

func TestRemoveFromHost(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRemoveInstanceFromHostSuccessfully(t)

	herbID := 2

	actual, err := ptpinstances.RemovePTPInstanceFromHost(client.ServiceClient(),
	    controllerHostID,
		ptpinstances.PTPInstToHostOpts{
			PTPInstanceID: &herbID,
		}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PTPInstanceHerp, *actual)
}

func TestListHostPTPInstances(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHostPTPInstanceListSuccessfully(t)

	allPages, err := ptpinstances.HostList(client.ServiceClient(), controllerHostID, ptpinstances.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ptpinstances.ExtractPTPInstances(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PTPInstanceHerp, actual[0])
	th.CheckDeepEquals(t, PTPInstanceDerp, actual[1])
}
