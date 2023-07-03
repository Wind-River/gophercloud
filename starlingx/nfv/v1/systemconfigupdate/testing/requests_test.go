/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/starlingx/nfv/v1/systemconfigupdate"
	"github.com/gophercloud/gophercloud/testhelper/client"

	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestCreateStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyCreationSuccessfully(t, SCUpdateCreateResponse)
	controllerApplyType := "serial"
	storageApplyType := "ignore"
	workerApplyType := "serial"
	defaultInstanceAction := "stop-start"
	alarmRestrictions := "strict"
	maxParallerWorkers := 2

	actual, err := systemconfigupdate.Create(client.ServiceClient(), systemconfigupdate.SystemConfigUpdateOpts{
		ControllerApplyType:   controllerApplyType,
		StorageApplyType:      storageApplyType,
		WorkerApplyType:       workerApplyType,
		DefaultInstanceAction: defaultInstanceAction,
		AlarmRestrictions:     alarmRestrictions,
		MaxParallerWorkers:    maxParallerWorkers,
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, *actual, SCUpdateHerp)
}

func TestShowStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyShowSuccessfully(t, SCUpdateCreateResponse)

	client := client.ServiceClient()
	actual, err := systemconfigupdate.Show(client)
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}
	th.CheckDeepEquals(t, SCUpdateHerp, *actual)
}

func TestApplyStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleStrategyApplySuccessfully(t, SCUpdateApplyResponse)
	action := "apply-all"
	actual, err := systemconfigupdate.ActionStrategy(client.ServiceClient(), systemconfigupdate.StrategyActionOpts{
		Action: &action,
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, *actual, SCUpdateDerp)
}
