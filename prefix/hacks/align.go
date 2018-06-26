package hacks

import (
	"github.com/gernest/gs"
	"github.com/gernest/gs/prefix/decl"
)

var _ decl.Declaration = (*AlignContent)(nil)
var _ decl.Declaration = (*AlignItems)(nil)

type AlignContent struct {
	decl.Decl
}

func (a *AlignContent) Names() []string {
	return []string{"align-content", "flex-line-pack"}
}

func (a *AlignContent) Normalize(_ string) string {
	return "align-content"
}

func FlexSpec(prefix string) (string, string) {
	var spec string
	switch prefix {
	case "-webkit- 2009", "-moz-":
		spec = "2009"
	case "-ms-":
		spec = "2012"
	case "-webkit-":
		spec = "final"
	}
	if prefix == "-webkit- 2009" {
		prefix = "-webkit-"
	}
	return spec, prefix
}

func (a *AlignContent) Prefixed(prop, prefix string) string {
	spec, pref := FlexSpec(prefix)
	if spec == "2012" {
		return prefix + "flex-line-pack"
	}
	return a.Decl.Prefixed(prop, pref)
}

func (a *AlignContent) Set(rule gs.CSSRule, prefix string) gs.CSSRule {
	e, ok := rule.(gs.SimpleRule)
	if !ok {
		return e
	}
	spec, _ := FlexSpec(prefix)
	if spec == "2012" {
		switch e.Value {
		case "flex-end":
			e.Value = "end"
		case "flex-start":
			e.Value = "start"
		case "space-between":
			e.Value = "justify"
		case "space-around":
			e.Value = "distribute"
		}
		return a.Decl.Set(e, prefix)
	}
	if spec == "final" {
		return a.Decl.Set(rule, prefix)
	}
	return e
}

type AlignItems struct {
	decl.Decl
}

func (a *AlignItems) Prefixed(prop, prefix string) string {
	spec, prefix := FlexSpec(prefix)
	switch spec {
	case "2009":
		return prefix + "box-align"
	case "2012":
		return prefix + "flex-align"
	default:
		return a.Decl.Prefixed(prop, prefix)
	}
}

func (a *AlignItems) Normalize(_ string) string {
	return "align-items"
}

func (a *AlignItems) Set(css gs.CSSRule, prefix string) gs.CSSRule {
	spec, _ := FlexSpec(prefix)
	if spec == "2009" || spec == "2012" {
		if e, ok := css.(gs.SimpleRule); ok {
			switch e.Value {
			case "flex-end":
				e.Value = "end"
			case "flex-start":
				e.Value = "start"
			}
			return a.Decl.Set(e, prefix)
		}
	}
	return a.Decl.Set(css, prefix)
}
