/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/clusters"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"testing"
)

func TestListClusters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleClusterListSuccessfully(t)

	pages := 0
	err := clusters.List(client.ServiceClient(), clusters.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := clusters.ExtractClusters(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 clusters, got %d", len(actual))
		}
		th.CheckDeepEquals(t, ClusterHerp, actual[0])
		th.CheckDeepEquals(t, ClusterDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllClusters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleClusterListSuccessfully(t)

	allPages, err := clusters.List(client.ServiceClient(), clusters.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := clusters.ExtractClusters(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ClusterHerp, actual[0])
	th.CheckDeepEquals(t, ClusterDerp, actual[1])
}

func TestGetCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSystemGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := clusters.Get(client, "27caf12f-19af-4cd3-bc41-d4467ec80e39").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, ClusterDerp, *actual)
}
