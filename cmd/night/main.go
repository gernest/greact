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
	"github.com/gernest/prom/config"
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
			Name:   "test",
			Usage:  "runs the test suites",
			Flags:  config.FLags(),
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
		if err = buildPackage(cfg); err != nil {
			return err
		}
	}
	req := &api.TestRequest{
		Package:  cfg.Info.ImportPath,
		Path:     cfg.Info.Dir,
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
func generateTestPackage(cfg *config.Config) error {
	tsPkg, err := build.ImportDir(cfg.TestPath, 0)
	if err != nil {
		return err
	}
	var files []*ast.File
	out := filepath.Join(cfg.OutputPath, tsPkg.Name)
	os.MkdirAll(out, 0755)
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
	data := make(map[string]interface{})
	data["testPkg"] = cfg.TestUnitPkg
	data["funcs"] = funcs
	return writeMain(cfg.OutputPath, data)
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
				if err = buildPackage(cfg); err != nil {
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
	// println(pkg)
	q.Set("src", pkg+"/main.js")
	mainFIle := "http://localhost" + port + "/resource?" + q.Encode()
	println(mainFIle)
	ctx := map[string]string{
		"mainFile": mainFIle,
	}
	var buf bytes.Buffer
	err = idx.Execute(&buf, ctx)
	m := filepath.Join(cfg.OutputPath, "index.html")
	return ioutil.WriteFile(m, buf.Bytes(), 0600)
}

var mainUnitTpl = `package main

import(
	"{{.testPkg}}"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gernest/prom/report/text"
	"github.com/gernest/prom"
)

func main()  {
	js.Global.Set("startTest", startTest)
}

func startTest(){
	 v:= start()
	 text.JSON(v)
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
func buildPackage(cfg *config.Config) error {
	o := filepath.Join(cfg.OutputPath, "main.js")
	cmd := exec.Command("gopherjs", "build", "-o", o, cfg.OutputMainPkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	return cmd.Run()
}
