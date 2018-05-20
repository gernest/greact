package launcher

import (
	"regexp"
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
