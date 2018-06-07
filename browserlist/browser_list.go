package browserlist

import "strings"

type filter func(name, version string) bool

func query(str string) filter {
	str = strings.TrimSpace(str)
	switch str[0] {
	case '<':
		parts := strings.Split(str, " ")
		ver := strings.TrimSpace(parts[1])
		if len(parts[0]) == 2 {
			if parts[0][1] == '=' {
				return func(name, version string) bool {
					return version <= ver
				}
			}
			return noop
		}
	case '>':
		parts := strings.Split(str, " ")
		ver := strings.TrimSpace(parts[1])
		if len(parts[0]) == 2 {
			if parts[0][1] == '=' {
				return func(name, version string) bool {
					return version >= ver
				}
			}
			return noop
		}
	}
	parts := strings.Split(str, " ")
	switch parts[0] {
	case "cover":
	default:
		if n, ok := aliasReverse[parts[0]]; ok {
			return func(name, version string) bool {
				return name == n
			}
		}
	}
	return noop
}

func noop(_, _ string) bool {
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
