//go:build !go1.6
// +build !go1.6

package segmentio

import "net/http"

// http clients on versions of go before 1.6 only support timeout if the
// transport implements the `CancelRequest` method.
func supportsTimeout(transport http.RoundTripper) bool {
	_, ok := transport.(requestCanceler)
	return ok
}

type requestCanceler interface {
	CancelRequest(*http.Request)
}
