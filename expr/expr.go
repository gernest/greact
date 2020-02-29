package expr

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strconv"
	"strings"
)

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
	if e.Plain {
		return parser.ParseExpr(strconv.Quote(e.Text))
	}
	return Parse(e.Text)
}

func (e Expression) quoteExpr() (ast.Expr, error) {
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

// Parse returns ast.Expr wtih exp interpreted as function body of a
// func()interface{} {}
// WHich yields the last expression
//
// x:=1
// y:=2
// x+y
//
// will generate
//
// func() interface{} {
// 	x := 1
// 	y := 2
// 	return x + y
// }
func Parse(exp string) (ast.Expr, error) {
	s := `func() interface{}{
		%s
		return 1+1
	}
		`
	s = fmt.Sprintf(s, exp)
	a, err := parser.ParseExpr(s)
	if err != nil {
		return nil, err
	}
	f := a.(*ast.FuncLit)
	n := len(f.Body.List)
	if n > 1 {
		// The last item is the return statement
		ret := f.Body.List[n-1]
		body := f.Body.List[:n-1]
		last := body[len(body)-1]
		stmt := ret.(*ast.ReturnStmt)
		if e, ok := last.(*ast.ExprStmt); ok {
			stmt.Results[0] = e.X
			body[len(body)-1] = stmt
			f.Body.List = body
		}
	}
	return a, nil
}

func wrapExpr(exprs ...Expression) (ast.Expr, error) {
	var args []ast.Expr
	for _, a := range exprs {
		e, err := a.Expr()
		if err != nil {
			return nil, err
		}
		args = append(args, e)
	}
	return wrap(args...), nil
}

func WrapString(exprs ...Expression) (string, error) {
	a, err := wrapExpr(exprs...)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	printer.Fprint(&buf, token.NewFileSet(), a)
	return buf.String(), nil
}

func wrap(args ...ast.Expr) ast.Expr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X: &ast.Ident{
				Name: "expr",
			},
			Sel: &ast.Ident{
				Name: "Eval",
			},
		},
		Args: args,
	}
}

func Eval(args ...interface{}) string {
	var buf bytes.Buffer
	for _, a := range args {
		switch v := a.(type) {
		case string:
			buf.WriteString(v)
		case func() interface{}:
			g := v()
			if g != nil {
				buf.WriteString(toValue(g))
			}
		}
	}
	return buf.String()
}

func toValue(v interface{}) string {
	// TODO remove dependency on fmt
	return fmt.Sprint(v)
}
