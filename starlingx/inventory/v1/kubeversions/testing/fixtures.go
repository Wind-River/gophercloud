/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/kubeversions"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

// KubeVersionListBody is a sample GET /kube_versions response.
const KubeVersionListBody = `
{
	"kube_versions": [
		{"version": "v1.32.2", "state": "unavailable", "target": false},
		{"version": "v1.33.0", "state": "available", "target": true},
		{"version": "v1.34.1", "state": "active", "target": false}
	]
}
`

var (
	KubeVersionUnavailable = kubeversions.KubeVersion{
		Version: "v1.32.2",
		State:   "unavailable",
		Target:  false,
	}
	KubeVersionTarget = kubeversions.KubeVersion{
		Version: "v1.33.0",
		State:   "available",
		Target:  true,
	}
	KubeVersionActive = kubeversions.KubeVersion{
		Version: "v1.34.1",
		State:   "active",
		Target:  false,
	}
)

func HandleKubeVersionListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kube_versions", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, KubeVersionListBody)
	})
}
