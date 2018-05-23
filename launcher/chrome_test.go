package launcher

import (
	"os"
	"testing"
)

func TestFindChrome(t *testing.T) {
	v, err := resolveChromePath()
	if err != nil {
		t.Fatal(err)
	}
	if len(v) == 0 {
		t.Fatal("expected absolute path to chrome")
	}
	for _, path := range v {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			t.Errorf("expected file %s to exist", path)
		}
	}
}
