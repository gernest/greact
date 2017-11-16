package goss

import (
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

	str := o.ToString(&Options{})
	str = strings.TrimSpace(str)
	e := "background: blue;"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}
}
