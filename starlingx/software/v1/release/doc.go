/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

// Package release contains functionality for working with the StarlingX
// Software (USM) release REST API.
//
// The release resource is served by the usm/software service under
// v1/release. This package covers the read operations consumed by platform
// operators:
//
//	List            — GET v1/release                     (all releases)
//	ListMetapackages — GET v1/release/metapackage?all    (metapackage releases)
//	Get             — GET v1/release/{release_id}         (single release detail)
//
// These are the Software API calls used by the upgrade-operator. Deploy/commit
// and other mutating operations are intentionally not implemented here.
package release
