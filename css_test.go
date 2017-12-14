package goss

import (
	"fmt"
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
	b, err := ioutil.ReadFile("fixture/css/css_01.css")
	if err != nil {
		t.Fatal(err)
	}
	e = string(b)
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
	b, err = ioutil.ReadFile("fixture/css/css_02.css")
	if err != nil {
		t.Fatal(err)
	}
	e = string(b)
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
	b, err = ioutil.ReadFile("fixture/css/css_03.css")
	if err != nil {
		t.Fatal(err)
	}
	e = string(b)
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
	b, err = ioutil.ReadFile("fixture/css/css_4.css")
	if err != nil {
		t.Fatal(err)
	}
	e = string(b)
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
	b, err = ioutil.ReadFile("fixture/css/css_5.css")
	if err != nil {
		t.Fatal(err)
	}
	e = string(b)
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}

	s, err = ParseCSS(
		"a",
		CSS{
			"float": "left",
			"{{.a}} b": CSS{
				"float": "left",
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS(s, &Options{})
	b, err = ioutil.ReadFile("fixture/css/css_6.css")
	if err != nil {
		t.Fatal(err)
	}
	e = string(b)
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

func writeSample(s string, n int, t *testing.T) {
	err := ioutil.WriteFile(fmt.Sprintf("fixture/css/css_%d.css", n), []byte(s), 0600)
	if err != nil {
		t.Fatal(err)
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
	b, err := ioutil.ReadFile("fixture/css/css_7.css")
	if err != nil {
		t.Fatal(err)
	}
	e := string(b)
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}
	writeSample(str, 7, t)

	s, err = ParseCSS(
		"",
		CSS{
			"@media(max-width: 715px)": CSS{
				"a": CSS{
					Display: "none",
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	str = ToCSS(s, &Options{})
	b, err = ioutil.ReadFile("fixture/css/css_8.css")
	if err != nil {
		t.Fatal(err)
	}
	e = string(b)
	if str != e {
		t.Errorf("expected %s got %s", e, str)
	}
}
