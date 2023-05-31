/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package swpatch

import (
	"encoding/json"
	"github.com/gophercloud/gophercloud"
)

// DeleteResult represents the result of an delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// SystemConfigUpdateStrategy defines the data associated to a single System Config Update Strategy instance.
type SwPatch struct {
	// ID is the generated unique UUID for the System Config Update Strategy
	ID string `json:"uuid"`

	// StrategyName is the name of the strategy.
	StrategyName string `json:"name"`

	// ControllerApplyType is the apply type for controller hosts.
	ControllerApplyType string `json:"controller-apply-type"`

	// StorageApplyType is the apply type for storage hosts.
	StorageApplyType string `json:"storage-apply-type"`

	// SwiftApplyType is the apply type for storage hosts.
	// TODO: Delete when updating to system-config-update
	SwiftApplyType string `json:"swift-apply-type"`

	// WorkerApplyType is the apply type for worker hosts.
	WorkerApplyType string `json:"worker-apply-type"`

	// The maximum number of worker hosts to update in parallel; only applicable
	// if ``worker-apply-type = parallel``.
	MaxParallerWorkers int `json:"max-parallel-worker-hosts,omitempty"`

	// The default instance action.
	DefaultInstanceAction string `json:"default-instance-action"`

	// The strictness of alarm checks.
	AlarmRestrictions string `json:"alarm-restrictions,omitempty"`

	// The strictness of alarm checks.
	State string `json:"state"`
}

// GenerateSwPatch takes a JSON string and converts it into a SwPatch structure.
// It returns a pointer to SwPatch and an error if the JSON parsing fails.
func GenerateSwPatch(jsonData string) (*SwPatch, error) {
	var data map[string]SwPatch
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}

	swPatch := data["strategy"]
	return &swPatch, nil
}
