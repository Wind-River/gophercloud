/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package kernel

import (
	"github.com/gophercloud/gophercloud"
	common "github.com/gophercloud/gophercloud/starlingx"
)

type KernelOpts struct {
	Kernel *string `json:"kernel_provisioned" mapstructure:"kernel_provisioned"`
}

// Get retrieves a specific Kernel based on the host's unique ID.
func Get(c *gophercloud.ServiceClient, hostid string) (r GetResult) {
	// Send GET request to API
	_, r.Err = c.Get(getURL(c, hostid), &r.Body, nil)
	return r
}

// Update accepts an array of KernelOpts and updates the specified host kernel
func Update(c *gophercloud.ServiceClient, hostid string, opts KernelOpts) (r UpdateResult) {
	reqBody, err := common.ConvertToPatchMap(opts, common.ReplaceOp)
	if err != nil {
		r.Err = err
	} else {
		// Send PATCH request to API
		_, r.Err = c.Patch(updateURL(c, hostid), reqBody, &r.Body, nil)
	}
	return r
}
