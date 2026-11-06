/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package release

import (
	"github.com/gophercloud/gophercloud"
)

// List returns all releases known to the Software (USM) service
// (GET v1/release). The response is a JSON array of releases.
func List(c *gophercloud.ServiceClient) ([]Release, error) {
	var releases []Release
	_, err := c.Get(listURL(c), &releases, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	return releases, nil
}

// ListMetapackages returns all metapackage releases
// (GET v1/release/metapackage?all).
//
// The "?all" query parameter is required: without it USM returns only the
// DEPLOYED metapackages, which would make valid AVAILABLE releases appear
// unknown to callers performing override validation.
func ListMetapackages(c *gophercloud.ServiceClient) ([]Release, error) {
	var releases []Release
	_, err := c.Get(metapackageListAllURL(c), &releases, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	return releases, nil
}

// Get returns the detail for a single release by its release ID
// (GET v1/release/{release_id}). The detail response populates the Requires
// field with the release's dependencies.
func Get(c *gophercloud.ServiceClient, releaseID string) (*Release, error) {
	var r Release
	_, err := c.Get(getURL(c, releaseID), &r, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	return &r, nil
}
