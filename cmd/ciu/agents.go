package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"sort"

	"github.com/urfave/cli"
)

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
	// versionsInverted := invert(opts.versions)
	// b := invert(browsers.New())
	var keys []string
	for k := range opts.agents {
		keys = append(keys, k)
	}
	sort.Strings(keys)
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
			av[gk] = gv
		}
		mf := make(map[string]int64)
		for _, item := range opts.full[v].VersionList {
			mf[item.Version] = item.ReleaseData
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
				em[ek] = ev
			}
			o.D = em
		}
		m[v] = o
	}
	tpl, err := template.New("agents").Parse(agentsTpl)
	if err != nil {
		return err
	}
	return tpl.Execute(w, map[string]interface{}{
		"agents": m,
		"keys":   keys,
	})
}

const agentsTpl = `package agents

type Agent struct {
	UsageGlobal map[string]float64
	Prefix string
	Versions []string
	PrefixExceptions map[string]string
	Browser string
	ReleaseDate map[string]int64
}

// Keys returns a sorted slice of agents names.
func Keys()[]string  {
	return []string{
		{{- range $k,$v:= .keys -}}
		"{{$v}}",
		{{- end}}
	}
}

func New()map[string]Agent {
	return map[string]Agent{
		{{- range $ak,$av:=.agents}}
		"{{$ak}}": Agent{
			{{- with $av}}
				{{- with .A}}
				UsageGlobal: map[string]float64{
					{{- range $k,$v:=.}}
					"{{$k}}":{{$v}},
					{{- end}}
				},
				{{- end}}	
				{{- with .B}}
				Prefix: "{{.}}",
				{{- end}}
				{{- with .C}}
				Versions: []string{
					{{- range $k,$v:=. -}}
					"{{$v}}",
					{{- end}}
				},
				{{- end}}	
				{{- with .D}}
				PrefixExceptions: map[string]string{
					{{- range $k,$v:=.}}
					"{{$k}}":"{{$v}}",
					{{- end}}
				},
				{{- end}}	
				{{- with .E}}
				Browser: "{{.}}",
				{{- end}}
				{{- with .F}}
				ReleaseDate: map[string]int64{
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

	err = agentsCmd(&buf, agetntOptions{
		agents: data.Agents,
		full:   full.Agents,
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
