package tools

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"sort"

	"golang.org/x/tools/go/ast/astutil"
)

const (
	coverageID = "instrumentCodeID"
)

// TestNames is a collection of functions defining integration and unit tests.
type TestNames struct {
	Integration []string
	Unit        []string
}

type testNameMap struct {
	integration map[string]int
	unit        map[string]int
}

// AddFileNumber this will add line number markers to the file's t.Error,
// t.Errof, t.Fatal and t.Fatalf methods.
func AddFileNumber(set *token.FileSet, file *ast.File) *TestNames {
	match := &testNameMap{
		integration: make(map[string]int),
		unit:        make(map[string]int),
	}
	astutil.Apply(file,
		applyLineNumber(set, true, match),
		applyLineNumber(set, false, match),
	)
	o := &TestNames{}
	for k := range match.integration {
		o.Integration = append(o.Integration, k)
	}
	for k := range match.unit {
		o.Unit = append(o.Unit, k)
	}
	if o.Integration != nil {
		sort.Slice(o.Integration, func(i, j int) bool {
			a := o.Integration[i]
			b := o.Integration[j]
			return match.integration[a] < match.integration[b]
		})
	}
	if o.Unit != nil {
		sort.Slice(o.Unit, func(i, j int) bool {
			a := o.Unit[i]
			b := o.Unit[j]
			return match.unit[a] < match.unit[b]
		})
	}
	return o
}

func addStrLit(str, lit string) string {
	return fmt.Sprintf(`"%s%s`, str, lit[1:])
}

func checkName(name string) bool {
	return ast.IsExported(name) && testName.MatchString(name)
}

var testName = regexp.MustCompile(`^Test[[:upper:]].*`)

func checkSignature(typ *ast.FuncType) (isUnit, ok bool) {
	ret := typ.Results != nil && len(typ.Results.List) == 1
	args := typ.Params.List == nil
	if !(ret && args) {
		return
	}
	field := typ.Results.List[0]
	if a, ok := field.Type.(*ast.SelectorExpr); ok {
		id, ok := a.X.(*ast.Ident)
		if !ok {
			return false, false
		}
		n := a.Sel.Name
		if id.Name == "mad" && n == "Test" {
			return true, true
		}
		return false, id.Name == "mad" && n == "Integration"
	}
	return
}

func applyLineNumber(set *token.FileSet, pre bool, match *testNameMap) func(*astutil.Cursor) bool {
	units := 0
	integrations := 0
	return func(c *astutil.Cursor) bool {
		node := c.Node()
		switch e := node.(type) {
		case *ast.FuncDecl:
			if pre {
				if checkName(e.Name.Name) {
					if u, ok := checkSignature(e.Type); ok {
						if u {
							match.unit[e.Name.Name] = units
							units++
						} else {
							match.integration[e.Name.Name] = integrations
							integrations++
						}
					}
				}
			}
		case *ast.CallExpr:
			if s, ok := e.Fun.(*ast.SelectorExpr); ok {
				if id, ok := s.X.(*ast.Ident); ok {
					if id.Name == "mad" && s.Sel.Name == "It" {
						if len(e.Args) == 2 {
							a := e.Args[1].(*ast.FuncLit)
							selector := a.Type.Params.List[0].Names[0].Name
							insert(set, selector, a.Body)
							return false
						}
					}
				}
			}
		}
		return true
	}
}

func insert(set *token.FileSet, sel string, node *ast.BlockStmt) {
	astutil.Apply(node, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		if e, ok := n.(*ast.CallExpr); ok {
			if s, ok := e.Fun.(*ast.SelectorExpr); ok {
				if id, ok := s.X.(*ast.Ident); ok {
					if id.Name == sel {
						file := set.File(e.Pos())
						line := file.Line(e.Pos())
						k := fmt.Sprintf("%s:%v ", file.Name(), line)
						switch s.Sel.Name {
						case "Error":
							e.Args = append([]ast.Expr{
								&ast.BasicLit{
									Value: fmt.Sprintf(`"%s"`, k),
								},
							}, e.Args...)
							return false
						case "Errorf":
							b := e.Args[0].(*ast.BasicLit)
							b.Value = addStrLit(k, b.Value)
							return false
						case "Fatal":
							e.Args = append([]ast.Expr{
								&ast.BasicLit{
									Value: fmt.Sprintf(`"%s"`, k),
								},
							}, e.Args...)
							return false
						case "Fatalf":
							b := e.Args[0].(*ast.BasicLit)
							b.Value = addStrLit(k, b.Value)
							return false
						}
					}
				}
			}
		}
		return true
	})
}
