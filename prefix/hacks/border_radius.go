package hacks

import (
	"github.com/gernest/gs/prefix/decl"
)

type BorderRadius struct {
	decl.Decl
}

func (BorderRadius) Names() []string {
	return []string{
		"border-radius",
		"border-top-left-radius",
		"border-radius-top-left",
		"border-top-right-radius",
		"border-radius-top-right",
		"border-bottom-left-radius",
		"border-radius-bottom-left",
		"border-bottom-right-radius",
		"border-radius-bottom-right",
	}
}

func (b BorderRadius) Prefixed(prop, prefix string) string {
	if prefix == "-moz-" {
		switch prop {
		case "border-top-left-radius":
			prop = "border-radius-top-left"
		case "border-top-right-radius":
			prop = "border-radius-top-right"
		case "border-bottom-left-radius":
			prop = "border-radius-bottom-left"
		case "border-bottom-right-radius":
			prop = "border-radius-bottom-right"
		}
		return prefix + prop
	}
	return b.Decl.Prefixed(prop, prefix)
}

func (BorderRadius) Normalize(prop string) string {
	switch prop {
	case "border-radius-top-left":
		prop = "border-top-left-radius"
	case "border-radius-top-right":
		prop = "border-top-right-radius"
	case "border-radius-bottom-left":
		prop = "border-bottom-left-radius"
	case "border-radius-bottom-right":
		prop = "border-bottom-right-radius"
	}
	return prop
}
