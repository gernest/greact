package icon

import (
	"errors"
	"fmt"

	"github.com/gernest/vected"
	"github.com/gernest/vected/lib/cn"
	"github.com/gernest/vected/lib/props"
)

// ErrMissingType is the error returned when trying to render icon component
// without specifying the icon type.
var ErrMissingType = errors.New("Icon: missing type prop")

// Icon is a vected component for rendering antd icons.
type Icon struct {
	vected.Core
}

// Opts is collection of props that the icon component uses. These are expected
// to be passed down by the parent component.
type Opts struct {

	// This is the type of icon that you wish to render.
	//
	// @prop(type)
	Type props.NullString

	// You can optionally pass the class name to be used.
	//
	// @prop(className)
	ClassName props.NullString

	// when true the icon will rotate with animation.
	//
	// @prop(spin)
	Spin props.NullBool

	//TODO
	// Add Style
}

func options(p props.Props) Opts {
	return Opts{
		Type:      p.String("type"),
		ClassName: p.String("className"),
		Spin:      p.Bool("spin"),
	}
}

// New returns a new Icon component or an error when type prop is missing.
func (c *Icon) New(p props.Props) (vected.Component, error) {
	opts := options(p)
	if opts.Type.IsNull || opts.Type.Value == "" {
		return nil, ErrMissingType
	}
	return &Icon{}, nil
}

// Context returns props to be passed to the icon template for rendering.
func (c *Icon) Context(p props.Props) props.Props {
	opts := options(p)
	klass := cn.Join(
		cn.N("anticon", true),
		cn.N("anticon-spin", !opts.Spin.Value || opts.Type.Value == "loading"),
		cn.N(fmt.Sprintf("anticon-%s", opts.Type.Value), true),
		opts.ClassName.Value,
	)
	return props.Merge(p, props.Props{
		"classString": klass,
	})
}

// Template renders html for antd icon.
func (c *Icon) Template() string {
	return `<i {omit . "type" "spin"|attr} className={classString} />`
}
