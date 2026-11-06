/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/nfv/v1/swupgrade"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreateStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyCreationSuccessfully(t, SwUpgradeCreateResponse)

	actual, err := swupgrade.Create(client.ServiceClient(), swupgrade.SwUpgradeOpts{
		Release:                []string{"starlingx-27.03.0"},
		ControllerApplyType:    "serial",
		StorageApplyType:       "ignore",
		WorkerApplyType:        "serial",
		DefaultInstanceAction:  "stop-start",
		AlarmRestrictions:      "strict",
		MaxParallelWorkerHosts: intPtr(2),
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SwUpgradeReadyToApply, *actual)
}

func TestShowStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyShowSuccessfully(t, SwUpgradeCreateResponse)

	actual, err := swupgrade.Show(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SwUpgradeReadyToApply, *actual)
}

func TestApplyStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyApplySuccessfully(t, SwUpgradeApplyResponse)

	action := "apply-all"
	actual, err := swupgrade.ActionStrategy(client.ServiceClient(), swupgrade.StrategyActionOpts{
		Action: &action,
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SwUpgradeBuilding, *actual)
}

func TestDeleteStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyDeleteSuccessfully(t)

	res := swupgrade.Delete(client.ServiceClient())
	th.AssertNoErr(t, res.Err)
}

// intPtr returns a pointer to the given int, for optional *int opts fields.
func intPtr(n int) *int { return &n }

// TestCreateRollbackStrategy verifies that a rollback strategy (nil Release)
// omits the "release" key from the request body while still sending
// rollback:true and delete:false.
func TestCreateRollbackStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRollbackStrategyCreationSuccessfully(t, SwUpgradeCreateResponse)

	_, err := swupgrade.Create(client.ServiceClient(), swupgrade.SwUpgradeOpts{
		ControllerApplyType:   "serial",
		StorageApplyType:      "serial",
		WorkerApplyType:       "serial",
		DefaultInstanceAction: "stop-start",
		AlarmRestrictions:     "relaxed",
		Rollback:              true,
		Delete:                false,
		// Release intentionally left nil (rollback target is implicit).
	})
	th.AssertNoErr(t, err)
}
