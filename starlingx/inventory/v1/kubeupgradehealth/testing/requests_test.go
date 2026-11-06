/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"strings"
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/kubeupgradehealth"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestKubeUpgradeHealthy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKubeUpgradeHealthSuccessfully(t, false, HealthyBody)

	body, err := kubeupgradehealth.KubeUpgrade(client.ServiceClient(), false)
	th.AssertNoErr(t, err)
	if !strings.Contains(body, "[OK]") || strings.Contains(body, "[Fail]") {
		t.Errorf("expected healthy body, got: %q", body)
	}
}

func TestKubeUpgradeHealthyRelaxed(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKubeUpgradeHealthSuccessfully(t, true, HealthyBody)

	body, err := kubeupgradehealth.KubeUpgrade(client.ServiceClient(), true)
	th.AssertNoErr(t, err)
	if !strings.Contains(body, "[OK]") {
		t.Errorf("expected healthy body, got: %q", body)
	}
}

func TestKubeUpgradeUnhealthy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleKubeUpgradeHealthSuccessfully(t, false, UnhealthyBody)

	body, err := kubeupgradehealth.KubeUpgrade(client.ServiceClient(), false)
	th.AssertNoErr(t, err)
	if !strings.Contains(body, "[Fail]") {
		t.Errorf("expected unhealthy body with [Fail], got: %q", body)
	}
}
