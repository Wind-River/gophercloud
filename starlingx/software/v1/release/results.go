/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package release

// Release represents a single StarlingX Software (USM) release as returned by
// the release API. The list endpoints return an array of these; the detail
// endpoint returns a single one (and populates Requires).
type Release struct {
	// ReleaseID is the unique release identifier, e.g. "starlingx-26.10.0" or
	// a metapackage id such as "infra_26.10.3".
	ReleaseID string `json:"release_id"`

	// State is the USM release state, e.g. "available", "deployed",
	// "deploy-selected", "unavailable".
	State string `json:"state"`

	// SwVersion is the product release version the release belongs to,
	// e.g. "26.10.3".
	SwVersion string `json:"sw_version"`

	// Component is the component name (used by metapackage releases),
	// e.g. "infra".
	Component string `json:"component"`

	// RebootRequired indicates whether deploying the release requires a reboot.
	RebootRequired bool `json:"reboot_required"`

	// Packages is the list of package names included in the release. Used by
	// callers to derive supported Kubernetes versions from meta-k8s-* entries.
	Packages []string `json:"packages"`

	// PrepatchedIso indicates the release was delivered as a pre-patched ISO.
	PrepatchedIso bool `json:"prepatched_iso"`

	// PreUpgradeDeploy carries pre-upgrade-deploy metadata when present. It is
	// left as a generic map because callers only test for its presence.
	PreUpgradeDeploy map[string]interface{} `json:"pre_upgrade_deploy"`

	// Requires lists the release IDs this release depends on. Populated by the
	// detail endpoint (GET v1/release/{release_id}); typically empty in list
	// responses.
	Requires []string `json:"requires"`
}
