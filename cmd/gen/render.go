package gen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gernest/vected"
	"github.com/urfave/cli"
)

func RenderCMD() cli.Command {
	return cli.Command{
		Name:   "render",
		Usage:  "generates Render functions for components",
		Action: render,
	}
}

// render generates component's render functions from a given package.
func render(ctx *cli.Context) error {
	path := ctx.Args().First()
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return renderDir(ctx)
	}
	return renderFile(ctx)
}

func renderDir(ctx *cli.Context) error {
	fs := token.NewFileSet()
	path := ctx.Args().First()
	pkgs, err := parser.ParseDir(fs, path, func(i os.FileInfo) bool {
		if strings.HasSuffix(i.Name(), "_test.go") {
			return false
		}
		return !strings.HasSuffix(i.Name(), "_render_gen.go")
	}, 0)
	if err != nil {
		return err
	}
	for pkg := range pkgs {
		err = processPackage(path, pkgs[pkg])
		if err != nil {
			return err
		}
	}
	return nil
}

func renderFile(ctx *cli.Context) error {
	return nil
}

func processPackage(path string, pkg *ast.Package) error {
	ctxs := make(map[string]vected.Context)

	// First we collect all structs that implements that emebds vected.Core. Then
	// we check for the Template method which we then use to generate the render
	// functions.
	//
	// The two iterations are important to allow the user to define the Template
	// function in a separate file than the one which defines the component struct.
	for _, file := range pkg.Files {
		for _, v := range file.Decls {
			if g, ok := v.(*ast.GenDecl); ok && g.Tok == token.TYPE {
				for _, spec := range g.Specs {
					vs := spec.(*ast.TypeSpec)
					if typ, ok := vs.Type.(*ast.StructType); ok {
						if len(typ.Fields.List) > 0 {
							for _, f := range typ.Fields.List {
								if x, ok := f.Type.(*ast.SelectorExpr); ok {
									if id, ok := x.X.(*ast.Ident); ok {
										if f.Names == nil && id.Name == "vected" &&
											x.Sel.Name == "Core" {
											ctx := vected.Context{
												StructName: vs.Name.Name,
											}
											ctxs[ctx.StructName] = ctx
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	for _, file := range pkg.Files {
		for _, v := range file.Decls {
			if fn, ok := v.(*ast.FuncDecl); ok {
				if fn.Recv != nil && fn.Name.Name == "Template" &&
					fn.Recv.NumFields() == 1 {
					// pretty.Println(fn)
					fd := fn.Recv.List[0]
					recv := ""
					if fd.Names != nil {
						recv = fd.Names[0].Name
					}
					if typ, ok := fd.Type.(*ast.Ident); ok {
						if ctx, ok := ctxs[typ.Name]; ok {
							ctx.Recv = recv
							if fn.Type.Results.NumFields() == 1 {
								o := fn.Type.Results.List[0]
								if xt, ok := o.Type.(*ast.Ident); ok && xt.Name == "string" {
									if len(fn.Body.List) == 1 {
										rs := fn.Body.List[0].(*ast.ReturnStmt)
										if len(rs.Results) == 1 {
											if ret, ok := rs.Results[0].(*ast.BasicLit); ok {
												v := strings.TrimPrefix(ret.Value, "`")
												v = strings.TrimSuffix(v, "`")
												fmt.Println(v)
												n, err := vected.ParseString(v)
												if err != nil {
													return err
												}
												ctx.Node = n
												ctxs[ctx.StructName] = ctx
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	var c []vected.Context
	for _, v := range ctxs {
		c = append(c, v)
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i].StructName < c[j].StructName
	})
	v, err := vected.GenerateRenderMethod(pkg.Name, c...)
	if err != nil {
		return err
	}
	n := filepath.Join(path, fmt.Sprintf("%s_render_gen.go", pkg.Name))
	return ioutil.WriteFile(n, v, 0600)
}
