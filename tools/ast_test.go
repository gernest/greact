package tools

import (
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"testing"
)

var sanpl11 = `package test

import "github.com/gernest/prom"

func hello(){
	if true{
		markTen("some fist", 10)
	}
	if empty{

	}
}
`

func TestProcess(t *testing.T) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "test.go", sanpl11, 0)
	if err != nil {
		t.Fatal(err)
	}
	Process(fs, f, true)
	// ast.Print(fs, f)
	printer.Fprint(os.Stdout, fs, f)

	t.Error("yay")
}
