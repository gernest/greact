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

	// The absolute path to the the package source files. This field is optional.
	// When this is empty the path will be calculated relative to GOPATH using the
	// Package name.
	Path string `json:"path"`
}
