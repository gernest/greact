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
	"html/template"
	"io"
)

// Component is an interface for reneding components. ctx is the data that will
// be passed to the component as context.
type Component interface {
	Identifier
	Templater
	Render(w io.Writer, ctx interface{}) error
}

// Templater is an interface for exposing component's tempates.
type Templater interface {
	Template() string
}

// Identifier is an interface for component identification.
type Identifier interface {
	ID() string
}

var templateCache *template.Template

// Register compiles the components templates and register them. This must be
// called only once in the application life cycle.
//
// The component.ID is used to register a template func that can be used to in
func Register(cmp ...Component) error {
	funcs := make(template.FuncMap)
	for _, v := range cmp {
		funcs[v.ID()] = compile(v)
	}
	tpl := template.New("Vected").Funcs(funcs)
	for _, v := range cmp {
		id := v.ID()
		e := tpl.New(id)
		_, err := e.Parse(v.Template())
		if err != nil {
			return err
		}
	}
	return nil
}

func compile(cmp Component) func(...interface{}) template.HTML {
	return nil
}
