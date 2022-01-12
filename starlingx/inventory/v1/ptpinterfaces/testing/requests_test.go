/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2022 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/ptpinterfaces"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"testing"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestListPTPInterfaces(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInterfaceListSuccessfully(t)

	pages := 0
	err := ptpinterfaces.List(client.ServiceClient(), ptpinterfaces.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := ptpinterfaces.ExtractPTPInterfaces(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 ptp interfaces, got %d", len(actual))
		}
		th.CheckDeepEquals(t, PTPInterfaceHerp, actual[0])
		th.CheckDeepEquals(t, PTPInterfaceDerp, actual[1])

		return true, nil
	})
	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllPTPInterfaces(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInterfaceListSuccessfully(t)

	allPages, err := ptpinterfaces.List(client.ServiceClient(), ptpinterfaces.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ptpinterfaces.ExtractPTPInterfaces(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PTPInterfaceHerp, actual[0])
	th.CheckDeepEquals(t, PTPInterfaceDerp, actual[1])
}

func TestListHostPTPInterfaces(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHostPTPInterfaceListSuccessfully(t)

	allPages, err := ptpinterfaces.HostList(client.ServiceClient(),
		controllerHostID,
		ptpinterfaces.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ptpinterfaces.ExtractPTPInterfaces(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PTPInterfaceHerp, actual[0])
	th.CheckDeepEquals(t, PTPInterfaceDerp, actual[1])
}

func TestListInterfacePTPInterfaces(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleIntPTPInterfaceListSuccessfully(t)

	allPages, err := ptpinterfaces.InterfaceList(client.ServiceClient(),
	    interfaceUUID,
		ptpinterfaces.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ptpinterfaces.ExtractPTPInterfaces(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PTPInterfaceHerp, actual[0])
	th.CheckDeepEquals(t, PTPInterfaceDerp, actual[1])
}

func TestCreatePTPInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInterfaceCreationSuccessfully(t, PTPInterfaceSingleBody)

	ptpInterfaceName := "ptpint1"
	ptpInstanceUUID := "53041360-451f-49ea-8843-44fab16f6628"
	actual, err := ptpinterfaces.Create(client.ServiceClient(), ptpinterfaces.PTPInterfaceOpts{
		Name:    		 &ptpInterfaceName,
		PTPInstanceUUID: &ptpInstanceUUID,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, PTPInterfaceHerp, *actual)
}

func TestGetPTPInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInterfaceGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := ptpinterfaces.Get(client, "b7d51ba0-35d7-4bab-9e27-a8b701587c54").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, PTPInterfaceHerp, *actual)
}

func TestDeletePTPInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePTPInterfaceDeletionSuccessfully(t)

	res := ptpinterfaces.Delete(client.ServiceClient(), "b7d51ba0-35d7-4bab-9e27-a8b701587c54")
	th.AssertNoErr(t, res.Err)
}

func TestAddPTPParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddPTPParameterSuccessfully(t)

	newValue := "masterOnly=0"

	actual, err := ptpinterfaces.AddPTPParamToPTPInt(client.ServiceClient(),
		herpUUID,
		ptpinterfaces.PTPParamToPTPIntOpts{
			Parameter: &newValue,
		}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PTPInterfaceHerpUpdated, *actual)
}

func TestRemovePTPParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRemovePTPParameterSuccessfully(t)

	newValue := "masterOnly=0"

	actual, err := ptpinterfaces.RemovePTPParamFromPTPInt(client.ServiceClient(),
		herpUUID,
		ptpinterfaces.PTPParamToPTPIntOpts{
			Parameter: &newValue,
		}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PTPInterfaceHerp, *actual)
}

func TestAddPTPIntToInt(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddPTPIntToIntSuccessfully(t)

	PTPinterfaceID := 3

	actual, err := ptpinterfaces.AddPTPIntToInt(client.ServiceClient(),
	    interfaceUUID,
		ptpinterfaces.PTPIntToIntOpt{
			PTPinterfaceID: &PTPinterfaceID,
		}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PTPIntHerpAssignedToInt, *actual)
}

func TestRemovePTPIntFromInt(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRemovePTPIntFromIntSuccessfully(t)

	PTPinterfaceID := 3

	actual, err := ptpinterfaces.RemovePTPIntFromInt(client.ServiceClient(),
	    interfaceUUID,
		ptpinterfaces.PTPIntToIntOpt{
			PTPinterfaceID: &PTPinterfaceID,
		}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PTPInterfaceHerp, *actual)
}
