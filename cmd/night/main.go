package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"

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
				tools.ProcessCoverage(set, f),
			)
		}
		for _, v := range files {
			err := writeFile(dst, set, v)
			if err != nil {
				return err
			}
		}
	}

	// for _, v := range pkg.Imports {
	// 	fmt.Println(v)
	// }
	// cover := ctx.Bool("cover")
	// out := filepath.Join(pkg, testsOutDir)
	// base := filepath.Join(pkg, testsDir)
	// _, err := os.Stat(base)
	// if os.IsNotExist(err) {
	// 	return fmt.Errorf("missing %s directory", base)
	// }
	// set := token.NewFileSet()
	// tpkg, err := parser.ParseDir(set, pkg, nil, parser.ParseComments)
	// if err != nil {
	// 	return err
	// }
	// for name := range tpkg {
	// 	processTestPackage(set, tpkg[name])
	// 	// fmt.Println(pkg, name, tpkg[name].Imports)
	// }
	// if p, ok := tpkg[testsDir]; ok {
	// os.MkdirAll(out, 0755)
	// p.Name = testsOutDir
	// var imports []*ast.ImportSpec
	// for name, f := range p.Files {
	// 	on, err := filepath.Rel(base, name)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	ou := filepath.Join(out, on)
	// 	tools.AddFileNumber(set, f)
	// 	var buf bytes.Buffer
	// 	printer.Fprint(&buf, set, f)
	// 	err = ioutil.WriteFile(ou, buf.Bytes(), 0600)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// return writeHelper(out)
	// }
	// return errors.New("can't find test package")
	return nil
}

func processTestPackage(set *token.FileSet, testPkg *ast.Package) {
	for _, file := range testPkg.Files {
		for _, i := range file.Imports {
			println(i.Path.Value)
		}
	}
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

func writeHelper(out string) error {
	h := fmt.Sprintf(helperTpl, tools.MarkCoverName, tools.HitCoverName)
	m := filepath.Join(out, "prom_test_helper.go")
	b, err := format.Source([]byte(h))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(m, b, 0600)
}

const helperTpl = `package main

import "github.com/gernest/prom/helper"

func %s(file string,n int)  {
	helper.Mark(file,n)
}
func %s(file string,n int)  {
	helper.Hit(file,n)
}
`

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
