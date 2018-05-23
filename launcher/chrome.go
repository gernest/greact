package launcher

import (
	"fmt"
	"regexp"
	"runtime"
	"sort"
)

const chromePath = "CHROME_PATH"

type priority struct {
	regex  *regexp.Regexp
	weight int
}
type installPriority struct {
	path   string
	weight int
}

func sortStuff(install []string, priorities []*priority) []string {
	defaultPolicy := 10
	var m []*installPriority
	for _, v := range install {
		for _, p := range priorities {
			if p.regex.MatchString(v) {
				m = append(m, &installPriority{
					path: v, weight: p.weight,
				})
				continue
			}
		}
		m = append(m, &installPriority{
			path: v, weight: defaultPolicy,
		})
	}
	sort.Slice(m, func(a, b int) bool {
		return m[a].weight < m[b].weight
	})
	var o []string
	for _, v := range m {
		o = append(o, v.path)
	}
	return o
}

func defaultFlags() []string {
	return []string{
		"--disable-translate",
		"--disable-extensions",
		"--disable-background-networking",
		"--safebrowsing-disable-auto-update",
		"--disable-sync",
		"--metrics-recording-only",
		"--disable-default-apps",
		"--mute-audio",
		"--no-first-run",
	}
}

type Options struct {
	StartingURL string
	Flags       []string
}

type Launcher struct {
	Opts *Options
}

func New() {

}

func getPlatform() (string, error) {
	v := runtime.GOOS
	switch v {
	case "darwin", "linux":
		return v, nil
	default:
		return "", fmt.Errorf("platform %s is not supported", v)
	}
}

func darwin() ([]string, error) {
	return resolveChromePathDarwin()
}

func resolveChromePath() ([]string, error) {
	platform, err := getPlatform()
	if err != nil {
		return nil, err
	}
	switch platform {
	case "darwin":
		return darwin()
	default:
		return nil, fmt.Errorf("platform %s is not supported", platform)
	}
}
