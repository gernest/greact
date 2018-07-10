package ws

import (
	"encoding/json"
	"net"
	"net/url"

	"github.com/gernest/vected/lib/mad"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket"
)

func New(id string) (*WS, error) {
	h := js.Global.Get("location").Get("href").String()
	u, err := url.Parse(h)
	if err != nil {
		return nil, err
	}
	u.Path = "/" + id
	u.Scheme = "ws"
	u.RawQuery = ""
	conn, err := websocket.Dial(u.String())
	if err != nil {
		return nil, err
	}
	w := &WS{Conn: conn}
	w.enc = json.NewEncoder(w)
	w.enc.SetIndent("", "  ")
	return w, nil
}

type WS struct {
	net.Conn
	enc *json.Encoder
}

func (w *WS) Report(ts mad.Test, pkg, id string) error {
	rs := toResult(ts, pkg, id)
	for _, v := range rs {
		if err := w.WriteResponse(v); err != nil {
			return err
		}
	}
	return nil
}

func (w *WS) WriteResponse(rs *mad.SpecResult) error {
	return w.enc.Encode(rs)
}

func toResult(rs mad.Test, pkg, id string) []*mad.SpecResult {
	var results []*mad.SpecResult
	switch e := rs.(type) {
	case *mad.Suite:
		e.ID = id
		e.Package = pkg
		results = append(results, e.Result())
	case mad.List:
		for _, v := range e {
			results = append(results, toResult(v, pkg, id)...)
		}
	}
	return results
}
