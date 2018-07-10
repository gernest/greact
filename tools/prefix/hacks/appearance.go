package hacks

import (
	"github.com/gernest/vected/tools/prefix/decl"
)

var _ decl.Declaration = Appearance{}

type Appearance struct {
	decl.Decl
}

func (Appearance) Names() []string {
	return []string{"appearance"}
}
