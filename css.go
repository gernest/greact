package gs

import (
	"bytes"
	"strings"
)

type RuleList []CSSRule

func (RuleList) isRule() {}

func (ls RuleList) write(f func(string), opts ...Options) {
	for _, v := range ls {
		v.write(f, opts...)
	}
}

func CSS(rules ...CSSRule) CSSRule {
	return RuleList(rules)
}

type simple struct {
	key   string
	value string
}

func (simple) isRule() {}
func (s simple) write(f func(string), opts ...Options) {
	if len(opts) > 0 {
		o := opts[0]
		if o.NoPretty {
			f(s.key + ":" + s.value + ";")
			return
		}
	}
	f(s.key + " : " + s.value + ";")
}

func P(key, value string) CSSRule {
	return simple{key: key, value: value}
}

type style struct {
	selector string
	rules    RuleList
}

func (s style) write(f func(string), opts ...Options) {
	if s.rules == nil {
		return
	}
	f(s.selector)
	if len(opts) > 0 {
		o := opts[0]
		if o.NoPretty {
			f("{")
			s.rules.write(func(v string) {
				f(v)
			}, opts...)
			f("}")
			return
		}
	}
	f(" {")
	s.rules.write(func(v string) {
		f("\n   " + v)
	}, opts...)
	f("\n}\n")
}

func (style) isRule() {}

func S(selector string, rules ...CSSRule) CSSRule {
	return style{selector: selector, rules: RuleList(rules)}
}

type fontFace struct {
	key   string
	style RuleList
}

func (fontFace) isRule()                               {}
func (fontFace) write(f func(string), opts ...Options) {}

func FontFace(rules ...CSSRule) CSSRule {
	return fontFace{key: "@font-face", style: RuleList(rules)}
}

type CSSRule interface {
	//we don't want users to implement this.

	write(func(string), ...Options)
	isRule()
}

type Transformer func(CSSRule) CSSRule

func fLattern(rule CSSRule) CSSRule {
	switch e := rule.(type) {
	case RuleList:
		return flatternRuleList(e)
	case style:
		return flatternStyle(e)
	default:
		return e
	}
}

func flatternRuleList(list RuleList) RuleList {
	var o RuleList
	for _, v := range list {
		switch e := v.(type) {
		case RuleList:
			o = append(o, flatternRuleList(e)...)
		case style:
			o = append(o, flatternStyle(e)...)
		default:
			o = append(o, e)
		}
	}
	return o
}

func flatternStyle(s style) RuleList {
	var o RuleList
	baseStyle := style{selector: s.selector}
	for _, v := range s.rules {
		switch e := v.(type) {
		case style:
			ls := flatternStyle(
				style{
					selector: replaceParent(baseStyle.selector, e.selector),
					rules:    e.rules,
				})
			for _, value := range ls {
				o = append(o, value)
			}
		default:
			baseStyle.rules = append(baseStyle.rules, e)
		}
	}
	o = append(o, baseStyle)
	return o
}

func replaceParent(parent, selector string) string {
	if strings.Contains(selector, "&,") {
		return strings.Replace(selector, "&,", parent+",\n", -1)
	}
	return strings.Replace(selector, "&", parent, -1)
}

func ToString(rule CSSRule, ts ...Transformer) string {
	rule = process(rule, ts...)
	return toString(rule)
}

type Options struct {
	NoPretty bool
}

func toString(rule CSSRule, opts ...Options) string {
	var buf bytes.Buffer
	rule.write(func(v string) {
		buf.WriteString(v)
	}, opts...)
	return buf.String()
}

func process(rule CSSRule, ts ...Transformer) CSSRule {
	ts = append(ts, fLattern)
	for _, v := range ts {
		rule = v(rule)
	}
	return rule
}
