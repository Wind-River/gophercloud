/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/ports"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListPorts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePortListSuccessfully(t)

	pages := 0
	err := ports.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", ports.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := ports.ExtractPorts(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 Ports, got %d", len(actual))
		}
		th.CheckDeepEquals(t, PortHerp, actual[0])
		th.CheckDeepEquals(t, PortDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllPorts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePortListSuccessfully(t)

	allPages, err := ports.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		ports.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ports.ExtractPorts(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PortHerp, actual[0])
	th.CheckDeepEquals(t, PortDerp, actual[1])
}

func TestGetPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePortGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := ports.Get(client, "c3896f57-59c0-4362-b04e-62cbcf76e5c9").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, PortDerp, *actual)
}
