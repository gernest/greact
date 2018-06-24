package prefix

import (
	"regexp"

	"github.com/gernest/gs"
)

var re = regexp.MustCompile(`^(-\w+-)`)
var reSel = regexp.MustCompile(`:(-\w+-)`)

// Prefix returns a vendor prefix extracted from s.
//
// example -moz-tab-size => -moz-
func Prefix(s string) string {
	sub := re.FindStringSubmatch(s)
	if sub != nil {
		return sub[0]
	}
	return ""
}

// UnPrefix returns the input string stripped of its vendor prefix.
//
//-moz-tab-size => tab-size
func UnPrefix(s string) string {
	return string(re.ReplaceAll([]byte(s), []byte("")))
}

// RulePrefixer this is an interface for processing css rules by adding the
// prefix.
type RulePrefixer interface {
	Name() []string
}

type Normal interface {
	Normalize() string
}

type Setter interface {
	Set(gs.CSSRule) gs.CSSRule
}

// FindPrefix returns vendor  prefix for the given css rule.
func FindPrefix(b *Browser, rule gs.CSSRule) string {
	var prefix string
	switch e := rule.(type) {
	case gs.SimpleRule:
		prefix = Prefix(e.Key)
	case gs.StyleRule:
		sub := reSel.FindStringSubmatch(e.Selector)
		if sub != nil {
			prefix = sub[1]
		}
	case gs.Conditional:
		if len(e.Key) > 1 && e.Key[1] == '-' {
			prefix = Prefix(e.Key[1:])
		}
	}
	if prefix != "" {
		if b.WithPrefix(prefix) {
			return prefix
		}
	}
	return ""
}
