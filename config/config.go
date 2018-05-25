package config

import (
	"go/build"
	"path/filepath"

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

	OutputDirName string

	OutputPath string

	OutputMainPkg string

	Build bool

	UUID string

	UnitFuncs   []string
	Integration []string

	// Port is the port on which to run the websocket server.
	Port int

	// Browser details to use.
	Browsers BrowserList

	// if true tells the runner to generate coverage profile.
	Cover bool

	// the name of the file containing the generated coverage profile.
	Coverfile string

	Chrome *ChromeOptions
}

type ChromeOptions struct {
	Port  int
	Flags []string
}

type BrowserList struct {
	Browsers []*Browser
}

// Browser defines a browser that implements chrome debugging protocol.
type Browser struct {
	Name    string
	Options map[string]interface{}
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
		cli.BoolFlag{
			Name: "build",
		},
		cli.BoolFlag{
			Name: "cover",
		},
		cli.StringFlag{
			Name:  "coverfile",
			Value: "coverage.json",
		},
		cli.IntFlag{
			Name:  "port",
			Value: 1956,
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
		Build:         ctx.Bool("build"),
		Cover:         ctx.Bool("cover"),
		Coverfile:     ctx.String("coverfile"),
		Port:          ctx.Int("port"),
	}
	if !filepath.IsAbs(c.Root) {
		p, err := filepath.Abs(c.Root)
		if err != nil {
			return nil, err
		}
		c.Root = p
	}
	pkg, err := build.ImportDir(c.Root, build.FindOnly)
	if err != nil {
		return nil, err
	}
	c.Info = pkg
	c.TestPath = filepath.Join(c.Info.Dir, c.TestDirName)
	c.OutputPath = filepath.Join(c.Info.Dir, c.OutputDirName)
	c.GeneratedTestPath = filepath.Join(c.OutputPath, c.TestDirName)
	c.GeneratedTestPkg = filepath.Join(c.Info.ImportPath, c.OutputDirName, c.TestDirName)
	c.OutputMainPkg = filepath.Join(c.Info.ImportPath, c.OutputDirName)
	c.UUID = uuid.NewV4().String()
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
