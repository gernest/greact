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
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
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
	a.Name = serviceName
	a.Usage = desc
	a.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "starts the test-runner daemon service",
			Action: startDaemon,
		},
		{
			Name:   "stop",
			Usage:  "stops the test-runner daemon service",
			Action: stopDaemon,
		},
		{
			Name:   "install",
			Usage:  "installs the test-runner daemon service",
			Action: installDaemon,
		},
		{
			Name:   "remove",
			Usage:  "uninstall the test-runner daemon service",
			Action: removeDaemon,
		},
		{
			Name:   "status",
			Usage:  "shows status of the test-runner daemon service",
			Action: removeDaemon,
		},
		{
			Name:  "test",
			Usage: "runs the test suites",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "pkg",
					Value: ".",
				},
				cli.BoolFlag{
					Name: "build",
				},
			},
			Action: runTestSuites,
		},
	}
	a.Action = daemonService
	if err := a.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func runTestSuites(ctx *cli.Context) error {
	pkgPath := ctx.String("string")
	buildPkg := ctx.Bool("build")
	out := filepath.Join(pkgPath, testsOutDir)
	rootPkg, err := calcPkgPath(pkgPath)
	if err != nil {
		return err
	}
	if err = generateTestPackage(pkgPath, rootPkg); err != nil {
		return err
	}
	o := filepath.Join(rootPkg, testsOutDir)
	if buildPkg {
		return buildPackage(out, o)
	}
	return nil
}

func generateTestPackage(pkgPath, rootPkg string) error {
	out := filepath.Join(pkgPath, testsOutDir)
	tdir := filepath.Join(pkgPath, testsDir)
	tsPkg, err := build.ImportDir(tdir, 0)
	if err != nil {
		return err
	}
	var files []*ast.File
	dst := out
	os.MkdirAll(filepath.Join(out, tsPkg.Name), 0755)
	set := token.NewFileSet()
	var funcs []string
	for _, v := range tsPkg.GoFiles {
		f, err := parser.ParseFile(set, filepath.Join(tsPkg.Dir, v), nil, 0)
		if err != nil {
			return err
		}
		fn := tools.AddFileNumber(set, f)
		if fn != nil {
			funcs = append(funcs, fn...)
		}
		files = append(files, f)
	}
	for _, v := range files {
		err := writeFile(dst, set, v)
		if err != nil {
			return err
		}
	}
	data := make(map[string]interface{})
	data["testPkg"] = filepath.Join(rootPkg, testsOutDir, testsDir)
	data["funcs"] = funcs
	if err = writeMain(out, data); err != nil {
		return err
	}
	return writeIndex(out)
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

func writeMain(dst string, ctx interface{}) error {
	tpl, err := template.New("main").Parse(mainTpl)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, ctx)
	if err != nil {
		return err
	}
	m := filepath.Join(dst, "main.go")
	b, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(m, b, 0600)
}

func writeIndex(dst string) error {
	m := filepath.Join(dst, "index.html")
	return ioutil.WriteFile(m, []byte(idxTpl), 0600)
}

var mainTpl = `package main

import(
	"{{.testPkg}}"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gernest/prom/report/text"
	"github.com/gernest/prom"
)

func main()  {
	js.Global.Set("startTest", startTest)
	js.Global.Set("start", start)
}

func startTest() string   {
	 v:= start()
	 text.Report(v)
	 return v.ToJson()
}

func start()*prom.ResultCtx  {
	return prom.Exec(
		{{range .funcs -}}
		prom.NewTest("{{.}}").Cases(test.{{.}}()),
		{{end -}}
	)
}

`

const idxTpl = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>prom test runner</title>
</head>

<body>

</body>
<script src="main.js"></script>

</html>`

func buildPackage(dst string, pkg string) error {
	o := filepath.Join(dst, "main.js")
	cmd := exec.Command("gopherjs", "build", "-o", o, pkg)
	return cmd.Run()
}
