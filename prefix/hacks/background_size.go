package hacks

import (
	"strings"

	"github.com/gernest/gs"
	"github.com/gernest/gs/prefix/decl"
)

var _ decl.Declaration = BackgroundSize{}

type BackgroundSize struct {
	decl.Decl
}

func (BackgroundSize) Names() []string {
	return []string{"background-size"}
}

func (b BackgroundSize) Set(rule gs.CSSRule, prefix string) gs.CSSRule {
	if e, ok := rule.(gs.SimpleRule); ok {
		v := strings.ToLower(e.Value)
		if prefix == "-webkit-" && !strings.Contains(v, " ") && v != "contain" && v != "cover" {
			e.Value = e.Value + " " + e.Value
		}
		return b.Decl.Set(e, prefix)
	}
	return rule
}
