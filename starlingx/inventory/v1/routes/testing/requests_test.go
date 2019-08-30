/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/routes"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListRoutes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRouteListSuccessfully(t)

	pages := 0
	err := routes.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", routes.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := routes.ExtractRoutes(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 Routes, got %d", len(actual))
		}
		th.CheckDeepEquals(t, RouteHerp, actual[0])
		th.CheckDeepEquals(t, RouteDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllRoutes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRouteListSuccessfully(t)

	allPages, err := routes.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		routes.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := routes.ExtractRoutes(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, RouteHerp, actual[0])
	th.CheckDeepEquals(t, RouteDerp, actual[1])
}

func TestGetRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRouteGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := routes.Get(client, "354968fc-6f18-46dc-93a1-6118280e3cee").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, RouteDerp, *actual)
}

func TestCreateRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRouteCreationSuccessfully(t, RouteSingleBody)

	InterfaceUUID := "9a01253c-17bd-453b-8dcf-038da4636100"
	Network := "0.0.0.0"
	Prefix := 0
	Gateway := "192.168.59.1"
	Metric := 1
	actual, err := routes.Create(client.ServiceClient(), routes.RouteOpts{
		InterfaceUUID: &InterfaceUUID,
		Network:       &Network,
		Prefix:        &Prefix,
		Gateway:       &Gateway,
		Metric:        &Metric,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, RouteDerp, *actual)
}

func TestDeleteRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRouteDeletionSuccessfully(t)

	res := routes.Delete(client.ServiceClient(), "354968fc-6f18-46dc-93a1-6118280e3cee")
	th.AssertNoErr(t, res.Err)
}
