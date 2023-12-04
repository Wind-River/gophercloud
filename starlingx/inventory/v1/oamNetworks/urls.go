/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package oamNetworks

import "github.com/gophercloud/gophercloud"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("iextoam", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("iextoam")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
