package main

import (
	"fmt"
	"os"

	"github.com/gernest/vected/cmd/gen"
	"github.com/urfave/cli"
)

func main() {
	a := cli.NewApp()
	a.Name = "vected_gen"
	a.Usage = "provides various commands that generate code for vected project"
	a.Commands = []cli.Command{
		gen.RuntimeCMD(),
		gen.AgentsCommand(),
		gen.DataCommand(),
		gen.AttrCMD(),
	}
	if err := a.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
