/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package kernel

import (
	"github.com/gophercloud/gophercloud"
)

// Extract interprets any commonResult as an Kernel.
func (r commonResult) Extract() (*Kernel, error) {
	var s Kernel
	err := r.ExtractInto(&s)
	return &s, err
}

type commonResult struct {
	gophercloud.Result
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation.
type UpdateResult struct {
	commonResult
}

type Kernel struct {
	// ID of the host
	HostID string `json:"ihost_uuid"`

	// The hostname
	Hostname string `json:"hostname"`

	// The provisioned kernel
	ProvisionedKernel string `json:"kernel_provisioned"`

	// The running kernel
	RunningKernel string `json:"kernel_running"`
}
