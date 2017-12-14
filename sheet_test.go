package goss

import (
	"strings"
	"testing"
)

func TestStyleSheets(t *testing.T) {
	sh := &StyleSheet{}
	shit := sh.NewSheet()
	err := shit.Parse(CSS{
		"a": CSS{
			"display": "run-in",
			"fallbacks": []CSS{
				{
					"display": "inline",
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	g := shit.Class["a"]
	if g != "a-1" {
		t.Errorf("expected a-1 got %s", g)
	}
}

func IDNamer(c string) string {
	if strings.HasPrefix(c, "@") || c == "" || strings.Contains(c, "{{") {
		return c
	}
	return c + "-id"
}

func TestSheet_ClassName(t *testing.T) {
	sh := &StyleSheet{}
	shit := sh.NewSheet()
	shit.ClassFunc = IDNamer
	err := shit.Parse(CSS{
		"a": CSS{
			"display": "run-in",
			"fallbacks": []CSS{
				{
					"display": "inline",
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	g := shit.Class["a"]
	if g != "a-id" {
		t.Errorf("expected a-id got %s", g)
	}
}

func TestStyleSheet_ClassName(t *testing.T) {
	sh := &StyleSheet{
		Namer: IDNamer,
	}
	shit := sh.NewSheet()
	err := shit.Parse(CSS{
		"a": CSS{
			"display": "run-in",
			"fallbacks": []CSS{
				{
					"display": "inline",
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	g := shit.Class["a"]
	if g != "a-id" {
		t.Errorf("expected a-id got %s", g)
	}
}
