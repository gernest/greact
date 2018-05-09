package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gernest/alien"
	"github.com/gernest/prom/api"
	"github.com/gorilla/websocket"

	"github.com/takama/daemon"
	"github.com/urfave/cli"
)

const (
	serviceName  = "promnight"
	desc         = "Treat your vecty tests like your first date"
	port         = ":1955"
	testEndpoint = "test"
)

func startDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Start()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	} else {
		fmt.Println(msg)
	}
	return nil
}

func stopDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Stop()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	} else {
		fmt.Println(msg)
	}
	return nil
}

func installDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Install()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	} else {
		fmt.Println(msg)
	}
	return nil
}

func removeDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Remove()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	} else {
		fmt.Println(msg)
	}
	return nil
}

func statusDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Status()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	} else {
		fmt.Println(msg)
	}
	return nil
}

func daemonService(ctx *cli.Context) (err error) {
	host := ctx.String("host")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	rctx, cancel := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    port,
		Handler: apiServer(rctx, host),
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Printf("started test-server on %s\n", port)
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

var upgrader = websocket.Upgrader{}

func apiServer(ctx context.Context, host string) *alien.Mux {
	mux := alien.New()
	stats := &api.TestStats{}
	queue := make(chan *api.TestRequest, 50)
	cache := &sync.Map{}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case req := <-queue:
				ts := &api.TestSuite{
					Package: req.Package,
					Path:    req.Path,
				}
				ts.Status = "queued"
				cache.Store(ts.Package, ts)
				stats.Queued = append(stats.Queued, ts)
			}
		}

	}()
	// we only display the server's test stats on GET /
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(stats)
		fmt.Fprint(w, string(b))
	})

	// Accept test-running request on POST /. On success we return a json object
	// with websocket url to connect for the test running events.
	mux.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		req := &api.TestRequest{}
		err = json.Unmarshal(b, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res, err := homeResponse(host, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		queue <- req
		b, _ = json.Marshal(res)
		fmt.Fprint(w, string(b))
	})
	mux.Get("/"+testEndpoint, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		pkg := query.Get("pkg")
		pkg, err := url.QueryUnescape(pkg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tsv, ok := cache.Load(pkg)
		if !ok {
			http.Error(w, "package not found", http.StatusNotFound)
			return
		}
		ts := tsv.(*api.TestSuite)
		ts.Status = "running"

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		go func() {
			for {
				select {
				case <-ctx.Done():
					break
				default:
					if _, _, err := conn.NextReader(); err != nil {
						conn.Close()
						break
					}
				}
			}
		}()
		for {
			err := conn.WriteJSON(tsv)
			if err != nil {
				//log error
				fmt.Println(err)
				break
			}
		}
	})
	return mux
}

func homeResponse(base string, req *api.TestRequest) (*api.TestResponse, error) {
	u, err := websocketURL(base, req.Package)
	if err != nil {
		return nil, err
	}
	return &api.TestResponse{WebsocketURL: u}, nil
}

func websocketURL(base string, pkg string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	query := make(url.Values)
	query.Set("pkg", pkg)
	u.Path = testEndpoint
	u.Scheme = "ws"
	u.RawQuery = query.Encode()
	return u.String(), nil
}
