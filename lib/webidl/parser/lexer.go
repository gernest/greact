package lexer

import (
	"io"
	"text/scanner"
)

type Ast struct {
	Nodes []*Node
}

type Position struct {
	Start scanner.Position
	End   scanner.Position
}

type Node struct {
	Position           Position
	ExtendedAttributes *ExtendedAttributeList
}

type ExtendedAttributeList struct {
	Position Position
	Body     []*ExtendedAttribute
}

type AttributeArgument struct {
	Position   Position
	Type       []string
	Name       string
	Identifier *Identifier
}

type Identifier struct {
	Position Position
	Name     string
}

type ExtendedAttribute struct {
	Position Position
	Name     string
	Args     []AttributeArgument
}

type Parser struct {
	scan *scanner.Scanner
	ast  *Ast
}

func (p *Parser) Parse(src io.Reader) (*Ast, error) {
	if p.scan == nil {
		p.scan = &scanner.Scanner{}
	}
	p.scan.Init(src)
	p.ast = &Ast{}
	return p.parse()
}

func (p *Parser) parse() (*Ast, error) {
	return nil, nil
}
