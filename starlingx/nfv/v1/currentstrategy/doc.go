/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

// Package currentstrategy contains functionality for querying the VIM
// current-strategy orchestration endpoint.
//
// The current-strategy endpoint (api/orchestration/current-strategy/strategy)
// is read-only. It reports the single orchestration strategy currently known
// to the VIM sw-manager — of whatever type (sw-upgrade, kube-upgrade,
// sw-patch, system-config-update, fw-update, ...) — along with its name and
// state. It is a lightweight way to detect whether any strategy is active
// without probing each strategy type individually.
package currentstrategy
