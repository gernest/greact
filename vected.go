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
	"github.com/gernest/vected/lib/state"
)

// Component is an interface which defines a unit of user interface.
type Component interface {

	// New must return an initialized component. This acts as a constructor, the
	// props passed to the component from the parent are passed as arguments,
	//
	// Initializing state should happen here.
	New(props.Props) (Component, error)

	// Template this is the vected template that is rendered by the component.
	Template() string
	core()
}

type Core struct{}

func (c *Core) core() {}

// InitState is an interface for exposing initial state. The returned map state
// will available to the component through Core.GetState.
//
// Component should implement this interface if they want to set initial state
// when the component is first created before being rendered.
type InitState interface {
	InitState(props.Props) map[string]interface{}
}

// WillMount is an interface defining a callback which is invoked before the
// component is mounted on the dom.
type WillMount interface {
	ComponentWillMount(UpdateOptions)
}

// DidMount is an interface defining a callback that is invoked after the
// component has been mounted to the dom.
type DidMount interface {
	ComponentDidMount(UpdateOptions)
}

// WillUnmount is an interface defining a callback that is invoked prior to
// removal of the rendered component from the dom.
type WillUnmount interface {
	ComponentWillUnmount(UpdateOptions)
}

// WillReceiveProps is an interface defining a callback that will be called with
// the new props before they are accepted and passed to be rendered.
type WillReceiveProps interface {
	ComponentWillReceiveProps(UpdateOptions)
}

// UpdateOptions these are options passed to lifecycle hooks.
type UpdateOptions struct {
	State     state.State
	NextState state.State
	Props     props.Props
	NextProps props.Props
}

// ShouldUpdate is an interface defining callback that is called before render
// determine if re render is necessary.
type ShouldUpdate interface {
	// If this returns false then re rendering for the component is skipped.
	ShouldComponentUpdate(UpdateOptions) bool
}

// WillUpdate is an interface defining a callback that is called before rendering
type WillUpdate interface {
	// If returned props are not nil, then it will be merged with nextprops then
	// passed to render for rendering.
	ComponentWillUpdate(UpdateOptions) props.Props
}

// DidUpdate defines a callback that is invoked after rendering.
type DidUpdate interface {
	ComponentDidUpdate()
}
