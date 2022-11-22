/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019-2022 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/hosts"
	"github.com/gophercloud/gophercloud/testhelper/client"

	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestListHosts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHostListSuccessfully(t)

	pages := 0
	err := hosts.List(client.ServiceClient(), hosts.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := hosts.ExtractHosts(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 hosts, got %d", len(actual))
		}
		th.CheckDeepEquals(t, HostHerp, actual[0])
		th.CheckDeepEquals(t, HostDerp, actual[1])
		th.CheckDeepEquals(t, HostMerp, actual[2])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllHosts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHostListSuccessfully(t)

	allPages, err := hosts.List(client.ServiceClient(), hosts.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := hosts.ExtractHosts(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, HostHerp, actual[0])
	th.CheckDeepEquals(t, HostDerp, actual[1])
}

func TestGetHost(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHostGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := hosts.Get(client, "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, HostDerp, *actual)
}

func TestCreateHost(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHostCreationSuccessfully(t, SingleHostBody)

	hostname := HostDerp.Hostname
	personality := HostDerp.Personality
	subfunctions := HostDerp.SubFunctions
	location := hosts.LocationOpts{Name: icewallLocation}
	installOutput := HostDerp.InstallOutput
	console := HostDerp.Console
	maxCPUMhzConfigured := HostDerp.MaxCPUMhzConfigured
	actual, err := hosts.Create(client.ServiceClient(), hosts.HostOpts{
		Hostname:            &hostname,
		Personality:         &personality,
		SubFunctions:        &subfunctions,
		Location:            &location,
		InstallOutput:       &installOutput,
		Console:             &console,
		MaxCPUMhzConfigured: &maxCPUMhzConfigured,
		BootIP:              nil,
		BootMAC:             nil,
		RootDevice:          nil,
		BootDevice:          nil,
		BMAddress:           nil,
		BMType:              nil,
		BMUsername:          nil,
		BMPassword:          nil,
		Action:              nil,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, HostDerp, *actual)
}

func TestUpdateHost(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHostUpdateSuccessfully(t)

	newName := "new-name"
	actual, err := hosts.Update(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		hosts.HostOpts{Hostname: &newName}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, HostDerp, *actual)
}

func TestDeleteHost(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHostDeletionSuccessfully(t)

	res := hosts.Delete(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e")
	th.AssertNoErr(t, res.Err)
}
