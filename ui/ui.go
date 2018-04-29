package ui

import (
	"github.com/gernest/gs"
	"github.com/gopherjs/gopherjs/js"
)

var Registry = gs.NewSimpleRegistry(create())

var Global = Registry.NewSheet()

func create() func() gs.SheetObject {
	if js.Global == nil ||
		js.Global != nil && js.Global.Get("document") == nil {
		return func() gs.SheetObject {
			return &mockSheetObject{}
		}
	}
	return createStyle()
}

type CSSSheet interface {
	Attach()
	Detach()
}

func createStyle() func() gs.SheetObject {
	doc := js.Global.Get("document")
	s := doc.Call("createElement", "style")
	s.Call("appendChild", doc.Call("createTextNode", ""))
	doc.Get("head").Call("appendChild", s)
	sheet := s.Get("sheet")
	return func() gs.SheetObject {
		return &sheetObject{Object: sheet}
	}
}

//implements gs.SheetObject but uses real dom node
type sheetObject struct {
	*js.Object

	indexes []int64
}

func (s *sheetObject) InsertRule(rule string) {
	n := s.Get("cssRules").Get("length").Int64()
	g := s.Call("insertRule", rule, n).Int64()
	s.indexes = append(s.indexes, g)
}

func (s *sheetObject) Detach() {
	for _, v := range s.indexes {
		s.Call("deleteRule", v)
	}
}

func NewSheet() *gs.Sheet {
	return Registry.NewSheet()
}

type mockSheetObject struct {
	rules    []string
	detached bool
}

func (m *mockSheetObject) InsertRule(rule string) {
	m.rules = append(m.rules, rule)
}

func (m *mockSheetObject) Detach() {
	m.detached = true
}
