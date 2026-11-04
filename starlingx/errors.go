/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2026 Wind River Systems, Inc. */

package starlingx

import (
	"encoding/json"
	"strings"

	"github.com/gophercloud/gophercloud"
)

// remoteTracebackMarker marks a Python traceback. sysinv-api returns the
// remote traceback of conductor RPC failures appended to the faultstring.
const remoteTracebackMarker = "Traceback (most recent call last):"

// FaultError wraps a gophercloud.ErrUnexpectedResponseCode and overrides the
// error message with the human readable reason reported by a StarlingX API in
// the response body.
//
// It is installed as the ServiceClient.DefaultErrorContext by the StarlingX
// service constructors, so that only StarlingX requests parse the server
// reason. The core gophercloud error handling is left untouched for every
// other service. parseFaultString understands both the sysinv
// {"error_message": "{...faultstring...}"} body and the top-level
// {"faultstring": ...} WSME body returned by nfv-vim.
type FaultError struct {
	gophercloud.ErrUnexpectedResponseCode
}

// GetStatusCode returns the HTTP status code of the rejected response.
func (e FaultError) GetStatusCode() int { return e.Actual }

// GetBody returns the raw body of the rejected response.
func (e FaultError) GetBody() []byte { return e.Body }

// FaultString returns the reason reported by the server, or an empty string if
// none could be extracted.
func (e FaultError) FaultString() string { return parseFaultString(e.Body) }

func (e FaultError) Error() string {
	if reason := e.FaultString(); reason != "" {
		return reason
	}
	return e.ErrUnexpectedResponseCode.Error()
}

// faultErrorContext implements the gophercloud ErrXXXer interfaces so that a
// StarlingX ServiceClient can substitute FaultError for the default error
// types on the status codes a StarlingX API is likely to return with a body.
type faultErrorContext struct{}

// FaultErrorContext returns an error context value suitable for assigning to a
// gophercloud.ServiceClient.DefaultErrorContext (or RequestOpts.ErrorContext).
// Requests made with it installed return a FaultError, whose message is the
// server reported faultstring when the response body carries one.
func FaultErrorContext() error { return faultErrorContext{} }

// Error is only needed so faultErrorContext satisfies the error interface that
// RequestOpts.ErrorContext expects; it is never surfaced to callers.
func (faultErrorContext) Error() string { return "starlingx fault error context" }

// faultError is the single place that turns a core ErrUnexpectedResponseCode
// into a StarlingX FaultError. The ErrorXXX methods below exist only because
// gophercloud declares a separate ErrXXXer interface per status code, so each
// must be implemented by its exact name; they all delegate here.
func (faultErrorContext) faultError(e gophercloud.ErrUnexpectedResponseCode) error {
	return FaultError{e}
}

func (c faultErrorContext) Error400(e gophercloud.ErrUnexpectedResponseCode) error {
	return c.faultError(e)
}
func (c faultErrorContext) Error403(e gophercloud.ErrUnexpectedResponseCode) error {
	return c.faultError(e)
}
func (c faultErrorContext) Error404(e gophercloud.ErrUnexpectedResponseCode) error {
	return c.faultError(e)
}
func (c faultErrorContext) Error405(e gophercloud.ErrUnexpectedResponseCode) error {
	return c.faultError(e)
}
func (c faultErrorContext) Error408(e gophercloud.ErrUnexpectedResponseCode) error {
	return c.faultError(e)
}
func (c faultErrorContext) Error409(e gophercloud.ErrUnexpectedResponseCode) error {
	return c.faultError(e)
}
func (c faultErrorContext) Error429(e gophercloud.ErrUnexpectedResponseCode) error {
	return c.faultError(e)
}
func (c faultErrorContext) Error500(e gophercloud.ErrUnexpectedResponseCode) error {
	return c.faultError(e)
}
func (c faultErrorContext) Error503(e gophercloud.ErrUnexpectedResponseCode) error {
	return c.faultError(e)
}

// parseFaultString extracts the reason from a StarlingX error body. sysinv
// APIs encode errors as:
//
//	{"error_message": "{\"faultcode\": \"Client\", \"faultstring\": \"<reason>\"}"}
//
// Other OpenStack style bodies, {"error": {"message": "..."}} and
// {"message": "..."}, are also handled. Any appended Python traceback is
// removed.
func parseFaultString(body []byte) string {
	text := strings.TrimSpace(string(body))
	if text == "" {
		return ""
	}
	var outer map[string]interface{}
	if err := json.Unmarshal([]byte(text), &outer); err != nil {
		return ""
	}
	str := func(m map[string]interface{}, k string) string {
		v, _ := m[k].(string)
		return v
	}
	reason := ""
	if msg := str(outer, "error_message"); msg != "" {
		var inner map[string]interface{}
		if err := json.Unmarshal([]byte(msg), &inner); err == nil {
			reason = str(inner, "faultstring")
		} else {
			reason = msg
		}
	}
	if reason == "" {
		reason = str(outer, "faultstring")
	}
	if reason == "" {
		if nested, ok := outer["error"].(map[string]interface{}); ok {
			reason = str(nested, "message")
		}
	}
	if reason == "" {
		reason = str(outer, "message")
	}
	if idx := strings.Index(reason, remoteTracebackMarker); idx >= 0 {
		reason = reason[:idx]
	}
	return strings.TrimSpace(reason)
}
