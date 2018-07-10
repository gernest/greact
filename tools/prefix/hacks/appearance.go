package hacks

import (
	"github.com/gernest/gs/prefix/decl"
)

var _ decl.Declaration = Appearance{}

type Appearance struct {
	decl.Decl
}

func (Appearance) Names() []string {
	return []string{"appearance"}
}
