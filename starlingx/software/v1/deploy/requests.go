/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package deploy

import (
	"io"
	"net/url"
	"strings"

	"github.com/gophercloud/gophercloud"
)

// precheckOkCodes are the HTTP statuses the precheck endpoint uses to return a
// result BODY rather than a transport error. 200 = precheck passed, 406 =
// precheck failed (the normal "not acceptable" response carrying the failure
// details). Both are treated as acceptable so the caller always receives the
// body to parse; only other statuses surface as errors.
var precheckOkCodes = []int{200, 406}

// PrecheckOpts is the JSON body for the componentized precheck request
// (product version > 26.03): POST v1/deploy/precheck with
// {"releases": [...], "force": "true"?, "pre_upgrade_deploy": true?}.
//
// Force is a STRING ("true") rather than a bool to match the USM API, and is
// omitted when empty. PreUpgradeDeploy requests pre-upgrade-deploy metapackage
// validation only and is omitted when false.
type PrecheckOpts struct {
	Releases         []string `json:"releases"`
	Force            string   `json:"force,omitempty"`
	PreUpgradeDeploy bool     `json:"pre_upgrade_deploy,omitempty"`
}

// readBody performs a request and returns the raw response body. precheck
// always returns a text/JSON body (on 200 and 406); JSONResponse is left nil so
// gophercloud does not attempt to decode, and the body is read directly.
func readBody(resp *gophercloud.ServiceClient, method, reqURL string, opts *gophercloud.RequestOpts) (string, error) {
	r, err := resp.Request(method, reqURL, opts)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Precheck runs the deploy precheck using the componentized JSON-body format
// (product version > 26.03). It returns the raw response body (JSON with an
// "error" field on failure, or plain text), which the caller parses. A 406
// response is NOT an error — it is the normal precheck-failed response and its
// body is returned.
func Precheck(c *gophercloud.ServiceClient, opts PrecheckOpts) (string, error) {
	return readBody(c, "POST", precheckURL(c), &gophercloud.RequestOpts{
		JSONBody: opts,
		OkCodes:  precheckOkCodes,
	})
}

// PrecheckLegacy runs the deploy precheck using the pre-componentization query
// parameter format (product version <= 26.03):
// POST v1/deploy/precheck?release=<release_id>[&force=true] with no body.
// release may be empty (no release parameter is added). It returns the raw
// response body like Precheck; a 406 is treated as the normal failure response.
func PrecheckLegacy(c *gophercloud.ServiceClient, release string, force bool) (string, error) {
	reqURL := precheckURL(c)
	q := url.Values{}
	if release != "" {
		q.Set("release", release)
	}
	if force {
		q.Set("force", "true")
	}
	if encoded := q.Encode(); encoded != "" {
		reqURL += "?" + encoded
	}

	return readBody(c, "POST", reqURL, &gophercloud.RequestOpts{
		// Legacy precheck takes no JSON body; send an empty body so the request
		// is well-formed for servers that expect one.
		RawBody: strings.NewReader(""),
		OkCodes: precheckOkCodes,
	})
}

// Deploy is a single entry from the USM deploy list (GET v1/deploy). Only the
// fields the operator reads are modeled; the endpoint returns a JSON array of
// these (or null/empty when there is no active deploy).
type Deploy struct {
	FromRelease string `json:"from_release"`
	ToRelease   string `json:"to_release"`
	State       string `json:"state"`
}

// List returns the current USM deploy entries (GET v1/deploy). The response is
// a flat JSON array (null/empty when no deploy is in progress), so it is
// decoded directly into []Deploy.
func List(c *gophercloud.ServiceClient) ([]Deploy, error) {
	var result []Deploy
	_, err := c.Get(listURL(c), &result, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
