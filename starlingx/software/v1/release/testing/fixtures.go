/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/software/v1/release"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

// ListResponse is a sample GET /v1/release array response.
const ListResponse = `
[
	{
		"release_id": "starlingx-26.10.0",
		"state": "deployed",
		"sw_version": "26.10.0",
		"reboot_required": true,
		"prepatched_iso": false,
		"packages": ["meta-k8s-v1.30.6_26.10.0_all.deb"]
	},
	{
		"release_id": "starlingx-27.03.0",
		"state": "available",
		"sw_version": "27.03.0",
		"reboot_required": true,
		"prepatched_iso": false,
		"packages": ["meta-k8s-v1.32.2_27.03.0_all.deb"],
		"pre_upgrade_deploy": {"enabled": true}
	}
]
`

// ExpectedReleases is the decoded form of ListResponse.
var ExpectedReleases = []release.Release{
	{
		ReleaseID:      "starlingx-26.10.0",
		State:          "deployed",
		SwVersion:      "26.10.0",
		RebootRequired: true,
		PrepatchedIso:  false,
		Packages:       []string{"meta-k8s-v1.30.6_26.10.0_all.deb"},
	},
	{
		ReleaseID:        "starlingx-27.03.0",
		State:            "available",
		SwVersion:        "27.03.0",
		RebootRequired:   true,
		PrepatchedIso:    false,
		Packages:         []string{"meta-k8s-v1.32.2_27.03.0_all.deb"},
		PreUpgradeDeploy: map[string]interface{}{"enabled": true},
	},
}

// MetapackageResponse is a sample GET /v1/release/metapackage?all response.
const MetapackageResponse = `
[
	{
		"release_id": "infra_26.10.3",
		"state": "deployed",
		"sw_version": "26.10.3",
		"component": "infra"
	},
	{
		"release_id": "monitor_26.10.3",
		"state": "available",
		"sw_version": "26.10.3",
		"component": "monitor"
	}
]
`

// ExpectedMetapackages is the decoded form of MetapackageResponse.
var ExpectedMetapackages = []release.Release{
	{
		ReleaseID: "infra_26.10.3",
		State:     "deployed",
		SwVersion: "26.10.3",
		Component: "infra",
	},
	{
		ReleaseID: "monitor_26.10.3",
		State:     "available",
		SwVersion: "26.10.3",
		Component: "monitor",
	},
}

// GetResponse is a sample GET /v1/release/{id} detail response.
const GetResponse = `
{
	"release_id": "starlingx-27.03.0",
	"state": "available",
	"sw_version": "27.03.0",
	"requires": ["starlingx-26.10.0", "starlingx-26.10.3"]
}
`

// ExpectedRelease is the decoded form of GetResponse.
var ExpectedRelease = release.Release{
	ReleaseID: "starlingx-27.03.0",
	State:     "available",
	SwVersion: "27.03.0",
	Requires:  []string{"starlingx-26.10.0", "starlingx-26.10.3"},
}

func HandleListSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/v1/release", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

func HandleListMetapackagesSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/v1/release/metapackage", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		// The "?all" query parameter must be present (bare flag, no value).
		if _, ok := r.URL.Query()["all"]; !ok {
			t.Errorf("expected ?all query parameter, got %q", r.URL.RawQuery)
		}
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

func HandleGetSuccessfully(t *testing.T, releaseID, response string) {
	th.Mux.HandleFunc("/v1/release/"+releaseID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}
