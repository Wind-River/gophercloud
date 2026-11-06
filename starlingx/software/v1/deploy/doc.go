/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

// Package deploy contains functionality for the StarlingX Software (USM) deploy
// REST API, specifically the deploy precheck endpoint used by platform
// operators to validate a release deploy before executing it.
//
//	Precheck — POST v1/deploy/precheck
//
// The request format changed with componentization (product version > 26.03):
//
//   - Product version > 26.03:  JSON body
//     {"releases": [...], "force": "true"?, "pre_upgrade_deploy": true?}
//   - Product version <= 26.03: query parameters
//     ?release=<release_id>[&force=true]   (no JSON body)
//
// This package exposes both request shapes; the caller selects the correct one
// for the current product version.
//
// The precheck endpoint returns the precheck result in the RESPONSE BODY on
// BOTH success (HTTP 200) AND failure (HTTP 406 Not Acceptable) — 406 is the
// normal "precheck failed" response, not a transport error. This package treats
// both 200 and 406 as acceptable and returns the raw body to the caller, which
// is responsible for parsing it (the body may be JSON with an "error" field, or
// plain text with "[Fail]" lines on older USM versions).
package deploy
