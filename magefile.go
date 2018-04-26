//+build mage

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"

	"github.com/magefile/mage/sh"
)

const pkg = "github.com/gernest/gs/demo"

func Demo() {
	sh.RunV("gopherjs", "build", "-o", "demo/main.js", pkg)
}

type Stats struct {
	Stats map[string]map[string]string `json:"stats"`
}

func TestSample() error {
	b, err := ioutil.ReadFile("caniuse/sample-data.json")
	if err != nil {
		return err
	}
	s := Stats{}
	err = json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	return nil
}

func getstate(file string) Stats {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	s := Stats{}
	err = json.Unmarshal(b, &s)
	if err != nil {
		panic(err)
	}
	return s
}

var match = regexp.MustCompile(`\sx($|\s)`)

func f(s Stats, opts ...*regexp.Regexp) []string {
	matcher := match
	if len(opts) > 0 {
		matcher = opts[0]
	}
	var need []string
	for browser := range s.Stats {
		versions := s.Stats[browser]
		for version := range versions {
			support := versions[version]
			if matcher.Match([]byte(support)) {
				need = append(need, fmt.Sprintf("%s %s", browser, version))
			}
		}
	}
	return need
}

type data struct {
	Props    []string
	Mistakes []string
	Feature  string
	Browsers []string
}

func prefix(d data, names ...string) map[string]data {
	o := make(map[string]data)
	for _, v := range names {
		o[v] = d
	}
	return o
}

func mergeData(d ...map[string]data) map[string]data {
	o := make(map[string]data)
	for _, m := range d {
		for k, v := range m {
			o[k] = v
		}
	}
	return o
}

func Gen() error {
	fmt.Println(mergeData(
		prefix(data{
			Mistakes: []string{"-khtml-", "-ms-", "-o-"},
			Feature:  "border-radius",
			Browsers: f(getstate("caniuse/features-json/border-radius.json")),
		},
			"border-radius", "border-top-left-radius", "border-top-right-radius",
			"border-bottom-right-radius", "border-bottom-left-radius",
		),
		prefix(data{
			Mistakes: []string{"-khtml-"},
			Feature:  "css-boxshadow",
			Browsers: f(getstate("caniuse/features-json/css-boxshadow.json")),
		},
			"box-shadow",
		),
		prefix(data{
			Mistakes: []string{"-khtml-", "-ms-"},
			Feature:  "css-animation",
			Browsers: f(getstate("caniuse/features-json/css-animation.json")),
		},
			"animation", "animation-name", "animation-duration",
			"animation-delay", "animation-direction", "animation-fill-mode",
			"animation-iteration-count", "animation-play-state",
			"animation-timing-function", "@keyframes",
		),
		prefix(data{
			Mistakes: []string{"-khtml-", "-ms-"},
			Feature:  "css-transitions",
			Browsers: f(getstate("caniuse/features-json/css-transitions.json")),
		},
			"transition", "transition-property", "transition-duration",
			"transition-delay", "transition-timing-function",
		),
		prefix(data{
			Feature:  "transforms2d",
			Browsers: f(getstate("caniuse/features-json/transforms2d.json")),
		},
			"transform", "transform-origin",
		),
		prefix(data{
			Feature:  "transforms3d",
			Browsers: f(getstate("caniuse/features-json/transforms3d.json")),
		},
			"perspective", "perspective-origin",
		),
		prefix(data{
			Mistakes: []string{"-ms-", "-o-"},
			Feature:  "transforms3d",
			Browsers: f(getstate("caniuse/features-json/transforms3d.json")),
		},
			"transform-style",
		),
		prefix(data{
			Mistakes: []string{"-ms-", "-o-"},
			Feature:  "transforms3d",
			Browsers: f(getstate("caniuse/features-json/transforms3d.json"),
				regexp.MustCompile(`y\sx|y\s#2/`),
			),
		},
			"backface-visibility",
		),
		gradients(),
		prefix(data{
			Feature:  "css3-boxsizing",
			Browsers: f(getstate("caniuse/features-json/css3-boxsizing.json")),
		},
			"box-sizing",
		),
		prefix(data{
			Feature:  "css-filters",
			Browsers: f(getstate("caniuse/features-json/css-filters.json")),
		},
			"filter",
		),
		prefix(data{
			Props: []string{
				"background", "background-image", "border-image", "mask",
				"list-style", "list-style-image", "content", "mask-image",
			},
			Feature:  "css-filter-function",
			Browsers: f(getstate("caniuse/features-json/css-filter-function.json")),
		},
			"filter-function",
		),
		prefix(data{
			Feature:  "css-backdrop-filter",
			Browsers: f(getstate("caniuse/features-json/css-backdrop-filter.json")),
		},
			"backdrop-filter",
		),
		prefix(data{
			Props: []string{
				"background", "background-image", "border-image", "mask",
				"list-style", "list-style-image", "content", "mask-image",
			},
			Feature:  "css-element-function",
			Browsers: f(getstate("caniuse/features-json/css-element-function.json")),
		},
			"element",
		),
	))
	return nil
}

func gradients() map[string]data {
	origin := prefix(data{
		Mistakes: []string{"-ms-"},
		Props: []string{"background", "background-image", "border-image", "mask",
			"list-style", "list-style-image", "content", "mask-image"},
		Feature: "css-gradients",
		Browsers: f(getstate("caniuse/features-json/css-gradients.json"),
			regexp.MustCompile(`y\sx`),
		),
	},
		"linear-gradient", "repeating-linear-gradient",
		"radial-gradient", "repeating-radial-gradient",
	)

	browsers := f(getstate("caniuse/features-json/css-gradients.json"),
		regexp.MustCompile(`a\sx`),
	)
	for i := range browsers {
		if !strings.Contains(browsers[i], "op") {
			browsers[i] = fmt.Sprintf("%s old", browsers[i])
		}
	}

	names := []string{
		"linear-gradient", "repeating-linear-gradient",
		"radial-gradient", "repeating-radial-gradient",
	}
	for _, name := range names {
		odata := origin[name]
		odata.Browsers = append(odata.Browsers, browsers...)
		sort.Strings(odata.Browsers)
		origin[name] = odata
	}
	return origin
}
