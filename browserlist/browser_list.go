package browserlist

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gernest/gs/ciu/agents"
)

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
			match:  regexp.MustCompile(`^unreleased\s+(\w+)\s+versions?$`),
			filter: unreleasedName,
		},
		{
			match:  regexp.MustCompile(`^last\s+(\d+)\s+years?$`),
			filter: lastYears,
		},
		{
			match:  regexp.MustCompile(`^since (\d+)(?:-(\d+))?(?:-(\d+))?$`),
			filter: sinceDate,
		},
		{
			match:  regexp.MustCompile(`^(>=?|<=?)\s*(\d*\.?\d+)%$`),
			filter: popularitySign,
		},
	}
}

func popularitySign(dataCtx map[string]data, v []string) ([]string, error) {
	sign := v[0]
	popularity, err := strconv.ParseFloat(v[1], 64)
	if err != nil {
		return nil, err
	}
	var o []string
	for _, name := range agents.Keys() {
		d, ok := dataCtx[name]
		if !ok {
			continue
		}
		var vers []string
		for rel, cov := range d.usage {
			switch sign {
			case ">":
				if cov > popularity {
					vers = append(vers, rel)
				}
			case "<":
				if cov < popularity {
					vers = append(vers, rel)
				}
			case ">=":
				if cov >= popularity {
					vers = append(vers, rel)
				}
			case "<=":
				if cov <= popularity {
					vers = append(vers, rel)
				}
			}
		}
		o = append(o, mapNames(name, vers...)...)
	}
	return o, nil
}

func lastYears(dataCtx map[string]data, v []string) ([]string, error) {
	i, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, err
	}
	now := time.Now()
	year, month, day := now.Date()
	year -= i
	n := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).UnixNano() * 1000
	return filterByYear(dataCtx, n)
}

func sinceDate(dataCtx map[string]data, v []string) ([]string, error) {
	year, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, err
	}
	month := 1
	if v[1] != "" {
		month, err = strconv.Atoi(v[1])
		if err != nil {
			return nil, err
		}
	}
	day := 1
	if v[2] != "" {
		day, err = strconv.Atoi(v[2])
		if err != nil {
			return nil, err
		}
	}
	n := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Unix()
	return filterByYear(dataCtx, n)
}

func filterByYear(dataCTx map[string]data, since int64) ([]string, error) {
	var o []string
	fmt.Println(since)
	for _, name := range agents.Keys() {
		d, ok := dataCTx[name]
		if !ok {
			continue
		}
		var vers []string
		for key, value := range d.releaseDate {
			if value >= since {
				vers = append(vers, key)
			}
		}
		sort.Strings(vers)
		o = append(o, mapNames(name, vers...)...)
	}
	return o, nil
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
	if v[0] != "" {
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
	name        string
	versions    []string
	released    []string
	releaseDate map[string]int64
	usage       map[string]float64
}

func getData() map[string]data {
	m := make(map[string]data)
	for k, v := range agents.New() {
		ve := normalize(v.Versions...)
		d := data{
			name:        k,
			versions:    ve,
			releaseDate: v.ReleaseDate,
			usage:       v.UsageGlobal,
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
