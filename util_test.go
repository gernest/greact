package goss

import (
	"io/ioutil"
	"testing"

	"github.com/kr/pretty"
)

func TestFormatCSS(t *testing.T) {
	s, err := ParseCSS("", testSTyle())
	if err != nil {
		t.Fatal(err)
	}
	c := FormatCSS(s, nil, NewOpts())
	ioutil.WriteFile("nested.css", []byte(c.Print(0)), 0600)
	t.Error(pretty.Sprint(c))
}

func testSTyle() CSS {
	return CSS{
		"root": CSS{
			LineHeight:   "1.4em",
			BoxSizing:    "border-box",
			MinWidth:     88,
			MinHeight:    36,
			BorderRadius: 2,
			"&:hover": CSS{
				TextDecoration: "none",
				"& .disabled": CSS{
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
