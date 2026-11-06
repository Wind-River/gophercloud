/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package kubeversions

import (
	"github.com/gophercloud/gophercloud/pagination"
)

// KubeVersion defines the data associated with a single Kubernetes version
// entry as reported by the sysinv kube_versions API.
type KubeVersion struct {
	// Version is the Kubernetes version string, e.g. "v1.34.1".
	Version string `json:"version"`

	// State is the version's state on the system, e.g. "active", "available",
	// "unavailable", "partial".
	State string `json:"state"`

	// Target indicates whether this version is the current upgrade target.
	Target bool `json:"target"`

	// UpgradeFrom lists the versions from which this version can be upgraded.
	UpgradeFrom []string `json:"upgrade_from,omitempty"`

	// DowngradeTo lists the versions to which this version can be downgraded.
	DowngradeTo []string `json:"downgrade_to,omitempty"`

	// AppliedPatches lists patches that must be applied for this version.
	AppliedPatches []string `json:"applied_patches,omitempty"`

	// AvailablePatches lists patches available for this version.
	AvailablePatches []string `json:"available_patches,omitempty"`
}

// KubeVersionPage is the page returned by a pager when traversing over a
// collection of Kubernetes versions.
type KubeVersionPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether a KubeVersionPage is empty.
func (r KubeVersionPage) IsEmpty() (bool, error) {
	is, err := ExtractKubeVersions(r)
	return len(is) == 0, err
}

// ExtractKubeVersions accepts a Page struct, specifically a KubeVersionPage,
// and extracts the elements into a slice of KubeVersion structs.
func ExtractKubeVersions(r pagination.Page) ([]KubeVersion, error) {
	var s struct {
		KubeVersions []KubeVersion `json:"kube_versions"`
	}

	err := (r.(KubeVersionPage)).ExtractInto(&s)

	return s.KubeVersions, err
}
