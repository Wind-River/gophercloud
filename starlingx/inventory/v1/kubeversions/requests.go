/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package kubeversions

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// List returns a Pager which allows iterating over a collection of Kubernetes
// versions. The kube_versions endpoint returns a single (non-paginated) page.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, listURL(c), func(r pagination.PageResult) pagination.Page {
		return KubeVersionPage{pagination.SinglePageBase(r)}
	})
}

// ListKubeVersions is a convenience function that lists and extracts the entire
// collection of Kubernetes versions in one call.
func ListKubeVersions(c *gophercloud.ServiceClient) ([]KubeVersion, error) {
	pages, err := List(c).AllPages()
	if err != nil {
		return nil, err
	}

	empty, err := pages.IsEmpty()
	if empty || err != nil {
		return nil, err
	}

	return ExtractKubeVersions(pages)
}
