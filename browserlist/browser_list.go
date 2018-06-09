package browserlist

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gernest/gs/ciu/agents"
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
		switch {
		case signs[x]:
			if s.Scan() {
				return compare(x, s.Text())
			}
			return noop
		}
	}
	return noop
}

var signs = map[string]bool{
	">":  true,
	">=": true,
	"<":  true,
	"<=": true,
	"==": true,
}

func compare(sign string, ref string) filter {
	v := version(ref)
	return func(name string, ver version, usage float64) bool {
		if strings.HasSuffix(ref, "%") {
			n := ref[:len(ref)-1]
			v, err := strconv.ParseFloat(n, 64)
			if err != nil {
				panic(err)
			}
			nv := v * 0.01
			switch sign {
			case ">":
				fmt.Println(nv)
				return usage > nv
			case ">=":
				return usage >= nv
			case "<":
				return usage < nv
			case "<=":
				return usage <= nv
			case "==":
				return usage == nv
			default:
				return false
			}
		}
		switch sign {
		case ">":
			return ver.gt(v)
		case ">=":
			return ver.ge(v)
		case "<":
			return ver.lt(v)
		case "<=":
			return ver.le(v)
		case "==":
			return ver == v
		default:
			return false
		}
	}
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

func (v version) eq(v2 version) bool {
	return v == v2
}

func (v version) gt(v2 version) bool {
	m1 := v.getMajor()
	m2 := v2.getMajor()
	return m1 > m2
}

func (v version) lt(v2 version) bool {
	m1 := v.getMajor()
	m2 := v2.getMajor()
	return m1 < m2
}

func (v version) ge(v2 version) bool {
	m1 := v.getMajor()
	m2 := v2.getMajor()
	return m1 >= m2
}

func (v version) le(v2 version) bool {
	m1 := v.getMajor()
	m2 := v2.getMajor()
	return m1 <= m2
}

func (v version) getMajor() int {
	if string(v) == "" {
		return 0
	}
	p := strings.Split(string(v), ".")[0]
	b, err := strconv.Atoi(p)
	if err == nil {
		return b
	}
	return 0
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

func defaultQuery() []string {
	return []string{
		"> 0.5%",
		"last 2 versions",
		"Firefox ESR",
		"not dead",
	}
}

type handler struct {
	match  *regexp.Regexp
	filter func([]string) []string
}

var lastRegexp = regexp.MustCompile(`^last\s+(\d+)\s+major versions?$`)

type data struct {
	name     string
	versions []string
	released []string
}

func getData() map[string]data {
	m := make(map[string]data)
	for k, v := range agents.New() {
		ve := normalize(v.Versions...)
		d := data{
			name:     k,
			versions: ve,
		}
		if len(ve) > 2 {
			d.released = ve[len(ve)-2:]
		}
		m[k] = d
	}
	return m
}

func normalize(s ...string) []string {
	var o []string
	for _, v := range s {
		if v != "" {
			o = append(o, v)
		}
	}
	return o
}
