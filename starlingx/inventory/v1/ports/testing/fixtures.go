/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/ports"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

var (
	PortHerp = ports.Port{
		ID:          "ac331b85-b2a1-4f49-9f1a-2888e20d2d59",
		Name:        "enp0s3",
		PCIAddress:  "0000:00:03.0",
		InterfaceID: "fa7721c8-cb24-4f91-bb2e-736cabeff4e2",
	}
	PortDerp = ports.Port{
		ID:          "c3896f57-59c0-4362-b04e-62cbcf76e5c9",
		Name:        "eth1000",
		PCIAddress:  "0000:00:08.0",
		InterfaceID: "2b9fba21-7b7c-42c1-bcab-50841370f049",
	}
)

const PortListBody = `
{
    "ethernet_ports": [
        {
            "dev_id": 0,
            "namedisplay": null,
            "bootp": null,
            "sriov_vf_pdevice_id": null,
            "autoneg": null,
            "speed": 1000,
            "uuid": "ac331b85-b2a1-4f49-9f1a-2888e20d2d59",
            "duplex": null,
            "pdevice": "82540EM Gigabit Ethernet Controller [100e]",
            "capabilities": {},
            "host_uuid": "47dd83d7-1f6b-4d09-b31c-7a271e2f54f8",
            "psdevice": "PRO/1000 MT Desktop Adapter [001e]",
            "link_mode": 0,
            "type": "ethernet",
            "pvendor": "Intel Corporation [8086]",
            "sriov_numvfs": 0,
            "driver": "e1000",
            "updated_at": "2019-11-07T20:52:02.227471+00:00",
            "mac": "08:00:27:00:1d:10",
            "psvendor": "Intel Corporation [8086]",
            "node_uuid": "fa183823-c83a-4f90-befd-cdd6bb61468d",
            "name": "enp0s3",
            "numa_node": -1,
            "created_at": "2019-11-07T20:52:00.532918+00:00",
            "pclass": "Ethernet controller [0200]",
            "mtu": 1500,
            "sriov_vfs_pci_address": "",
            "sriov_totalvfs": null,
            "pciaddr": "0000:00:03.0",
            "dpdksupport": false,
            "sriov_vf_driver": null,
            "interface_uuid": "fa7721c8-cb24-4f91-bb2e-736cabeff4e2"
        },
        {
            "dev_id": 0,
            "namedisplay": null,
            "bootp": null,
            "sriov_vf_pdevice_id": null,
            "autoneg": null,
            "speed": null,
            "uuid": "c3896f57-59c0-4362-b04e-62cbcf76e5c9",
            "duplex": null,
            "pdevice": "Virtio network device [1000]",
            "capabilities": {},
            "host_uuid": "47dd83d7-1f6b-4d09-b31c-7a271e2f54f8",
            "psdevice": "Device [0001]",
            "link_mode": 0,
            "type": "ethernet",
            "pvendor": "Red Hat, Inc. [1af4]",
            "sriov_numvfs": 0,
            "driver": "virtio-pci",
            "updated_at": "2019-11-07T20:52:02.346453+00:00",
            "mac": "08:00:27:a2:f1:e9",
            "psvendor": "Red Hat, Inc. [1af4]",
            "node_uuid": "fa183823-c83a-4f90-befd-cdd6bb61468d",
            "name": "eth1000",
            "numa_node": -1,
            "created_at": "2019-11-07T20:52:00.966096+00:00",
            "pclass": "Ethernet controller [0200]",
            "mtu": 1500,
            "sriov_vfs_pci_address": "",
            "sriov_totalvfs": null,
            "pciaddr": "0000:00:08.0",
            "dpdksupport": true,
            "sriov_vf_driver": null,
            "interface_uuid": "2b9fba21-7b7c-42c1-bcab-50841370f049"
		}
	]
}
`

const PortSingleBody = `
{
	"dev_id": 0,
	"namedisplay": null,
	"bootp": null,
	"sriov_vf_pdevice_id": null,
	"autoneg": null,
	"speed": null,
	"uuid": "c3896f57-59c0-4362-b04e-62cbcf76e5c9",
	"duplex": null,
	"pdevice": "Virtio network device [1000]",
	"capabilities": {},
	"host_uuid": "47dd83d7-1f6b-4d09-b31c-7a271e2f54f8",
	"psdevice": "Device [0001]",
	"link_mode": 0,
	"type": "ethernet",
	"pvendor": "Red Hat, Inc. [1af4]",
	"sriov_numvfs": 0,
	"driver": "virtio-pci",
	"updated_at": "2019-11-07T20:52:02.346453+00:00",
	"mac": "08:00:27:a2:f1:e9",
	"psvendor": "Red Hat, Inc. [1af4]",
	"node_uuid": "fa183823-c83a-4f90-befd-cdd6bb61468d",
	"name": "eth1000",
	"numa_node": -1,
	"created_at": "2019-11-07T20:52:00.966096+00:00",
	"pclass": "Ethernet controller [0200]",
	"mtu": 1500,
	"sriov_vfs_pci_address": "",
	"sriov_totalvfs": null,
	"pciaddr": "0000:00:08.0",
	"dpdksupport": true,
	"sriov_vf_driver": null,
	"interface_uuid": "2b9fba21-7b7c-42c1-bcab-50841370f049"
}
`

func HandlePortListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ihosts/f757b5c7-89ab-4d93-bfd7-a97780ec2c1e/ethernet_ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, PortListBody)
	})
}

func HandlePortGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/ethernet_ports/c3896f57-59c0-4362-b04e-62cbcf76e5c9", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, PortSingleBody)
	})
}
