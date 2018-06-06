package tools

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"

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

// AddFileNumber this will add line number markers to the file's t.Error,
// t.Errof, t.Fatal and t.Fatalf methods.
func AddFileNumber(set *token.FileSet, file *ast.File) TestNames {
	v := New(set)
	ast.Walk(v, file)
	return v.Names
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

func insert(set *token.FileSet, sel string, node *ast.BlockStmt) {
	for _, v := range node.List {
		astutil.Apply(v, nil, func(c *astutil.Cursor) bool {
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

}

func New(set *token.FileSet) *Visitor {
	return &Visitor{
		Set:    set,
		number: &numbers{set: set}}
}

type Visitor struct {
	Names  TestNames
	Set    *token.FileSet
	number *numbers
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch e := node.(type) {
	case *ast.FuncDecl:
		if checkName(e.Name.Name) {
			if u, ok := checkSignature(e.Type); ok {
				if u {
					v.Names.Unit = append(v.Names.Unit, e.Name.Name)
				} else {
					v.Names.Integration = append(v.Names.Integration, e.Name.Name)
				}
			}
		}
		for _, field := range e.Type.Params.List {
			if sel, ok := field.Type.(*ast.SelectorExpr); ok {
				if sel.Sel.Name == "T" {
					if id, ok := sel.X.(*ast.Ident); ok {
						if id.Name == "mad" {
							selector := field.Names[0].Name
							insert(v.Set, selector, e.Body)
							return nil
						}
					}
				}
			}
		}
		return v.number
	}
	return v
}

type numbers struct {
	set *token.FileSet
}

func (n *numbers) Visit(node ast.Node) ast.Visitor {
	switch e := node.(type) {
	case *ast.FuncLit:
		for _, field := range e.Type.Params.List {
			if sel, ok := field.Type.(*ast.SelectorExpr); ok {
				if sel.Sel.Name == "T" {
					if id, ok := sel.X.(*ast.Ident); ok {
						if id.Name == "mad" {
							selector := field.Names[0].Name
							insert(n.set, selector, e.Body)
						}
					}
				}
			}
		}
	}
	return n
}
