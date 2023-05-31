/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/nfv/v1/swpatch"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var (
	SwPatchHerp = swpatch.SwPatch{
		ID:                    "5dd16d94-dfc5-4029-bfcb-d815e7c2dc3d",
		ControllerApplyType:   "serial",
		StrategyName:          "sw-patch",
		StorageApplyType:      "ignore",
		SwiftApplyType: 	   "serial", // TODO: Delete when updating to system-config-update
		WorkerApplyType:       "serial",
		MaxParallerWorkers:    2,
		DefaultInstanceAction: "stop-start",
		AlarmRestrictions:     "strict",
		State:                 "ready-to-apply",
	}
	SwPatchDerp = swpatch.SwPatch{
		ID:                    "5dd16d94-dfc5-4029-bfcb-d815e7c2dc3d",
		ControllerApplyType:   "serial",
		StrategyName:          "sw-patch",
		StorageApplyType:      "ignore",
		SwiftApplyType: 	   "serial", // TODO: Delete when updating to system-config-update
		WorkerApplyType:       "serial",
		MaxParallerWorkers:    2,
		DefaultInstanceAction: "stop-start",
		AlarmRestrictions:     "strict",
		State:                 "building",
	}
)

const SwPatchCreateResponse = `
{
	"strategy":
	{
		"controller-apply-type": "serial",
		"storage-apply-type": "ignore",
		"worker-apply-type": "serial",
		"state": "ready-to-apply",
		"swift-apply-type": "serial", // TODO: Delete when updating to system-config-update
		"default-instance-action": "stop-start",
		"current-phase": "build",
		"current-phase-completion-percentage": 0,
		"max-parallel-worker-hosts": 2,
		"alarm-restrictions": "strict",
		"uuid": "5dd16d94-dfc5-4029-bfcb-d815e7c2dc3d",
		"name": "sw-patch"
	}
}
`
const SwPatchApplyResponse = `
{
	"strategy":
	{
		"controller-apply-type": "serial",
		"storage-apply-type": "ignore",
		"worker-apply-type": "serial",
		"current-phase": "build",
		"swift-apply-type": "serial", // TODO: Delete when updating to system-config-update
		"current-phase-completion-percentage": 0,
		"state": "building",
		"default-instance-action": "stop-start",
		"max-parallel-worker-hosts": 2,
		"alarm-restrictions": "strict",
		"uuid": "5dd16d94-dfc5-4029-bfcb-d815e7c2dc3d",
		"name": "sw-patch"
	}
}
`

func HandleStrategyCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/api/orchestration/sw-patch/strategy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, `{
			"AlarmRestrictions": "strict",
			"ControllerApplyType": "serial",
			"SwiftApplyType": "serial", // TODO: Delete when updating to system-config-update
			"DefaultInstanceAction": "stop-start",
			"MaxParallerWorkers": 2,
			"StorageApplyType": "ignore",
			"WorkerApplyType": "serial"
		  }`)

			w.WriteHeader(http.StatusAccepted)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, response)
		})
}

func HandleStrategyApplySuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/api/orchestration/sw-patch/strategy/actions",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, `{
			"action": "apply-all"
		  }`)
			w.WriteHeader(http.StatusAccepted)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, response)
		})
}

func HandleStrategyShowSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/api/orchestration/sw-patch/strategy", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, SwPatchCreateResponse)
	})
}
