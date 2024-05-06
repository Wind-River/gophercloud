/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2024 Wind River Systems, Inc. */

package networkAddressPools

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Extract interprets any commonResult as an Image.
func (r commonResult) Extract() (*NetworkAddressPool, error) {
	var s NetworkAddressPool
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

// NetworkAddressPool defines the data associated to a single network address pool instance.
type NetworkAddressPool struct {
	// UUID is a system generated unique UUID for the network address pool
	UUID string `json:"uuid"`

	// NetworkUUID is the UUID of the network resource
	NetworkUUID string `json:"network_uuid"`

	// AddressPoolUUID is the UUID of the address pool resource
	AddressPoolUUID string `json:"address_pool_uuid"`

	// NetworkName is the name of network resource
	NetworkName string `json:"network_name"`

	// AddressPoolName is the name of the address pool resource
	AddressPoolName string `json:"address_pool_name"`
}

// NetworkAddressPool is the page returned by a pager when traversing over a
// collection of network address pools.
type NetworkAddressPoolPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether a NetworkAddressPool struct is empty.
func (r NetworkAddressPoolPage) IsEmpty() (bool, error) {
	is, err := ExtractNetworkAddressPools(r)
	return len(is) == 0, err
}

// ExtractNetworkAddressPools accepts a Page struct, specifically a NetworkAddressPool struct,
// and extracts the elements into a slice of NetworkAddressPool structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractNetworkAddressPools(r pagination.Page) ([]NetworkAddressPool, error) {
	var s struct {
		NetworkAddressPools []NetworkAddressPool `json:"network_addresspools"`
	}

	err := (r.(NetworkAddressPoolPage)).ExtractInto(&s)

	return s.NetworkAddressPools, err
}
