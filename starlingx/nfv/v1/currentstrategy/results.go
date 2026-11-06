/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package currentstrategy

// CurrentStrategy describes the orchestration strategy currently known to the
// VIM sw-manager, independent of its type.
type CurrentStrategy struct {
	// Name is the strategy name (e.g. "sw-upgrade", "kube-upgrade",
	// "sw-patch", "system-config-update", "fw-update").
	Name string `json:"name"`

	// State is the current strategy state (e.g. building, ready-to-apply,
	// applying, applied, build-failed, apply-failed, aborted).
	State string `json:"state"`
}
