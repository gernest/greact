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
		dir, pkg, rel, abs string
	}{
		{"/tests/", "tests", "", "/output/tests"},
		{"/tests/web", "web", "tests/web", "/output/tests/web"},
	}
	for _, v := range sample {
		g, err := outputInfo(cfg, v.dir, v.pkg)
		if err != nil {
			t.Error(err)
			continue
		}
		if g.RelativePath != v.rel {
			t.Errorf("expected %s got %s", v.rel, g.RelativePath)
		}
		if g.OutputPath != v.abs {
			t.Errorf("expected %s got %s", v.abs, g.OutputPath)
		}
	}
}
