package main

import (
	"testing"

	"github.com/gernest/mad/config"
)

func TestOutputPath(t *testing.T) {
	cfg := &config.Config{
		OutputPath:  "/output/",
		TestPath:    "/tests/",
		TestDirName: "tests",
	}
	sample := []struct {
		dir, pkg, expect string
	}{
		{"/tests/", "tests", "/output/tests"},
		{"/tests/web", "web", "/output/tests/web"},
	}
	for _, v := range sample {
		g, err := outputPath(cfg, v.dir, v.pkg)
		if err != nil {
			t.Error(err)
			continue
		}
		if g != v.expect {
			t.Errorf("expected %s got %s", v.expect, g)
		}
	}
}
