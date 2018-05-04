package helper

import "sync"

var state = &State{
	files: make(map[string]*CoverStats),
}

type State struct {
	files map[string]*CoverStats
	mu    sync.RWMutex
}

type CoverStats struct {
	Filename string
	Lines    map[int]int
	mu       sync.RWMutex
}

func (c *CoverStats) Mark(line int) {
	c.mu.Lock()
	c.Lines[line] = 0
	c.mu.Unlock()
}

func (c *CoverStats) Hit(line int) {
	c.mu.Lock()
	ln := c.Lines[line]
	ln++
	c.Lines[line] = ln
	c.mu.Unlock()
}

func Mark(name string, line int) {
	state.mu.Lock()
	f, ok := state.files[name]
	if !ok {
		f = &CoverStats{
			Filename: name,
			Lines:    make(map[int]int)}
		f.Mark(line)
		state.files[name] = f
	} else {
		f.Mark(line)
	}
}

func Hit(name string, line int) {
	state.mu.Lock()
	if f, ok := state.files[name]; ok {
		f.Hit(line)
	}
}

func Stats() []*CoverStats {
	state.mu.RLock()
	var o []*CoverStats
	for _, v := range state.files {
		o = append(o, v)
	}
	state.mu.RUnlock()
	return o
}
