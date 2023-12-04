/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package oamNetworks

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Extract interprets any commonResult as an Image.
func (r commonResult) Extract() (*OAMNetwork, error) {
	var s OAMNetwork
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

// UpdateResult represents the result of an update operation.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of an delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Network defines the data associated to a single network instance.
type OAMNetwork struct {
	// Unique UUID for this extoam
	UUID string `json:"uuid"`

	// Represent the oam subnet.
	OAMSubnet string `json:"oam_subnet"`

	// Represent the oam gateway IP.
	OAMGatewayIP string `json:"oam_gateway_ip"`

	// Represent the oam floating IP.
	OAMFloatingIP string `json:"oam_floating_ip"`

	// Represent the oam controller-0 IP address.
	OAMC0IP string `json:"oam_c0_ip"`

	// Represent the oam controller-1 IP address.
	OAMC1IP string `json:"oam_c1_ip"`

	// Represent the oam network start IP address.
	OAMStartIP string `json:"oam_start_ip"`

	// Represent the oam network end IP address.
	OAMEndIP string `json:"oam_end_ip"`

	// Represents whether in region_config. True=region_config
	RegionConfig bool `json:"region_config"`

	// Represent the action on the OAM network.
	Action string `json:"action"`

	// The isystemid that this iextoam belongs to.
	ForISystemID int `json:"forisystemid"`

	// The UUID of the system this extoam belongs to.
	ForISystemUUID string `json:"isystem_uuid"`

	// A list containing a self link and associated extoam links.
	Links []gophercloud.Link `json:"links"`

	// CreatedAt defines the timestamp at which the resource was created.
	CreatedAt string `json:"created_at"`

	// UpdatedAt defines the timestamp at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// OAMNetworkPage is the page returned by a pager when traversing over a
// collection of networks.
type OAMNetworkPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a OAMNetworkPage struct is empty.
func (r OAMNetworkPage) IsEmpty() (bool, error) {
	is, err := ExtractOAMNetworks(r)
	return len(is) == 0, err
}

// NextPageURL is invoked when a paginated collection of oam networks has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r OAMNetworkPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"iextoams_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// ExtractOAMNetworks accepts a Page struct, specifically a OAMNetworkPage struct,
// and extracts the elements into a slice of OAM Network structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractOAMNetworks(r pagination.Page) ([]OAMNetwork, error) {
	var s struct {
		OAMNetwork []OAMNetwork `json:"iextoams"`
	}
	err := (r.(OAMNetworkPage)).ExtractInto(&s)
	return s.OAMNetwork, err
}

func ExtractOAMNetworksInto(r pagination.Page, v interface{}) error {
	return r.(OAMNetworkPage).Result.ExtractIntoSlicePtr(v, "iextoams")
}
