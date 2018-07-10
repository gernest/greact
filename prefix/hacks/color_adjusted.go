package hacks

import "github.com/gernest/gs/prefix/decl"

var _ decl.Declaration = ColorAdjust{}

type ColorAdjust struct {
	decl.Decl
}

func (ColorAdjust) Names() []string {
	return []string{
		"color-adjust", "print-color-adjust",
	}
}

func (ColorAdjust) Prefixed(prop, prefix string) string {
	return prefix + "print-color-adjust"
}

func (ColorAdjust) Normalize(prop string) string {
	return "color-adjust"
}
