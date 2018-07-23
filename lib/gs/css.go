// Package gs allows describing css using go functions.
package gs

import (
	"io"
	"strings"
)

// RuleList is a list of css style rules.
type RuleList []CSSRule

func (RuleList) isRule() {}

func (r RuleList) String() string {
	s := ""
	for _, v := range r {
		s += v.String()
	}
	return s
}

// filter out nil values
func toRuleList(src []CSSRule) RuleList {
	var ls RuleList
	for _, v := range src {
		if v != nil {
			ls = append(ls, v)
		}
	}
	return ls
}

// CSS defines a list of style rules. Nil values will be removed.
func CSS(rules ...CSSRule) CSSRule {
	return toRuleList(rules)
}

type SimpleRule struct {
	Key   string
	Value string
}

func (SimpleRule) isRule() {}

func (s SimpleRule) String() string {
	return s.Key + ":" + s.Value + ";"
}

func (s SimpleRule) Print(o io.Writer) (int64, error) {
	v, err := o.Write([]byte(s.Key + ":" + s.Value + ";"))
	return int64(v), err
}

func P(key, value string) CSSRule {
	return SimpleRule{Key: key, Value: value}
}

func If(cond bool, c CSSRule) CSSRule {
	if cond {
		return c
	}
	return nil
}

type StyleRule struct {
	Selector string
	Rules    RuleList
}

func (s StyleRule) String() string {
	if len(s.Rules) == 0 {
		return ""
	}
	o := ""
	b := ""
	for _, v := range s.Rules {
		switch v.(type) {
		case SimpleRule:
			b += indent(v.String(), 2)
		default:
			o += v.String()
		}
	}
	if b != "" {
		return s.Selector + " {\n" + b + "}\n" + o
	}
	return o
}

func (StyleRule) isRule() {}

func S(selector string, rules ...CSSRule) CSSRule {
	return StyleRule{Selector: selector, Rules: toRuleList(rules)}
}

type Conditional struct {
	Key   string
	Rules RuleList
}

func (c Conditional) String() string {
	if len(c.Rules) == 0 {
		return ""
	}
	b := ""
	for _, v := range c.Rules {
		b += v.String()
	}
	return c.Key + " {\n" + indent(b, 2) + "}\n"
}

func (Conditional) isRule() {}

func indent(s string, by int) string {
	s = strings.TrimSpace(s)
	p := strings.Split(s, "\n")
	o := ""
	idx := ""
	for i := 0; i < by; i++ {
		idx += " "
	}
	for _, v := range p {
		o += idx + v + "\n"
	}
	return o
}

func FontFace(rules ...CSSRule) CSSRule {
	return S("@font-face", rules...)
}

func Cond(cond string, rules ...CSSRule) CSSRule {
	return Conditional{Key: cond, Rules: RuleList(rules)}
}

func KeyFrame(name string, rules ...CSSRule) CSSRule {
	return Cond("@keyframes "+name, rules...)
}

func Media(name string, rules ...CSSRule) CSSRule {
	return Cond("@media "+name, rules...)
}

type CSSRule interface {
	//we don't want users to implement this.
	String() string
	isRule()
}

type Transformer func(CSSRule) CSSRule

func Flattern(rule CSSRule) CSSRule {
	return flattern("", rule)
}
func flattern(parent string, rule CSSRule) RuleList {
	var o RuleList
	switch e := rule.(type) {
	case RuleList:
		for _, v := range e {
			o = append(o, flattern(parent, v)...)
		}
	case StyleRule:
		sel := e.Selector
		if parent != "" {
			sel = replaceParent(parent, sel)
		}
		o = append(o, StyleRule{
			Selector: sel,
			Rules:    flattern(sel, e.Rules),
		})
	case Conditional:
		key := e.Key
		if parent != "" {
			key = replaceParent(parent, key)
		}
		o = append(o, Conditional{
			Key:   key,
			Rules: flattern(key, e.Rules),
		})
	default:
		o = append(o, e)
	}
	return o
}

func replaceParent(parent, selector string) string {
	if strings.Contains(selector, "&,") {
		return strings.Replace(selector, "&,", parent+",\n", -1)
	}
	return strings.Replace(selector, "&", parent, -1)
}

func ToString(rule CSSRule, ts ...Transformer) string {
	rule = Process(rule, ts...)
	return strings.TrimSpace(rule.String())
}

type Options struct {
	NoPretty bool
}

// Process this applies any transformation to the rule. It automatically
// flatterns the rule tree.
func Process(rule CSSRule, ts ...Transformer) CSSRule {
	ts = append(ts, Flattern)
	for _, v := range ts {
		rule = v(rule)
	}
	return rule
}
