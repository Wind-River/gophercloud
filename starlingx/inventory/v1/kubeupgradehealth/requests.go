/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package kubeupgradehealth

import (
	"io"

	"github.com/gophercloud/gophercloud"
)

// getText performs a GET and returns the raw plain-text response body. The
// health endpoints always return HTTP 200 with a text report; JSONResponse is
// left nil so gophercloud does not attempt to decode, and the body is read
// directly.
func getText(c *gophercloud.ServiceClient, url string) (string, error) {
	resp, err := c.Get(url, nil, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// KubeUpgrade returns the plain-text kube-upgrade health report
// (GET /v1/health/kube-upgrade). When relaxed is true it queries the relaxed
// variant (GET /v1/health/kube-upgrade/relaxed), which bypasses non-critical
// checks.
func KubeUpgrade(c *gophercloud.ServiceClient, relaxed bool) (string, error) {
	url := kubeUpgradeURL(c)
	if relaxed {
		url = kubeUpgradeRelaxedURL(c)
	}
	return getText(c, url)
}
