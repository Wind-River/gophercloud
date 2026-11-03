/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/nfv/v1/swupgrade"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var (
	// SwUpgradeReadyToApply is the expected strategy after creation.
	SwUpgradeReadyToApply = swupgrade.SwUpgrade{
		ID:                    "5dd16d94-dfc5-4029-bfcb-d815e7c2dc3d",
		StrategyName:          "sw-upgrade",
		ControllerApplyType:   "serial",
		StorageApplyType:      "ignore",
		WorkerApplyType:       "serial",
		MaxParallelWorkers:    2,
		DefaultInstanceAction: "stop-start",
		AlarmRestrictions:     "strict",
		State:                 "ready-to-apply",
	}
	// SwUpgradeBuilding is the expected strategy after apply.
	SwUpgradeBuilding = swupgrade.SwUpgrade{
		ID:                    "5dd16d94-dfc5-4029-bfcb-d815e7c2dc3d",
		StrategyName:          "sw-upgrade",
		ControllerApplyType:   "serial",
		StorageApplyType:      "ignore",
		WorkerApplyType:       "serial",
		MaxParallelWorkers:    2,
		DefaultInstanceAction: "stop-start",
		AlarmRestrictions:     "strict",
		State:                 "building",
	}
)

const SwUpgradeCreateResponse = `
{
	"strategy":
	{
		"controller-apply-type": "serial",
		"storage-apply-type": "ignore",
		"worker-apply-type": "serial",
		"state": "ready-to-apply",
		"default-instance-action": "stop-start",
		"current-phase": "build",
		"current-phase-completion-percentage": 0,
		"max-parallel-worker-hosts": 2,
		"alarm-restrictions": "strict",
		"uuid": "5dd16d94-dfc5-4029-bfcb-d815e7c2dc3d",
		"name": "sw-upgrade"
	}
}
`

const SwUpgradeApplyResponse = `
{
	"strategy":
	{
		"controller-apply-type": "serial",
		"storage-apply-type": "ignore",
		"worker-apply-type": "serial",
		"current-phase": "build",
		"current-phase-completion-percentage": 0,
		"state": "building",
		"default-instance-action": "stop-start",
		"max-parallel-worker-hosts": 2,
		"alarm-restrictions": "strict",
		"uuid": "5dd16d94-dfc5-4029-bfcb-d815e7c2dc3d",
		"name": "sw-upgrade"
	}
}
`

// HandleStrategyCreationSuccessfully verifies the create request body. The
// release field is sent as a JSON list (post-componentization format) and the
// optional pointer fields are omitted because they are left nil.
func HandleStrategyCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/api/orchestration/sw-upgrade/strategy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, `{
			"alarm-restrictions": "strict",
			"controller-apply-type": "serial",
			"default-instance-action": "stop-start",
			"delete": false,
			"max-parallel-worker-hosts": 2,
			"release": ["starlingx-27.03.0"],
			"rollback": false,
			"storage-apply-type": "ignore",
			"worker-apply-type": "serial"
		  }`)

			w.WriteHeader(http.StatusAccepted)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, response)
		})
}

func HandleStrategyApplySuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/api/orchestration/sw-upgrade/strategy/actions",
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
	th.Mux.HandleFunc("/api/orchestration/sw-upgrade/strategy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, response)
		})
}

func HandleStrategyDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/api/orchestration/sw-upgrade/strategy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.WriteHeader(http.StatusNoContent)
		})
}

// HandleRollbackStrategyCreationSuccessfully verifies a rollback create request.
// For a rollback strategy the "release" field MUST be omitted entirely (VIM
// rejects "Cannot set both release and rollback"), while "rollback" is true and
// "delete" is false (both required-present).
func HandleRollbackStrategyCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/api/orchestration/sw-upgrade/strategy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, `{
			"alarm-restrictions": "relaxed",
			"controller-apply-type": "serial",
			"default-instance-action": "stop-start",
			"delete": false,
			"rollback": true,
			"storage-apply-type": "serial",
			"worker-apply-type": "serial"
		  }`)

			w.WriteHeader(http.StatusAccepted)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, response)
		})
}
