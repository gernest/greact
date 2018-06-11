package browserlist

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gernest/gs/ciu/agents"
)

// func compare(sign string, ref string) filter {
// 	v := version(ref)
// 	return func(name string, ver version, usage float64) bool {
// 		if strings.HasSuffix(ref, "%") {
// 			n := ref[:len(ref)-1]
// 			v, err := strconv.ParseFloat(n, 64)
// 			if err != nil {
// 				panic(err)
// 			}
// 			nv := v * 0.01
// 			switch sign {
// 			case ">":
// 				fmt.Println(nv)
// 				return usage > nv
// 			case ">=":
// 				return usage >= nv
// 			case "<":
// 				return usage < nv
// 			case "<=":
// 				return usage <= nv
// 			case "==":
// 				return usage == nv
// 			default:
// 				return false
// 			}
// 		}
// 		switch sign {
// 		case ">":
// 			return ver.gt(v)
// 		case ">=":
// 			return ver.ge(v)
// 		case "<":
// 			return ver.lt(v)
// 		case "<=":
// 			return ver.le(v)
// 		case "==":
// 			return ver == v
// 		default:
// 			return false
// 		}
// 	}
// }

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
	filter func(map[string]data, []string) ([]string, error)
}

func allHandlers() []handler {
	return []handler{
		{
			match:  regexp.MustCompile(`^last\s+(\d+)\s+major versions?$`),
			filter: lastMajorVersions,
		},
		{
			match:  regexp.MustCompile(`^last\s+(\d+)\s+versions?$`),
			filter: lastVersions,
		},
		{
			match:  regexp.MustCompile(`^last\s+(\d+)\s+(\w+)\s+major versions?$`),
			filter: lastMajorVersionsName,
		},
		{
			match:  regexp.MustCompile(`^last\s+(\d+)\s+(\w+)\s+versions?$`),
			filter: lastVersionsName,
		},
		{
			match:  regexp.MustCompile(`^unreleased\s+versions$`),
			filter: unreleased,
		},
		{
			match:  regexp.MustCompile(`^unreleased\s+(\w+)\s+versions?$/`),
			filter: unreleasedName,
		},
	}
}
func unreleased(dataCtx map[string]data, v []string) ([]string, error) {
	var o []string
	for _, key := range agents.Keys() {
		d, ok := dataCtx[key]
		if !ok {
			continue
		}
		var vers []string
		for _, ver := range d.versions {
			for _, rel := range d.released {
				if rel == ver {
					continue
				}
			}
			vers = append(vers, ver)
		}
		o = append(o, mapNames(key, vers...)...)
	}
	return o, nil
}

func unreleasedName(dataCtx map[string]data, v []string) ([]string, error) {
	if len(v) != 1 {
		return []string{}, nil
	}
	name := v[0]
	d, ok := dataCtx[name]
	if !ok {
		return []string{}, nil
	}
	var vers []string
	for _, ver := range d.versions {
		for _, rel := range d.released {
			if rel == ver {
				continue
			}
		}
		vers = append(vers, ver)
	}
	return mapNames(name, vers...), nil
}

func lastMajorVersions(dataCtx map[string]data, v []string) ([]string, error) {
	var o []string
	ver := 1
	if len(v) == 1 {
		i, err := strconv.Atoi(v[0])
		if err != nil {
			return nil, err
		}
		ver = i
	}
	for _, k := range agents.Keys() {
		d, ok := dataCtx[k]
		if !ok {
			return []string{}, nil
		}
		i, err := getMajorVersions(d.released, ver)
		if err != nil {
			return nil, err
		}
		o = append(o, mapNames(d.name, i...)...)
	}
	return o, nil
}
func lastMajorVersionsName(dataCtx map[string]data, v []string) ([]string, error) {
	if len(v) != 2 {
		return []string{}, nil
	}
	ver, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, err
	}
	name := v[1]
	d, ok := dataCtx[name]
	if !ok {
		return []string{}, nil
	}
	i, err := getMajorVersions(d.released, ver)
	if err != nil {
		return nil, err
	}
	return mapNames(name, i...), nil
}

func lastVersions(dataCtx map[string]data, v []string) ([]string, error) {
	if len(v) != 1 {
		return []string{}, nil
	}
	ver, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, err
	}
	var o []string
	for _, k := range agents.Keys() {
		d, ok := dataCtx[k]
		if !ok {
			continue
		}
		if len(d.released) > ver {
			idx := len(d.released) - ver
			o = append(o, mapNames(d.name, d.released[idx:]...)...)
		} else {
			o = append(o, mapNames(d.name, d.released...)...)
		}
	}
	return o, nil
}

func lastVersionsName(dataCtx map[string]data, v []string) ([]string, error) {
	if len(v) != 2 {
		return []string{}, nil
	}
	ver, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, err
	}
	name := v[1]
	d, ok := dataCtx[name]
	if !ok {
		return []string{}, nil
	}
	if len(d.released) > ver {
		idx := len(d.released) - ver
		return mapNames(d.name, d.released[idx:]...), nil
	}
	return mapNames(d.name, d.released...), nil
}

func mapNames(base string, s ...string) (o []string) {
	for _, v := range s {
		o = append(o, base+" "+v)
	}
	return
}

func getMajorVersions(released []string, number int) ([]string, error) {
	if len(released) == 0 {
		return []string{}, nil
	}
	min := version(released[0]).getMajor() - number + 1
	var o []string
	for k, v := range released {
		if k == 0 {
			continue
		}
		if version(v).getMajor() > min {
			o = append(o, v)
		}
	}
	return o, nil
}

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

func Query(s ...string) ([]string, error) {
	return QueryWith(getData(), s...)
}

func QueryWith(dataCtx map[string]data, s ...string) ([]string, error) {
	h := allHandlers()
	var o []string
	for _, v := range s {
		v = strings.ToLower(v)
		for _, c := range h {
			if c.match.MatchString(v) {
				i, err := c.filter(dataCtx, c.match.FindStringSubmatch(v)[1:])
				if err != nil {
					return nil, err
				}
				o = append(o, i...)
			}
		}
	}
	return o, nil
}
