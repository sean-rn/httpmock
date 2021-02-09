// Package httpmock provides boilerplate code necessary for testing functions using http.Client by the Transport method.
package httpmock

import (
	"net/http"
	"net/http/httptest"
)

// The RoundTripperFunc type is an adapter to allow the use of ordinary functions as HTTP Client transports.
// If f is a function with the appropriate signature, RoundTripperFunc(f) is a RoundTripper that calls f.
// This really ought to already be in the net/http/httptest package, it's just like http.HandlerFunc
type RoundTripperFunc func(req *http.Request) (*http.Response, error)

// RoundTrip makes RoundTripFunc implement http.RoundTripper. It calls f(req)
func (f RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// HandlerTransport returns a new RoundTripFunc that invokes `handler` and returns the response it wrote.
func HandlerTransport(handler http.HandlerFunc) RoundTripperFunc {
	return func(req *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		return w.Result(), nil
	}
}

// NewHandlerClient returns *http.Client with Transport replaced with one that only invokes handler to avoid making real calls.
func NewHandlerClient(handler http.HandlerFunc) *http.Client {
	return &http.Client{Transport: HandlerTransport(handler)}
}

// NewTransportClient is a convenience function that returns *http.Client with Transport replaced to avoid making real calls
func NewTransportClient(fn RoundTripperFunc) *http.Client {
	return &http.Client{Transport: fn}
}
