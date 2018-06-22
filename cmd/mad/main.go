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

	"github.com/gernest/mad/annotate"
	"github.com/gernest/mad/config"
	"github.com/gernest/mad/launcher/chrome"
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
	mainCoverTpl    = template.Must(template.New("cover").Parse(coverTpl))
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
			Name:  "coverage",
			Usage: "calculate code coverage",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "text",
					Usage: "Formats output to be like what go test -cover does",
				},
			},
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
	executionContext, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()
	os.RemoveAll(cfg.OutputPath)
	os.MkdirAll(cfg.OutputPath, 0755)
	if err = generateTestPackage(executionContext, cfg); err != nil {
		return err
	}
	if cfg.Dry {
		return nil
	}
	if cfg.Build {
		o := filepath.Join(cfg.OutputPath, "main.js")
		err = buildPackage(executionContext, o, cfg.OutputMainPkg)
		if err != nil {
			return err
		}
	}
	for _, v := range cfg.Browsers {
		switch v {
		case "chrome":
			chrome, err := chrome.New(chrome.Options{
				Verbose:     cfg.Verbose,
				Port:        cfg.DevtoolPort,
				ChromeFlags: []string{"--headless"},
			})
			if err != nil {
				return err
			}
			var h respHandler
			if cfg.JSON != "" {
				h = console.NewJSON(cfg.JSON)
			} else {
				h = console.New(cfg.Verbose)
			}
			err = streamResponse(executionContext,
				cfg, chrome, h)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type respHandler interface {
	Handle(*mad.SpecResult)
	Done() error
}

// generateTestPackage process the test directory and generate processed files
// in the promtest directory.
func generateTestPackage(ctx context.Context, cfg *config.Config) error {
	for _, v := range cfg.TestInfo {
		err := createTestPackage(cfg, v)
		if err != nil {
			return err
		}
	}
	wsImport := cfg.ImportMap[websocketImportPath]
	if wsImport == "" {
		wsImport = websocketImportPath
	}
	madImport := cfg.ImportMap[madImportPath]
	if madImport == "" {
		madImport = madImportPath
	}
	err := writeMain(cfg.OutputPath, map[string]interface{}{
		"config":    cfg,
		"wsImport":  wsImport,
		"madImport": madImport,
	})
	if err != nil {
		return err
	}
	err = generateIntegrationPackages(ctx, cfg)
	if err != nil {
		return err
	}
	return writeIndex(cfg)
}

// This loads files found in path, process them and generate processed package
// into the directory specified by cfg.OutputPath.
func createTestPackage(cfg *config.Config, out *config.Info) error {
	var files []*ast.File
	os.MkdirAll(out.OutputPath, 0755)
	set := token.NewFileSet()
	// we need to keep track of the defined unit and integration test functions.
	// This collects functions from all files.
	funcs := &tools.TestNames{}
	if cfg.Cover {
		for _, v := range out.Package.Imports {
			err := instrumentPackage(cfg, v)
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
				err := instrumentPackage(cfg, v)
				if err != nil {
					return err
				}
			}
		}
	}
	for _, v := range out.Package.GoFiles {
		f, err := parser.ParseFile(set, filepath.Join(out.Package.Dir, v), nil, 0)
		if err != nil {
			return err
		}
		fn := tools.AddFileNumber(set, f)
		funcs.Integration = append(funcs.Integration, fn.Integration...)
		funcs.Unit = append(funcs.Unit, fn.Unit...)
		for old, newImport := range cfg.ImportMap {
			astutil.RewriteImport(set, f, old, newImport)
		}
		files = append(files, f)
	}
	for _, v := range files {
		err := writeFile(out.OutputPath, set, v)
		if err != nil {
			return err
		}
	}
	hasTests := len(funcs.Unit) > 0 || len(funcs.Integration) > 0
	if hasTests {
		cfg.TestNames[out] = funcs
	}
	return nil
}

const coverTpl = `
package {{.pkgName}}
import(
	"github.com/gernest/mad/cover"
)

func coverage()[]cover.Profile  {
	return []cover.Profile{
	{{- $mode:=.mode}}
	{{- range $k,$v:=.vars}}
	cover.File("{{$k}}","{{$mode}}", {{$v}}.Count[:], {{$v}}.Pos[:], {{$v}}.NumStmt[:]),
	{{- end}}
	}
}
func init()  {
	cover.Register("{{.pkg}}",coverage)
}



`

// instrumentPackage processes pkg and adds instrumentation for coverage analysis.
func instrumentPackage(cfg *config.Config, pkg string) error {
	if !strings.HasPrefix(pkg, cfg.Info.ImportPath) {
		return nil
	}
	if _, ok := cfg.ImportMap[pkg]; ok {
		return nil
	}
	if pkg == coverImportPath {
		cfg.ImportMap[pkg] = pkg
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
		err := instrumentPackage(cfg, v)
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
	out := filepath.Join(cfg.OutputPath, base)
	outPkg := filepath.Join(cfg.Info.ImportPath, cfg.OutputDirName, base)
	os.MkdirAll(out, 0755)
	var buf bytes.Buffer
	coverVarNames := make(map[string]string)
	for k, v := range info.GoFiles {
		f, err := parser.ParseFile(set, filepath.Join(info.Dir, v), nil, parser.ParseComments)
		if err != nil {
			return err
		}
		buf.Reset()
		varName := fmt.Sprintf("coverStats%d", k)
		fullName := info.ImportPath + "/" + v
		coverVarNames[fullName] = varName
		err = annotate.Annotate(&buf, set, f, annotate.Options{
			Name:    fullName,
			VarName: varName,
			Mode:    cfg.Covermode,
		})
		if err != nil {
			return err
		}
		varDef := buf.String()
		buf.Reset()
		for old, newImport := range cfg.ImportMap {
			astutil.RewriteImport(set, f, old, newImport)
		}
		filename := filepath.Join(out, v)
		err = printer.Fprint(&buf, set, f)
		if err != nil {
			return err
		}
		buf.WriteString(varDef)
		err = ioutil.WriteFile(filename, buf.Bytes(), 0600)
		if err != nil {
			return err
		}
	}
	buf.Reset()
	data := map[string]interface{}{
		"pkgName": info.Name,
		"vars":    coverVarNames,
		"mode":    cfg.Covermode,
		"pkg":     pkg,
	}
	err := mainCoverTpl.Execute(&buf, data)
	if err != nil {
		return err
	}
	filename := filepath.Join(out, "coverrage_stats.go")
	b, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0600)
	if err != nil {
		return err
	}
	cfg.ImportMap[pkg] = outPkg
	return nil
}

// This genegrates and compiles packages with integration tests. Each
// integration test lives in it's own separate package.
func generateIntegrationPackages(ctx context.Context, cfg *config.Config) error {
	for info, funcs := range cfg.TestNames {
		if len(funcs.Integration) > 0 {
			data := make(map[string]interface{})
			data["config"] = cfg
			data["info"] = info
			madImport := cfg.ImportMap[madImportPath]
			if madImport == "" {
				madImport = madImportPath
			}
			interImport := cfg.ImportMap[integrationImportPath]
			if interImport == "" {
				interImport = integrationImportPath
			}
			data["madImport"] = madImport
			data["interImport"] = interImport
			var buf bytes.Buffer
			for _, v := range funcs.Integration {
				name := strings.ToLower(v)
				data["PkgName"] = name
				pkg := info.ImportPath + "/" + name
				data["IntegrationPkg"] = pkg
				e := filepath.Join(info.OutputPath, name)
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
				q.Set("src", filepath.Join(info.RelativePath, name, "main.js"))
				mainFIle := fmt.Sprintf("%s:%d%s?%s",
					localhost, cfg.Port, resourcePath, q.Encode())
				var buf bytes.Buffer
				err = indexHTMLTpl.Execute(&buf, map[string]interface{}{
					"mainFile": mainFIle,
					"config":   cfg,
				})
				m := filepath.Join(e, "index.html")
				err = ioutil.WriteFile(m, buf.Bytes(), 0600)
				if err != nil {
					return err
				}
				query := make(url.Values)
				query.Set("src", filepath.Join(info.RelativePath, name, "index.html"))
				cfg.IntegrationIndexPages = append(cfg.IntegrationIndexPages,
					fmt.Sprintf("%s:%d%s?%s",
						localhost, cfg.Port, resourcePath, query.Encode()))
				if cfg.Build {
					err = buildPackage(ctx, filepath.Join(e, "main.js"), pkg)
					if err != nil {
						return err
					}
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
	o := cfg.OutputPath
	var buf bytes.Buffer
	err := indexHTMLTpl.Execute(&buf, ctx)
	m := filepath.Join(o, "index.html")
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
	{{range $k,$v:=.config.TestNames}}
	"{{$k.ImportPath}}"
	{{- end}}
	"{{.wsImport}}"
	"{{.madImport}}"
	{{if .config.Cover}}
	"github.com/gernest/mad/cover"
	{{end}}
)

func main()  {
	startTest()
}

const testID ="{{.config.UUID}}"
const testPkg ="{{.config.Info.ImportPath}}"

func startTest(){
	go func ()  {
	{{if .config.Cover}}
	defer func ()  {
		println(cover.Key+cover.JSON())
	}()
	{{end}}
	 w,err:=ws.New(testID)
	 if err!=nil{
		 panic(err)
	 }
	 defer w.Close()
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
func allTests()[]mad.Test  {
	return []mad.Test{
		{{- range $k, $v:=.config.TestNames}}
		{{- range $v.Unit}}
		mad.Describe("{{$k.Desc .}}",{{$k.FormatName .}}()),
		{{- end}}
		{{- end}}
	}
}
`

var mainIntegrationTpl = `package main
import (
	"{{.info.ImportPath}}"
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
			Cover: {{.config.Cover}},
			Component: testFunc().(*mad.Component),
		},
	)
}
func testFunc() mad.Integration {
	return mad.SetupIntegration("{{.info.Desc .FuncName}}",{{.info.FormatName .FuncName}}() )
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

func buildPackage(ctx context.Context, out, pkg string) error {
	cmd := exec.CommandContext(ctx, "gopherjs", "build", "-o", out, pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	return cmd.Run()
}
