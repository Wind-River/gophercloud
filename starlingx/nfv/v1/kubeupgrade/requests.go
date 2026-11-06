/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package kubeupgrade

import (
	"github.com/gophercloud/gophercloud"
	common "github.com/gophercloud/gophercloud/starlingx"
)

// StrategyActionOpts is the body for POST .../strategy/actions.
// Action is typically "apply-all" or "abort". StageID is optional and only
// meaningful for staged apply/abort.
type StrategyActionOpts struct {
	Action  *string `json:"action" mapstructure:"action"`
	StageID *string `json:"stage-id,omitempty" mapstructure:"stage-id"`
}

// KubeUpgradeOpts is the body for creating a kube-upgrade strategy.
//
// Unlike sw-upgrade, the kube-upgrade strategy has no controller-apply-type
// and no release field; the target Kubernetes version is given by ToVersion.
//
// MaxParallelWorkerHosts is a pointer so it is omitted when nil: VIM returns
// 500 if the field is present on a system with no worker hosts (e.g. AIO-SX).
type KubeUpgradeOpts struct {
	StorageApplyType       string `json:"storage-apply-type" mapstructure:"storage-apply-type"`
	WorkerApplyType        string `json:"worker-apply-type" mapstructure:"worker-apply-type"`
	DefaultInstanceAction  string `json:"default-instance-action" mapstructure:"default-instance-action"`
	AlarmRestrictions      string `json:"alarm-restrictions,omitempty" mapstructure:"alarm-restrictions"`
	MaxParallelWorkerHosts *int   `json:"max-parallel-worker-hosts,omitempty" mapstructure:"max-parallel-worker-hosts"`
	ToVersion              string `json:"to-version" mapstructure:"to-version"`
}

func performRequest(c *gophercloud.ServiceClient, url string, reqBody interface{}) (*KubeUpgrade, error) {
	var resp struct {
		Strategy *KubeUpgrade `json:"strategy"`
	}
	_, err := c.Post(url, reqBody, &resp, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}
	// resp.Strategy is nil when "strategy" is absent or explicitly null, so
	// callers can treat "no strategy" distinctly from a populated one.
	return resp.Strategy, nil
}

// Show returns the current kube-upgrade strategy, or a nil pointer if no
// strategy is active. The VIM reports {"strategy": null} (or omits the key)
// when none exists; a nil pointer and nil error are returned in that case so
// callers can distinguish "no active strategy" from a parse failure.
func Show(c *gophercloud.ServiceClient) (*KubeUpgrade, error) {
	var resp struct {
		Strategy *KubeUpgrade `json:"strategy"`
	}
	if _, err := c.Get(showURL(c), &resp, nil); err != nil {
		return nil, err
	}
	// resp.Strategy is nil when "strategy" is absent or explicitly null.
	return resp.Strategy, nil
}

// Create accepts a KubeUpgradeOpts struct and creates a new kube-upgrade
// strategy using the values provided.
func Create(c *gophercloud.ServiceClient, opts KubeUpgradeOpts) (*KubeUpgrade, error) {
	reqBody, err := common.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}

	return performRequest(c, createURL(c), reqBody)
}

// Delete deletes the current strategy.
func Delete(c *gophercloud.ServiceClient) (r DeleteResult) {
	// kube-upgrade DELETE needs an empty JSON body. The VIM returns 200 on a
	// successful delete, which is not in gophercloud's default DELETE codes
	// (202, 204), so OkCodes is set explicitly.
	m := make(map[string]interface{})
	_, r.Err = c.Delete(deleteURL(c), &gophercloud.RequestOpts{
		JSONBody: m,
		OkCodes:  []int{200, 202, 204},
	})
	return r
}

// ActionStrategy applies or aborts the current strategy.
func ActionStrategy(c *gophercloud.ServiceClient, opts StrategyActionOpts) (*KubeUpgrade, error) {
	reqBody, err := common.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}

	return performRequest(c, actionURL(c), reqBody)
}
