package prefix

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gernest/vected/tools/prefix/data"
)

type SelectedOptions struct {
	Add    map[string][]string
	Remove map[string][]string
}

type Prefixes struct {
	Browsers *Browser
	Data     map[string]data.Data
	Opts     *PrefixesOptions
}

type PrefixesOptions struct {
	FlexBox string
}

type addOpts struct {
	browser string
	note    string
}

func (p *Prefixes) Select(list map[string]data.Data) *SelectedOptions {
	selected := &SelectedOptions{
		Add:    make(map[string][]string),
		Remove: make(map[string][]string),
	}
	for name := range list {
		data := list[name]
		var add []addOpts
		for _, v := range data.Browsers {
			parts := strings.Split(v, " ")
			o := addOpts{
				browser: fmt.Sprintf("%s %s", parts[0], parts[1]),
			}
			if len(parts) == 3 {
				o.note = parts[2]
			}
			add = append(add, o)
		}
		var notes []string
		for _, v := range add {
			if v.note != "" {
				notes = append(notes,
					fmt.Sprintf("%s %s", p.Browsers.Prefix(v.browser), v.note),
				)
			}
		}
		notes = uniq(notes)
		var addList []string
		fadd := filterAddOptions(add, func(v addOpts) bool {
			sl := p.Browsers.IsSelected(v.browser)
			return sl
		})
		for _, v := range fadd {
			prefx := p.Browsers.Prefix(v.browser)
			if v.note != "" {
				addList = append(addList,
					fmt.Sprintf("%s %s", prefx, v.note),
				)
			} else {
				addList = append(addList, prefx)
			}
		}
		addList = uniq(addList)
		sort.Strings(addList)
		if p.Opts != nil && p.Opts.FlexBox == "no-2009" {
			addList = filter(addList, func(v string) bool {
				return !strings.Contains(v, "2009")
			})
		}
		all := mapSlice(data.Browsers, func(v string) string {
			return p.Browsers.Prefix(v)
		})
		if data.Mistakes != nil {
			all = append(all, data.Mistakes...)
		}
		if notes != nil {
			all = append(all, notes...)
		}
		all = uniq(all)
		if len(addList) > 0 {
			selected.Add[name] = addList
			if len(addList) < len(all) {
				rm := filter(all, func(v string) bool {
					return !sliceContains(addList, v)
				})
				selected.Remove[name] = rm
			}
		} else {
			selected.Remove[name] = all
		}
	}
	return selected
}

func sliceContains(s []string, v string) bool {
	for k := range s {
		if s[k] == v {
			return true
		}
	}
	return false
}

func filter(src []string, fn func(string) bool) []string {
	var ns []string
	for _, v := range src {
		if fn(v) {
			ns = append(ns, v)
		}
	}
	return ns
}

func filterAddOptions(src []addOpts, fn func(addOpts) bool) []addOpts {
	var ns []addOpts
	for _, v := range src {
		if fn(v) {
			ns = append(ns, v)
		}
	}
	return ns
}

func mapSlice(m []string, fn func(string) string) []string {
	var ns []string
	for _, v := range m {
		ns = append(ns, fn(v))
	}
	return ns
}
