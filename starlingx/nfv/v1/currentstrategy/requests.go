/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package currentstrategy

import (
	"encoding/json"

	"github.com/gophercloud/gophercloud"
)

// Show returns the orchestration strategy currently known to the VIM, or a nil
// pointer if no strategy is active.
func Show(c *gophercloud.ServiceClient) (*CurrentStrategy, error) {
	var respBody map[string]interface{}
	_, err := c.Get(showURL(c), &respBody, nil)
	if err != nil {
		return nil, err
	}

	respJSON, err := json.Marshal(respBody)
	if err != nil {
		return nil, err
	}

	return GenerateCurrentStrategy(string(respJSON))
}
