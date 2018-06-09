package main

import (
	"sort"

	"github.com/gernest/gs/ciu/browsers"

	"github.com/gernest/gs/cmd/ciu/base62"
)

func browserVersions(agents map[string]Agent) map[string]string {
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
	return out
}

func invert(m map[string]string) map[string]string {
	o := make(map[string]string)
	for k, v := range m {
		o[v] = k
	}
	return o
}

func agentsCmd(agents map[string]Agent, versions map[string]string, full map[string]Agent) {
	versionsInverted := invert(versions)
	b := invert(browsers.New())
	var keys []string
	for k := range agents {
		keys = append(keys, k)
	}
	type agetObj struct {
		A map[string]float64
		B string
		C []string
		E string
		F map[string]int64
		D map[string]string
	}
	m := make(map[string]agetObj)
	for _, v := range keys {
		a := agents[v]
		av := make(map[string]float64)
		for gk, gv := range a.UsageGlobal {
			av[versionsInverted[gk]] = gv
		}
		mf := make(map[string]int64)
		for _, item := range full[v].VersionList {
			mf[versionsInverted[item.Version]] = item.ReleaseData
		}
		m[b[v]] = agetObj{
			A: av,
			B: a.Prefix,
			C: a.Versions,
			E: a.Browser,
			F: mf,
		}
	}
}
