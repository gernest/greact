package goss

import (
	"testing"

	"github.com/kr/pretty"
)

func TestSelector(t *testing.T) {
	s := "div"
	o := C(Selector(s))
	if o.selector != s {
		t.Errorf("expected %s got %s", s, o.selector)
	}
}

func TestRenderObject(t *testing.T) {
	s := "div"
	o := C(Selector(s),
		Prop("height", "10px"),
		Prop("width", "10px"),
		C(Selector("p"),
			Prop("background", "blue"),
		),
	)
	ctx := NewSheet()
	err := renderObject(nil, ctx, o)
	if err != nil {
		t.Fatal(err)
	}
	pretty.Println(o)
	pretty.Println(ctx)
	t.Error("yay")
}

func TestPrintSelector(t *testing.T) {
	s := []struct {
		sel    []string
		expect string
	}{
		{
			[]string{".intro"}, ".intro",
		},
		{
			[]string{"#firstname"}, "#firstname",
		},
		{
			[]string{"*"}, "*",
		},
		{
			[]string{"div", ",p"}, "div,p",
		},
		{
			[]string{"div", ">", "p"}, "div > p",
		},
	}
	for _, v := range s {
		g := printSelectors(v.sel)
		if g != v.expect {
			t.Errorf("expected %s got %s", v.expect, g)
		}
	}
}
