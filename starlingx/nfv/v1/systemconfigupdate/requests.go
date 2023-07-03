/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package systemconfigupdate

import (
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	common "github.com/gophercloud/gophercloud/starlingx"
)

/* POST /api/orchestration/system-config-update/strategy/apply  supports an optional "stage-id" request parameter  */
type StrategyActionOpts struct {
	Action  *string `json:"action" mapstructure:"action"`
	StageID *string `json:"stage-id,omitempty" mapstructure:"stage-id"`
}

type SystemConfigUpdateOpts struct {
	ControllerApplyType   string `json:"controller-apply-type" mapstructure:"controller-apply-type"`
	StorageApplyType      string `json:"storage-apply-type" mapstructure:"storage-apply-type"`
	WorkerApplyType       string `json:"worker-apply-type" mapstructure:"worker-apply-type"`
	MaxParallerWorkers    int    `json:"max-parallel-worker-hosts,omitempty" mapstructure:"max-parallel-worker-hosts"`
	DefaultInstanceAction string `json:"default-instance-action" mapstructure:"default-instance-action"`
	AlarmRestrictions     string `json:"alarm-restrictions,omitempty" mapstructure:"alarm-restrictions,omitempty"`
}

func performRequest(c *gophercloud.ServiceClient, url string, reqBody interface{}) (*SystemConfigUpdate, error) {
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

	scUpdate, err := GenerateSCUpdate(string(respJSON))
	if err != nil {
		return nil, err
	}

	return scUpdate, nil
}

func Show(c *gophercloud.ServiceClient) (*SystemConfigUpdate, error) {
	var respBody map[string]interface{}
	_, err := c.Get(showURL(c), &respBody, nil)
	if err != nil {
		return nil, err
	}

	respJSON, err := json.Marshal(respBody)
	if err != nil {
		return nil, err
	}

	scUpdate, err := GenerateSCUpdate(string(respJSON))
	if err != nil {
		return nil, err
	}

	return scUpdate, nil
}

// Create accepts a SystemConfigUpdateOpts struct and creates a new SystemConfigUpdate using the
// values provided.
func Create(c *gophercloud.ServiceClient, opts SystemConfigUpdateOpts) (*SystemConfigUpdate, error) {
	reqBody, err := common.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}

	return performRequest(c, createURL(c), reqBody)
}

// Delete current strategy.
func Delete(c *gophercloud.ServiceClient) (r DeleteResult) {
	// systemconfigupdate DELETE needs empty body in json
	m := make(map[string]interface{})
	_, r.Err = c.Delete(deleteURL(c), &gophercloud.RequestOpts{
		JSONBody: m,
	})
	return r
}

// Apply or Abort current strategy.
func ActionStrategy(c *gophercloud.ServiceClient, opts StrategyActionOpts) (*SystemConfigUpdate, error) {
	reqBody, err := common.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}

	return performRequest(c, actionURL(c), reqBody)
}
