/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

// Package kubeupgrade contains functionality for working with the VIM
// kube-upgrade orchestration strategy.
//
// The strategy is managed through the VIM sw-manager REST API under
// api/orchestration/kube-upgrade/strategy. A single strategy may exist at a
// time; it is created, shown, applied/aborted via the actions endpoint, and
// deleted.
package kubeupgrade
