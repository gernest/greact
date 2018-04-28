package prefix

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gernest/gs/agents"
)

var re = regexp.MustCompile(`^(-\w+-)`)

func Prefix(s string) string {
	sub := re.FindSubmatch([]byte(s))
	if sub != nil {
		return string(sub[0])
	}
	return ""
}

func UnPrefix(s string) string {
	return string(re.ReplaceAll([]byte(s), []byte("")))
}

type Broswer struct {
	Prefixcache  []string
	prefixRegexp *regexp.Regexp
	Selected     []string
}

func NewBrowser(filter ...func(name, version string) bool) *Broswer {
	b := &Broswer{}
	re := ""
	all := agents.All()
	for k, a := range all {
		pre := "-" + a.Prefix + "-"
		if k == 0 {
			re = pre
		} else {
			re += "|" + pre
		}
		b.Prefixcache = append(b.Prefixcache, pre)
	}
	b.Prefixcache = uniq(b.Prefixcache)
	sort.Strings(b.Prefixcache)
	reg := regexp.MustCompile(re)
	b.prefixRegexp = reg
	if len(filter) > 0 {
		f := filter[0]
		for _, v := range all {
			for _, version := range v.Versions {
				if f(v.Name, version) {
					b.Selected = append(b.Selected, v.Name+" "+version)
				}
			}
		}
		if len(b.Selected) > 0 {
			sortBrowsers(b.Selected)
		}
	}
	return b
}

func uniq(s []string) []string {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}
	var o []string
	for k := range m {
		o = append(o, k)
	}
	return o
}

func splitBrowser(b string) (string, float64) {
	parts := strings.Split(b, " ")
	k := parts[0]
	if len(parts) == 1 || parts[1] == "" {
		return k, 0
	}
	v, _ := strconv.ParseFloat(parts[1], 64)
	return k, v
}

func sortBrowsers(v []string) {
	sort.SliceStable(v, func(i, j int) bool {
		a, av := splitBrowser(v[i])
		b, bv := splitBrowser(v[j])
		return a > b || math.Signbit(av-bv)
	})
}

func (b *Broswer) WithPrefix(value string) bool {
	return b.prefixRegexp.Match([]byte(value))
}

func (b *Broswer) Prefix(name string) string {
	p := strings.Split(name, " ")
	name, version := p[0], p[1]
	a := agents.AgentsMap[name]
	prefix := ""
	if a.DataPrefixEceptions != nil {
		prefix = a.DataPrefixEceptions[version]
	}
	if prefix == "" {
		prefix = a.Prefix
	}
	return "-" + prefix + "-"
}

func (b *Broswer) IsSelected(name string) bool {
	return sort.SearchStrings(b.Selected, name) != len(b.Selected)
}

type Prefixes struct {
}
