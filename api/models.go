package api

import (
	"time"

	"github.com/gernest/prom"
)

// TestSuite is an object representing a single test suite. A single test suite
// means a single compiled test package.
//
type TestSuite struct {

	// The name of the package being tested. This is a valid Go import path.
	Package string `json:"package"`

	Path string `json:"-"`

	// completed, running,queued
	Status string `json:"status"`

	// The time it took to complete running the test suite.
	Duration time.Time `json:"duration"`

	// This is the response from running startTest function on the browser.
	Result *prom.ResultCtx `json:"results"`
}

// TestRequest is the object sent to the server to initiate the test runner.
type TestRequest struct {

	// The name of the package to be tested. This must be a valid package import
	// path.
	Package string `json:"package"`

	// The absolute path to the the package source files. This must be provided,
	// the daemon has no idea where GOPATH is.
	Path string `json:"path"`

	// True if the test package has already been compiled. By default the server
	// will compile the package after receiving the the test request.
	//
	// Use this to tell the server to just run the tests and skip the compilation
	// step. Note that, if the main.js file is not found then the server will try
	// to compile and completely ignore this field.
	Compiled bool `json:"compiled"`
}

// TestResponse is the object sent to the client after a successful test request.
type TestResponse struct {
	WebsocketURL string `json:"websocket"`
}

// TestStats store statistics about the tests happening in the server.
type TestStats struct {
	Completed []*TestSuite `json:"completed,omitempty"`
	Running   []*TestSuite `json:"running,omitempty"`
	Queued    []*TestSuite `json:"queued,omitempty"`
}
