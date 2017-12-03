package goss

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestCSS(t *testing.T) {
	s, err := ParseCSS(
		"",
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
		"a",
		CSS{
			"float": "left",
			"width": "1px",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS(s, &Options{})
	e = "a {\n  float: left;\n  width: 1px;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		"a",
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
	str = ToCSS(s, &Options{})
	e = "a {\n  display: inline;\n  display: run-in;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		"a",
		CSS{
			"border": []string{"1px solid red", "1px solid blue"},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS(s, &Options{})
	e = "a {\n  border: 1px solid red, 1px solid blue;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		"a",
		CSS{
			"fallbacks": CSS{
				"border": []string{"1px solid red", "1px solid blue"},
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS(s, &Options{})
	e = "a {\n  border: 1px solid red, 1px solid blue;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		"a",
		CSS{
			"margin": [][]string{
				[]string{"5px", "10px"},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS(s, &Options{})
	e = "a {\n  margin: 5px 10px;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		"a",
		CSS{
			"float": "left",
			"& b": CSS{
				"float": "left",
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS(s, &Options{})
	e = "a {\n  float: left;\n}\na b {\n  float: left;\n}"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		"",
		CSS{
			"@import": []string{
				`url("something") print`,
				`url("something") screen`,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS(s, &Options{})
	e = "@import url(\"something\") print;\n@import url(\"something\") screen;"
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

}

func TestConditional(t *testing.T) {
	s, err := ParseCSS(
		"",
		CSS{
			"@media print": CSS{
				"a": CSS{
					Display: "none",
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str := ToCSS(s, &Options{})
	b, err := ioutil.ReadFile("fixture/css/media.css")
	if err != nil {
		t.Fatal(err)
	}
	e := string(b)
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}
}
