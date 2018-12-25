// Package spncore is a client library for communicating with an
// instance of spn-core using through the server's HTTP port.
package spncore

import "net/http"

// SetCursorDone is the success message returned by spn-core when a cursor
// update succeeds.
const SetCursorDone = "Done"

// HTTP represents the http client that a spncore client uses to make http
// requests.
type HTTP interface {
	Do(req *http.Request) (*http.Response, error)
}

// confirm interface conformity
var _ HTTP = http.DefaultClient
