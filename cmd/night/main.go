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

	"github.com/gernest/mad"

	"github.com/gernest/mad/api"
	"github.com/gernest/mad/config"
	"github.com/gernest/mad/report/console"
	"github.com/gernest/mad/tools"
	"github.com/urfave/cli"
)

const (
	testsDir     = "tests"
	testsOutDir  = "madness"
	serverURL    = "http://localhost:1955"
	resourcePath = "/resource"
	desc         = "Treat your vecty tests like your first date"
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
			Name:   "test",
			Usage:  "runs the test suites",
			Flags:  config.FLags(),
			Action: runTestSuites,
		},
	}
	a.Action = deployDaemonService
	if err := a.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func runTestSuites(ctx *cli.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		return err
	}
	os.RemoveAll(cfg.OutputPath)
	os.MkdirAll(cfg.OutputPath, 0755)
	if err = generateTestPackage(cfg); err != nil {
		return err
	}
	if cfg.Build {
		if err = buildGeneratedTestPackage(cfg); err != nil {
			return err
		}
	}
	req := &api.TestRequest{
		ID:       cfg.UUID,
		Package:  cfg.Info.ImportPath,
		Path:     cfg.Info.Dir,
		Compiled: true,
	}
	res, err := sendTestRequest(cfg, req)
	if err != nil {
		return err
	}
	return streamResponse(context.Background(),
		cfg, res, &console.ResponseHandler{})
}

type respHandler interface {
	Handle(*mad.SpecResult)
	Done()
}

func handleResponse(ts *mad.SpecResult) {
	console.Report(ts)
}

func sendTestRequest(cfg *config.Config, req *api.TestRequest) (*api.TestResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(cfg.ServerURL, "application/json", bytes.NewReader(b))
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
	r := &api.TestResponse{}
	err = json.Unmarshal(b, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// generateTestPackage process the test directory and generate processed files
// in the promtest directory.
//
// Position information is injected in all calls to Error,Errorf,Fatal,FatalF.
// Tis is the simpleset way to provide informative error messages on test failure.
func generateTestPackage(cfg *config.Config) error {
	tsPkg, err := build.ImportDir(cfg.TestPath, 0)
	if err != nil {
		return err
	}
	var files []*ast.File
	out := filepath.Join(cfg.OutputPath, tsPkg.Name)
	os.MkdirAll(out, 0755)
	set := token.NewFileSet()

	// we need to keet ptrack of the defined unit and integration test functions.
	// This collects functions from all files.
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
	if funcs.Unit != nil {
		cfg.UnitFuncs = append(cfg.UnitFuncs, funcs.Unit...)
	}
	if funcs.Integration != nil {
		cfg.Integration = append(cfg.Integration, funcs.Integration...)
	}
	for _, v := range files {
		err := writeFile(out, set, v)
		if err != nil {
			return err
		}
	}
	if err = writeUnitMain(cfg, funcs); err != nil {
		return err
	}
	// if err = writeIntegrationMain(cfg, funcs); err != nil {
	// 	return err
	// }
	return writeIndex(cfg)
}

// generates main package for running all unit tests.
func writeUnitMain(cfg *config.Config, funcs *tools.TestNames) error {
	return writeMain(cfg.OutputPath, map[string]interface{}{
		"config": cfg,
		"funcs":  funcs,
	})
}

var itpl = template.Must(template.New("i").Parse(mainIntegrationTpl))

func writeIntegrationMain(cfg *config.Config, funcs *tools.TestNames) error {
	if len(funcs.Integration) > 0 {
		data := make(map[string]interface{})
		data["testPkg"] = cfg.TestUnitPkg
		var buf bytes.Buffer
		for _, v := range funcs.Integration {
			name := strings.ToLower(v)
			e := filepath.Join(cfg.TestUnitDir, name)
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
			if cfg.Build {
				if err = buildGeneratedTestPackage(cfg); err != nil {
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
	o := filepath.Join(to, filepath.Base(file.Name()))
	// println(o)
	err = ioutil.WriteFile(o, buf.Bytes(), 0600)
	if err != nil {
		return err
	}
	return nil
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
func writeIndex(cfg *config.Config) error {
	idx, err := template.New("idx").Parse(idxTpl)
	if err != nil {
		return err
	}
	pkg := cfg.OutputMainPkg
	q := make(url.Values)
	q.Set("src", pkg+"/main.js")
	q.Set("id", cfg.UUID)
	mainFIle := cfg.ServerURL + resourcePath + "?" + q.Encode()
	ctx := map[string]interface{}{
		"mainFile": mainFIle,
		"config":   cfg,
	}
	var buf bytes.Buffer
	err = idx.Execute(&buf, ctx)
	m := filepath.Join(cfg.OutputPath, "index.html")
	return ioutil.WriteFile(m, buf.Bytes(), 0600)
}

// This is the template for the main entrypoint of the generated unit test
// package.
//
// This goes to the madness/main.go what will eventual be compiled by gopherjs.
// And loaded for execution in the browser.
var mainUnitTpl = `package main

import(
	"{{.config.TestUnitPkg}}"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gernest/mad/ws"
	"github.com/gernest/mad"
)

func main()  {
	startTest()
}

const testID ="{{.config.UUID}}"
const testPkg ="{{.config.Info.ImportPath}}"

func startTest(){
	go func ()  {
	 w,err:=ws.New()
	 if err!=nil{
		 panic(err)
	 }
	 for _,ts:=range allTests(){
		 v:=mad.Exec(ts)
		 err=w.Report(v,testPkg,testID)
		 if err!=nil{
			 println(err)
		 }
	 }
	}()
}
{{$n:=.config.TestDirName}}
func start()mad.Test  {
	return mad.Exec(
		{{range .funcs.Unit -}}
		{{$n}}.{{.}}(),
		{{end -}}
	)
}
func allTests()[]mad.Test  {
	return []mad.Test{
		{{range .funcs.Unit -}}
		mad.Describe("{{.}}",{{$n}}.{{.}}()),
		{{end -}}
	}
}
`

var mainIntegrationTpl = `package main

func main()  {
	// Integration tests are not supported yet
}

`

const idxTpl = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>mad test runner</title>
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
func buildGeneratedTestPackage(cfg *config.Config) error {
	o := filepath.Join(cfg.OutputPath, "main.js")
	cmd := exec.Command("gopherjs", "build", "-o", o, cfg.OutputMainPkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	return cmd.Run()
}
