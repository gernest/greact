package config

import (
	"testing"
)

func TestGetOutputPath(t *testing.T) {
	cfg := &Config{
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
		g, err := getOutputInfo(cfg, v.dir, v.pkg)
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
