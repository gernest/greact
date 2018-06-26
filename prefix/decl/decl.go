package decl

import (
	"github.com/gernest/gs"
)

var _ Declaration = Decl{}

type Decl struct{}

func (d Decl) Check(gs.CSSRule) bool {
	return true
}

func (d Decl) Prefixed(prop, prefix string) string {
	return prefix + prop
}

func (d Decl) Normalize(prop string) string {
	return prop
}

func (d Decl) Set(rule gs.CSSRule, prefix string) gs.CSSRule {
	switch e := rule.(type) {
	case gs.SimpleRule:
		e.Key = d.Prefixed(e.Key, prefix)
		return e
	case gs.StyleRule:
		e.Selector = d.Prefixed(e.Selector, prefix)
		return e
	default:
		return e
	}
}

type Declaration interface {
	Check(gs.CSSRule) bool
	Prefixed(prop, prefix string) string
	Normalize(prop string) string
	Set(rule gs.CSSRule, prefix string) gs.CSSRule
}
