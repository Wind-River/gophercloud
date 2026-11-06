/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package currentstrategy

import (
	"github.com/gophercloud/gophercloud"
)

// Show returns the orchestration strategy currently known to the VIM, or a nil
// pointer if no strategy is active. The VIM reports {"strategy": null} when no
// strategy exists; in that case a nil pointer and nil error are returned so
// callers can treat "no active strategy" distinctly from a parse failure.
func Show(c *gophercloud.ServiceClient) (*CurrentStrategy, error) {
	var resp struct {
		Strategy *CurrentStrategy `json:"strategy"`
	}
	if _, err := c.Get(showURL(c), &resp, nil); err != nil {
		return nil, err
	}
	// resp.Strategy is nil when "strategy" is absent or explicitly null.
	return resp.Strategy, nil
}
