package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli"
)

// Serve defines serve command which builds and serves wasm modules.
func Serve() cli.Command {
	return cli.Command{
		Name:   "serve",
		Usage:  "builds and starts web server to serve wasm modules",
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
	out := filepath.Join(a, "main.wasm")
	cmd := exec.Command("go", "build", "-o", out, a)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GOARCH=wasm")
	cmd.Env = append(cmd.Env, "GOOS=js")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
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
			v, err := httputil.DumpRequest(r, true)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			} else {
				fmt.Println(string(v))
				w.Write(v)
			}
		}
	})
	msg := fmt.Sprint("serving main.wasm from http://localhost:8099")
	fmt.Println(msg)
	return http.ListenAndServe(":8099", h)
}
