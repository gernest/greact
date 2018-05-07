package tools

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"sort"

	"golang.org/x/tools/go/ast/astutil"
)

func AddCoverage(set *token.FileSet, file *ast.File) *ast.File {
	astutil.AddImport(set, file, "github.com/gernest/prom/helper")
	astutil.AddImport(set, file, "go/token")
	astutil.Apply(file,
		applyCoverage(set, true),
		applyCoverage(set, false),
	)
	return file
}

func AddFileNumber(set *token.FileSet, file *ast.File) []string {
	match := make(map[string]int)
	astutil.Apply(file,
		applyLineNumber(set, true, match),
		applyLineNumber(set, false, match),
	)
	if len(match) == 0 {
		return nil
	}
	var ls []string
	for k := range match {
		ls = append(ls, k)
	}
	sort.Slice(ls, func(i, j int) bool {
		a := ls[i]
		b := ls[j]
		return match[a] < match[b]
	})
	return ls
}

func addStrLit(str, lit string) string {
	return fmt.Sprintf(`"%s%s`, str, lit[1:])
}

func matchTestName(name string, typ *ast.FuncType) bool {
	ret := typ.Results != nil && len(typ.Results.List) == 1
	args := typ.Params.List == nil
	return ast.IsExported(name) &&
		testName.MatchString(name) &&
		ret && args && checkSignature(typ.Results.List[0])
}

var testName = regexp.MustCompile(`^Test[[:upper:]].*`)

func checkSignature(field *ast.Field) bool {
	if a, ok := field.Type.(*ast.SelectorExpr); ok {
		id, ok := a.X.(*ast.Ident)
		if !ok {
			return false
		}
		if id.Name != "prom" {
			return false
		}
		if a.Sel.Name != "Test" {
			return false
		}
		return true
	}
	return false
}

func mark(num int, pos token.Position) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{
				Name: "idx",
			},
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "helper"},
					Sel: &ast.Ident{Name: "Mark"},
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: fmt.Sprint(num),
					},
					&ast.UnaryExpr{
						Op: token.AND,
						X: &ast.CompositeLit{
							Type: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: "token",
								},
								Sel: &ast.Ident{
									Name: "Position",
								},
							},
							Elts: []ast.Expr{
								&ast.KeyValueExpr{
									Key: &ast.Ident{
										Name: "Filename",
									},
									Value: &ast.BasicLit{
										Kind:  token.STRING,
										Value: fmt.Sprintf(`"%s"`, pos.Filename),
									},
								},
								&ast.KeyValueExpr{
									Key: &ast.Ident{
										Name: "Offset",
									},
									Value: &ast.BasicLit{
										Kind:  token.INT,
										Value: fmt.Sprint(pos.Offset),
									},
								},
								&ast.KeyValueExpr{
									Key: &ast.Ident{
										Name: "Column",
									},
									Value: &ast.BasicLit{
										Kind:  token.INT,
										Value: fmt.Sprint(pos.Line),
									},
								},
								&ast.KeyValueExpr{
									Key: &ast.Ident{
										Name: "Line",
									},
									Value: &ast.BasicLit{
										Kind:  token.INT,
										Value: fmt.Sprint(pos.Line),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func hit(pos token.Position) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "helper"},
				Sel: &ast.Ident{Name: "Hit"},
			},
			Args: []ast.Expr{
				&ast.Ident{
					Name: "idx",
				},
				&ast.UnaryExpr{
					Op: token.AND,
					X: &ast.CompositeLit{
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: "token",
							},
							Sel: &ast.Ident{
								Name: "Position",
							},
						},
						Elts: []ast.Expr{
							&ast.KeyValueExpr{
								Key: &ast.Ident{
									Name: "Filename",
								},
								Value: &ast.BasicLit{
									Kind:  token.STRING,
									Value: fmt.Sprintf(`"%s"`, pos.Filename),
								},
							},
							&ast.KeyValueExpr{
								Key: &ast.Ident{
									Name: "Offset",
								},
								Value: &ast.BasicLit{
									Kind:  token.INT,
									Value: fmt.Sprint(pos.Offset),
								},
							},
							&ast.KeyValueExpr{
								Key: &ast.Ident{
									Name: "Column",
								},
								Value: &ast.BasicLit{
									Kind:  token.INT,
									Value: fmt.Sprint(pos.Line),
								},
							},
							&ast.KeyValueExpr{
								Key: &ast.Ident{
									Name: "Line",
								},
								Value: &ast.BasicLit{
									Kind:  token.INT,
									Value: fmt.Sprint(pos.Line),
								},
							},
						},
					},
				},
			},
		},
	}
}

func applyLineNumber(set *token.FileSet, pre bool, match map[string]int) func(*astutil.Cursor) bool {
	n := 0
	return func(c *astutil.Cursor) bool {
		node := c.Node()
		switch e := node.(type) {
		case *ast.FuncDecl:
			if pre {
				if matchTestName(e.Name.Name, e.Type) {
					match[e.Name.Name] = n
					n++
				}
			}
		case *ast.CallExpr:
			if s, ok := e.Fun.(*ast.SelectorExpr); ok {
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

				// if s.Sel.Name == "Error" {

				// 	if a, ok := e.Args[0].(*ast.BasicLit); ok {
				// 		k := fmt.Sprintf("%s:%v ", file.Name(), line)
				// 		a.Value = addStrLit(k, a.Value)
				// 	}
				// }
			}
		}
		// if e, ok := node.(*ast.CallExpr); ok {
		// 	if s, ok := e.Fun.(*ast.SelectorExpr); ok {
		// 		if s.Sel.Name == "Error" {
		// 			file := set.File(e.Pos())
		// 			line := file.Line(e.Pos())
		// 			if a, ok := e.Args[0].(*ast.BasicLit); ok {
		// 				k := fmt.Sprintf("%s:%v ", file.Name(), line)
		// 				a.Value = addStrLit(k, a.Value)
		// 			}
		// 		}
		// 	}
		// }
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
				size := len(e.Body.List)
				start := set.Position(e.Body.Lbrace)
				end := set.Position(e.Body.Rbrace)
				list := append([]ast.Stmt{mark(size, start)}, e.Body.List...)
				list = append(list, hit(end))
				e.Body.List = list
			}
		}
		return true
	}
}
