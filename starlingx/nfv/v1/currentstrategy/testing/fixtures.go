/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/nfv/v1/currentstrategy"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

// CurrentStrategyActive is the expected result when a strategy is in progress.
var CurrentStrategyActive = currentstrategy.CurrentStrategy{
	Name:  "sw-upgrade",
	State: "applying",
}

// CurrentStrategyActiveResponse reports an active sw-upgrade strategy.
const CurrentStrategyActiveResponse = `
{
	"strategy":
	{
		"name": "sw-upgrade",
		"state": "applying"
	}
}
`

// CurrentStrategyNoneResponse is what the VIM returns when no strategy exists.
const CurrentStrategyNoneResponse = `
{
	"strategy": null
}
`

func HandleShowActiveSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/api/orchestration/current-strategy/strategy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, response)
		})
}
