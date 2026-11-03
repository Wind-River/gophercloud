/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

const HealthyBody = "System Health:\n" +
	"All hosts are provisioned: [OK]\n" +
	"All hosts are unlocked/enabled: [OK]\n" +
	"No alarms: [OK]\n"

const UnhealthyBody = "System Health:\n" +
	"All hosts are provisioned: [OK]\n" +
	"All hosts are unlocked/enabled: [Fail]\n" +
	"Locked or disabled hosts: controller-1\n"

func HandleKubeUpgradeHealthSuccessfully(t *testing.T, relaxed bool, body string) {
	path := "/health/kube-upgrade"
	if relaxed {
		path += "/relaxed"
	}
	th.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "text/plain")
		fmt.Fprint(w, body)
	})
}
