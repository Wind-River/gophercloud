/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package swupgrade

import "github.com/gophercloud/gophercloud"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("api", "orchestration", "sw-upgrade", "strategy")
}

func showURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func deleteURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func actionURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("api", "orchestration", "sw-upgrade", "strategy", "actions")
}
