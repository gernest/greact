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

type SimpleRule struct {
	Key   string
	Value string
}

func (SimpleRule) isRule() {}
func (s SimpleRule) write(f func(string), opts ...Options) {
	if len(opts) > 0 {
		o := opts[0]
		if o.NoPretty {
			f(s.Key + ":" + s.Value + ";")
			return
		}
	}
	f(s.Key + " : " + s.Value + ";")
}

func P(key, value string) CSSRule {
	return SimpleRule{Key: key, Value: value}
}

type StyleRule struct {
	Selector string
	Rules    RuleList
}

func (s StyleRule) write(f func(string), opts ...Options) {
	if s.Rules == nil {
		return
	}
	f(s.Selector)
	if len(opts) > 0 {
		o := opts[0]
		if o.NoPretty {
			f("{")
			s.Rules.write(func(v string) {
				f(v)
			}, opts...)
			f("}")
			return
		}
	}
	f(" {")
	s.Rules.write(func(v string) {
		f("\n   " + v)
	}, opts...)
	f("\n}\n")
}

func (StyleRule) isRule() {}

func S(selector string, rules ...CSSRule) CSSRule {
	return StyleRule{Selector: selector, Rules: RuleList(rules)}
}

type Conditional struct {
	Key   string
	Rules RuleList
}

func (Conditional) isRule() {}
func (c Conditional) write(f func(string), opts ...Options) {
	f(c.Key)
	var buf bytes.Buffer
	c.Rules.write(func(v string) {
		buf.WriteString(v)
	}, opts...)
	if len(opts) > 0 {
		o := opts[0]
		if o.NoPretty {
			f("{")
			f(buf.String())
			f("}")
			return
		}
	}
	f(" {\n")
	f(indent(buf.String(), 2))
	f("\n}")
}

func indent(s string, by int) string {
	p := strings.Split(s, "\n")
	var o bytes.Buffer
	idx := ""
	for i := 0; i < by; i++ {
		idx += " "
	}
	for _, v := range p {
		o.WriteString(idx)
		o.WriteString(v)
		o.WriteString("\n")
	}
	return o.String()
}
func FontFace(rules ...CSSRule) CSSRule {
	return S("@font-face", rules...)
}

func Cond(cond string, rules ...CSSRule) CSSRule {
	return Conditional{Key: cond, Rules: RuleList(rules)}
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
	case StyleRule:
		return flatternStyle(e)
	case Conditional:
		return Conditional{Key: e.Key, Rules: flatternRuleList(e.Rules)}
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
		case StyleRule:
			o = append(o, flatternStyle(e)...)
		case Conditional:
			o = append(o, Conditional{
				Key:   e.Key,
				Rules: flatternRuleList(e.Rules),
			})
		default:
			o = append(o, e)
		}
	}
	return o
}

func flatternStyle(s StyleRule) RuleList {
	var o RuleList
	baseStyle := StyleRule{Selector: s.Selector}
	for _, v := range s.Rules {
		switch e := v.(type) {
		case StyleRule:
			ls := flatternStyle(
				StyleRule{
					Selector: replaceParent(baseStyle.Selector, e.Selector),
					Rules:    e.Rules,
				})
			for _, value := range ls {
				o = append(o, value)
			}
		default:
			baseStyle.Rules = append(baseStyle.Rules, e)
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
