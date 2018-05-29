package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

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
	println(cfg.Info.SrcRoot)
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
	_, err = coverToXCover(cfg, c)
	return err
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
			Mode:     "set",
		}
		for _, block := range v.Blocks {
			p.Blocks = append(p.Blocks, xcover.ProfileBlock{
				StartLine: block.StartPosition.Line,
				StartCol:  block.StartPosition.Column,
				// EndLine:   block.EndPosition.Line,
				// EndCol:    block.EndPosition.Column,
				NumStmt: block.NumStmt,
				Count:   block.Count,
			})
		}
		profiles = append(profiles, p)
	}
	return
}
