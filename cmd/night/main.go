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

	"golang.org/x/tools/go/ast/astutil"

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
	importMap := make(map[string]string)
	importMap[rootPkg] = filepath.Join(rootPkg, testsOutDir, pkg.Name)
	if cover {
		var files []*ast.File
		dst := filepath.Join(out, pkg.Name)
		os.MkdirAll(dst, 0755)
		set := token.NewFileSet()
		for _, v := range pkg.GoFiles {
			fn := filepath.Join(pkg.Dir, v)
			f, err := parser.ParseFile(set, fn, nil, 0)
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
	tdir := filepath.Join(pkgPath, testsDir)
	tsPkg, err := build.ImportDir(tdir, 0)
	if err != nil {
		return err
	}
	var files []*ast.File
	dst := out
	os.MkdirAll(filepath.Join(out, tsPkg.Name), 0755)
	set := token.NewFileSet()
	for _, v := range tsPkg.GoFiles {
		f, err := parser.ParseFile(set, filepath.Join(tsPkg.Dir, v), nil, 0)
		if err != nil {
			return err
		}
		files = append(files,
			tools.AddFileNumber(set, f),
		)
	}
	for _, v := range files {
		for key, value := range importMap {
			astutil.DeleteImport(set, v, key)
			astutil.AddImport(set, v, value)
		}
		err := writeFile(dst, set, v)
		if err != nil {
			return err
		}
	}
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
	err = ioutil.WriteFile(dst, buf.Bytes(), 0600)
	if err != nil {
		return err
	}
	fmt.Printf("prom: written %s\n", dst)
	return nil
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
