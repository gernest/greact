package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	a := cli.NewApp()
	a.Commands = []cli.Command{
		{
			Name:   "browser",
			Action: BrowserCMD,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "o",
					Value: "ciu/browser/browser.go",
				},
			},
		},
	}
	if err := a.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
