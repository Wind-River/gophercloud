/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package swpatch

import (
	"encoding/json"
	"github.com/gophercloud/gophercloud"
	common "github.com/gophercloud/gophercloud/starlingx"
)

/* POST /api/orchestration/sw-patch/strategy/apply  supports an optional "stage-id" request parameter  */
type StrategyActionOpts struct {
	Action  *string `json:"action" mapstructure:"action"`
	StageID *string `json:"stage-id,omitempty" mapstructure:"stage-id"`
}

type SwPatchOpts struct {
	ControllerApplyType   string `json:"controller-apply-type"`
	StorageApplyType      string `json:"storage-apply-type"`
	WorkerApplyType       string `json:"worker-apply-type"`
	SwiftApplyType 		  string `json:"swift-apply-type"` // TODO: Delete when updating to system-config-update
	MaxParallerWorkers    int    `json:"max-parallel-worker-hosts,omitempty"`
	DefaultInstanceAction string `json:"default-instance-action"`
	AlarmRestrictions     string `json:"alarm-restrictions,omitempty"`
}

func performRequest(c *gophercloud.ServiceClient, url string, reqBody interface{}) (*SwPatch, error) {
	var respBody map[string]interface{}
	_, err := c.Post(url, reqBody, &respBody, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	respJSON, err := json.Marshal(respBody)
	if err != nil {
		return nil, err
	}

	swPatch, err := GenerateSwPatch(string(respJSON))
	if err != nil {
		return nil, err
	}

	return swPatch, nil
}

func Show(c *gophercloud.ServiceClient) (*SwPatch, error) {
	var respBody map[string]interface{}
	_, err := c.Get(showURL(c), &respBody, nil)
	if err != nil {
		return nil, err
	}

	respJSON, err := json.Marshal(respBody)
	if err != nil {
		return nil, err
	}

	swPatch, err := GenerateSwPatch(string(respJSON))
	if err != nil {
		return nil, err
	}

	return swPatch, nil
}

// Create accepts a SwPatchOpts struct and creates a new SwPatch using the
// values provided.
func Create(c *gophercloud.ServiceClient, opts SwPatchOpts) (*SwPatch, error) {
	reqBody, err := common.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}

	return performRequest(c, createURL(c), reqBody)
}

// Delete current strategy.
func Delete(c *gophercloud.ServiceClient) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c), nil)
	return r
}

// Apply or Abort current strategy.
func ActionStrategy(c *gophercloud.ServiceClient, opts StrategyActionOpts) (*SwPatch, error) {
	reqBody, err := common.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}

	return performRequest(c, actionURL(c), reqBody)
}
