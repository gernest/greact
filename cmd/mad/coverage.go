package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

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
	printCoverage(p)
	return nil
}

func printCoverage(p []cover.Profile) {
	m := make(map[string][]cover.Profile)
	var keys []string
	for _, v := range p {
		b := filepath.Base(v.FileName)
		pkgName := strings.TrimSuffix(v.FileName, b)
		pkgName = strings.TrimSuffix(pkgName, "/")
		if pkg, ok := m[pkgName]; ok {
			pkg = append(pkg, v)
			m[pkgName] = pkg
			continue
		}
		keys = append(keys, pkgName)
		m[pkgName] = []cover.Profile{v}
	}
	sort.Strings(keys)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
	for _, v := range keys {
		c := cover.Calc(m[v])
		fmt.Fprintf(w, "%s\t %.1f%%\n", v, 100*c)
	}
	v := cover.Calc(p)
	fmt.Fprintf(w, "Total\t %.1f%%\n", 100*v)
	w.Flush()
}

func mergeProfiles(a, b []cover.Profile) []cover.Profile {
	cache := make(map[string]cover.Profile)
	for _, v := range a {
		cache[v.FileName] = v
	}
	for _, v := range b {
		if c, ok := cache[v.FileName]; ok {
			for key, value := range c.Blocks {
				c.Blocks[key].Count = c.Blocks[key].Count + value.Count
			}
			cache[v.FileName] = c
		} else {
			cache[v.FileName] = v
		}
	}
	var o []cover.Profile
	for _, v := range cache {
		o = append(o, v)
	}
	return o
}
