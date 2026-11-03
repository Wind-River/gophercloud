/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/kubeversions"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListKubeVersions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKubeVersionListSuccessfully(t)

	pages := 0
	err := kubeversions.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := kubeversions.ExtractKubeVersions(page)
		if err != nil {
			return false, err
		}
		if len(actual) != 3 {
			t.Fatalf("expected 3 versions, got %d", len(actual))
		}
		th.CheckDeepEquals(t, KubeVersionUnavailable, actual[0])
		th.CheckDeepEquals(t, KubeVersionTarget, actual[1])
		th.CheckDeepEquals(t, KubeVersionActive, actual[2])
		return true, nil
	})
	th.AssertNoErr(t, err)
	if pages != 1 {
		t.Errorf("expected 1 page, saw %d", pages)
	}
}

func TestListKubeVersionsConvenience(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKubeVersionListSuccessfully(t)

	actual, err := kubeversions.ListKubeVersions(client.ServiceClient())
	th.AssertNoErr(t, err)
	if len(actual) != 3 {
		t.Fatalf("expected 3 versions, got %d", len(actual))
	}
	th.CheckDeepEquals(t, KubeVersionActive, actual[2])
}
