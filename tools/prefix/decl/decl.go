package decl

import (
	"github.com/gernest/vected/lib/gs"
)

var _ Declaration = Decl{}

type Decl struct{}

func (d Decl) Check(gs.CSSRule) bool {
	return true
}

func (d Decl) Prefixed(prop, prefix string) string {
	return prefix + prop
}

func (d Decl) Normalize(prop string) string {
	return prop
}

// Set returns a gs.CSSRule with  the prefix. For gs.SimpleRule the prefix is
// added to the Key, and for StyleRule the prefix is added to the selector.
//
// This calls Decl.Prefixed to create the new value that is assigned to Key(for
// gs.SimpleRule) or Selector(for gs.StyleRule).
func (d Decl) Set(rule gs.CSSRule, prefix string) gs.CSSRule {
	switch e := rule.(type) {
	case gs.SimpleRule:
		e.Key = d.Prefixed(e.Key, prefix)
		return e
	case gs.StyleRule:
		e.Selector = d.Prefixed(e.Selector, prefix)
		return e
	default:
		return e
	}
}

// Insert returns a list of prefixed rules, the supplied rule is added as the
// last item of the list.
func (d Decl) Insert(rule gs.CSSRule, prefixes ...string) gs.CSSRule {
	if len(prefixes) > 0 {
		var list gs.RuleList
		for _, k := range prefixes {
			list = append(list, d.Set(rule, k))
		}
		list = append(list, rule)
		return list
	}
	return rule
}

func UpdateValue(rule gs.CSSRule, value string) gs.CSSRule {
	if value == "" {
		return rule
	}
	if e, ok := rule.(gs.SimpleRule); ok {
		e.Value = value
		return e
	}
	return rule
}

type Declaration interface {
	Check(gs.CSSRule) bool
	Prefixed(prop, prefix string) string
	Normalize(prop string) string
	Set(rule gs.CSSRule, prefix string) gs.CSSRule
}
