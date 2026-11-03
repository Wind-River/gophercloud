/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

// PassBody is a sample successful (HTTP 200) precheck response.
const PassBody = `{"info":"Deploy precheck passed","error":""}`

// FailBody is a sample failed (HTTP 406) precheck response. 406 is the normal
// "precheck failed" status; the body carries the failure details.
const FailBody = `{"info":"","error":"Deploy precheck failed: alarms present"}`

// capturedRequest records what the handler saw, for assertions.
type capturedRequest struct {
	Method   string
	RawQuery string
	Body     map[string]interface{}
}

// HandlePrecheckJSON installs a handler for the componentized (JSON-body)
// precheck. It captures the decoded JSON body and responds with the given
// status and body.
func HandlePrecheckJSON(t *testing.T, status int, response string) *capturedRequest {
	captured := &capturedRequest{}
	th.Mux.HandleFunc("/v1/deploy/precheck", func(w http.ResponseWriter, r *http.Request) {
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		captured.Method = r.Method
		captured.RawQuery = r.URL.RawQuery
		_ = json.NewDecoder(r.Body).Decode(&captured.Body)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)
		fmt.Fprint(w, response)
	})
	return captured
}

// HandlePrecheckLegacy installs a handler for the legacy (query-param)
// precheck. It captures the raw query and responds with the given status/body.
func HandlePrecheckLegacy(t *testing.T, status int, response string) *capturedRequest {
	captured := &capturedRequest{}
	th.Mux.HandleFunc("/v1/deploy/precheck", func(w http.ResponseWriter, r *http.Request) {
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		captured.Method = r.Method
		captured.RawQuery = r.URL.RawQuery
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)
		fmt.Fprint(w, response)
	})
	return captured
}

// DeployListBody is a sample GET /v1/deploy response with an in-progress deploy.
const DeployListBody = `[{"from_release":"starlingx-26.10.0","to_release":"starlingx-26.10.2","state":"host-rollback-done"}]`

// DeployListEmptyBody is a sample GET /v1/deploy response with no deploy.
const DeployListEmptyBody = `[]`

// HandleDeployList installs a handler for GET /v1/deploy returning the given body.
func HandleDeployList(t *testing.T, response string) {
	th.Mux.HandleFunc("/v1/deploy", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}
