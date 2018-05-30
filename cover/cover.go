package cover

import (
	"encoding/json"
	"fmt"
	"go/token"
	"sort"
	"sync"

	"github.com/gopherjs/gopherjs/js"
)

const (
	Key = "mad_coverage_stats"
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
	files: &sync.Map{},
}

type State struct {
	files *sync.Map
}

type CoverStats struct {
	Profile *Profile
	blocks  *sync.Map
}

func (c *CoverStats) FillBlocks() {
	var keys []int
	c.blocks.Range(func(k, _ interface{}) bool {
		keys = append(keys, k.(int))
		return true
	})
	sort.Ints(keys)
	c.Profile.Blocks = nil
	for _, k := range keys {
		if v, ok := c.blocks.Load(k); ok {
			c.Profile.Blocks = append(c.Profile.Blocks, v.(*ProfileBlock))
		}
	}
}

func (c *CoverStats) Mark(p *ProfileBlock) int {
	if idx, ok := c.blocks.Load(p.StartPosition.Line); ok {
		idx.(*ProfileBlock).Count++
		return p.StartPosition.Line
	}
	c.blocks.Store(p.StartPosition.Line, p)
	return p.StartPosition.Line
}

func (c *CoverStats) Hit(pos *token.Position) {
	if block, ok := c.blocks.Load(pos.Line); ok {
		block.(*ProfileBlock).Count++
	}
}

func Mark(numStmt int, start, end *token.Position) int {
	p := &ProfileBlock{
		StartPosition: start,
		EndPosition:   end,
		NumStmt:       numStmt,
	}
	fileName := p.StartPosition.Filename
	f, ok := state.files.Load(fileName)
	if !ok {
		f := &CoverStats{
			blocks: &sync.Map{},
			Profile: &Profile{
				FileName: fileName,
			},
		}
		state.files.Store(fileName, f)
		return f.Mark(p)
	}
	return f.(*CoverStats).Mark(p)
}

func Hit(pos *token.Position) {
	if f, ok := state.files.Load(pos.Filename); ok {
		f.(*CoverStats).Hit(pos)
	}
}

func Stats() []*CoverStats {
	var o []*CoverStats
	state.files.Range(func(_, v interface{}) bool {
		c := v.(*CoverStats)
		c.FillBlocks()
		o = append(o, c)
		return true
	})
	sort.Slice(o, func(i, j int) bool {
		return o[i].Profile.FileName < o[j].Profile.FileName
	})
	return o
}

// JSON marshals current coverage state to json string.
func JSON() string {
	b, _ := json.MarshalIndent(Stats(), "", "  ")
	return string(b)
}

// Dump calls console.info("coverage",$COVERAGE_DATA_IN_JSON)
func Dump() (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = fmt.Errorf("%v", perr)
		}
	}()
	js.Global.Get("console").Call("info", "coverage", JSON())
	return
}
