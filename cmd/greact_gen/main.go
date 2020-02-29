package main

import (
	"fmt"
	"os"

	"github.com/gernest/greact/cmd/internal/gen"
	"github.com/gernest/greact/cmd/internal/server"
	"github.com/urfave/cli"
)

func main() {
	a := cli.NewApp()
	a.Name = "greact_gen"
	a.Usage = "provides various commands that generate code for greact project"
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
