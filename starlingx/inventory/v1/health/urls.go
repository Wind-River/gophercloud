/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package health

import "github.com/gophercloud/gophercloud"

func kubeUpgradeURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("health", "kube-upgrade")
}

func kubeUpgradeRelaxedURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("health", "kube-upgrade", "relaxed")
}
