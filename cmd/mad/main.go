package main

import (
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gernest/mad"
	"golang.org/x/tools/go/ast/astutil"

	"github.com/gernest/mad/config"
	"github.com/gernest/mad/report/console"
	"github.com/gernest/mad/tools"
	"github.com/urfave/cli"
)

const (
	testsDir           = "tests"
	testsOutDir        = "madness"
	localhost          = "http://localhost"
	resourcePath       = "/resource"
	projectDescription = "Inter galactic test runner for Go frontend projects"
	serviceName        = "madtitan"

	// hardcoded import paths
	madImportPath         = "github.com/gernest/mad"
	integrationImportPath = "github.com/gernest/mad/integration"
	coverImportPath       = "github.com/gernest/mad/cover"
	websocketImportPath   = "github.com/gernest/mad/ws"
)

// precompile templates
var (
	integrationTpl  = template.Must(template.New("i").Parse(mainIntegrationTpl))
	indexHTMLTpl    = template.Must(template.New("idx").Parse(idxTpl))
	mainUnitTestTpl = template.Must(template.New("main").Parse(mainUnitTpl))
)

func main() {
	a := cli.NewApp()
	a.Name = serviceName
	a.Usage = projectDescription
	a.Commands = []cli.Command{
		{
			Name:   "test",
			Usage:  "runs the test suites",
			Flags:  config.FLags(),
			Action: runTestsCommand,
		},
		{
			Name:   "coverage",
			Usage:  "calculate code coverage",
			Flags:  config.FLags(),
			Action: runCoverage,
		},
	}
	if err := a.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runTestsCommand(ctx *cli.Context) error {
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
	return streamResponse(context.Background(),
		cfg, &console.ResponseHandler{Verbose: cfg.Verbose})
}

type respHandler interface {
	Handle(*mad.SpecResult)
	Done()
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

	// we need to kee track of the defined unit and integration test functions.
	// This collects functions from all files.
	funcs := &tools.TestNames{}
	importMap := make(map[string]string)
	if cfg.Cover {
		for _, v := range tsPkg.Imports {
			err := instrumentImport(cfg, importMap, v)
			if err != nil {
				return err
			}
		}
		if cfg.Info.ImportPath == madImportPath {
			//WORKAROUND : when we testing the mad package
			imports := []string{
				integrationImportPath,
			}
			for _, v := range imports {
				err := instrumentImport(cfg, importMap, v)
				if err != nil {
					return err
				}
			}
		}

	}
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
		for old, newImport := range importMap {
			astutil.RewriteImport(set, f, old, newImport)
		}
		files = append(files, f)
	}
	if funcs.Unit != nil {
		cfg.UnitFuncs = append(cfg.UnitFuncs, funcs.Unit...)
	}
	if funcs.Integration != nil {
		cfg.IntegrationFuncs = append(cfg.IntegrationFuncs, funcs.Integration...)
	}
	for _, v := range files {
		err := writeFile(out, set, v)
		if err != nil {
			return err
		}
	}
	madImport := importMap[madImportPath]
	if madImport == "" {
		madImport = madImportPath
	}
	wsImport := importMap[websocketImportPath]
	if wsImport == "" {
		wsImport = websocketImportPath
	}
	ctx := map[string]interface{}{
		"config":    cfg,
		"funcs":     funcs,
		"madImport": madImport,
		"wsImport":  wsImport,
	}
	if err = writeMain(cfg.OutputPath, ctx); err != nil {
		return err
	}
	if err = writeIntegrationMain(cfg, importMap); err != nil {
		return err
	}
	return writeIndex(cfg)
}

// instrumentImport processes pkg and adds instrumentation for coverage analysis.
func instrumentImport(cfg *config.Config, importMap map[string]string, pkg string) error {
	if !strings.HasPrefix(pkg, cfg.Info.ImportPath) {
		return nil
	}
	if _, ok := importMap[pkg]; ok {
		return nil
	}
	if pkg == coverImportPath {
		importMap[pkg] = pkg
		return nil
	}
	info := cfg.Info
	if pkg != cfg.Info.ImportPath {
		path := strings.TrimPrefix(pkg, cfg.Info.ImportPath)
		dir := filepath.Join(cfg.Info.Dir, path)
		newPkg, err := build.ImportDir(dir, 0)
		if err != nil {
			return err
		}
		info = newPkg
	}
	for _, v := range info.Imports {
		err := instrumentImport(cfg, importMap, v)
		if err != nil {
			return err
		}
	}
	set := token.NewFileSet()
	base := filepath.Base(pkg)
	if pkg != cfg.Info.ImportPath {
		rel, err := filepath.Rel(cfg.Info.ImportPath, pkg)
		if err != nil {
			return err
		}
		base = rel
	}
	out := filepath.Join(cfg.GeneratedTestPath, base)
	outPkg := filepath.Join(cfg.Info.ImportPath, cfg.OutputDirName, cfg.TestDirName, base)
	os.MkdirAll(out, 0755)
	var buf bytes.Buffer
	for _, v := range info.GoFiles {
		f, err := parser.ParseFile(set, filepath.Join(info.Dir, v), nil, 0)
		if err != nil {
			return err
		}

		f = tools.AddCoverage(set, f)
		for old, newImport := range importMap {
			astutil.RewriteImport(set, f, old, newImport)
		}
		filename := filepath.Join(out, v)
		err = printer.Fprint(&buf, set, f)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filename, buf.Bytes(), 0600)
		if err != nil {
			return err
		}
	}
	importMap[pkg] = outPkg
	return nil
}

func writeIntegrationMain(cfg *config.Config, importMap map[string]string) error {
	if len(cfg.IntegrationFuncs) > 0 {
		data := make(map[string]interface{})
		data["config"] = cfg
		madImport := importMap[madImportPath]
		if madImport == "" {
			madImport = madImportPath
		}
		interImport := importMap[integrationImportPath]
		if interImport == "" {
			interImport = integrationImportPath
		}
		data["madImport"] = madImport
		data["interImport"] = interImport
		var buf bytes.Buffer
		for _, v := range cfg.IntegrationFuncs {
			name := strings.ToLower(v)
			data["PkgName"] = name
			pkg := cfg.GeneratedTestPkg + "/" + name
			data["IntegrationPkg"] = pkg
			e := filepath.Join(cfg.GeneratedTestPath, name)
			os.MkdirAll(e, 0755)
			data["FuncName"] = v
			buf.Reset()
			err := integrationTpl.Execute(&buf, data)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filepath.Join(e, "main.go"), buf.Bytes(), 0600)
			if err != nil {
				return err
			}
			q := make(url.Values)
			q.Set("src", filepath.Join(cfg.TestDirName, name, "main.js"))
			mainFIle := fmt.Sprintf("%s:%d%s?%s",
				localhost, cfg.Port, resourcePath, q.Encode())
			ctx := map[string]interface{}{
				"mainFile": mainFIle,
				"config":   cfg,
			}
			var buf bytes.Buffer
			err = indexHTMLTpl.Execute(&buf, ctx)
			m := filepath.Join(e, "index.html")
			err = ioutil.WriteFile(m, buf.Bytes(), 0600)
			if err != nil {
				return err
			}
			query := make(url.Values)
			query.Set("src", filepath.Join(cfg.TestDirName, name, "index.html"))
			cfg.IntegrationIndexPages = append(cfg.IntegrationIndexPages,
				fmt.Sprintf("%s:%d%s?%s",
					localhost, cfg.Port, resourcePath, query.Encode()))
			if cfg.Build {
				o := filepath.Join(e, "main.js")
				cmd := exec.Command("gopherjs", "build", "-o", o, pkg)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stdout
				if err := cmd.Run(); err != nil {
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
	var buf bytes.Buffer
	err := mainUnitTestTpl.Execute(&buf, ctx)
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
	q := make(url.Values)
	q.Set("src", "main.js")
	mainFIle := fmt.Sprintf("%s:%d%s?%s",
		localhost, cfg.Port, resourcePath, q.Encode())
	ctx := map[string]interface{}{
		"mainFile": mainFIle,
		"config":   cfg,
	}
	var buf bytes.Buffer
	err := indexHTMLTpl.Execute(&buf, ctx)
	m := filepath.Join(cfg.OutputPath, "index.html")
	err = ioutil.WriteFile(m, buf.Bytes(), 0600)
	if err != nil {
		return err
	}
	query := make(url.Values)
	query.Set("src", "index.html")
	cfg.UnitIndexPage = fmt.Sprintf("%s:%d%s?%s",
		localhost, cfg.Port, resourcePath, query.Encode())
	return nil
}

// This is the template for the main entrypoint of the generated unit test
// package.
//
// This goes to the madness/main.go what will eventual be compiled by gopherjs.
// And loaded for execution in the browser.
var mainUnitTpl = `package main

import(
	"{{.config.GeneratedTestPkg}}"
	"{{.wsImport}}"
	"{{.madImport}}"
	"github.com/gernest/mad/cover"
)

func main()  {
	startTest()
}

const testID ="{{.config.UUID}}"
const testPkg ="{{.config.Info.ImportPath}}"

func startTest(){
	go func ()  {
	defer func ()  {
		println(cover.Key+cover.JSON())
	}()
	 w,err:=ws.New(testID)
	 if err!=nil{
		 panic(err)
	 }
	 for _,ts:=range allTests(){
		 v:=mad.Exec(ts)
		 err=w.Report(v,testPkg,testID)
		 if err!=nil{
			println("error "+testID+" "+testPkg+" "+err.Error())
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
import (
	"{{.config.GeneratedTestPkg}}"
	"{{.interImport}}"
	"{{.madImport}}"
	"github.com/gopherjs/vecty"
)

const testID ="{{.config.UUID}}"
const testPkg ="{{.config.Info.ImportPath}}"
{{$n:=.config.TestDirName}}
func main()  {
	vecty.RenderBody(
		&integration.Integration{
			UUID: testID,
			Pkg: testPkg,
			Component: testFunc().(*mad.Component),
		},
	)
}
func testFunc() mad.Integration {
	return mad.SetupIntegration("{{.FuncName}}",{{$n}}.{{.FuncName}}() )
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
func buildGeneratedTestPackage(cfg *config.Config) error {
	o := filepath.Join(cfg.OutputPath, "main.js")
	cmd := exec.Command("gopherjs", "build", "-o", o, cfg.OutputMainPkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	return cmd.Run()
}
