package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"sort"

	"github.com/gernest/gs/ciu/browsers"
	"github.com/urfave/cli"

	"github.com/gernest/gs/cmd/ciu/base62"
)

func browserVersions(w io.Writer, agents map[string]Agent) (map[string]string, error) {
	var keys []string
	for k := range agents {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	type bstat struct {
		version string
		count   int
	}
	var m []*bstat
	var exist = func(ver string) bool {
		for _, v := range m {
			if v.version == ver {
				v.count++
				return true
			}
		}
		return false
	}
	for _, v := range keys {
		var vers []string
		for ver := range agents[v].UsageGlobal {
			vers = append(vers, ver)
		}
		sort.Strings(vers)
		for _, ver := range vers {
			if !exist(ver) {
				m = append(m, &bstat{
					ver, 1,
				})
			}
		}
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i].count < m[j].count
	})
	out := make(map[string]string)
	for i, v := range m {
		out[base62.Encode(int64(i))] = v.version
	}
	tpl, err := template.New("browserList").Parse(browserVersionsTpl)
	if err != nil {
		return nil, err
	}
	err = tpl.Execute(w, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

const browserVersionsTpl = `package version

func New()map[string]string  {
	return map[string]string{
		{{- range $k,$v:=.}}
		"{{$k}}": "{{$v}}",
		{{- end}}
	}
}
`

func invert(m map[string]string) map[string]string {
	o := make(map[string]string)
	for k, v := range m {
		o[v] = k
	}
	return o
}

type agetntOptions struct {
	agents, full map[string]Agent
	versions     map[string]string
}

func agentsCmd(w io.Writer, opts agetntOptions) error {
	versionsInverted := invert(opts.versions)
	b := invert(browsers.New())
	var keys []string
	for k := range opts.agents {
		keys = append(keys, k)
	}
	type agetObj struct {
		A map[string]float64
		B string
		C []string
		D map[string]string
		E string
		F map[string]int64
	}
	m := make(map[string]agetObj)
	for _, v := range keys {
		a := opts.agents[v]
		av := make(map[string]float64)
		for gk, gv := range a.UsageGlobal {
			av[versionsInverted[gk]] = gv
		}
		mf := make(map[string]int64)
		for _, item := range opts.full[v].VersionList {
			mf[versionsInverted[item.Version]] = item.ReleaseData
		}
		o := agetObj{
			A: av,
			B: a.Prefix,
			C: a.Versions,
			E: a.Browser,
			F: mf,
		}
		if a.DataPrefixEceptions != nil {
			em := make(map[string]string)
			for ek, ev := range a.DataPrefixEceptions {
				em[versionsInverted[ek]] = ev
			}
			o.D = em
		}
		m[b[v]] = o
	}
	tpl, err := template.New("agents").Parse(agentsTpl)
	if err != nil {
		return err
	}
	return tpl.Execute(w, m)
}

const agentsTpl = `package agents

type Agent struct {
	A map[string]float64
	B string
	C []string
	D map[string]string
	E string
	F map[string]int64
}

func New()map[string]Agent {
	return map[string]Agent{
		{{- range $ak,$av:=.}}
		"{{$ak}}": Agent{
			{{- with $av}}
				{{- with .A}}
				A: map[string]float64{
					{{- range $k,$v:=.}}
					"{{$k}}":{{$v}},
					{{- end}}
				},
				{{- end}}	
				{{- with .B}}
				B: "{{.}}",
				{{- end}}
				{{- with .C}}
				C: []string{
					{{- range $k,$v:=. -}}
					"{{$v}}",
					{{- end}}
				},
				{{- end}}	
				{{- with .D}}
				D: map[string]string{
					{{- range $k,$v:=.}}
					"{{$k}}":"{{$v}}",
					{{- end}}
				},
				{{- end}}	
				{{- with .E}}
				E: "{{.}}",
				{{- end}}
				{{- with .F}}
				F: map[string]int64{
					{{- range $k,$v:=.}}
					"{{$k}}":{{$v}},
					{{- end}}
				},
				{{- end}}	
			{{- end}}
		},
		{{- end}}
	}
}

`

func AgentCMD(ctx *cli.Context) error {
	b, err := ioutil.ReadFile(ctx.String("data"))
	if err != nil {
		return err
	}
	data := &Data{}
	err = json.Unmarshal(b, data)
	if err != nil {
		return err
	}
	b, err = ioutil.ReadFile(ctx.String("full"))
	if err != nil {
		return err
	}
	full := &Data{}
	err = json.Unmarshal(b, full)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	v, err := browserVersions(&buf, data.Agents)
	if err != nil {
		return err
	}
	b, err = format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ctx.String("list-file"), b, 0600)
	if err != nil {
		return err
	}
	buf.Reset()
	err = agentsCmd(&buf, agetntOptions{
		agents:   data.Agents,
		versions: v,
		full:     full.Agents,
	})
	if err != nil {
		return err
	}
	b, err = format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ctx.String("agents-file"), b, 0600)
}
