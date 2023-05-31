/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package swpatch

import (
	"errors"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	inventoryv1 "github.com/gophercloud/gophercloud/starlingx/inventory/v1"
)

/* POST /api/orchestration/sw-patch/strategy/apply  supports an optional "stage-id" request parameter  */
type StrategyApplyOpts struct {
	StageID *string `json:"stage-id,omitempty" mapstructure:"stage-id"`
}

/* POST /api/orchestration/sw-patch/strategy/abort  supports an optional "stage-id" request parameter  */
type StrategyAbortOpts struct {
	StageID *string `json:"stage-id,omitempty" mapstructure:"stage-id"`
}

type SwPatchOpts struct {
	ControllerApplyType   string `json:"controller-apply-type"`
	StorageApplyType      string `json:"storage-apply-type"`
	SwiftApplyType      string `json:"swift-apply-type"`
	WorkerApplyType       string `json:"worker-apply-type"`
	MaxParallerWorkers    int    `json:"max-parallel-worker-hosts,omitempty"`
	DefaultInstanceAction string `json:"default-instance-action"`
	AlarmRestrictions     string `json:"alarm-restrictions,omitempty"`
}

// Get retrieves the created SwPatch.
func Show(c *gophercloud.ServiceClient) (swp *SwPatch, error) {
	var respBody map[string]interface{}
	_, err := c.Get(showURL(c), nil &respBody, nil)
	if err != nil {
		return nil, err
	}

	swPatch, err := GenerateSwPatch(respBody)
	if err != nil {
		return nil, err
	}

	return swPatch, nil
}

// Create accepts a SwPatchOpts struct and creates a new SwPatch using the
// values provided.
func Create(c *gophercloud.ServiceClient, opts SwPatchOpts) (swp *SwPatch, error) {
	reqBody, err := inventoryv1.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}
	var respBody map[string]interface{}
	_, err = c.Post(createURL(c), reqBody, &respBody,  &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	swPatch, err := GenerateSwPatch(respBody)
	if err != nil {
		return nil, err
	}

	return swPatch, nil
}

// Deletes the only related resource.
func Delete(c *gophercloud.ServiceClient ) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c), nil)
	return r
}

func ApplyStrategy(c *gophercloud.ServiceClient, opts StrategyApplyOpts ) (swp *SwPatch, error) {
	reqBody, err := inventoryv1.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}
	var respBody map[string]interface{}
	_, err = c.Post(applyURL(c), reqBody, &respBody,  &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	swPatch, err := GenerateSwPatch(respBody)
	if err != nil {
		return nil, err
	}

	return swPatch, nil
}

func AbortStrategy(c *gophercloud.ServiceClient, opts StrategyAbortOpts ) (swp *SwPatch, error) {
	reqBody, err := inventoryv1.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}
	var respBody map[string]interface{}
	_, err = c.Post(abortURL(c), reqBody, &respBody,  &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	swPatch, err := GenerateSwPatch(respBody)
	if err != nil {
		return nil, err
	}

	return swPatch, nil
}