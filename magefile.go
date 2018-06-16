//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Test executes mad test command.
func Test() error {
	return sh.RunV("mad", "test")
}
