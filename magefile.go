//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	return sh.RunV("go", "build", "-o", "prom", "./cmd/night/")
}

func Install() error {
	return sh.RunV("go", "install", "./cmd/night/")
}
