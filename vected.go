// Package vected is a component based frontend framework for golang. This
// framework delivers high performance and sleek ui's, that works both on the
// serverside and the frontend.
//
// Go templates are used as the main templating system. The framework is
// inspired by react, especially preact which I used to learn more about how
// react works.
//
// Also, this borrows from vue js. The templates are just normal go templates so
// no need to learn a different syntax.
//
// The user intrface styles is ant design
// see https://github.com/ant-design/ant-design to learn more about ant design.
package vected

import (
	"github.com/gernest/vected/lib/props"
)

// Component is an interface which defines a unit of user interface.
type Component interface {
	New(props.Props) Component
	ID() string
	Template() string

	// all components must embed the Core struct to satisfy this interface.xw
	core()
}

type Core struct{}

func (c *Core) core() {}
