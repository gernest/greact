package gs

import (
	"bytes"
	"strings"
)

type stringer interface {
	String() string
}

// RuleList is a list of css style rules.
type RuleList []CSSRule

func (RuleList) isRule() {}

func (r RuleList) String() string {
	var buf bytes.Buffer
	for k, v := range r {
		if s, ok := v.(stringer); ok {
			sc := s.String()
			if sc == "" {
				continue
			}
			if k != 0 {
				buf.WriteByte('\n')
				buf.WriteByte('\n')
			}
			buf.WriteString(sc)
		}
	}
	return buf.String()
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
	var buf bytes.Buffer
	buf.WriteString(s.Selector + " {\n")
	for _, v := range s.Rules {
		buf.WriteString(indent(v.String(), 2))
	}
	buf.WriteString("}")
	return buf.String()
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
	var buf bytes.Buffer
	buf.WriteString(c.Key + " {\n")
	for k, v := range c.Rules {
		if s, ok := v.(stringer); ok {
			if k != 0 {
				buf.WriteByte('\n')
			}
			buf.WriteString(indent(s.String(), 2))
		}
	}
	return buf.String()
}

func (Conditional) isRule() {}

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
	stringer
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
	return rule.String()
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
