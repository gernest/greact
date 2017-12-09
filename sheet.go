package goss

import (
	"bytes"
	"io"
	"strconv"
	"sync"
)

type Sheet struct {
	Class     ClassMap
	ClassFunc func(string) string
	Src       bytes.Buffer
}

func (s *Sheet) Parse(css CSS) error {
	style, err := ParseCSS("", css)
	if err != nil {
		return err
	}
	opts := NewOpts()
	opts.ClassNamer = s.ClassFunc
	out := ToCSS(style, opts)
	if s.Src.Len() == 0 {
		s.Src.WriteString(out)
	} else {
		s.Src.WriteString("\n" + out)
	}
	s.Class.Merge(opts.ClassMap)
	return nil
}

type StyleSheet struct {
	Sheets []*Sheet
	index  int64
	mu     sync.RWMutex
}

func (s *StyleSheet) incr() {
	s.mu.Lock()
	s.index++
	s.mu.Unlock()
}

func (s *StyleSheet) getIndex() int64 {
	s.mu.RLock()
	v := s.index
	s.mu.RUnlock()
	return v
}

func (s *StyleSheet) ClassNamer(c string) string {
	if hasPrefix(c, "@") {
		return c
	}
	s.incr()
	id := strconv.FormatInt(s.getIndex(), 10)
	return c + "-" + id
}

func (s *StyleSheet) NewSheet() *Sheet {
	shit := &Sheet{
		Class:     make(ClassMap),
		ClassFunc: s.ClassNamer,
	}
	s.Sheets = append(s.Sheets, shit)
	return shit
}

func (s *StyleSheet) String() string {
	var o bytes.Buffer
	for _, shit := range s.Sheets {
		if o.Len() == 0 {
			io.Copy(&o, &shit.Src)
		} else {
			o.WriteRune('\n')
			io.Copy(&o, &shit.Src)
		}
	}
	return o.String()
}
