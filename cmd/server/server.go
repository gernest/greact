package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func Serve() cli.Command {
	return cli.Command{
		Name:   "serve",
		Usage:  "starts a web server that serves main.wasm file",
		Action: serve,
	}
}

func serve(ctx *cli.Context) error {
	a := ctx.Args().First()
	if a == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		a = wd
	}
	idx := "cmd/server/index.html"
	v, err := Asset(idx)
	if err != nil {
		return err
	}
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Write(v)
		case "/main.wasm":
			http.ServeFile(w, r, filepath.Join(a, "main.wasm"))
		default:
			http.Error(w, r.URL.Path, http.StatusNotFound)
		}
	})
	msg := fmt.Sprint("serving main.wasm from http://localhost:8099")
	fmt.Println(msg)
	return http.ListenAndServe(":8099", h)
}
