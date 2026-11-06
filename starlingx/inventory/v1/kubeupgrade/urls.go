/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package kubeupgrade

import "github.com/gophercloud/gophercloud"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("kube_upgrade")
}
