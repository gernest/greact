package tools

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

const (
	MarkCoverName = "markLineForCOverage"
	HitCoverName  = "hitLineForCOverage"
)

func ProcessCoverage(set *token.FileSet, file *ast.File) *ast.File {
	astutil.AddImport(set, file, "github.com/gernest/prom/helper")
	astutil.Apply(file,
		applyCoverage(set, true),
		applyCoverage(set, false),
	)
	return file
}

func AddFileNumber(set *token.FileSet, file *ast.File) ast.Node {
	return astutil.Apply(file,
		applyLineNumber(set, true),
		applyLineNumber(set, false),
	)
}

type cover struct {
	cover bool
	set   *token.FileSet
}

func (c *cover) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return c
	}
	switch e := node.(type) {
	case *ast.FuncDecl:
	case *ast.IfStmt:
		if c.cover && len(e.Body.List) > 0 {
			ast.Print(c.set, e.Body)
			file := c.set.File(e.Pos())
			line := file.Line(e.Pos())
			name := file.Name()
			list := append([]ast.Stmt{mark(name, line)}, e.Body.List...)
			list = append(list, hit(name, line))
			e.Body.List = list
		}
	}
	return c
}

type fileNum struct {
	set *token.FileSet
}

func (f *fileNum) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return f
	}
	switch e := node.(type) {
	case *ast.FuncDecl:
		name := e.Name.Name
		if !matchTestName(name, e.Type) {
			return nil
		}
	case *ast.CallExpr:
		if s, ok := e.Fun.(*ast.SelectorExpr); ok {
			if s.Sel.Name == "Error" {
				file := f.set.File(e.Pos())
				line := file.Line(e.Pos())
				if a, ok := e.Args[0].(*ast.BasicLit); ok {
					k := fmt.Sprintf("%s:%v ", file.Name(), line)
					a.Value = addStrLit(k, a.Value)
				}
			}
		}
	}
	return f
}

func addStrLit(str, lit string) string {
	return fmt.Sprintf(`"%s%s`, str, lit[1:])
}

func matchTestName(name string, typ *ast.FuncType) bool {
	ret := typ.Results == nil
	args := len(typ.Params.List) == 1
	return ast.IsExported(name) &&
		name != "Test" && strings.HasSuffix(name, "Test") &&
		ret && args && checkSignature(typ.Params.List[0])
}

func checkSignature(field *ast.Field) bool {
	if a, ok := field.Type.(*ast.StarExpr); ok {
		if s, ok := a.X.(*ast.SelectorExpr); ok {
			id, ok := s.X.(*ast.Ident)
			if !ok {
				return false
			}
			if id.Name != "prom" {
				return false
			}
			if s.Sel.Name != "T" {
				return false
			}
			return true
		}
	}
	return false
}

func mark(file string, n int) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.Ident{
				Name: MarkCoverName,
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%s"`, file),
				},
				&ast.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprint(n),
				},
			},
		},
	}
}

func hit(file string, n int) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.Ident{
				Name: HitCoverName,
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%s"`, file),
				},
				&ast.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprint(n),
				},
			},
		},
	}
}

func applyLineNumber(set *token.FileSet, pre bool) func(*astutil.Cursor) bool {
	return func(c *astutil.Cursor) bool {
		node := c.Node()
		if pre {
			if fn, ok := node.(*ast.FuncDecl); ok {
				name := fn.Name.Name
				return matchTestName(name, fn.Type)
			}
			return true
		}
		if e, ok := node.(*ast.CallExpr); ok {
			if s, ok := e.Fun.(*ast.SelectorExpr); ok {
				if s.Sel.Name == "Error" {
					file := set.File(e.Pos())
					line := file.Line(e.Pos())
					if a, ok := e.Args[0].(*ast.BasicLit); ok {
						k := fmt.Sprintf("%s:%v ", file.Name(), line)
						a.Value = addStrLit(k, a.Value)
					}
				}
			}
		}
		return true
	}
}

func applyCoverage(set *token.FileSet, pre bool) func(*astutil.Cursor) bool {
	return func(c *astutil.Cursor) bool {
		node := c.Node()
		if pre {
			return true
		}
		if e, ok := node.(*ast.IfStmt); ok {
			if len(e.Body.List) > 0 {
				file := set.File(e.Pos())
				line := file.Line(e.Pos())
				name := file.Name()
				list := append([]ast.Stmt{mark(name, line)}, e.Body.List...)
				list = append(list, hit(name, line))
				e.Body.List = list
			}
		}
		return true
	}
}
