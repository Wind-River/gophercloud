/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/labels"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListLabels(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLabelListSuccessfully(t)

	pages := 0
	err := labels.List(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", labels.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := labels.ExtractLabels(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 Labels, got %d", len(actual))
		}
		th.CheckDeepEquals(t, LabelHerp, actual[0])
		th.CheckDeepEquals(t, LabelDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllLabels(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLabelListSuccessfully(t)

	allPages, err := labels.List(client.ServiceClient(),
		"f757b5c7-89ab-4d93-bfd7-a97780ec2c1e",
		labels.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := labels.ExtractLabels(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, LabelHerp, actual[0])
	th.CheckDeepEquals(t, LabelDerp, actual[1])
}

func TestGetLabel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLabelGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := labels.Get(client, "c0940c5f-dedf-4010-9cca-4d9b7f5dacec").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, LabelDerp, *actual)
}

func TestAssignLabel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLabelAssignSuccessfully(t, LabelSingleBody)

	m := make(map[string]string)
	m["openstack-compute-node"] = "enabled"

	actual, err := labels.Create(client.ServiceClient(), "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e", m).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, LabelDerp, *actual)
}

func TestDeleteLabel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLabelDeletionSuccessfully(t)

	res := labels.Delete(client.ServiceClient(), "c0940c5f-dedf-4010-9cca-4d9b7f5dacec")
	th.AssertNoErr(t, res.Err)
}
