//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	return sh.RunV("go", "build", "-o", "prom", "./cmd/mad/")
}

func Install() error {
	return sh.RunV("go", "install", "./cmd/mad/")
}

func Test() error {
	return sh.RunV("mad", "test")
}

func V() error {
	return sh.RunV("mad", "test", "--v")
}
