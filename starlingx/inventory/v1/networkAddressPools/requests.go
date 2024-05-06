/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2024 Wind River Systems, Inc. */

package networkAddressPools

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	common "github.com/gophercloud/gophercloud/starlingx"
)

type NetworkAddressPoolOpts struct {
	NetworkUUID     *string `json:"network_uuid,omitempty" mapstructure:"network_uuid"`
	AddressPoolUUID *string `json:"address_pool_uuid,omitempty" mapstructure:"address_pool_uuid"`
	NetworkName     *string `json:"network_name,omitempty" mapstructure:"network_name"`
	AddressPoolName *string `json:"address_pool_name,omitempty" mapstructure:"address_pool_name"`
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToNetworkAddressPoolListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the network address pool attributes you want to see returned. SortKey allows
// you to sort by a particular network address pool attribute. SortDir sets the
// direction, and is either `asc' or `desc'. Marker and Limit are used for
// pagination.
type ListOpts struct {
	Marker  string `q:"marker"`
	Limit   int    `q:"limit"`
	SortKey string `q:"sort_key"`
	SortDir string `q:"sort_dir"`
}

// ToNetworkAddressPoolListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToNetworkAddressPoolListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// network address pools. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToNetworkAddressPoolListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return NetworkAddressPoolPage{pagination.SinglePageBase(r)}
	})
}

// Get retrieves a specific network address pool based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = c.Get(getURL(c, id), &res.Body, nil)
	return res
}

// Create accepts a CreateOpts struct and creates a new NetworkAddressPool using the
// values provided.
func Create(c *gophercloud.ServiceClient, opts NetworkAddressPoolOpts) (r CreateResult) {
	reqBody, err := common.ConvertToCreateMap(opts)
	if err != nil {
		r.Err = err
		return r
	}

	_, r.Err = c.Post(createURL(c), reqBody, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return r
}

// Delete accepts a unique ID and deletes the related resource.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), nil)
	return r
}

// ListNetworkAddressPools is a convenience function to list and extract the entire list
// of network address pools.
func ListNetworkAddressPools(c *gophercloud.ServiceClient) ([]NetworkAddressPool, error) {
	pages, err := List(c, nil).AllPages()
	if err != nil {
		return nil, err
	}

	empty, err := pages.IsEmpty()
	if empty || err != nil {
		return nil, err
	}

	objs, err := ExtractNetworkAddressPools(pages)
	if err != nil {
		return nil, err
	}

	return objs, err
}
