/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/kubeupgrade"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

// KubeUpgradeListBody is a sample GET /kube_upgrade response with an in-progress
// upgrade.
const KubeUpgradeListBody = `
{
	"kube_upgrades": [
		{"state": "upgrading-first-master", "from_version": "v1.33.0", "to_version": "v1.34.1"}
	]
}
`

// KubeUpgradeEmptyBody is a sample GET /kube_upgrade response with no upgrade in
// progress.
const KubeUpgradeEmptyBody = `{"kube_upgrades": []}`

var KubeUpgradeInProgress = kubeupgrade.KubeUpgrade{
	State:       "upgrading-first-master",
	FromVersion: "v1.33.0",
	ToVersion:   "v1.34.1",
}

func HandleKubeUpgradeListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kube_upgrade", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, KubeUpgradeListBody)
	})
}

func HandleKubeUpgradeListEmpty(t *testing.T) {
	th.Mux.HandleFunc("/kube_upgrade", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, KubeUpgradeEmptyBody)
	})
}

func HandleKubeUpgradeDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kube_upgrade", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}
