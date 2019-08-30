/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/dns"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"testing"
)

func TestListDNSs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDNSListSuccessfully(t)

	pages := 0
	err := dns.List(client.ServiceClient(), dns.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := dns.ExtractDNSs(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 DNSs, got %d", len(actual))
		}
		th.CheckDeepEquals(t, DNSHerp, actual[0])
		th.CheckDeepEquals(t, DNSDerp, actual[1])
		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllDNSs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDNSListSuccessfully(t)

	allPages, err := dns.List(client.ServiceClient(), dns.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := dns.ExtractDNSs(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, DNSHerp, actual[0])
	th.CheckDeepEquals(t, DNSDerp, actual[1])
}

func TestUpdateDNS(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDNSUpdateSuccessfully(t)

	newNameServer := "128.224.144.130"
	actual, err := dns.Update(client.ServiceClient(),
		"e60b7d12-7585-486e-9c27-3d16e0daba09",
		dns.DNSOpts{Nameservers: &newNameServer}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, DNSDerp, *actual)
}

func TestGetDNS(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDNSGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := dns.Get(client, "e60b7d12-7585-486e-9c27-3d16e0daba09").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, DNSDerp, *actual)
}
