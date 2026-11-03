/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"strings"
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/software/v1/deploy"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestPrecheckPass(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	captured := HandlePrecheckJSON(t, 200, PassBody)

	body, err := deploy.Precheck(client.ServiceClient(), deploy.PrecheckOpts{
		Releases: []string{"starlingx-27.03.0"},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, PassBody, body)
	th.AssertEquals(t, "POST", captured.Method)

	releases, ok := captured.Body["releases"].([]interface{})
	if !ok || len(releases) != 1 || releases[0] != "starlingx-27.03.0" {
		t.Errorf("releases = %v, want [starlingx-27.03.0]", captured.Body["releases"])
	}
	// force and pre_upgrade_deploy omitted when unset.
	if _, present := captured.Body["force"]; present {
		t.Errorf("force should be omitted, got %v", captured.Body["force"])
	}
	if _, present := captured.Body["pre_upgrade_deploy"]; present {
		t.Errorf("pre_upgrade_deploy should be omitted, got %v", captured.Body["pre_upgrade_deploy"])
	}
}

// TestPrecheckFail406 verifies that a 406 response is NOT an error — it is the
// normal precheck-failed response and its body is returned to the caller.
func TestPrecheckFail406(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePrecheckJSON(t, 406, FailBody)

	body, err := deploy.Precheck(client.ServiceClient(), deploy.PrecheckOpts{
		Releases: []string{"starlingx-27.03.0"},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, FailBody, body)
	if !strings.Contains(body, "error") {
		t.Errorf("expected error body, got %q", body)
	}
}

func TestPrecheckForceAndPreUpgrade(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	captured := HandlePrecheckJSON(t, 200, PassBody)

	_, err := deploy.Precheck(client.ServiceClient(), deploy.PrecheckOpts{
		Releases:         []string{"infra_27.03.0", "monitor_27.03.0"},
		Force:            "true",
		PreUpgradeDeploy: true,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "true", captured.Body["force"])
	if v, _ := captured.Body["pre_upgrade_deploy"].(bool); !v {
		t.Errorf("pre_upgrade_deploy = %v, want true", captured.Body["pre_upgrade_deploy"])
	}
	releases, _ := captured.Body["releases"].([]interface{})
	if len(releases) != 2 {
		t.Errorf("releases = %v, want 2 entries", captured.Body["releases"])
	}
}

func TestPrecheckLegacy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	captured := HandlePrecheckLegacy(t, 200, PassBody)

	body, err := deploy.PrecheckLegacy(client.ServiceClient(), "starlingx-26.03.0", true)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, PassBody, body)
	th.AssertEquals(t, "POST", captured.Method)
	if !strings.Contains(captured.RawQuery, "release=starlingx-26.03.0") {
		t.Errorf("query = %q, want release param", captured.RawQuery)
	}
	if !strings.Contains(captured.RawQuery, "force=true") {
		t.Errorf("query = %q, want force=true", captured.RawQuery)
	}
}

func TestPrecheckLegacyNoForce(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	captured := HandlePrecheckLegacy(t, 406, FailBody)

	body, err := deploy.PrecheckLegacy(client.ServiceClient(), "starlingx-26.03.0", false)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, FailBody, body)
	if strings.Contains(captured.RawQuery, "force") {
		t.Errorf("query = %q, should not contain force", captured.RawQuery)
	}
}

func TestPrecheckError500(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePrecheckJSON(t, 500, `{"error":"boom"}`)

	_, err := deploy.Precheck(client.ServiceClient(), deploy.PrecheckOpts{
		Releases: []string{"starlingx-27.03.0"},
	})
	if err == nil {
		t.Fatal("expected error for 500 response, got nil")
	}
}

func TestDeployList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeployList(t, DeployListBody)

	deploys, err := deploy.List(client.ServiceClient())
	th.AssertNoErr(t, err)
	if len(deploys) != 1 {
		t.Fatalf("expected 1 deploy, got %d", len(deploys))
	}
	th.AssertEquals(t, "host-rollback-done", deploys[0].State)
	th.AssertEquals(t, "starlingx-26.10.0", deploys[0].FromRelease)
	th.AssertEquals(t, "starlingx-26.10.2", deploys[0].ToRelease)
}

func TestDeployListEmpty(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeployList(t, DeployListEmptyBody)

	deploys, err := deploy.List(client.ServiceClient())
	th.AssertNoErr(t, err)
	if len(deploys) != 0 {
		t.Fatalf("expected 0 deploys, got %d", len(deploys))
	}
}
