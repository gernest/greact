package base62

import "testing"

func TestEncode(t *testing.T) {
	for i := 0; i < 5000; i++ {
		g := Decode(Encode(int64(i)))
		v := int(g)
		if v != i {
			t.Errorf("expected %d got %d", i, v)
		}
	}
}
