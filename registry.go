package gs

import (
	"strconv"
)

type SheetObject interface {
	InsertRule(rule string)
	Detach()
}

type SimpleRegistry struct {
	idx             int64
	sheets          map[int64]*registryNode
	attached        map[int64]bool
	CreateSheetNode func() SheetObject
}

func NewSimpleRegistry(fn func() SheetObject) *SimpleRegistry {
	return &SimpleRegistry{
		sheets:          make(map[int64]*registryNode),
		attached:        make(map[int64]bool),
		CreateSheetNode: fn,
	}
}

type registryNode struct {
	sheet *Sheet
	node  SheetObject
}

func (m *SimpleRegistry) NewSheet() *Sheet {
	node := m.CreateSheetNode()
	id := m.getID()
	s := &Sheet{
		id:       id,
		CLasses:  make(ClassMap),
		idGen:    genericIDGen(id),
		registry: m,
	}
	m.sheets[s.id] = &registryNode{
		sheet: s, node: node,
	}
	return s
}

func (m *SimpleRegistry) getID() int64 {
	m.idx++
	return m.idx
}

func genericIDGen(base int64) func() string {
	counter := int64(0)
	return func() string {
		counter++
		return join(formatInt64(base), "-", formatInt64(counter))
	}
}

func formatInt64(n int64) string {
	return strconv.FormatInt(n, 10)
}

func join(s ...string) string {
	o := ""
	for _, v := range s {
		o += v
	}
	return o
}

func (m *SimpleRegistry) Attach(s *Sheet) {
	if !m.attached[s.id] {
		r := m.sheets[s.id]
		for _, v := range s.ListRules() {
			if v != "" {
				r.node.InsertRule(v)
			}
		}
		m.attached[s.id] = true
	}
}

func (m *SimpleRegistry) Detach(s *Sheet) {
	if m.attached[s.id] {
		r := m.sheets[s.id]
		r.node.Detach()
		delete(m.attached, s.id)

		//TODO: recycle the ID to be used by another new sheet
	}
}

func (*SimpleRegistry) isRegistry() {}
