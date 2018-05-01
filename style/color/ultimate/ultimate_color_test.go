package ultimate

import (
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	s := []struct {
		src    interface{}
		expect Color
		err    string
	}{
		{
			src:    [3]uint8{255, 255, 0},
			expect: Color{RGB: [3]uint8{255, 255, 0}},
		},
		{
			src: 255,
			err: "unsupported type",
		},
		{
			src:    "#fff",
			expect: Color{RGB: [3]uint8{255, 255, 255}},
		},
	}

	for _, v := range s {
		if v.err != "" {
			checkError(t, v.err, func() {
				New(v.src)
			})
		} else {
			g := New(v.src)
			compareColors(t, g, &v.expect)
		}
	}
}

func checkError(ts *testing.T, msg string, fn func()) {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			if !strings.Contains(e.Error(), msg) {
				ts.Errorf("expected %v to contain %v", e, msg)
			}
		} else {
			ts.Error("expected an error")
		}
	}()
	fn()
}

func compareColors(ts *testing.T, a, b *Color) {
	if !reflect.DeepEqual(a.RGB, b.RGB) {
		ts.Errorf("rgb: expected %v to equal %v", a.RGB, b.RGB)
	}
}

func TestColor_Hex(t *testing.T) {
	for k, v := range commonColors {
		n := New(v)
		g := n.Hex()
		if g != v {
			t.Errorf("%s: expected %s got %s", k, v, g)
		}
	}
}
