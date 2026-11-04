/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package swupgrade

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

// SwUpgradeOpts is the body for creating a sw-upgrade strategy.
//
// Release carries the release ID(s) to deploy. The VIM REST API changed the
// "release" field type with componentization: product versions > 26.03 expect
// a JSON list, while <= 26.03 expect a string. Release is therefore typed as
// interface{} so the caller decides the wire format (pass a string or a
// []string). The field is required and is not omitted.
//
// NOTE on optional fields: the request body is built with
// starlingx.ConvertToCreateMap, which decodes via mapstructure and does NOT
// honor "omitempty" in mapstructure tags — it only skips nil pointers.
//
// REQUIRED-PRESENT fields (verified against a live VIM): Rollback and Delete
// MUST be present in the request body, even when false. If either is omitted,
// VIM leaves its internal attribute as a WSME UnsetType and then crashes while
// serializing the response with HTTP 500 "Object of type UnsetType is not JSON
// serializable". They are therefore plain bool (always serialized by
// ConvertToCreateMap), NOT pointers.
//
// TRULY-OPTIONAL fields (VIM accepts them absent): MaxParallelWorkerHosts,
// Snapshot, PreUpgradeDeploy, KubeUpgrade are pointers so a nil value is
// omitted. (MaxParallelWorkerHosts absent is fine on AIO-SX; it is not the
// cause of the UnsetType 500 — that is Rollback/Delete.)
type SwUpgradeOpts struct {
	Release                interface{} `json:"release" mapstructure:"release"`
	ControllerApplyType    string      `json:"controller-apply-type" mapstructure:"controller-apply-type"`
	StorageApplyType       string      `json:"storage-apply-type" mapstructure:"storage-apply-type"`
	WorkerApplyType        string      `json:"worker-apply-type" mapstructure:"worker-apply-type"`
	DefaultInstanceAction  string      `json:"default-instance-action" mapstructure:"default-instance-action"`
	AlarmRestrictions      string      `json:"alarm-restrictions,omitempty" mapstructure:"alarm-restrictions"`
	MaxParallelWorkerHosts *int        `json:"max-parallel-worker-hosts,omitempty" mapstructure:"max-parallel-worker-hosts"`
	Rollback               bool        `json:"rollback" mapstructure:"rollback"`
	Delete                 bool        `json:"delete" mapstructure:"delete"`
	Snapshot               *bool       `json:"snapshot,omitempty" mapstructure:"snapshot"`
	PreUpgradeDeploy       *bool       `json:"pre-upgrade-deploy,omitempty" mapstructure:"pre-upgrade-deploy"`
	KubeUpgrade            *string     `json:"kube-upgrade,omitempty" mapstructure:"kube-upgrade"`
}

func performRequest(c *gophercloud.ServiceClient, url string, reqBody interface{}) (*SwUpgrade, error) {
	var resp struct {
		Strategy *SwUpgrade `json:"strategy"`
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

// Show returns the current sw-upgrade strategy, or a nil pointer if no strategy
// is active. The VIM reports {"strategy": null} (or omits the key) when none
// exists; a nil pointer and nil error are returned in that case so callers can
// distinguish "no active strategy" from a parse failure.
func Show(c *gophercloud.ServiceClient) (*SwUpgrade, error) {
	var resp struct {
		Strategy *SwUpgrade `json:"strategy"`
	}
	if _, err := c.Get(showURL(c), &resp, nil); err != nil {
		return nil, err
	}
	// resp.Strategy is nil when "strategy" is absent or explicitly null.
	return resp.Strategy, nil
}

// Create accepts a SwUpgradeOpts struct and creates a new sw-upgrade strategy
// using the values provided.
//
// When Release is nil (a rollback strategy, where the target is implicit) the
// "release" key is dropped from the request body entirely: VIM rejects a
// rollback that also carries a "release" field ("Cannot set both release and
// rollback"), and a literal "release": null still counts as set. Rollback and
// Delete are always sent (see the SwUpgradeOpts doc for the UnsetType 500).
func Create(c *gophercloud.ServiceClient, opts SwUpgradeOpts) (*SwUpgrade, error) {
	reqBody, err := common.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}
	if opts.Release == nil {
		delete(reqBody, "release")
	}

	return performRequest(c, createURL(c), reqBody)
}

// Delete deletes the current strategy.
func Delete(c *gophercloud.ServiceClient) (r DeleteResult) {
	// sw-upgrade DELETE needs an empty JSON body. OkCodes must be set
	// explicitly: with no OkCodes gophercloud treats every status (including a
	// successful 200/204) as unexpected and returns an error.
	m := make(map[string]interface{})
	_, r.Err = c.Delete(deleteURL(c), &gophercloud.RequestOpts{
		JSONBody: m,
		OkCodes:  []int{200, 202, 204},
	})
	return r
}

// ActionStrategy applies or aborts the current strategy.
func ActionStrategy(c *gophercloud.ServiceClient, opts StrategyActionOpts) (*SwUpgrade, error) {
	reqBody, err := common.ConvertToCreateMap(opts)
	if err != nil {
		return nil, err
	}

	return performRequest(c, actionURL(c), reqBody)
}
