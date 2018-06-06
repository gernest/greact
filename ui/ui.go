package ui

import (
	"github.com/gernest/gs"
)

var Registry = gs.NewSimpleRegistry(gs.DomRegistry())

var Global = Registry.NewSheet()

func NewSheet() *gs.Sheet {
	return Registry.NewSheet()
}
