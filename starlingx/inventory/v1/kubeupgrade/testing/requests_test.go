/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/kubeupgrade"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListKubeUpgrade(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKubeUpgradeListSuccessfully(t)

	actual, err := kubeupgrade.List(client.ServiceClient())
	th.AssertNoErr(t, err)
	if len(actual) != 1 {
		t.Fatalf("expected 1 kube upgrade, got %d", len(actual))
	}
	th.CheckDeepEquals(t, KubeUpgradeInProgress, actual[0])
}

func TestListKubeUpgradeEmpty(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKubeUpgradeListEmpty(t)

	actual, err := kubeupgrade.List(client.ServiceClient())
	th.AssertNoErr(t, err)
	if len(actual) != 0 {
		t.Fatalf("expected 0 kube upgrades, got %d", len(actual))
	}
}

func TestDeleteKubeUpgrade(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKubeUpgradeDeleteSuccessfully(t)

	err := kubeupgrade.Delete(client.ServiceClient()).ExtractErr()
	th.AssertNoErr(t, err)
}
