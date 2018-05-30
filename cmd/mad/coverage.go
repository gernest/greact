package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gernest/mad/config"
	"github.com/gernest/mad/cover"
	"github.com/urfave/cli"
)

func runCoverage(ctx *cli.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		return err
	}
	f := filepath.Join(cfg.OutputPath, cfg.Coverfile)
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	p := []cover.Profile{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		return err
	}
	v := cover.Calc(p)
	fmt.Printf("coverage: %.1f%%\n", 100*v)
	return nil
}
