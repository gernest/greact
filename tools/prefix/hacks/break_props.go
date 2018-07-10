package hacks

import (
	"strings"

	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/tools/prefix/decl"
)

var _ decl.Declaration = BreakProps{}

type BreakProps struct {
	decl.Decl
}

func (BreakProps) Names() []string {
	return []string{
		"break-inside", "page-break-inside", "column-break-inside",
		"break-before", "page-break-before", "column-break-before",
		"break-after", "page-break-after", "column-break-after",
	}
}

func (BreakProps) Prefixed(prop, prefix string) string {
	return prefix + "column-" + prop
}

func (BreakProps) Normalize(prop string) string {
	if strings.Contains(prop, "inside") {
		return "break-inside"
	}
	if strings.Contains(prop, "before") {
		return "break-before"
	}
	return "break-after"
}

func (b BreakProps) Set(rule gs.CSSRule, prefix string) gs.CSSRule {
	if e, ok := rule.(gs.SimpleRule); ok {
		if e.Key == "break-inside" && e.Value == "avoid-column" || e.Value == "avoid-page" {
			e.Value = "avoid"
		}
		return b.Decl.Set(e, prefix)
	}
	return rule
}

func (b BreakProps) Insert(rule gs.CSSRule, prefixes ...string) gs.CSSRule {
	if e, ok := rule.(gs.SimpleRule); ok {
		if strings.Contains(e.Value, "region") || strings.Contains(e.Value, "page") {
			return rule
		}
		return b.Decl.Insert(e, prefixes...)
	}
	return rule
}
