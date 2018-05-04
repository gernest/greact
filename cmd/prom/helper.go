package main

import "github.com/gernest/prom/helper"

func mark(file string, n int) {
	helper.Mark(file, n)
}

func hit(file string, n int) {
	helper.Hit(file, n)
}
