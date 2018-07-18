package webidl

import (
	"errors"

	"github.com/kr/pretty"

	"github.com/gernest/vected/lib/webidl/ast"
	"github.com/gernest/vected/lib/webidl/scanner"
	"github.com/gernest/vected/lib/webidl/token"
)

func Parse(fs *token.FileSet, filename string, src []byte) (ast.Node, error) {
	tokens, err := collectTokens(fs, filename, src)
	if err != nil {
		return nil, err
	}
	p := parser{tokens: tokens}
	return p.parse()
}

func collectTokens(fs *token.FileSet, filename string, src []byte) ([]*item, error) {
	var s scanner.Scanner
	file := fs.AddFile(filename, fs.Base(), len(src))
	errs := &scanner.ErrorList{}
	s.Init(file, src, func(pos token.Position, msg string) {
		errs.Add(pos, msg)
	}, scanner.ScanComments)
	var tokens []*item
	for {
		pos, tok, lit := s.Scan()
		tokens = append(tokens, &item{
			pos: pos,
			tok: tok,
			lit: lit,
		})
		if tok == token.EOF {
			break
		}
	}
	if s.ErrorCount > 0 {
		return nil, errs
	}
	return tokens, nil
}

type item struct {
	pos token.Pos
	tok token.Token
	lit string
}

type parser struct {
	tokens []*item
	last   *item
	offset int
	file   *ast.File
}

func (p *parser) parse() (ast.Node, error) {
	p.file = &ast.File{}
	for {
		n := p.next()
		if n.tok == token.EOF {
			break
		}
		switch n.tok {
		case token.LBRACK:
			node, err := p.parseExtendedAttributeList()
			if err != nil {
				return nil, err
			}
			p.file.Fragment = append(p.file.Fragment, node)
		}
	}
	return p.file, nil
}

func (p *parser) next() *item {
	var last *item
	if p.offset == 0 && p.last == nil {
		if p.offset > len(p.tokens)-1 {
			return &item{tok: token.EOF}
		}
		last = p.tokens[p.offset]
	} else {
		p.offset++
		if p.offset > len(p.tokens)-1 {
			return &item{tok: token.EOF}
		}
		last = p.tokens[p.offset]
	}
	p.last = last
	return last
}

func (p *parser) rewind() {
	p.offset--
}

func (p *parser) parseExtendedAttributeList() (ast.Node, error) {
	begin := p.last
	if begin.tok != token.LBRACK {
		return nil, errors.New("expected [ got " + begin.tok.String())
	}
	list := &ast.ExtendedAttributeList{}
	list.StartPos = begin.pos
	for {
		n := p.peek()
		if n.tok == token.EOF {
			return nil, errors.New("unexpected EOF ")
		}
		if n.tok == token.RBRACK {
			p.next() //consume the token
			break
		}
		node, err := p.parseExtendedAttr()
		if err != nil {
			return nil, err
		}
		list.List = append(list.List, node)
	}
	return list, nil
}

func (p *parser) parseExtendedAttr() (ast.Node, error) {
	n := p.next()
	if n.tok != token.IDENT {
		return nil, errors.New("expected IDENT " + n.tok.String())
	}
	peek := p.peek()
	switch peek.tok {
	case token.RBRACK:
		p.next()
		return &ast.ExtendedAttributeNoArgs{
			Name:     n.lit,
			Position: ast.Position{StartPos: n.pos},
		}, nil
	}
	return nil, nil
}

func (p *parser) peek() *item {
	if p.offset+1 > len(p.tokens)-1 {
		return &item{tok: token.EOF}
	}
	return p.tokens[p.offset+1]
}

func Dump(v interface{}) string {
	return pretty.Sprint(v)
}
