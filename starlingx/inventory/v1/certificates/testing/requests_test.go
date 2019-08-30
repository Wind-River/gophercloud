/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/certificates"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListCertificates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateListSuccessfully(t)

	pages := 0
	err := certificates.List(client.ServiceClient(), certificates.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := certificates.ExtractCertificates(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 certificates, got %d", len(actual))
		}
		th.CheckDeepEquals(t, CertificateHerp, actual[0])
		th.CheckDeepEquals(t, CertificateDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllCertificates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateListSuccessfully(t)

	allPages, err := certificates.List(client.ServiceClient(), certificates.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := certificates.ExtractCertificates(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, CertificateHerp, actual[0])
	th.CheckDeepEquals(t, CertificateDerp, actual[1])
}

func TestGetCertificate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := certificates.Get(client, "f757b5c7-89ab-4d93-bfd7-a97780ec2c1e").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, CertificateHerp, *actual)
}

func TestCreateCertificate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateCreationSuccessfully(t, CertificateCreateResponse)

	actual, err := certificates.Create(client.ServiceClient(), certificates.CertificateOpts{
		Type: "ssl_ca",
		File: []byte("foobar"),
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, CertificateHerp, *actual)
}
