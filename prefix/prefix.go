package prefix

import "regexp"

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
