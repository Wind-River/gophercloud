/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package oamNetworks

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	common "github.com/gophercloud/gophercloud/starlingx"
)

type OAMNetworkOpts struct {
	OAMSubnet     *string `json:"oam_subnet,omitempty" mapstructure:"oam_subnet,omitempty"`
	OAMGatewayIP  *string `json:"oam_gateway_ip,omitempty" mapstructure:"oam_gateway_ip,omitempty"`
	OAMFloatingIP *string `json:"oam_floating_ip,omitempty" mapstructure:"oam_floating_ip,omitempty"`
	OAMC0IP       *string `json:"oam_c0_ip,omitempty" mapstructure:"oam_c0_ip,omitempty"`
	OAMC1IP       *string `json:"oam_c1_ip,omitempty" mapstructure:"oam_c1_ip,omitempty"`
	OAMStartIP    *string `json:"oam_start_ip,omitempty" mapstructure:"oam_start_ip,omitempty"`
	OAMEndIP      *string `json:"oam_end_ip,omitempty" mapstructure:"oam_end_ip,omitempty"`
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToOAMNetworkListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the noam network attributes you want to see returned. SortKey allows you to sort
// by a particular oam network attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Marker  string `q:"marker"`
	Limit   int    `q:"limit"`
	SortKey string `q:"sort_key"`
	SortDir string `q:"sort_dir"`
}

// ToOAMNetworkListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToOAMNetworkListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List retrieves a list of iextoams
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToOAMNetworkListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return OAMNetworkPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific oam network based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = c.Get(getURL(c, id), &res.Body, nil)
	return res
}

// Update accepts a PatchOpts struct and updates an existing oam network using the
// values provided. For more information, see the Create function.
func Update(c *gophercloud.ServiceClient, id string, opts OAMNetworkOpts) (r UpdateResult) {
	reqBody, err := common.ConvertToPatchMap(opts, common.ReplaceOp)
	if err != nil {
		r.Err = err
		return r
	}

	// Send request to API
	_, r.Err = c.Patch(updateURL(c, id), reqBody, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	return r
}

// ListNetworks is a convenience function to list and extract the entire list
// of platform networks
func ListNetworks(c *gophercloud.ServiceClient) ([]OAMNetwork, error) {
	pages, err := List(c, nil).AllPages()
	if err != nil {
		return nil, err
	}

	empty, err := pages.IsEmpty()
	if empty || err != nil {
		return nil, err
	}

	objs, err := ExtractOAMNetworks(pages)
	if err != nil {
		return nil, err
	}

	return objs, err
}
