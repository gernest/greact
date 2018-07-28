// Package gs allows describing css using go functions. This loosely borrows
// ideas from less. This is low level, providing primitive blocks to effectively
// represent style rules in the go programming language.
//
// Selectors
//
// You can describe a css selector with gs.S function. This function takes the
// first argument as the selector name, followed by arbitrary number of
// attributes.
//
// For example you want to create the following css
//   a {
//     color:green;
//   }
//
// Here is the equivalent gs code
//   gs.S("a", gs.P("color", "green"))
//
// CSS attributes
//
// You can define css attributes using the P function. P is a shorthand for
// Property, this accept a key and a value.
//
// The reason this function only accept strings is to remove any assumptions
// from this library, you have the whole go ecosystem at your disposal to build
// cool abstractions on to p of this.
//
// From the previous example
//  gs.P("color", "green")
// is used to describe selector a color attribute. The P function is the only
// way to attach attributes to css elements.
//
// gs API is functional, favoring composition. This allows to have maximum
// flexibility to describe almost anything that you might be able to describe
// while writing css by hand.
//
// Nesting selectors to simplify repetitive styles
//
// Take for example you have this css in mind
//   a {
//     color:blue;
//   }
//   a:hover {
//     color:blue;
//   }
//
// There is a lot of typing involved here. Based on the understanding gs you
// have at this proint you would have written something like this.
//   gs.CSS(
//   	gs.S("a",
//   		gs.P("color", "blue"),
//   	),
//  	gs.S("a:hover",
//  		gs.P("color", "blue"),
//  	))
//
// Okay, I am introducing a a new function here CSS. This is way of groupting
// more than one CSSRule, we have a and a:hove rules so we can describe both by
// using CSS. Thing about this function that accepts arbitrary style rules and
// returns a single style sheet.
//
// As you have noticed the gs code is not visually appailing and also it doesnt
// convey our intention. hover is part of the a selector , so there is a short
// way to visually describe the styles through nesting.
//
// Here is gs code that achieves the same effect but uses nesting instead.
//   gs.S("a",
//  	gs.P("color", "blue"),
//  	gs.S("&:hover",
//  		gs.P("color", "blue"),
//   	),
//   )
//
//
// Referencing parent selector
//
// The above snippet might seem confusing, instead of `a:hover` we are having
// `&:hover` . The `&` symbol is a special character which means parent, in
// short it means just place the parent selector name here.
//
// So If you nest a S call inside the body of another S , you cn reference the
// parent selector name with &. There is no limit on the level of nesting, and &
// will always resolve to the immediate parent.
//
// css is ot nested , so the nested tree is flattened first before rendered to a
// string.
//
// Let's use this nested style to demonstrate parent reference.
//
// 	gs.S("root",
// 		gs.S(" & > child_1",
// 			gs.S("& > child2",
// 				gs.P("key", "value"),
// 			),
// 		),
// 	)
// This will give us
//   root > child_1 > child2 {
//     key:value;
//   }
//
// So in the selector with child_1 , parent node is root. After replacement this
// selector will become root > child_1, now in the selector with child_2 parent
// will be root > child_1, so after replacement we get root > child_1 > child2
//
//
package gs

import (
	"fmt"
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

// SimpleRule this is used to describe css attribute.
type SimpleRule struct {
	Key   string
	Value string
}

func (SimpleRule) isRule() {}

func (s SimpleRule) String() string {
	return s.Key + ":" + s.Value + ";"
}

// P returns a CSSRule definition for a css attribute.
func P(key, value string) CSSRule {
	return SimpleRule{Key: key, Value: value}
}

//If is a helper method that returns c if cond is true and nil otherwise.
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
		if _, ok := v.(SimpleRule); ok {
			b += "\n"
		}
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
	if parent == "" {
		return selector
	}
	if strings.Contains(selector, "&") {
		return strings.Replace(selector, "&", parent, -1)
	}
	return fmt.Sprintf("%s %s", parent, selector)
}

func ToString(rule CSSRule, ts ...Transformer) string {
	rule = Process(rule, ts...)
	return strings.TrimSpace(rule.String())
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
