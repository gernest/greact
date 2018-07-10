package prefix

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gernest/gs/browserlist"
	"github.com/gernest/gs/ciu/agents"
)

type Browser struct {
	Prefixcache  []string
	PrefixRegexp *regexp.Regexp
	Selected     []string
	Data         map[string]agents.Agent
}

func NewBrowser(data map[string]agents.Agent, queries ...string) (*Browser, error) {
	if queries == nil {
		queries = []string{"defaults"}
	}
	selected, err := browserlist.Query(queries...)
	if err != nil {
		return nil, err
	}
	re := ""
	for _, a := range agents.New() {
		pre := "-" + a.Prefix + "-"
		if re != "" {
			re += "|" + pre
		} else {
			re = pre
		}
	}
	reg := regexp.MustCompile(re)
	var prefixcache []string
	for _, v := range agents.New() {
		prefixcache = append(prefixcache, "-"+v.Prefix+"-")
	}
	prefixcache = uniq(prefixcache)
	sort.Strings(prefixcache)
	return &Browser{
		Selected:     selected,
		Data:         data,
		PrefixRegexp: reg,
		Prefixcache:  prefixcache,
	}, nil
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

func (b *Browser) WithPrefix(value string) bool {
	return b.PrefixRegexp.Match([]byte(value))
}

func (b *Browser) Prefix(name string) string {
	p := strings.Split(name, " ")
	name, version := p[0], p[1]
	d := b.Data[name]
	prefix := ""
	if d.PrefixExceptions != nil {
		prefix = d.PrefixExceptions[version]
	}
	if prefix == "" {
		prefix = d.Prefix
	}
	return "-" + prefix + "-"
}

func (b *Browser) Prefixes() []string {
	if b.Prefixcache != nil {
		return b.Prefixcache
	}
	for _, v := range agents.New() {
		b.Prefixcache = append(b.Prefixcache, "-"+v.Prefix+"-")
	}
	b.Prefixcache = uniq(b.Prefixcache)
	sort.Strings(b.Prefixcache)
	return b.Prefixcache
}

func (b *Browser) IsSelected(name string) bool {
	if b.Selected != nil {
		return sliceContains(b.Selected, name)
	}
	return false
}
