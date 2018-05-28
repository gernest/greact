package cover

import (
	"go/token"
	"sync"
)

type Profile struct {
	FileName string
	Mode     string
	Blocks   []*ProfileBlock
}

// ProfileBlock represents a single block of profiling data.
type ProfileBlock struct {
	StartPosition  *token.Position
	EndPosition    *token.Position
	NumStmt, Count int
}

var state = &State{
	files: make(map[string]*CoverStats),
}

type State struct {
	files map[string]*CoverStats
	mu    sync.RWMutex
}

type CoverStats struct {
	profile *Profile
	mu      sync.RWMutex
}

func (c *CoverStats) Mark(p *ProfileBlock) int {
	c.mu.Lock()
	c.profile.Blocks = append(c.profile.Blocks, p)
	idx := len(c.profile.Blocks) - 1
	c.mu.Unlock()
	return idx
}

func (c *CoverStats) Hit(idx int, pos *token.Position) {
	c.mu.Lock()
	b := c.profile.Blocks[idx]
	b.EndPosition = pos
	c.mu.Unlock()
}

func Mark(numStmt int, pos *token.Position) int {
	state.mu.Lock()
	defer state.mu.Unlock()
	p := &ProfileBlock{StartPosition: pos, NumStmt: numStmt}
	fileName := p.StartPosition.Filename
	f, ok := state.files[fileName]
	if !ok {
		f = &CoverStats{profile: &Profile{
			FileName: fileName,
		}}
		state.files[fileName] = f
		return f.Mark(p)

	}
	return f.Mark(p)
}

func Hit(idx int, pos *token.Position) {
	state.mu.Lock()
	defer state.mu.Unlock()
	if f, ok := state.files[pos.Filename]; ok {
		f.Hit(idx, pos)
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
