package ast

import "github.com/gernest/vected/lib/webidl/token"

var (
	_ Node = (*File)(nil)
	_ Node = (*ExtendedAttributeList)(nil)
	_ Node = (*ExtendedAttributeNoArgs)(nil)
)

type Node interface {
	Pos() token.Pos // position of first character belonging to the node
	End() token.Pos // position of first character immediately after the node
}

type File struct {
	Position
	Fragment []Fragment
}

type Fragment interface {
	Node
}

type Position struct {
	StartPos, EndPos token.Pos
}

func (p Position) Pos() token.Pos {
	return p.StartPos
}
func (p Position) End() token.Pos {
	return p.EndPos
}

type ExtendedAttributeList struct {
	Position
	List []Node
}

type ExtendedAttributeNoArgs struct {
	Position
	Name string
}
