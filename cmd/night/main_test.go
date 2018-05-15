package main

import (
	"net/url"
	"testing"
)

func TestWebsockerURL(t *testing.T) {
	base := "http://localhost:8080"
	pkg := "github.com/gernest/mad"
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

func TestInScooe(t *testing.T) {
	s := "http://localhost:1955/resource?src=github.com/gernest/mad/promtest/main.js"
	pkg := "github.com/gernest/mad"
	u, err := url.Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	src := u.Query().Get("src")
	src, err = url.QueryUnescape(src)
	if err != nil {
		t.Fatal(err)
	}
	if !inPkgScope(src, pkg) {
		t.Errorf("expected %s to be in scope of %s", src, pkg)
	}
}
