/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/oamNetworks"
	"github.com/gophercloud/gophercloud/testhelper/client"

	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestListOAMNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOAMNetworkListSuccessfully(t)

	pages := 0
	err := oamNetworks.List(client.ServiceClient(), oamNetworks.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := oamNetworks.ExtractOAMNetworks(page)
		if err != nil {
			return false, err
		}

		th.CheckDeepEquals(t, OAMNetworkHerp, actual[0])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}
func TestGetNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOAMNetworkGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := oamNetworks.Get(client, "fd5aaa82-b503-40e2-af45-9fc4411df7a0").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, OAMNetworkDerp, *actual)
}

func TestUpdateNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOAMNetworkUpdateSuccessfully(t)

	newFloatingIP := "10.10.20.10"
	actual, err := oamNetworks.Update(client.ServiceClient(), "727bd796-070f-40c2-8b9b-7ed674fd0fe7", oamNetworks.OAMNetworkOpts{OAMFloatingIP: &newFloatingIP}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, OAMNetworkDerp, *actual)
}
