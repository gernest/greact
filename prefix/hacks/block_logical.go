package hacks

import (
	"strings"

	"github.com/gernest/gs/prefix/decl"
)

var _ decl.Declaration = BlockLogical{}

type BlockLogical struct {
	decl.Decl
}

func (BlockLogical) Names() []string {
	return []string{
		"border-block-start", "border-block-end",
		"margin-block-start", "margin-block-end",
		"padding-block-start", "padding-block-end",
		"border-before", "border-after",
		"margin-before", "margin-after",
		"padding-before", "padding-after",
	}
}

func (BlockLogical) Prefixed(prop, prefix string) string {
	if strings.Contains(prop, "-start") {
		return prefix + strings.Replace(prop, "-block-start", "-before", -1)
	}
	return prefix + strings.Replace(prop, "-block-end", "-after", -1)
}

func (BlockLogical) Normalize(prop string) string {
	if strings.Contains(prop, "-before") {
		return strings.Replace(prop, "-before", "-block-start", -1)
	}
	return strings.Replace(prop, "-after", "-block-end", -1)
}
