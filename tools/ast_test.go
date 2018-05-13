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
	AddFileNumber(fs, f)
	AddFileNumber(fs, f)
	printer.Fprint(os.Stdout, fs, f)

	t.Error("yay")
}

// func TestMatchTestName(t *testing.T) {
// 	src := `package test

// func TestExample()prom.Test {}
// func TestF()prom.Test { }
// func TestT_M()prom.Test{}

// func Testexample()prom.Test {}
// func Test_Test()prom.Test {}
// `
// 	m := make(map[string]*ast.FuncDecl)
// 	fs := token.NewFileSet()
// 	f, err := parser.ParseFile(fs, "", src, 0)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	astutil.Apply(f, nil, func(c *astutil.Cursor) bool {
// 		node := c.Node()
// 		if f, ok := node.(*ast.FuncDecl); ok {
// 			m[f.Name.Name] = f
// 		}
// 		return true
// 	})
// 	s := []struct {
// 		name string
// 		pass bool
// 	}{
// 		{"TestExample", true},
// 		{"TestF", true},
// 		{"TestT_M", true},
// 		{"Testexample", false},
// 		{"Test_Test", false},
// 	}
// 	for _, v := range s {
// 		fn := m[v.name]
// 		println(v.name, fn == nil)
// 		g := matchTestName(fn.Name.Name, fn.Type)
// 		if g != v.pass {
// 			t.Errorf("%s: expected %v got %v", v.name, v.pass, g)
// 		}
// 	}

// }

// func TestTestName(t *testing.T) {
// 	s := []struct {
// 		name string
// 		pass bool
// 	}{
// 		{"TestExample", true},
// 		{"TestF", true},
// 		{"TestT", true},
// 		{"TestT_M", true},
// 		{"Testexample", false},
// 		{"Test_T", false},
// 	}
// 	for _, v := range s {
// 		g := testName.MatchString(v.name)
// 		if g != v.pass {
// 			t.Errorf("%s: expected %v got %v", v.name, v.pass, g)
// 		}
// 	}
// }
