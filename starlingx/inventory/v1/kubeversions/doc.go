/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

// Package kubeversions contains functionality for querying the StarlingX
// sysinv Kubernetes versions REST API (GET /v1/kube_versions).
//
// The endpoint reports the set of Kubernetes versions known to the system and,
// for each, its state (active, available, unavailable, ...) and whether it is
// the current upgrade target. It is read-only here: platform operators use it
// to determine the active and target Kubernetes versions.
package kubeversions
