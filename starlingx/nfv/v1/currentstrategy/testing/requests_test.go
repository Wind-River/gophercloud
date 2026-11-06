/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/starlingx/nfv/v1/currentstrategy"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestShowActiveStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleShowActiveSuccessfully(t, CurrentStrategyActiveResponse)

	actual, err := currentstrategy.Show(client.ServiceClient())
	th.AssertNoErr(t, err)
	if actual == nil {
		t.Fatalf("expected a strategy, got nil")
	}
	th.CheckDeepEquals(t, CurrentStrategyActive, *actual)
}

func TestShowNoStrategy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleShowActiveSuccessfully(t, CurrentStrategyNoneResponse)

	actual, err := currentstrategy.Show(client.ServiceClient())
	th.AssertNoErr(t, err)
	if actual != nil {
		t.Fatalf("expected nil strategy when none active, got %+v", *actual)
	}
}
