package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/gernest/mad"
	"github.com/gorilla/websocket"

	"github.com/gernest/alien"
	"github.com/gernest/mad/config"
)

var upgrade = websocket.Upgrader{}

func newServer(ctx context.Context, cfg *config.Config) <-chan *mad.SpecResult {
	mux := alien.New()
	out := make(chan *mad.SpecResult)
	mux.Get("/"+cfg.UUID, func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		for {
			select {
			case <-ctx.Done():
				break
			default:
				typ, data, err := conn.ReadMessage()
				if err != nil {
					if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
						fmt.Printf(" reading response %s\n", err)
					}
					return
				}
				switch typ {
				case websocket.CloseMessage:
					return
				case websocket.TextMessage, websocket.BinaryMessage:
					o := &mad.SpecResult{}
					err := json.Unmarshal(data, o)
					if err != nil {
						fmt.Println(err)
						continue
					}
					out <- o
				}
			}
		}
	})
	mux.Get(resourcePath, func(w http.ResponseWriter, r *http.Request) {
		src := r.URL.Query().Get("src")
		if src == "" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		sc, err := url.PathUnescape(src)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sc = filepath.Clean(sc)
		if filepath.IsAbs(sc) {
			http.Error(w, "only relative src are allowed", http.StatusBadRequest)
			return
		}
		path := filepath.Join(cfg.GetOutDir(), sc)
		http.ServeFile(w, r, path)
	})
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				s.Close()
			}
		}
	}()
	go func() {
		if cfg.Verbose {
			fmt.Printf("starting websocket server at :%d\n", cfg.Port)
		}
		err := s.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				fmt.Printf("exit websocket server with error %s\n", err)
			}
		}
	}()
	return out
}
