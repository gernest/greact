package hacks

import (
	"github.com/gernest/gs"
	"github.com/gernest/gs/prefix/decl"
)

type AlignContent struct {
	decl.Decl
	names []string
}

var alignContentOldValues = map[string]string{
	"flex-end":      "end",
	"flex-start":    "start",
	"space-between": "justify",
	"space-around":  "distribute",
}

func (a *AlignContent) Names() []string {
	if a.names != nil {
		return a.names
	}
	a.names = []string{"align-content", "flex-line-pack"}
	return a.names
}

func (a *AlignContent) Normalize() string {
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

func (a *AlignContent) Set(prefix string, rule gs.CSSRule) gs.CSSRule {
	e, ok := rule.(gs.SimpleRule)
	if !ok {
		return e
	}
	spec, _ := FlexSpec(prefix)
	if spec == "2012" {
		v, ok := alignContentOldValues[e.Value]
		if !ok {
			v = e.Value
		}
		return gs.SimpleRule{
			Key:   e.Key,
			Value: v,
		}
	}
	return e
}
