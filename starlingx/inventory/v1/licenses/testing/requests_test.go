/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/licenses"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestCreateLicense(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLicenseCreationSuccessfully(t, LicenseCreateResponse)

	err := licenses.Create(client.ServiceClient(), licenses.LicenseOpts{
		Contents: []byte(license1),
	}).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestGetLicense(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLicenseGetSuccessfully(t)

	expected := licenses.License{Content: license1}

	client := client.ServiceClient()
	actual, err := licenses.Get(client).Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, &expected, actual)
}
