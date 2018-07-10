package hacks

import (
	"strings"

	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/tools/prefix/decl"
)

var _ decl.Declaration = Animation{}

type Animation struct {
	decl.Decl
}

func (Animation) Names() []string {
	return []string{"animation", "animation-direction"}
}

func (Animation) Check(rule gs.CSSRule) bool {
	if e, ok := rule.(gs.SimpleRule); ok {
		for _, v := range strings.Split(e.Value, " ") {
			low := strings.ToLower(v)
			if low == "reverse" || low == "alternate-reverse" {
				return false
			}
		}
	}
	return true
}
