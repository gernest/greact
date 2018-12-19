package main

import (
	"fmt"
	"os"

	"github.com/gernest/greact/cmd/gen"
	"github.com/gernest/greact/cmd/server"
	"github.com/urfave/cli"
)

func main() {
	a := cli.NewApp()
	a.Name = "vected_gen"
	a.Usage = "provides various commands that generate code for vected project"
	a.Commands = []cli.Command{
		gen.AttrCMD(),
		gen.RenderCMD(),
		gen.ElementsCMD(),
		server.Serve(),
	}
	if err := a.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
