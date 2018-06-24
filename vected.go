package vected

import "io"

// Component is an interface for reneding components. ctx is the data that will
// be passed to the component as context.
type Component interface {
	Render(w io.Writer, ctx interface{}) error
}
