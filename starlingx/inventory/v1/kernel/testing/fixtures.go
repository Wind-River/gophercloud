/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/kernel"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

var (
	TestHostUUID = "68d437d4-7381-4ad7-92aa-a3fe018ee803"
	KernelHerp   = kernel.Kernel{
		HostID:            TestHostUUID,
		Hostname:          "controller-0",
		ProvisionedKernel: "standard",
		RunningKernel:     "standard",
	}
	KernelDerp = kernel.Kernel{
		HostID:            TestHostUUID,
		Hostname:          "controller-0",
		ProvisionedKernel: "lowlatency",
		RunningKernel:     "standard",
	}
)

const KernelBody = `
{
	"ihost_uuid": "68d437d4-7381-4ad7-92aa-a3fe018ee803",
	"hostname": "controller-0",
	"kernel_provisioned": "standard",
	"kernel_running": "standard",
	"links": [
		{
			"href": "http://192.168.204.1:6385/v1/ihosts/68d437d4-7381-4ad7-92aa-a3fe018ee803/kernel",
			"rel": "self"
		},
		{
			"href": "http://192.168.204.1:6385/ihosts/68d437d4-7381-4ad7-92aa-a3fe018ee803/kernel",
			"rel": "bookmark"
		}
	]
}
`
const KernelBodyUpdated = `
{
	"ihost_uuid": "68d437d4-7381-4ad7-92aa-a3fe018ee803",
	"hostname": "controller-0",
	"kernel_provisioned": "lowlatency",
	"kernel_running": "standard",
	"links": [
		{
			"href": "http://192.168.204.1:6385/v1/ihosts/68d437d4-7381-4ad7-92aa-a3fe018ee803/kernel",
			"rel": "self"
		},
		{
			"href": "http://192.168.204.1:6385/ihosts/68d437d4-7381-4ad7-92aa-a3fe018ee803/kernel",
			"rel": "bookmark"
		}
	]
}
`

const KernelUpdateRequest = `
[
	{
		"op": "replace",
		"path": "/kernel_provisioned",
		"value": "lowlatency"
	}
]
`

func makeUrl(hostid string) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "/ihosts/%s/kernel", hostid)
	return sb.String()
}

func HandleKernelGetSuccessfully(t *testing.T) {
	url := makeUrl(TestHostUUID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, KernelBody)
	})
}

func HandleKernelUpdateSuccessfully(t *testing.T) {
	url := makeUrl(TestHostUUID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, KernelUpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, KernelBodyUpdated)
	})
}
