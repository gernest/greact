package main

import (
	"fmt"

	"github.com/takama/daemon"
	"github.com/urfave/cli"
)

const (
	serviceName = "promnight"
	desc        = "Treat your vecty tests like your first date"
)

func startDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	if msg, err := srv.Start(); err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	return nil
}

func stopDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	if msg, err := srv.Stop(); err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	return nil
}

func installDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	if msg, err := srv.Install(); err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	return nil
}

func removeDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	if msg, err := srv.Remove(); err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	return nil
}

func statusDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	if msg, err := srv.Status(); err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	return nil
}

func daemonService(ctx *cli.Context) error {
	return nil
}
