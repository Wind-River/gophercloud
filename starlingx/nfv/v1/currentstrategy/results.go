/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package currentstrategy

import (
	"encoding/json"
)

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

// GenerateCurrentStrategy takes a JSON string and converts it into a
// CurrentStrategy structure. The VIM reports {"strategy": null} when no
// strategy exists; in that case a nil pointer and nil error are returned so
// callers can treat "no active strategy" distinctly from a parse failure.
func GenerateCurrentStrategy(jsonData string) (*CurrentStrategy, error) {
	var data map[string]*CurrentStrategy
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}

	// data["strategy"] is nil when the key is absent or explicitly null.
	return data["strategy"], nil
}
