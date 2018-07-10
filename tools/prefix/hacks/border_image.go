package hacks

import (
	"regexp"

	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/tools/prefix/decl"
)

type BorderImage struct {
	decl.Decl
}

func (BorderImage) Names() []string {
	return []string{"border-image"}
}

var imgRe = regexp.MustCompile(`\s+fill(\s)`)

func (b BorderImage) Set(rule gs.CSSRule, prefix string) gs.CSSRule {
	if e, ok := rule.(gs.SimpleRule); ok {
		m := imgRe.FindStringSubmatch(e.Value)
		if len(m) > 1 {
			e.Value = imgRe.ReplaceAllString(e.Value, m[1])
		}
		return b.Decl.Set(e, prefix)
	}
	return rule
}
