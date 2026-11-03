/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/nfv/v1/kubeupgrade"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var (
	// KubeUpgradeReadyToApply is the expected strategy after creation.
	KubeUpgradeReadyToApply = kubeupgrade.KubeUpgrade{
		ID:                    "7aa16d94-dfc5-4029-bfcb-d815e7c2dc99",
		StrategyName:          "kube-upgrade",
		StorageApplyType:      "ignore",
		WorkerApplyType:       "serial",
		MaxParallelWorkers:    2,
		DefaultInstanceAction: "stop-start",
		AlarmRestrictions:     "strict",
		ToVersion:             "v1.30.6",
		State:                 "ready-to-apply",
	}
	// KubeUpgradeBuilding is the expected strategy after apply.
	KubeUpgradeBuilding = kubeupgrade.KubeUpgrade{
		ID:                    "7aa16d94-dfc5-4029-bfcb-d815e7c2dc99",
		StrategyName:          "kube-upgrade",
		StorageApplyType:      "ignore",
		WorkerApplyType:       "serial",
		MaxParallelWorkers:    2,
		DefaultInstanceAction: "stop-start",
		AlarmRestrictions:     "strict",
		ToVersion:             "v1.30.6",
		State:                 "building",
	}
)

const KubeUpgradeCreateResponse = `
{
	"strategy":
	{
		"storage-apply-type": "ignore",
		"worker-apply-type": "serial",
		"state": "ready-to-apply",
		"default-instance-action": "stop-start",
		"current-phase": "build",
		"current-phase-completion-percentage": 0,
		"max-parallel-worker-hosts": 2,
		"alarm-restrictions": "strict",
		"to-version": "v1.30.6",
		"uuid": "7aa16d94-dfc5-4029-bfcb-d815e7c2dc99",
		"name": "kube-upgrade"
	}
}
`

const KubeUpgradeApplyResponse = `
{
	"strategy":
	{
		"storage-apply-type": "ignore",
		"worker-apply-type": "serial",
		"current-phase": "build",
		"current-phase-completion-percentage": 0,
		"state": "building",
		"default-instance-action": "stop-start",
		"max-parallel-worker-hosts": 2,
		"alarm-restrictions": "strict",
		"to-version": "v1.30.6",
		"uuid": "7aa16d94-dfc5-4029-bfcb-d815e7c2dc99",
		"name": "kube-upgrade"
	}
}
`

func HandleStrategyCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/api/orchestration/kube-upgrade/strategy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, `{
			"alarm-restrictions": "strict",
			"default-instance-action": "stop-start",
			"max-parallel-worker-hosts": 2,
			"storage-apply-type": "ignore",
			"to-version": "v1.30.6",
			"worker-apply-type": "serial"
		  }`)

			w.WriteHeader(http.StatusAccepted)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, response)
		})
}

func HandleStrategyApplySuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/api/orchestration/kube-upgrade/strategy/actions",
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
	th.Mux.HandleFunc("/api/orchestration/kube-upgrade/strategy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, response)
		})
}

func HandleStrategyDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/api/orchestration/kube-upgrade/strategy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.WriteHeader(http.StatusNoContent)
		})
}
