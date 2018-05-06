//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

const pkg = "github.com/gernest/vected"

func Grid() error {
	return sh.RunV("gopherjs", "build", "-o", "grid/demo/main.js", pkg+"/grid/demo/")
}

func Icon() error {
	return sh.RunV("gopherjs", "build", "-o", "icon/demo/main.js", pkg+"/icon/demo/")
}

func Divider() error {
	return sh.RunV("gopherjs", "build", "-o", "divider/demo/main.js", pkg+"/divider/demo/")
}

func Events() error {
	return sh.RunV("gopherjs", "build", "-o", "event/demo/main.js", pkg+"/event/demo/")
}
