package httpmock_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sean-rn/httpmock"
)

// Test that responses written by a handler function are in fact returned by the function returned from HandlerTransport
func TestHandlerTransport(t *testing.T) {
	testReq := httptest.NewRequest("GET", "/test/url", strings.NewReader("Hello Function!"))

	roundTripperFunc := httpmock.HandlerTransport(func(rw http.ResponseWriter, r *http.Request) {
		if r != testReq {
			t.Error("Expected testReq to be passed to the handler.")
		}
		rw.Write([]byte("I'm Responding!"))
	})

	resp, err := roundTripperFunc.RoundTrip(testReq)
	if err != nil {
		t.Errorf("roundTripperFunc threw: %v", err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Reading response body threw: %v", err)
	}

	const expectedBody = "I'm Responding!"
	if string(bodyBytes) != expectedBody {
		t.Errorf(`Expected %#v but got %#v`, expectedBody, string(bodyBytes))
	}
}
