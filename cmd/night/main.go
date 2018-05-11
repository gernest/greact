package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kr/pretty"

	"github.com/gernest/prom/api"
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
			Action: statusDaemon,
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
	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "hostname with port to bind the server",
			Value: "http://localhost:1955",
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
	if err = generateTestPackage(pkgPath, rootPkg, buildPkg); err != nil {
		return err
	}
	o := filepath.Join(rootPkg, testsOutDir)
	if buildPkg {
		if err = buildPackage(out, o); err != nil {
			return err
		}
	}
	abs, err := filepath.Abs(pkgPath)
	if err != nil {
		return err
	}
	req := &api.TestRequest{
		Package:  rootPkg,
		Path:     abs,
		Compiled: true,
	}
	_, err = sendTestRequest(req)
	return err
}

func sendTestRequest(req *api.TestRequest) (*api.TestResponse, error) {
	h := "http://localhost" + port
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(h, "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(b))
	}
	// println(string(b))
	r := &api.TestResponse{}
	err = json.Unmarshal(b, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func callDaemon(r *api.TestResponse) error {
	return streamResponse(context.Background(), r.WebsocketURL, func(rs *api.TestSuite) {
		pretty.Println(rs)
	})
}

// generateTestPackage process the test directory and generate processed files
// in the promtest directory.
//
// Position information is injected in all calls to Error,Errorf,Fatal,FatalF.
// Tis is the simpleset way to provide informative error messages on test failure.
func generateTestPackage(pkgPath, rootPkg string, buildPkg bool) error {
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
	funcs := &tools.TestNames{}
	for _, v := range tsPkg.GoFiles {
		f, err := parser.ParseFile(set, filepath.Join(tsPkg.Dir, v), nil, 0)
		if err != nil {
			return err
		}
		fn := tools.AddFileNumber(set, f)
		if fn != nil {
			funcs.Integration = append(funcs.Integration, fn.Integration...)
			funcs.Unit = append(funcs.Unit, fn.Unit...)
		}
		files = append(files, f)
	}
	for _, v := range files {
		err := writeFile(dst, set, v)
		if err != nil {
			return err
		}
	}
	tsUnitPkg := filepath.Join(rootPkg, testsOutDir, testsDir)
	if err = writeUnitMain(out, tsUnitPkg, funcs); err != nil {
		return err
	}
	if err = writeIntegrationMain(out, tsUnitPkg, funcs, buildPkg); err != nil {
		return err
	}
	return writeIndex(out, filepath.Join(rootPkg, testsOutDir))
}

// generates main package for running all unit tests.
func writeUnitMain(out, pkg string, funcs *tools.TestNames) error {
	data := make(map[string]interface{})
	data["testPkg"] = pkg
	data["funcs"] = funcs
	return writeMain(out, data)
}

var itpl = template.Must(template.New("i").Parse(mainIntegrationTpl))

func writeIntegrationMain(out, pkg string, funcs *tools.TestNames, buildPkg bool) error {
	if len(funcs.Integration) > 0 {
		data := make(map[string]interface{})
		data["testPkg"] = pkg
		var buf bytes.Buffer
		for _, v := range funcs.Integration {
			name := strings.ToLower(v)
			e := filepath.Join(out, name)
			os.MkdirAll(e, 0755)
			data["funcName"] = v
			buf.Reset()
			err := itpl.Execute(&buf, data)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filepath.Join(e, "main.go"), buf.Bytes(), 0600)
			if err != nil {
				return err
			}
			if buildPkg {
				ipkg, err := calcPkgPath(e)
				if err != nil {
					return err
				}
				if err = buildPackage(e, ipkg); err != nil {
					return err
				}
			}

		}
	}
	return nil
}

// writeFile prints the ast for f using the printer package. The file name is
// obtained from the fset.
//
// The file is created in the to directory.
func writeFile(to string, fset *token.FileSet, f *ast.File) error {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, f)
	if err != nil {
		return err
	}
	file := fset.File(f.Pos())
	dst := filepath.Join(to, file.Name())
	err = ioutil.WriteFile(dst, buf.Bytes(), 0600)
	if err != nil {
		return err
	}
	return nil
}

// calcPkgPath return valid import path for the package defined in base. This
// assumes GOPATH is set, as the path is relative to the GOPATH/src directory.
//
// base can either be relative or absolute.
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

// writeMain creates main.go file which wraps the compiled test functions with
// extra logic for running the tests.
func writeMain(dst string, ctx interface{}) error {
	tpl, err := template.New("main").Parse(mainUnitTpl)
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

//creates index.html file which loads the generated test suite js file.
func writeIndex(dst string, pkg string) error {
	idx, err := template.New("idx").Parse(idxTpl)
	if err != nil {
		return err
	}
	q := make(url.Values)
	println(pkg)
	q.Set("src", pkg+"/main.js")
	mainFIle := "http://localhost" + port + "/resource?" + q.Encode()
	println(mainFIle)
	ctx := map[string]string{
		"mainFile": mainFIle,
	}
	var buf bytes.Buffer
	err = idx.Execute(&buf, ctx)
	m := filepath.Join(dst, "index.html")
	return ioutil.WriteFile(m, buf.Bytes(), 0600)
}

var mainUnitTpl = `package main

import(
	"{{.testPkg}}"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gernest/prom/report/text"
	"github.com/gernest/prom/helper"
	"github.com/gernest/prom"
)

func main()  {
	js.Global.Set("startTest", startTest)
	js.Global.Set("runApp", helper.Run)
}

func startTest(){
	 v:= start()
	 text.Report(v)
}

func start()prom.Test  {
	return prom.Exec(
		{{range .funcs.Unit -}}
		test.{{.}}(),
		{{end -}}
	)
}

`

var mainIntegrationTpl = `package main

import(
	"{{.testPkg}}"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gernest/prom/helper"
	"github.com/gernest/prom"
	"github.com/gopherjs/vecty"
)

func main()  {
	js.Global.Set("startComponents", startComponents)
}

func componentsToRender()vecty.ComponentOrHTML{
	return helper.Wrap(
		test.{{.funcName}}(),
	)
}

func startComponents()  {
	c:=&helper.ComponentRunner{
		Next:componentsToRender,
		AfterFunc:afterComponentSuite,
	}
	vecty.RenderBody(c)
}

func afterComponentSuite(rs *prom.ResultCtx)  {
	parent := js.Global.Get("parent")
	if parent != nil {
		parent.Call("postMessage", rs)
	}
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
<script src="{{.mainFile}}"></script>

</html>`

// test package is compiked to javascript using the gopherjs command. This
// requites gopherjs to be installed and in PATH.
//
// source map is important for coverage computation. So nodejs is required and
// sourcemap module must be installed.
// Taken from the gopherjs README this command
// 	npm install --global source-map-support
// should take care of the sourcemap support.
//
// The output is main.js file in the root directory of the generated test
// package.
func buildPackage(dst string, pkg string) error {
	o := filepath.Join(dst, "main.js")
	cmd := exec.Command("gopherjs", "build", "-o", o, pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	return cmd.Run()
}
