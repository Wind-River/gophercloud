/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

// Package health contains functionality for querying the StarlingX sysinv
// system health REST API (GET /v1/health/...).
//
// Unlike the other inventory resources, the health endpoints return a
// PLAIN-TEXT report (not JSON) and always respond HTTP 200; pass/fail is
// conveyed by the presence of "[Fail]" lines in the body. This package returns
// the raw text to the caller, which is responsible for parsing it.
package health
