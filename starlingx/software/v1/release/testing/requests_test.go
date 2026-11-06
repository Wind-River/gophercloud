/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/software/v1/release"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, ListResponse)

	actual, err := release.List(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedReleases, actual)
}

func TestListMetapackages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListMetapackagesSuccessfully(t, MetapackageResponse)

	actual, err := release.ListMetapackages(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedMetapackages, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t, "starlingx-27.03.0", GetResponse)

	actual, err := release.Get(client.ServiceClient(), "starlingx-27.03.0")
	th.AssertNoErr(t, err)
	if actual == nil {
		t.Fatalf("expected a release, got nil")
	}
	th.CheckDeepEquals(t, ExpectedRelease, *actual)
}
