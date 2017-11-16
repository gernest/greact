package goss

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseSimpleRules(t *testing.T) {
	s, err := ParseCSS(
		CSS{
			"background": "blue",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	o := s.Rules[0]
	if o.Type() != StyleRule {
		t.Error("expected StyleRule")
	}

	str := ToCSS("simple", s, &Options{})
	str = strings.TrimSpace(str)
	b, err := ioutil.ReadFile("fixture/css/simple.css")
	if err != nil {
		t.Fatal(err)
	}
	e := string(b)
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}
}
