/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package currentstrategy

import "github.com/gophercloud/gophercloud"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("api", "orchestration", "current-strategy", "strategy")
}

func showURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}
