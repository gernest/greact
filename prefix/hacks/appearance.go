package hacks

import (
	"github.com/gernest/gs/prefix/decl"
)

type Appearance struct {
	decl.Decl
}

func (Appearance) Names() []string {
	return []string{"appearance"}
}
