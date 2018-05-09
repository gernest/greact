package main

import "testing"

func TestWebsockerURL(t *testing.T) {
	base := "http://localhost:8080"
	pkg := "github.com/gernest/prom"
	u, err := websocketURL(base, pkg)
	if err != nil {
		t.Fatal(err)
	}
	expect := "ws://localhost:8080/test?pkg=github.com%2Fgernest%2Fprom"
	if u != expect {
		t.Errorf("expected %s got %s", expect, u)
	}
	// t.Error(u)
}
