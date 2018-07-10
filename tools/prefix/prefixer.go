package prefix

import (
	"github.com/gernest/vected/lib/gs"
)

func addPrefix(b *Browser, prefixes []string, rule gs.CSSRule) gs.CSSRule {
	pre := FindPrefix(b, rule)
	p := filter(prefixes, func(v string) bool {
		return v != pre
	})
	return add(p, rule)
}

func add(prefix []string, rule gs.CSSRule) gs.CSSRule {
	return nil
}
