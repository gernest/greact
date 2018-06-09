package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"sort"

	"github.com/gernest/gs/cmd/ciu/base62"
	"github.com/urfave/cli"
)

func browser(w io.Writer, s map[string]map[string]string) error {
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	funcs := template.FuncMap{
		"encode": func(i int) string {
			return base62.Encode(int64(i))
		},
	}
	tpl, err := template.New("browsers").Funcs(funcs).Parse(browsersTpl)
	if err != nil {
		return err
	}
	return tpl.Execute(w, keys)
}

const browsersTpl = `package browsers

// New retuns a map of encoded browser keys to their values.
func New()map[string]string  {
	return map[string]string{
		{{- range $k,$v:=.}}
		"{{encode $k}}" :"{{$v}}",
		{{- end}}	
	}
}

`

func BrowserCMD(ctx *cli.Context) error {
	b, err := ioutil.ReadFile(ctx.String("data"))
	if err != nil {
		return err
	}
	data := &Data{}
	err = json.Unmarshal(b, data)
	if err != nil {
		return err
	}
	var fk []string
	for k := range data.Features {
		fk = append(fk, k)
	}
	if fk == nil {
		return nil
	}
	var buf bytes.Buffer
	err = browser(&buf, data.Features[fk[0]].Stats)
	if err != nil {
		return err
	}
	v, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ctx.String("o"), v, 0600)
}
