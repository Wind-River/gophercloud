/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package kernel

import "github.com/gophercloud/gophercloud"

func resourceURL(c *gophercloud.ServiceClient, hostid string) string {
	return c.ServiceURL("ihosts", hostid, "kernel")
}

func getURL(c *gophercloud.ServiceClient, hostid string) string {
	return resourceURL(c, hostid)
}

func updateURL(c *gophercloud.ServiceClient, hostid string) string {
	return resourceURL(c, hostid)
}
