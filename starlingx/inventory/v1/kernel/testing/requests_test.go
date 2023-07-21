/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/kernel"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestGetKernel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKernelGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := kernel.Get(client, TestHostUUID).Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, KernelHerp, *actual)
}

func TestModifyKernel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKernelUpdateSuccessfully(t)

	newKernel := "lowlatency"
	actual, err := kernel.Update(client.ServiceClient(), TestHostUUID,
		kernel.KernelOpts{Kernel: &newKernel}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, KernelDerp, *actual)
}
