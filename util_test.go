package goss

import (
	"io/ioutil"
	"testing"
)

func TestFormatCSS(t *testing.T) {
	s, err := ParseCSS("", testStyle())
	if err != nil {
		t.Fatal(err)
	}
	opts := NewOpts()
	opts.ClassNamer = IDNamer
	c := FormatCSS(s, nil, NewOpts())
	v, err := c.Print(opts)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadFile("fixture/css/format.css")
	if err != nil {
		t.Fatal(err)
	}
	e := string(b)
	if v != e {
		t.Errorf("expected %s got %s", e, v)
	}
}

func testStyle() CSS {
	return CSS{
		"root": CSS{
			LineHeight:   "1.4em",
			BoxSizing:    "border-box",
			MinWidth:     88,
			MinHeight:    36,
			BorderRadius: 2,
			"{{.root}}:hover": CSS{
				TextDecoration: "none",
				"{{.root}} {{.disabled}}": CSS{
					Background: "transparent",
				},
			},
		},
		"dense": CSS{
			MinWidth:  64,
			MinHeight: 32,
		},
		"label": CSS{
			Width:          "100%",
			Display:        "inherit",
			AlignItems:     "inherit",
			JustifyContent: "inherit",
		},
		"flat-primary": CSS{},
		"flat-accent":  CSS{},
		"color-inherit": CSS{
			"color": "inherit",
		},
	}
}
