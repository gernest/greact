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
	"bytes"
	"html/template"

	"github.com/gernest/vected/props"
)

// Component is an interface for reneding components. ctx is the data that will
// be passed to the component as context.
type Component interface {
	Identifier
	Templater
	Render(props.Props) props.Props
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

func init() {
	templateCache = template.New("VectedUI")
	templateCache.Delims("{", "}")
}

// Register compiles the components templates and register them. This must be
// called only once in the application life cycle.
//
// The component.ID is used to register a template func that can be used to in
func Register(cmp ...Component) error {
	funcs := make(template.FuncMap)
	for _, v := range cmp {
		funcs[v.ID()] = compile(v)
	}
	tpl := templateCache.Funcs(funcs)
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

func compile(cmp Component) func(props.Props) (template.HTML, error) {
	return func(ctx props.Props) (template.HTML, error) {
		tpl := templateCache.Lookup(cmp.ID())
		if tpl != nil {
			var buf bytes.Buffer
			err := tpl.Execute(&buf, cmp.Render(ctx))
			if err != nil {
				return "", err
			}
			return template.HTML(buf.String()), nil
		}
		return "", nil
	}
}

func RenderHTML(tpl string, ctx props.Props) (template.HTML, error) {
	tree, err := templateCache.Clone()
	if err != nil {
		return "", err
	}
	t, err := tree.Parse(tpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, ctx)
	if err != nil {
		return "", err
	}
	return template.HTML(buf.String()), nil
}
