package ui

import (
	"github.com/gernest/gs"
	"github.com/gopherjs/gopherjs/js"
)

var Registry = gs.NewSimpleRegistry(create)

var Global = Registry.NewSheet()

func create() gs.SheetObject {
	if js.Global == nil ||
		js.Global != nil && js.Global.Get("document") != nil {
		return &mockSheetObject{}
	}
	return createStyle()
}

type CSSSheet interface {
	Attach()
	Detach()
}

func createStyle() gs.SheetObject {
	doc := js.Global.Get("document")
	s := doc.Call("createElement", "style")
	s.Call("appendChild", doc.Call("createTextNode", ""))
	doc.Get("head").Call("appendChild", s)
	sheet := s.Get("sheet")
	return &sheetObject{Object: sheet}
}

//implements gs.SheetObject but uses real dom node
type sheetObject struct {
	*js.Object
}

func (s *sheetObject) InsertRule(rule string) {
	s.Call("insertRule", rule)
}

func (s *sheetObject) Detach() {
	s.Get("parent").Call("removeChild", s.Object)
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
