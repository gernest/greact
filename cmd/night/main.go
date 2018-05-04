package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kr/pretty"

	"github.com/gernest/prom/tools"
	"github.com/urfave/cli"
)

const (
	testsDir    = "test"
	testsOutDir = "promtest"
)

func main() {
	a := cli.NewApp()
	a.Name = "prom"
	a.Usage = "Treat your vecty tests like your first date"
	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "pkg",
			Value: ".",
		},
		cli.BoolFlag{
			Name: "cover",
		},
	}
	a.Action = run
	if err := a.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func run(ctx *cli.Context) error {
	pkgPath := ctx.String("pkg")
	cover := ctx.Bool("cover")
	out := filepath.Join(pkgPath, testsOutDir)
	pkg, err := build.ImportDir(pkgPath, 0)
	if err != nil {
		return err
	}
	rootPkg, err := calcPkgPath(pkgPath)
	if err != nil {
		return err
	}
	println(rootPkg)
	importMap := make(map[string]string)
	importMap[rootPkg] = filepath.Join(rootPkg, testsOutDir, pkg.Name)
	if cover {
		var files []*ast.File
		dst := filepath.Join(out, pkg.Name)
		os.MkdirAll(dst, 0755)
		set := token.NewFileSet()
		for _, v := range pkg.GoFiles {
			f, err := parser.ParseFile(set, v, nil, 0)
			if err != nil {
				return err
			}
			files = append(files,
				tools.AddCoverage(set, f),
			)
		}
		for _, v := range files {
			err := writeFile(dst, set, v)
			if err != nil {
				return err
			}
		}
	}
	pretty.Println(importMap)
	return nil
}

func writeFile(to string, fset *token.FileSet, f *ast.File) error {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, f)
	if err != nil {
		return err
	}
	fp := fset.File(f.Pos())
	dst := filepath.Join(to, fp.Name())
	fmt.Println("written to ", dst)
	return ioutil.WriteFile(dst, buf.Bytes(), 0600)
}

func calcPkgPath(base string) (string, error) {
	if filepath.IsAbs(base) {
		src := filepath.Join(build.Default.GOPATH, "src")
		rel, err := filepath.Rel(src, base)
		if err != nil {
			return "", err
		}
		return rel, nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	p := filepath.Join(wd, base)
	return calcPkgPath(p)
}
