//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

const pkg = "github.com/gernest/gs/demo"

func Demo() {
	sh.RunV("gopherjs", "build", "-o", "demo/main.js", pkg)
}
