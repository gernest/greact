package config

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"time"

	"github.com/gernest/mad/tools"

	uuid "github.com/satori/go.uuid"
	"github.com/urfave/cli"
)

// Config contains configuration testails about the test running environment.
type Config struct {
	// Information about the package.
	Info *build.Package

	// URL where the test runner service is running.
	ServerURL string

	// This is absolute path to the root of the package being tested,
	Root string

	// The directory where test files stay.
	TestDirName string

	// Absolute path to the directory containing the tests
	TestPath string

	// This the absolute path of processed test directory.
	GeneratedTestPath string

	// This is import path for the processed test package
	// example   github.com/gernest/mad/madness/tests
	GeneratedTestPkg string

	// This is the name of the directory in which generated test files are saved.
	// Default is madness.
	OutputDirName string

	// Absolute path to the directory in which generated test files are save.
	OutputPath string

	// Import path pointing to the main package that will be compiled by gopherjs
	// example github.com/gernest/mad/madness
	OutputMainPkg string

	// WHen true it will compile the generated test packages with gopherjs The
	// default value is true.
	Build bool

	// This is a uuid v4 string which is generated for every test run. It is used
	// internally to collect test results through websocket.
	UUID string

	TestNames map[*Info]*tools.TestNames

	// Port is the port on which to run the websocket server.
	Port int

	// if true tells the runner to generate coverage profile.
	Cover bool

	// the name of the file containing the generated coverage profile.
	Coverfile string

	// UnitIndexPage this is the url to the index.html page of the generated unit
	// test package.
	UnitIndexPage string

	// This is the list of urls of index.html pages of the generated integration
	// unit test.
	IntegrationIndexPages []string

	// When true, this will output a lot of text to stdout. Also it will print
	// console output from the test package.
	Verbose bool

	// Time to wait before stoping tests execution.
	Timeout time.Duration

	DevtoolURL  string
	DevtoolPort int

	Covermode string

	// If true this will only generate the packages and print out whet the test
	// runner will do without building or executing the tests.
	Dry bool

	TestInfo []*Info

	ImportMap map[string]string

	Browsers []Browser
}

type Browser struct {
	Name  string
	Port  int
	Flags []string
}

// FLags returns configuration flags.
func FLags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "server_url",
			EnvVar: "PEST_SERVER_URL",
			Value:  "http://localhost:1955",
		},
		cli.StringFlag{
			Name:  "root",
			Usage: "the root path of the package",
		},
		cli.StringFlag{
			Name:  "test_dir",
			Usage: "relative path to the tests directory",
			Value: "tests",
		},
		cli.StringFlag{
			Name:  "output_dir",
			Usage: "relative path to the generated tests directory",
			Value: "madness",
		},
		cli.BoolTFlag{
			Name: "build",
		},
		cli.BoolFlag{
			Name: "cover",
		},
		cli.StringFlag{
			Name: "coverprofile",
		},
		cli.StringFlag{
			Name:  "mode",
			Value: "set",
		},
		cli.IntFlag{
			Name:  "port",
			Value: 1956,
		},
		cli.BoolFlag{
			Name:  "v",
			Usage: "enables verbose output",
		},
		cli.DurationFlag{
			Name:  "timeout",
			Usage: "time before stoping test execution",
			Value: 30 * time.Second,
		},
		cli.IntFlag{
			Name:  "devtool-port",
			Value: 9222,
		},
		cli.StringFlag{
			Name:  "devtool-url",
			Value: "http://127.0.0.1",
		},
		cli.BoolFlag{
			Name: "dry",
		},
	}
}

// Load returns *Config instance with values populated from ctx.
func Load(ctx *cli.Context) (*Config, error) {
	c := &Config{
		ServerURL:     ctx.String("server_url"),
		Root:          ctx.String("root"),
		TestDirName:   ctx.String("test_dir"),
		OutputDirName: ctx.String("output_dir"),
		Build:         ctx.BoolT("build"),
		Cover:         ctx.Bool("cover"),
		Coverfile:     ctx.String("coverprofile"),
		Covermode:     ctx.String("mode"),
		Port:          ctx.Int("port"),
		Verbose:       ctx.Bool("v"),
		Timeout:       ctx.Duration("timeout"),
		DevtoolPort:   ctx.Int("devtool-port"),
		DevtoolURL:    ctx.String("devtool-url"),
		Dry:           ctx.Bool("dry"),
	}
	if !filepath.IsAbs(c.Root) {
		p, err := filepath.Abs(c.Root)
		if err != nil {
			return nil, err
		}
		c.Root = p
	}
	pkg, err := build.ImportDir(c.Root, 0)
	if err != nil {
		return nil, err
	}
	c.TestNames = make(map[*Info]*tools.TestNames)
	c.ImportMap = make(map[string]string)
	c.Info = pkg
	c.TestPath = filepath.Join(c.Info.Dir, c.TestDirName)
	c.OutputPath = filepath.Join(c.Info.Dir, c.OutputDirName)
	c.GeneratedTestPath = filepath.Join(c.OutputPath, c.TestDirName)
	c.GeneratedTestPkg = filepath.Join(c.Info.ImportPath, c.OutputDirName, c.TestDirName)
	c.OutputMainPkg = filepath.Join(c.Info.ImportPath, c.OutputDirName)
	c.UUID = uuid.NewV4().String()
	i, err := os.Stat(c.TestPath)
	if err != nil {
		return nil, err
	}
	if !i.IsDir() {
		return nil, fmt.Errorf("%s is not a  directory %v", c.TestPath, err)
	}
	var testDirs []string
	filepath.Walk(c.TestPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			testDirs = append(testDirs, path)
		}
		return nil
	})
	for _, v := range testDirs {
		i, err := OutputInfo(c, v)
		if err != nil {
			return nil, err
		}
		c.TestInfo = append(c.TestInfo, i)
	}
	c.Browsers = []Browser{
		{
			Name:  "chrome",
			Flags: []string{"--headless"},
			Port:  c.DevtoolPort,
		},
	}
	return c, nil
}

// GetOutDir returns absolute path to the directory where generated output
// stays.
func (c *Config) GetOutDir() string {
	return filepath.Join(c.Info.Dir, c.OutputDirName)
}

// GetTestDirName returns absolute path where the tests are.
func (c *Config) GetTestDirName() string {
	return filepath.Join(c.Info.Dir, c.TestDirName)
}

// Info contains information about a generated test package.
type Info struct {

	// This is the absolute path to the generated package.
	OutputPath string

	// Relative path to the root of generated directory. So for instance if the
	// generation directory is /madness, and the package was generated to
	// /madness/tests/pkg
	// then RelativePath value will be tests/pkg.
	RelativePath string

	Package *build.Package

	ImportPath string
}

func (i *Info) Desc(n string) string {
	return fmt.Sprintf("%s.%s", i.Package.Name, n)
}

func (i *Info) FormatName(n string) string {
	return i.Desc(n)
}

func OutputInfo(cfg *Config, testPath string) (*Info, error) {
	tsPkg, err := build.ImportDir(testPath, 0)
	if err != nil {
		return nil, err
	}
	i, err := getOutputInfo(cfg, testPath, tsPkg.Name)
	if err != nil {
		return nil, err
	}
	i.Package = tsPkg
	return i, nil
}

func getOutputInfo(cfg *Config, testPath string, packagename string) (*Info, error) {
	if cfg.TestPath == testPath {
		path := filepath.Join(cfg.OutputPath, packagename)
		return &Info{OutputPath: path,
			ImportPath:   cfg.GeneratedTestPkg,
			RelativePath: cfg.TestDirName,
		}, nil
	}
	rel, err := filepath.Rel(cfg.TestPath, testPath)
	if err != nil {
		return nil, err
	}
	path := filepath.Join(cfg.OutputPath, cfg.TestDirName, rel)
	relPath := filepath.Join(filepath.Base(cfg.TestPath), rel)
	return &Info{
		OutputPath:   path,
		RelativePath: relPath,
		ImportPath:   cfg.GeneratedTestPkg + "/" + rel,
	}, nil
}
