/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package addresses

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Extract interprets any commonResult as an Image.
func (r commonResult) Extract() (*SystemConfigUpdateStrategy, error) {
	var s SystemConfigUpdateStrategy
	err := r.ExtractInto(&s)
	return &s, err
}

type commonResult struct {
	gophercloud.Result
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// CreateResult represents the result of an update operation.
type CreateResult struct {
	commonResult
}

// DeleteResult represents the result of an delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// SystemConfigUpdateStrategy defines the data associated to a single System Config Update Strategy instance.
type SystemConfigUpdateStrategy struct {
	// ID is the generated unique UUID for the System Config Update Strategy
	ID string `json:"uuid"`

	// StrategyName is the name of the strategy.
	StrategyName string `json:"name"`

	// ControllerApplyType is the apply type for controller hosts.
	ControllerApplyType string `json:"controller-apply-type"`

	// StorageApplyType is the apply type for storage hosts.
	StorageApplyType string `json:"storage-apply-type"`

	// WorkerApplyType is the apply type for worker hosts.
	WorkerApplyType string `json:"worker-apply-type"`

	// The maximum number of worker hosts to update in parallel; only applicable
	// if ``worker-apply-type = parallel``.
	MaxParallerWorkers int `json:"max-parallel-worker-hosts,omitempty"`

	// The default instance action.
	DefaultInstanceAction string `json:"default-instance-action"`

	// The strictness of alarm checks.
	AlarmRestrictions string `json:"default-instance-action"`

	// The strictness of alarm checks.
	State string `json:"default-instance-action"`
}

// AddressPage is the page returned by a pager when traversing over a
// collection of addresss.
type AddressPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether a AddressPage struct is empty.
func (r AddressPage) IsEmpty() (bool, error) {
	is, err := ExtractAddresses(r)
	return len(is) == 0, err
}

// ExtractAddresses accepts a Page struct, specifically a AddressPage struct,
// and extracts the elements into a slice of SystemConfigUpdateStrategy structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractAddresses(r pagination.Page) ([]SystemConfigUpdateStrategy, error) {
	var s struct {
		SystemConfigUpdateStrategy []SystemConfigUpdateStrategy `json:"addresses"`
	}

	err := (r.(AddressPage)).ExtractInto(&s)

	return s.SystemConfigUpdateStrategy, err
}
