package browserlist

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/gernest/gs/agents"
)

type filter func(name string, version version, usage float64) bool

func query(str string) filter {
	str = strings.TrimSpace(str)
	if str == "" {
		return noop
	}
	s := bufio.NewScanner(strings.NewReader(str))
	s.Split(bufio.ScanWords)
	if s.Scan() {
		x := s.Text()
		switch x {
		case ">":
			if s.Scan() {
				txt := s.Text()
				if strings.HasSuffix(txt, "%") {
					n := txt[:len(txt)-1]
					v, err := strconv.ParseFloat(n, 64)
					if err != nil {
						panic(err)
					}
					nv := v * 0.01
					return func(name string, ver version, usage float64) bool {
						return usage > nv
					}
				}
			}
			return noop
		case ">=":
			if s.Scan() {
				txt := s.Text()
				if strings.HasSuffix(txt, "%") {
					n := txt[:len(txt)-1]
					v, err := strconv.ParseFloat(n, 64)
					if err != nil {
						panic(err)
					}
					nv := v * 0.01

					return func(name string, ver version, usage float64) bool {
						return usage >= nv
					}
				}
			}
			return noop
		case "<":
			if s.Scan() {
				txt := s.Text()
				if strings.HasSuffix(txt, "%") {
					n := txt[:len(txt)-1]
					v, err := strconv.ParseFloat(n, 64)
					if err != nil {
						panic(err)
					}
					nv := v * 0.01
					return func(name string, ver version, usage float64) bool {
						return usage < nv
					}
				}
			}
			return noop
		case "<=":
			if s.Scan() {
				txt := s.Text()
				if strings.HasSuffix(txt, "%") {
					n := txt[:len(txt)-1]
					v, err := strconv.ParseFloat(n, 64)
					if err != nil {
						panic(err)
					}
					nv := v * 0.01
					return func(name string, ver version, usage float64) bool {
						return usage <= nv
					}
				}
			}
			return noop
		case "cover":
			if s.Scan() {
				txt := s.Text()
				if strings.HasSuffix(txt, "%") {
					n := txt[:len(txt)-1]
					v, err := strconv.ParseFloat(n, 64)
					if err != nil {
						panic(err)
					}
					nv := v * 0.01
					return func(name string, ver version, usage float64) bool {
						return usage >= nv
					}
				}
			}
			return noop
		}
	}

	return noop
}

func noop(_ string, _ version, _ float64) bool {
	return false
}

var browserAlias = map[string]string{
	"and_chr": "ChromeForAndroid",
	"and_ff":  "FirefoxForAndroid",
	"and_qq":  "QQForAndroid",
	"and_uc":  "UCForAndroid",
	"android": "Android",
	"baidu":   "Baidu",
	"bb":      "BlackBerry",
	"chrome":  "Chrome",
	"edge":    "Edge",
	"firefox": "Firefox",
	"ie":      "InternetExplorer",
	"ie_mob":  "InternetExplorerMobile",
	"ios_saf": "IOSSafari",
	"op_mini": "OperaMini",
	"op_mob":  "OperaMobile",
	"opera":   "Opera",
	"safari":  "Safari",
	"samsung": "Samsung",
}

var aliasReverse map[string]string

func init() {
	aliasReverse = make(map[string]string)
	for k, v := range browserAlias {
		aliasReverse[strings.ToLower(v)] = k
	}
}

func not(f filter) filter {
	return func(name string, version version, usage float64) bool {
		return !f(name, version, usage)
	}
}

type version string

func (v version) eq(v2 string) bool {
	return string(v) == v2
}

func (v version) gt(v2 string) bool {
	return v.filter(v2, func(a, b int) bool {
		return a > b
	})
}

func (v version) lt(v2 string) bool {
	return v.filter(v2, func(a, b int) bool {
		return a < b
	})
}

func (v version) ge(v2 string) bool {
	return v.filter(v2, func(a, b int) bool {
		return a >= b
	})
}

func (v version) le(v2 string) bool {
	return v.filter(v2, func(a, b int) bool {
		return a <= b
	})
}

func (v version) filter(v2 string, fn func(a, b int) bool) bool {
	s := string(v)
	if s == "" {
		return false
	}
	p1 := strings.Split(s, ".")
	p2 := strings.Split(v2, ".")
	n := len(p1)
	if len(p2) < n {
		n = len(p2)
	}
	for i := 0; i < n; i++ {
		a, err := strconv.Atoi(p1[i])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(p2[i])
		if err != nil {
			panic(err)
		}
		if fn(a, b) {
			return true
		}
	}
	return false
}

// Query returns a list of browsers that matches the given queries.
func Query(q ...string) []string {
	o := []string{}
	all := agents.All()
	for _, a := range all {
		o = append(o, apply(a, allFilterQuery(q...))...)
	}
	return o
}

func allFilterQuery(q ...string) filter {
	var f []filter
	for _, v := range q {
		f = append(f, query(v))
	}
	return allFilter(f...)
}

func allFilter(f ...filter) filter {
	return func(name string, v version, usage float64) bool {
		for _, fn := range f {
			if !fn(name, v, usage) {
				return false
			}
		}
		return true
	}
}

func apply(a agents.Agent, fn filter) []string {
	o := []string{}
	for _, v := range a.Versions {
		if fn(a.Name, version(v), a.UsageGlobal[v]) {
			o = append(o, a.Name+" "+v)
		}
	}
	return o
}
