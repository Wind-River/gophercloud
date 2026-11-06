/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

// Package kubeupgrade contains functionality for querying and clearing the
// StarlingX sysinv kube upgrade record (GET/DELETE /v1/kube_upgrade).
//
// The endpoint reports the in-progress Kubernetes control-plane upgrade and its
// state (e.g. "upgrading-first-master", "post-updated-apps", "upgrade-aborted").
// It is used to decide whether a failed k8s step can still be aborted (rolled
// back to its from_version) and to remove the record once an abort completes.
package kubeupgrade
