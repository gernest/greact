package grid

import (
	"fmt"

	cn "github.com/gernest/classnames"
	"github.com/gernest/vected/props"
)

// Col implements vected.Component
type Col struct {
}

// ID returns component's name.
func (Col) ID() string {
	return "Col"
}

// Template returns go template for rendering grid Columns.
func (Col) Template() string {
	return `<div {{.others}} class={{.classes}}>{{.children}}</div>`
}

// Render returns props needed to render this component's template.
func (Col) Render(ctx props.Props) props.Props {
	cfg := getColProps(ctx)
	classes := cn.Join(
		cn.Name{
			Class: fmt.Sprintf("%s-%d", cfg.prefixClass.Value, cfg.span.Value),
			Skip:  cfg.span.IsNull,
		},
		cn.Name{
			Class: fmt.Sprintf("%s-order-%d", cfg.prefixClass.Value, cfg.order.Value),
			Skip:  cfg.order.IsNull,
		},
		cn.Name{
			Class: fmt.Sprintf("%s-order-%d", cfg.prefixClass.Value, cfg.offset.Value),
			Skip:  cfg.offset.IsNull,
		},
		cn.Name{
			Class: fmt.Sprintf("%s-push-%d", cfg.prefixClass.Value, cfg.push.Value),
			Skip:  cfg.push.IsNull,
		},
		cn.Name{
			Class: fmt.Sprintf("%s-pull-%d", cfg.prefixClass.Value, cfg.pull.Value),
			Skip:  cfg.pull.IsNull,
		},
	)
	return props.Props{
		"classes": classes,
	}
}

type colProps struct {
	span        props.NullInt
	order       props.NullInt
	offset      props.NullInt
	push        props.NullInt
	pull        props.NullInt
	prefixClass props.NullString
}

func getColProps(p props.Props) colProps {
	return colProps{
		span:        p.Int("span"),
		order:       p.Int("order"),
		offset:      p.Int("offset"),
		push:        p.Int("push"),
		pull:        p.Int("pull"),
		prefixClass: p.StringV("prefixClass", ".ant-col"),
	}
}
