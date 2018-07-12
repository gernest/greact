package cn

import "testing"

func TestJoin(t *testing.T) {

	sample := []struct {
		names  []interface{}
		expect string
	}{
		{
			names: []interface{}{
				Name{C: "a"}, Name{C: "b", S: true}, Name{C: "f"},
			},
			expect: "a f",
		},
		{
			names: []interface{}{
				Name{C: ""}, Name{C: "b"}, "",
			},
			expect: "b",
		},
		{
			names:  []interface{}{},
			expect: "",
		},
	}

	for _, s := range sample {
		n := Join(s.names...)
		if n != s.expect {
			t.Errorf("expected %s got %s", s.expect, n)
		}
	}
}
