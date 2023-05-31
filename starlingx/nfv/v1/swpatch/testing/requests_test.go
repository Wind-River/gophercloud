/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/starlingx/nfv/v1/swpatch"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestCreateStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyCreationSuccessfully(t, SwPatchCreateResponse)
	controllerApplyType := "serial"
	swiftApplyType := "serial" // TODO: Delete when updating to system-config-update
	storageApplyType := "ignore"
	workerApplyType := "serial"
	defaultInstanceAction := "stop-start"
	alarmRestrictions := "strict"
	maxParallerWorkers := 2

	actual, err := swpatch.Create(client.ServiceClient(), swpatch.SwPatchOpts{
		ControllerApplyType:   controllerApplyType,
		StorageApplyType:      storageApplyType,
		WorkerApplyType:       workerApplyType,
		SwiftApplyType:		   swiftApplyType, // TODO: Delete when updating to system-config-update
		DefaultInstanceAction: defaultInstanceAction,
		AlarmRestrictions:     alarmRestrictions,
		MaxParallerWorkers:    maxParallerWorkers,
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, *actual, SwPatchHerp)
}

func TestShowStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyShowSuccessfully(t, SwPatchCreateResponse)

	client := client.ServiceClient()
	actual, err := swpatch.Show(client)
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}
	th.CheckDeepEquals(t, SwPatchHerp, *actual)
}

func TestApplyStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyApplySuccessfully(t, SwPatchApplyResponse)
	action := "apply-all"
	actual, err := swpatch.ActionStrategy(client.ServiceClient(), swpatch.StrategyActionOpts{
		Action: &action,
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, *actual, SwPatchDerp)
}
