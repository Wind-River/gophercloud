/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package kubeupgrade

import (
	"github.com/gophercloud/gophercloud"
)

// List returns the current kube upgrade records. The collection endpoint
// returns {"kube_upgrades": [ {...} ]}; an empty list means no upgrade is in
// progress. It is decoded directly (the body is small and non-paginated).
func List(c *gophercloud.ServiceClient) ([]KubeUpgrade, error) {
	var resp struct {
		KubeUpgrades []KubeUpgrade `json:"kube_upgrades"`
	}
	_, err := c.Get(rootURL(c), &resp, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	return resp.KubeUpgrades, nil
}

// Delete removes the current kube upgrade record. Error statuses (e.g. the
// record no longer exists) are returned to the caller. sysinv returns 200 on a
// successful delete, which is not in gophercloud's default DELETE codes
// (202, 204), so OkCodes is set explicitly.
func Delete(c *gophercloud.ServiceClient) (r DeleteResult) {
	_, r.Err = c.Delete(rootURL(c), &gophercloud.RequestOpts{
		OkCodes: []int{200, 202, 204},
	})
	return r
}

// DeleteResult carries the error (if any) from a Delete call.
type DeleteResult struct {
	gophercloud.ErrResult
}
