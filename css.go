package gs

import (
	"bytes"
	"io"
	"strings"
)

// RuleList is a list of css style rules.
type RuleList []CSSRule

func (RuleList) isRule() {}

func (r RuleList) String() string {
	var buf bytes.Buffer
	r.Print(&buf)
	return buf.String()
}

func (r RuleList) Print(o io.Writer) (int64, error) {
	var buf bytes.Buffer
	for _, v := range r {
		_, err := v.Print(&buf)
		if err != nil {
			return 0, err
		}
	}
	return buf.WriteTo(o)
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
	var buf bytes.Buffer
	s.Print(&buf)
	return buf.String()
}

func (s StyleRule) Print(o io.Writer) (int64, error) {
	if len(s.Rules) == 0 {
		return 0, nil
	}
	var buf bytes.Buffer
	var body bytes.Buffer
	for _, v := range s.Rules {
		switch v.(type) {
		case SimpleRule:
			if body.Len() == 0 {
				body.WriteString(s.Selector + " {\n")
			}
			body.WriteString(indent(v.String(), 2))
		default:
			_, err := v.Print(&buf)
			if err != nil {
				return 0, err
			}
		}
	}
	if body.Len() > 0 {
		body.WriteString("}\n")
		buf.WriteTo(&body)
		return body.WriteTo(o)
	}
	return buf.WriteTo(o)
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
	c.Print(&buf)
	return buf.String()
}

func (c Conditional) Print(o io.Writer) (int64, error) {
	var body bytes.Buffer
	var buf bytes.Buffer
	for _, v := range c.Rules {
		_, err := v.Print(&body)
		if err != nil {
			return 0, err
		}
	}
	if body.Len() > 0 {
		buf.WriteString(c.Key + " {\n")
		buf.WriteString(indent(body.String(), 2))
		buf.WriteString("}\n")
		return buf.WriteTo(o)
	}
	return 0, nil
}

func (Conditional) isRule() {}

func indent(s string, by int) string {
	s = strings.TrimSpace(s)
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

func KeyFrame(name string, rules ...CSSRule) CSSRule {
	return Cond("@keyframes "+name, rules...)
}

type CSSRule interface {
	//we don't want users to implement this.
	Print(io.Writer) (int64, error)
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
	var buf bytes.Buffer
	rule.Print(&buf)
	return strings.TrimSpace(buf.String())
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
