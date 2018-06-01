// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cover

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"sync"
)

const (
	Key = "mad_coverage_stats"
)

var profileStats = &sync.Map{}

type ProfileBlock struct {
	StartLine, StartCol int
	EndLine, EndCol     int
	NumStmt, Count      int
}

type Profile struct {
	FileName string
	Mode     string
	Blocks   []ProfileBlock
}

// JSON marshals current coverage state to json string.
func JSON() string {
	var profiles []Profile
	profileStats.Range(func(_, v interface{}) bool {
		fn := v.(CoverFunc)
		profiles = append(profiles, fn()...)
		return true
	})
	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].FileName < profiles[j].FileName
	})
	b, _ := json.Marshal(profiles)
	return string(b)
}

func Register(packageName string, coverFunc CoverFunc) {
	if _, ok := profileStats.Load(packageName); !ok {
		profileStats.Store(packageName, coverFunc)
	}
}

func File(fileName, mode string, counter []uint32, pos []uint32, numStmts []uint16) Profile {
	if 3*len(counter) != len(pos) || len(counter) != len(numStmts) {
		panic("coverage: mismatched sizes")
	}
	block := make([]ProfileBlock, len(counter))
	for i, v := range counter {
		block[i] = ProfileBlock{
			StartLine: int(pos[3*i+0]),
			StartCol:  int(uint16(pos[3*i+2])),
			EndLine:   int(pos[3*i+1]),
			EndCol:    int(uint16(pos[3*i+2] >> 16)),
			NumStmt:   int(numStmts[i]),
			Count:     int(v),
		}
	}
	return Profile{
		FileName: fileName,
		Mode:     mode,
		Blocks:   block,
	}
}

type CoverFunc func() []Profile

// Calc returns the code covearge for the given profiles. We just take the total
// number of profile blocks with Count >0 divide by the total number of profile
// blocks.
func Calc(profiles []Profile) float64 {
	var total, covered int64
	for _, p := range profiles {
		for _, v := range p.Blocks {
			total += int64(v.NumStmt)
			if v.Count > 0 {
				covered += int64(v.NumStmt)
			}
		}
	}
	if total == 0 {
		return 0
	}
	return float64(covered) / float64(total)
}

// FormatLine will write the same output as go test -cover does to the out writer.
func FormatLine(out io.Writer, mode string, profiles []Profile) error {
	_, err := fmt.Fprintf(out, "mode: %s\n", mode)
	if err != nil {
		return err
	}
	for _, profile := range profiles {
		for _, block := range profile.Blocks {
			_, err = fmt.Fprintf(out, "%s:%d.%d,%d.%d %d %d\n", profile.FileName,
				block.StartLine, block.StartCol,
				block.EndLine, block.EndCol,
				block.NumStmt,
				block.Count)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
