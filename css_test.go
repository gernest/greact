package goss

import (
	"strings"
	"testing"
)

func TestCSS(t *testing.T) {
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

	s, err = ParseCSS(
		CSS{
			"float": "left",
			"width": "1px",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS("a", s, &Options{})
	e = "a {\n  float: left;\n  width: 1px;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		CSS{
			"display": "run-in",
			"fallbacks": []CSS{
				{
					"display": "inline",
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS("a", s, &Options{})
	e = "a {\n  display: inline;\n  display: run-in;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		CSS{
			"border": []string{"1px solid red", "1px solid blue"},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS("a", s, &Options{})
	e = "a {\n  border: 1px solid red, 1px solid blue;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		CSS{
			"fallbacks": CSS{
				"border": []string{"1px solid red", "1px solid blue"},
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS("a", s, &Options{})
	e = "a {\n  border: 1px solid red, 1px solid blue;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		CSS{
			"margin": [][]string{
				[]string{"5px", "10px"},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS("a", s, &Options{})
	e = "a {\n  margin: 5px 10px;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

}
