/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package deploy

import "github.com/gophercloud/gophercloud"

func precheckURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("v1", "deploy", "precheck")
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("v1", "deploy")
}
