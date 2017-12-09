package goss

import "testing"

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
