/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/interfaces"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListInterfaces(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceListSuccessfully(t)

	pages := 0
	err := interfaces.List(client.ServiceClient(),
		"f73dda8e-be3c-4704-ad1e-ed99e44b846e",
		interfaces.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := interfaces.ExtractInterfaces(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 4 {
			t.Fatalf("Expected 4 interfaces, got %d", len(actual))
		}
		th.CheckDeepEquals(t, InterfaceHerp, actual[0])
		th.CheckDeepEquals(t, InterfaceDerp, actual[1])
		th.CheckDeepEquals(t, InterfaceMerp, actual[2])
		th.CheckDeepEquals(t, InterfaceBerp, actual[3])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllInterfaces(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceListSuccessfully(t)

	allPages, err := interfaces.List(client.ServiceClient(),
		"f73dda8e-be3c-4704-ad1e-ed99e44b846e",
		interfaces.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := interfaces.ExtractInterfaces(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, InterfaceHerp, actual[0])
	th.CheckDeepEquals(t, InterfaceDerp, actual[1])
}

func TestGetInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := interfaces.Get(client, "a5965fee-dc60-40dc-a234-edf87f1f9380").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, InterfaceDerp, *actual)
}

func TestCreateInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceCreationSuccessfully(t, InterfaceSingleBody)

	interfaceHostID := "f73dda8e-be3c-4704-ad1e-ed99e44b846e"
	interfaceName := "Derp"
	interfaceType := interfaces.IFTypeVirtual
	interfaceClass := interfaces.IFClassData
	interfaceMTU := 1400
	interfaceIPv4Mode := interfaces.AddressModePool
	interfaceIPv4Pool := "dbfa4c2e-4526-4aaf-b07b-a3da7aeb6c26"
	interfaceIPv6Mode := interfaces.AddressModePool
	interfaceIPv6Pool := "934d8341-5114-46d2-9560-7c47618892c7"
	interfaceVFCount := 1
	interfaceUses := []string{}
	interfaceUsers := []string{}
	ptpRole := interfaces.PTPRoleMaster
	actual, err := interfaces.Create(client.ServiceClient(), interfaces.InterfaceOpts{
		HostUUID:         &interfaceHostID,
		Type:             &interfaceType,
		Name:             &interfaceName,
		Class:            &interfaceClass,
		MTU:              &interfaceMTU,
		VID:              nil,
		IPv4Mode:         &interfaceIPv4Mode,
		IPv4Pool:         &interfaceIPv4Pool,
		IPv6Mode:         &interfaceIPv6Mode,
		IPv6Pool:         &interfaceIPv6Pool,
		Networks:         nil,
		NetworksToAdd:    nil,
		NetworksToDelete: nil,
		DataNetworks:     nil,
		AEMode:           nil,
		AETransmitHash:   nil,
		VFCount:          &interfaceVFCount,
		Uses:             &interfaceUses,
		UsesModify:       &interfaceUsers,
		PTPRole:          &ptpRole,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, InterfaceDerp, *actual)
}

func TestUpdateInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceUpdateSuccessfully(t)

	newName := "new-name"
	actual, err := interfaces.Update(client.ServiceClient(), "a5965fee-dc60-40dc-a234-edf87f1f9380", interfaces.InterfaceOpts{Name: &newName}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, InterfaceDerp, *actual)
}

func TestDeleteInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInterfaceDeletionSuccessfully(t)

	res := interfaces.Delete(client.ServiceClient(), "a5965fee-dc60-40dc-a234-edf87f1f9380")
	th.AssertNoErr(t, res.Err)
}
