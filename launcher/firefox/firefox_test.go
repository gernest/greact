package firefox

import "testing"

func TestGet(t *testing.T) {
	v := Find(Firefox)
	if v == "" {
		t.Fatal("expected firefox")
	}
}
