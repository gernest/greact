package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

type Feature struct {
	Stats map[string]map[string]string `json:"stats"`
}
type Agent struct {
	Name                string
	Browser             string
	Abbr                string
	Prefix              string
	Type                string
	UsageGlobal         map[string]float64
	Versions            []string
	DataPrefixEceptions map[string]string
}

type Data struct {
	Agents   map[string]Agent   `json:"agents"`
	Features map[string]Feature `json:"data"`
}

func main() {
	a := cli.NewApp()
	a.Commands = []cli.Command{
		{
			Name:   "browser",
			Action: BrowserCMD,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "o",
					Value: "ciu/browsers/browsers.go",
				},
			},
		},
	}
	if err := a.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
