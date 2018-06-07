package browserlist

import (
	"strconv"
	"strings"
)

type filter func(name string, version version) bool

func query(str string) filter {
	str = strings.TrimSpace(str)
	switch str[0] {
	case '<':
		parts := strings.Split(str, " ")
		ver := strings.TrimSpace(parts[1])
		if len(parts[0]) == 2 {
			if parts[0][1] == '=' {
				return func(name string, v version) bool {
					return v.le(ver)
				}
			}
			return noop
		}
		return func(name string, v version) bool {
			return v.lt(ver)
		}
	case '>':
		parts := strings.Split(str, " ")
		ver := strings.TrimSpace(parts[1])
		if len(parts[0]) == 2 {
			if parts[0][1] == '=' {
				return func(name string, v version) bool {
					return v.ge(ver)
				}
			}
			return noop
		}
		return func(name string, v version) bool {
			return v.gt(ver)
		}
	}
	parts := strings.Split(str, " ")
	switch parts[0] {
	case "cover":
	default:
		if n, ok := aliasReverse[parts[0]]; ok {
			return func(name string, v version) bool {
				return name == n
			}
		}
	}
	return noop
}

func noop(_ string, _ version) bool {
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
	return func(name string, version version) bool {
		return !f(name, version)
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
