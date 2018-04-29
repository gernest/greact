//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

const pkg = "github.com/gernest/vected"

func Grid() error {
	return sh.RunV("gopherjs", "build", "-o", "grid/demo/main.js", pkg+"/grid/demo/")
}
