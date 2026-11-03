/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package kubeupgrade

// KubeUpgrade describes a single in-progress Kubernetes upgrade record as
// reported by the sysinv kube_upgrade API. Only the fields the platform
// operator reads are modeled.
type KubeUpgrade struct {
	// State is the kube upgrade state, e.g. "upgrading-first-master",
	// "post-updated-apps", "upgrade-complete", "upgrade-aborted".
	State string `json:"state"`

	// FromVersion is the Kubernetes version the upgrade started from.
	FromVersion string `json:"from_version,omitempty"`

	// ToVersion is the Kubernetes version the upgrade targets.
	ToVersion string `json:"to_version,omitempty"`
}
