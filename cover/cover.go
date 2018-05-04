package cover

import "go/token"

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

func NewBlpck(pos *token.Position, numStmt int) *ProfileBlock {
	return &ProfileBlock{StartPosition: pos, NumStmt: numStmt}
}
