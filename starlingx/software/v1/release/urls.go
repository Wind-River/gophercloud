/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package release

import "github.com/gophercloud/gophercloud"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("v1", "release")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, releaseID string) string {
	return c.ServiceURL("v1", "release", releaseID)
}

func metapackageURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("v1", "release", "metapackage")
}

// metapackageListAllURL is the metapackage list URL with the bare "all" query
// flag. Without it USM returns only DEPLOYED metapackages.
func metapackageListAllURL(c *gophercloud.ServiceClient) string {
	return metapackageURL(c) + "?all"
}
