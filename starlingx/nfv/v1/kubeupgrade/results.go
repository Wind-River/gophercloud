/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package kubeupgrade

import (
	"github.com/gophercloud/gophercloud"
)

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Step is a single step within a strategy stage.
type Step struct {
	StartDateTime string   `json:"start-date-time,omitempty"`
	EndDateTime   string   `json:"end-date-time,omitempty"`
	Timeout       int      `json:"timeout,omitempty"`
	EntityType    string   `json:"entity-type,omitempty"`
	StepId        int      `json:"step-id,omitempty"`
	EntityUuids   []string `json:"entity-uuids,omitempty"`
	StepName      string   `json:"step-name,omitempty"`
	Result        string   `json:"result,omitempty"`
	EntityNames   []string `json:"entity-names,omitempty"`
	Reason        string   `json:"reason,omitempty"`
}

// Stage is a single stage within a strategy phase.
type Stage struct {
	StartDateTime string `json:"start-date-time,omitempty"`
	EndDateTime   string `json:"end-date-time,omitempty"`
	StageId       int    `json:"stage-id,omitempty"`
	Reason        string `json:"reason,omitempty"`
	CurrentStep   int    `json:"current-step,omitempty"`
	Steps         []Step `json:"steps,omitempty"`
	Result        string `json:"result,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	TotalSteps    int    `json:"total-steps,omitempty"`
	Inprogress    bool   `json:"inprogress,omitempty"`
	StageName     string `json:"stage-name,omitempty"`
}

// Phase is a build, apply, or abort phase of a strategy.
type Phase struct {
	StartDateTime        string  `json:"start-date-time,omitempty"`
	EndDateTime          string  `json:"end-date-time,omitempty"`
	PhaseName            string  `json:"phase-name,omitempty"`
	CompletionPercentage int     `json:"completion-percentage,omitempty"`
	TotalStages          int     `json:"total-stages,omitempty"`
	StopAtStage          int     `json:"stop-at-stage,omitempty"`
	Result               string  `json:"result,omitempty"`
	Timeout              int     `json:"timeout,omitempty"`
	Reason               string  `json:"reason,omitempty"`
	Inprogress           bool    `json:"inprogress,omitempty"`
	Stages               []Stage `json:"stages,omitempty"`
	CurrentStage         int     `json:"current-stage,omitempty"`
}

// KubeUpgrade defines the data associated with a single kube-upgrade strategy.
type KubeUpgrade struct {
	// ID is the generated unique UUID for the kube-upgrade strategy.
	ID string `json:"uuid"`

	// StrategyName is the name of the strategy.
	StrategyName string `json:"name"`

	// Release is the target identifier the VIM reports for the strategy. For
	// kube-upgrade this may carry the Kubernetes version. Kept generic so both
	// string and list forms round-trip. Populated in strategy responses.
	Release interface{} `json:"release,omitempty"`

	// ReleaseID is an alternative single identifier observed in some sw-manager
	// strategy show output.
	ReleaseID string `json:"release-id,omitempty"`

	// StorageApplyType is the apply type for storage hosts.
	StorageApplyType string `json:"storage-apply-type"`

	// WorkerApplyType is the apply type for worker hosts.
	WorkerApplyType string `json:"worker-apply-type"`

	// MaxParallelWorkers is the maximum number of worker hosts to update in
	// parallel; only applicable if worker-apply-type = parallel.
	MaxParallelWorkers int `json:"max-parallel-worker-hosts,omitempty"`

	// DefaultInstanceAction is the default instance action.
	DefaultInstanceAction string `json:"default-instance-action"`

	// AlarmRestrictions is the strictness of alarm checks.
	AlarmRestrictions string `json:"alarm-restrictions,omitempty"`

	// ToVersion is the target Kubernetes version for the upgrade.
	ToVersion string `json:"to-version,omitempty"`

	// State is the current strategy state (e.g. building, ready-to-apply,
	// applying, applied, build-failed, apply-failed, aborted).
	State string `json:"state"`

	// BuildPhase describes the build phase of the strategy.
	BuildPhase Phase `json:"build-phase,omitempty"`

	// ApplyPhase describes the apply phase of the strategy.
	ApplyPhase Phase `json:"apply-phase,omitempty"`

	// AbortPhase describes the abort phase of the strategy.
	AbortPhase Phase `json:"abort-phase,omitempty"`
}
