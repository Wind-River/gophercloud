/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

// Package kubeupgradehealth contains functionality for querying the StarlingX
// sysinv kube-upgrade health REST API (GET /v1/health/kube-upgrade[/relaxed]).
//
// The general /v1/health endpoint is intentionally not covered here; the
// "health" package name is reserved for it.
//
// Unlike the other inventory resources, these endpoints return a PLAIN-TEXT
// report (not JSON) and always respond HTTP 200; pass/fail is conveyed by the
// presence of "[Fail]" lines in the body. This package returns the raw text to
// the caller, which is responsible for parsing it.
package kubeupgradehealth
