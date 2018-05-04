//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	return sh.RunV("go", "build", "./cmd/prom/")
}
