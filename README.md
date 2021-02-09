
# Httpmock
Package httpmock provides boilerplate code necessary for testing functions using http.Client by the Transport method.

## Install
```
go get -u github.com/sean-rn/httpmock
```

## Usage

#### func  NewHandlerClient

```go
func NewHandlerClient(handler http.HandlerFunc) *http.Client
```
NewHandlerClient returns *http.Client with Transport replaced with one that only
invokes handler to avoid making real calls.

#### func  NewTransportClient

```go
func NewTransportClient(fn RoundTripperFunc) *http.Client
```
NewTransportClient is a convenience function that returns *http.Client with
Transport replaced to avoid making real calls

#### type RoundTripperFunc

```go
type RoundTripperFunc func(req *http.Request) (*http.Response, error)
```

The RoundTripperFunc type is an adapter to allow the use of ordinary functions
as HTTP Client transports. If f is a function with the appropriate signature,
RoundTripperFunc(f) is a RoundTripper that calls f. This really ought to already
be in the net/http/httptest package, it's just like http.HandlerFunc

#### func  HandlerTransport

```go
func HandlerTransport(handler http.HandlerFunc) RoundTripperFunc
```
HandlerTransport returns a new RoundTripFunc that invokes `handler` and returns
the response it wrote.

#### func (RoundTripperFunc) RoundTrip

```go
func (f RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error)
```
RoundTrip makes RoundTripFunc implement http.RoundTripper. It calls f(req)
