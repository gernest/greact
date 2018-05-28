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

func TestT_Before() prom.Test {
	return prom.Describe("T.Before",
		prom.It("be called before the testcase", func(rs prom.Result) {
			before := 500
			ts := prom.NewTest("TestT_Before")
			ts.Before(func() {
				before = 200
			})
			ts.Describe("Set before",
				prom.It("must set before value", func(rs prom.Result) {
					before += 100
				}),
			)
			_, err := prom.Exec(ts)
			if err != nil {
				rs.Errorf("expected no error got %v instead", err)
			} else {
				if before != 200 {
					rs.Errorf("expected %v got %v", 200, before)
				}
			}
		}),
	)
}

`

func TestProcess(t *testing.T) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", sanpl11, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Process(fs, f, true)
	// ast.Print(fs, f)
	AddCoverage(fs, f)
	printer.Fprint(os.Stdout, fs, f)

	t.Error("yay")
}
