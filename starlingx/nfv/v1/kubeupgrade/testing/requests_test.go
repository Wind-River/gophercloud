/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/nfv/v1/kubeupgrade"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreateStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyCreationSuccessfully(t, KubeUpgradeCreateResponse)

	actual, err := kubeupgrade.Create(client.ServiceClient(), kubeupgrade.KubeUpgradeOpts{
		StorageApplyType:       "ignore",
		WorkerApplyType:        "serial",
		DefaultInstanceAction:  "stop-start",
		AlarmRestrictions:      "strict",
		MaxParallelWorkerHosts: intPtr(2),
		ToVersion:              "v1.30.6",
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, KubeUpgradeReadyToApply, *actual)
}

func TestShowStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyShowSuccessfully(t, KubeUpgradeCreateResponse)

	actual, err := kubeupgrade.Show(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, KubeUpgradeReadyToApply, *actual)
}

func TestApplyStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyApplySuccessfully(t, KubeUpgradeApplyResponse)

	action := "apply-all"
	actual, err := kubeupgrade.ActionStrategy(client.ServiceClient(), kubeupgrade.StrategyActionOpts{
		Action: &action,
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, KubeUpgradeBuilding, *actual)
}

func TestDeleteStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyDeleteSuccessfully(t)

	res := kubeupgrade.Delete(client.ServiceClient())
	th.AssertNoErr(t, res.Err)
}

// intPtr returns a pointer to the given int, for optional *int opts fields.
func intPtr(n int) *int { return &n }
