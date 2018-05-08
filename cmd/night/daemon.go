package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gernest/alien"

	"github.com/takama/daemon"
	"github.com/urfave/cli"
)

const (
	serviceName = "promnight"
	desc        = "Treat your vecty tests like your first date"
	port        = ":1955"
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

func daemonService(ctx *cli.Context) (err error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	rctx, cancel := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    port,
		Handler: api(rctx),
	}
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		err = server.ListenAndServe()
		wg.Done()
	}()
	go func() {
		for {
			select {
			case <-rctx.Done():
				err = server.Shutdown(rctx)
			case <-interrupt:
				cancel()
			}
		}
	}()
	wg.Wait()
	return
}

func api(ctx context.Context) *alien.Mux {
	return nil
}
