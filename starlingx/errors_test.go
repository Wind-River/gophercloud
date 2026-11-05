/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package starlingx

import (
	"errors"
	"testing"

	"github.com/gophercloud/gophercloud"
)

func TestFaultStringErrors(t *testing.T) {
	body := []byte(`{"error_message": "{\"faultcode\": \"Client\", \"faultstring\": \"bad thing\\nTraceback (most recent call last):\\n  File x\"}"}`)
	base := gophercloud.ErrUnexpectedResponseCode{Method: "PATCH", URL: "http://x", Actual: 400, Body: body}

	ctx := faultErrorContext{}

	// A sysinv body yields the faultstring, with the traceback stripped.
	var err error = ctx.Error400(base)
	if got := err.Error(); got != "bad thing" {
		t.Fatalf("unexpected message %q", got)
	}

	// Status code and body remain accessible via errors.As.
	var re interface {
		GetStatusCode() int
		GetBody() []byte
	}
	if !errors.As(err, &re) || re.GetStatusCode() != 400 || len(re.GetBody()) == 0 {
		t.Fatalf("accessor not found via errors.As")
	}

	// Without a parsable body, the underlying default message is kept.
	plain := gophercloud.ErrUnexpectedResponseCode{Actual: 500, Body: []byte("oops")}
	err = ctx.Error500(plain)
	if got := err.Error(); got != plain.Error() {
		t.Fatalf("expected fallback to default, got %q", got)
	}

	// 409 is parsed the same way as other StarlingX statuses.
	err = ctx.Error409(base)
	if got := err.Error(); got != "bad thing" {
		t.Fatalf("unexpected message %q", got)
	}
}

func TestParseFaultStringVariants(t *testing.T) {
	cases := []struct {
		name string
		body string
		want string
	}{
		{"sysinv nested", `{"error_message": "{\"faultstring\": \"inner one\"}"}`, "inner one"},
		{"wsme top level", `{"faultstring": "top level reason"}`, "top level reason"},
		{"openstack nested", `{"error": {"message": "nested message"}}`, "nested message"},
		{"plain message", `{"message": "plain message"}`, "plain message"},
		{"not json", `not json`, ""},
		{"empty", ``, ""},
	}
	for _, tc := range cases {
		if got := parseFaultString([]byte(tc.body)); got != tc.want {
			t.Errorf("%s: parseFaultString(%q) = %q, want %q", tc.name, tc.body, got, tc.want)
		}
	}
}
