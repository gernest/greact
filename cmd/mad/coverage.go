package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gernest/mad/config"
	"github.com/gernest/mad/cover"
	"github.com/urfave/cli"
	xcover "golang.org/x/tools/cover"
)

func runCoverage(ctx *cli.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		return err
	}
	file := filepath.Join(cfg.OutputPath, cfg.Coverfile)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	var c []*cover.Profile
	err = json.Unmarshal(b, &c)
	if err != nil {
		return err
	}
	cov, err := coverToXCover(cfg, c)
	if err != nil {
		return err
	}
	ext := filepath.Ext(cfg.Coverfile)
	name := strings.TrimSuffix(cfg.Coverfile, ext)
	pprofName := filepath.Join(name + ".pprof" + ext)
	data, err := json.Marshal(cov)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(cfg.OutputPath, pprofName), data, 0600)
}

func coverToXCover(cfg *config.Config, c []*cover.Profile) (profiles []xcover.Profile, err error) {
	root := cfg.Info.SrcRoot
	for _, v := range c {

		name, err := filepath.Rel(root, v.FileName)
		if err != nil {
			return nil, err
		}
		p := xcover.Profile{
			FileName: name,
		}
		for _, block := range v.Blocks {
			p.Blocks = append(p.Blocks, xcover.ProfileBlock{
				StartLine: block.StartPosition.Line,
				StartCol:  block.StartPosition.Column,
				EndLine:   block.EndPosition.Line,
				EndCol:    block.EndPosition.Column,
				NumStmt:   block.NumStmt,
				Count:     block.Count,
			})
		}
		profiles = append(profiles, p)
	}
	return
}
