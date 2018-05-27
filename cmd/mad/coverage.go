package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/gernest/mad/config"
	"github.com/gernest/mad/cover"
	"github.com/mafredri/cdp/protocol/profiler"
	"github.com/urfave/cli"
)

func runCoverage(ctx *cli.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		return err
	}
	coverageFile := filepath.Join(cfg.OutputPath, cfg.Coverfile)
	b, err := ioutil.ReadFile(coverageFile)
	if err != nil {
		return err
	}
	p := []profiler.Profile{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		return err
	}
	return cover.Process(cfg, p)
}
