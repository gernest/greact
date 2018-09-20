package expr

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

func Parse(src string) (ast.Expr, error) {
	a, err := parser.ParseExpr(src)
	if err != nil {
		return nil, err
	}
	a = Wrap(a)
	fset := token.NewFileSet()
	ast.Print(fset, a)
	return a, nil
}

func Wrap(args ...ast.Expr) ast.Expr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X: &ast.Ident{
				Name: "fmt",
			},
			Sel: &ast.Ident{
				Name: "Println",
			},
		},
		Args: args,
	}
}

func printExpr(a ast.Expr) {
	printer.Fprint(os.Stdout, token.NewFileSet(), a)
	fmt.Print("\n")
}

// Expression represent part of text that should be evaluated. Something between
// { }
//
// example { 1+2 }
type Expression struct {
	Text  string
	Plain bool
}

// Expr returns ast of the Text field.
func (e Expression) Expr() (ast.Expr, error) {
	return parser.ParseExpr(e.Text)
}

func (e Expression) QuoteExpr() (ast.Expr, error) {
	return parser.ParseExpr(fmt.Sprintf("%q", e.Text))
}

// ExtractExpressions given a src string. This will find text that is within
// begin and end marker returning it as Expression.
//
// For text within begin,end marker Expression.Plain field will be set to false,
// to preserve order of expression, parsing is done from left to right, and any
// other text will be returned in an expression whose Plain field is set to
// true.
//
// Leading and trailing space is trimmed. So, It is not possible to reconstruct
// original text from the returned expressions.
//
// Note that the expression must be valid go expressions.
func ExtractExpressions(src string, begin, end rune) (result []Expression, err error) {
	var buf bytes.Buffer
	line, col, count := 0, 0, 0
	for _, v := range src {
		switch v {
		case begin:
			if buf.Len() > 0 {
				txt := strings.TrimSpace(buf.String())
				if txt != "" {
					result = append(result, Expression{
						Text:  txt,
						Plain: true,
					})
				}
				buf.Reset()
			}
			if count != 0 {
				err = fmt.Errorf("unexpected %s at %d:%d", string(v), line, col)
				return
			}
			count++
			col++
		case end:
			result = append(result, Expression{
				Text: buf.String(),
			})
			buf.Reset()
			count--
			col++
		default:
			switch v {
			case '\n', '\r':
				line++
			default:
				col++
			}
			buf.WriteRune(v)
		}
	}
	if buf.Len() > 0 {
		txt := strings.TrimSpace(buf.String())
		if txt != "" {
			result = append(result, Expression{
				Text:  txt,
				Plain: true,
			})
		}
	}
	return
}
