//+build mage

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

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

func f(s Stats) []string {
	var need []string
	for browser := range s.Stats {
		versions := s.Stats[browser]
		for version := range versions {
			support := versions[version]
			if match.Match([]byte(support)) {
				need = append(need, fmt.Sprintf("%s %s", browser, version))
			}
		}
	}
	return need
}

type data struct {
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
	))
	return nil
}
