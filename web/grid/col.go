package grid

import (
	"fmt"

	"github.com/gernest/vected/lib/cn"
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
	return `<div {.others} class="{.classes}">{.children}</div>`
}

// Context returns props needed to render this component's template.
func (Col) Context(ctx props.Props) props.Props {
	cfg := getColProps(ctx)
	classes := cn.Join(
		cn.Name{
			C: fmt.Sprintf("%s-%d", cfg.prefixClass.Value, cfg.span.Value),
			S: cfg.span.IsNull,
		},
		cn.Name{
			C: fmt.Sprintf("%s-order-%d", cfg.prefixClass.Value, cfg.order.Value),
			S: cfg.order.IsNull,
		},
		cn.Name{
			C: fmt.Sprintf("%s-order-%d", cfg.prefixClass.Value, cfg.offset.Value),
			S: cfg.offset.IsNull,
		},
		cn.Name{
			C: fmt.Sprintf("%s-push-%d", cfg.prefixClass.Value, cfg.push.Value),
			S: cfg.push.IsNull,
		},
		cn.Name{
			C: fmt.Sprintf("%s-pull-%d", cfg.prefixClass.Value, cfg.pull.Value),
			S: cfg.pull.IsNull,
		},
	)
	others := ctx.Filter(otherColValues).Attr()
	return props.Props{
		"classes": classes,
		"others":  others,
	}
}

func otherColValues(k, v interface{}) bool {
	if s, ok := k.(string); ok {
		switch s {
		case "span", "order", "offset", "push", "pull", "prefixClass":
			return false
		}
	}
	return true
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
