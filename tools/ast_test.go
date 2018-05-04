package tools

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"testing"
)

var sanpl11 = `package test

import "github.com/gernest/prom"

func mark(file string, n int) {
	helper.Mark(file, n)
}

func hit(file string, n int) {
	helper.Hit("file", 10)
	if true{
		z:=1
	}
}
`

func TestProcess(t *testing.T) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "test.go", sanpl11, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Process(fs, f, true)
	ast.Print(fs, f)
	ProcessCoverage(fs, f)
	printer.Fprint(os.Stdout, fs, f)

	t.Error("yay")
}
