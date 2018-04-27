//+build mage

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
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
	Selector bool
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
		multiColumn(),
		prefix(data{
			Mistakes: []string{"-khtml-"},
			Feature:  "user-select-none",
			Browsers: f(getstate("caniuse/features-json//user-select-none.json")),
		},
			"user-select",
		),
		flexbox(),
		prefix(data{
			Props:    []string{"*"},
			Feature:  "calc",
			Browsers: f(getstate("caniuse/features-json//calc.json")),
		},
			"calc",
		),
		prefix(data{
			Feature:  "background-img-opts",
			Browsers: f(getstate("caniuse/features-json//background-img-opts.json")),
		},
			"background-clip", "background-origin", "background-size",
		),
		prefix(data{
			Feature:  "font-feature",
			Browsers: f(getstate("caniuse/features-json//font-feature.json")),
		},
			"font-feature-settings", "font-variant-ligatures",
			"font-language-override",
		),
		prefix(data{
			Feature:  "font-kerning",
			Browsers: f(getstate("caniuse/features-json//font-kerning.json")),
		},
			"font-kerning",
		),
		prefix(data{
			Feature:  "border-image",
			Browsers: f(getstate("caniuse/features-json//border-image.json")),
		},
			"border-image",
		),
		prefix(data{
			Selector: true,
			Feature:  "css-selection",
			Browsers: f(getstate("caniuse/features-json//css-selection.json")),
		},
			"::selection",
		),
		placeHolderSelector(),
		prefix(data{
			Feature:  "css-hyphens",
			Browsers: f(getstate("caniuse/features-json//css-hyphens.json")),
		},
			"hyphens",
		),
		prefix(data{
			Selector: true,
			Feature:  "fullscreen",
			Browsers: f(getstate("caniuse/features-json//fullscreen.json")),
		},
			":fullscreen",
		),
		prefix(data{
			Selector: true,
			Feature:  "fullscreen",
			Browsers: f(getstate("caniuse/features-json//fullscreen.json"), regexp.MustCompile(`x(\s#2|$)`)),
		},
			"::backdrop",
		),
		prefix(data{
			Feature:  "css3-tabsize",
			Browsers: f(getstate("caniuse/features-json//css3-tabsize.json")),
		},
			"tab-size",
		),
		prefix(data{
			Props: []string{
				"width", "min-width", "max-width",
				"height", "min-height", "max-height",
				"inline-size", "min-inline-size", "max-inline-size",
				"block-size", "min-block-size", "max-block-size",
				"grid", "grid-template",
				"grid-template-rows", "grid-template-columns",
				"grid-auto-columns", "grid-auto-rows",
			},
			Feature:  "intrinsic-width",
			Browsers: f(getstate("caniuse/features-json/intrinsic-width.json")),
		},
			"max-content", "min-content", "fit-content",
			"fill", "fill-available", "stretch",
		),
		prefix(data{
			Props:    []string{"cursor"},
			Feature:  "css3-cursors-newer",
			Browsers: f(getstate("caniuse/features-json/css3-cursors-newer.json")),
		},
			"zoom-in", "zoom-out",
		),
		prefix(data{
			Props:    []string{"cursor"},
			Feature:  "css3-cursors-grab",
			Browsers: f(getstate("caniuse/features-json/css3-cursors-grab.json")),
		},
			"grab", "grabbing",
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

func multiColumn() map[string]data {
	base := getstate("caniuse/features-json//multicolumn.json")
	browsers := f(base)

	od := prefix(data{
		Feature:  "multicolumn",
		Browsers: browsers,
	},
		"columns", "column-width", "column-gap",
		"column-rule", "column-rule-color", "column-rule-width",
		"column-count", "column-rule-style", "column-span", "column-fill",
	)
	var noff []string
	for _, v := range browsers {
		if !strings.Contains(v, "firefox") {
			continue
		}
		noff = append(noff, v)
	}
	return mergeData(od, prefix(data{
		Feature:  "multicolumn",
		Browsers: noff,
	},
		"break-before", "break-after", "break-inside",
	))
}

func flexbox() map[string]data {
	ostat := getstate("caniuse/features-json/flexbox.json")
	browsers := f(ostat, regexp.MustCompile(`a\sx`))
	for i := range browsers {
		v := browsers[i]
		if strings.Contains(v, "ie") || strings.Contains(v, "firefox") {
			continue
		}
		browsers[i] = fmt.Sprintf("%s 2009", v)
	}
	feature := "flexbox"
	origin := mergeData(
		prefix(data{
			Props:    []string{"display"},
			Feature:  feature,
			Browsers: browsers,
		},
			"display-flex", "inline-flex",
		),
		prefix(data{
			Feature:  feature,
			Browsers: browsers,
		},
			"flex", "flex-grow", "flex-shrink", "flex-basis",
		),
		prefix(data{
			Feature:  feature,
			Browsers: browsers,
		},
			"flex-direction", "flex-wrap", "flex-flow", "justify-content",
			"order", "align-items", "align-self", "align-content",
		),
	)
	newBrowsers := f(ostat, regexp.MustCompile(`y\sx`))

	add(origin, newBrowsers, "display-flex", "inline-flex")
	add(origin, newBrowsers, "flex", "flex-grow", "flex-shrink", "flex-basis")
	add(origin, newBrowsers,
		"flex-direction", "flex-wrap", "flex-flow", "justify-content",
		"order", "align-items", "align-self", "align-content",
	)
	return origin
}

func add(m map[string]data, browsers []string, names ...string) {
	for _, name := range names {
		v := m[name]
		v.Browsers = append(v.Browsers, browsers...)
		sort.Strings(v.Browsers)
		m[name] = v
	}
}

func placeHolderSelector() map[string]data {
	browsers := f(getstate("caniuse/features-json/css-placeholder.json"))

	for i := range browsers {
		v := browsers[i]
		p := strings.Split(v, " ")
		name, version := p[0], p[1]
		f, _ := strconv.ParseFloat(version, 64)
		if name == "firefox" && f <= 18 {
			browsers[i] = fmt.Sprintf("%s old", v)
		} else if name == "ie" {
			browsers[i] = fmt.Sprintf("%s old", v)
		}
	}
	return prefix(data{
		Selector: true,
		Feature:  "css-placeholder",
		Browsers: browsers,
	},
		"::placeholder",
	)
}
