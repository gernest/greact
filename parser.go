package vected

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"strings"

	"github.com/gernest/vected/expr"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// ToNode recursively transform n to a *Node.
func ToNode(n *html.Node) *Node {
	node := &Node{
		Type:      NodeType(uint32(n.Type)),
		Data:      n.Data,
		Namespace: n.Namespace,
	}
	for _, v := range n.Attr {
		if v.Key == "" {
			continue
		}
		node.Attr = append(node.Attr, Attribute{
			Namespace: v.Namespace,
			Key:       v.Key,
			Val:       v.Val,
		})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode && strings.TrimSpace(c.Data) == "" {
			continue
		}
		node.Children = append(node.Children, ToNode(c))
	}
	return node
}

// Parse parses src
func Parse(r io.Reader) (*Node, error) {
	base := root()
	n, err := html.ParseFragment(r, base)
	if err != nil {
		return nil, err
	}
	var rst []*Node
	for _, v := range n {
		node := ToNode(v)
		if node.Type == TextNode && strings.TrimSpace(node.Data) == "" {
			continue
		}
		rst = append(rst, node)
	}
	container := &Node{
		Type: ElementNode,
		Data: "div",
	}
	switch len(rst) {
	case 0:
		return container, nil
	case 1:
		return rst[0], nil
	default:
		container.Children = rst
		return container, nil
	}
}

func root() *html.Node {
	return &html.Node{
		DataAtom: atom.Div,
		Type:     html.ElementNode,
		Data:     "div",
	}
}

// ParseString helper that wraps s to io.Reader.
func ParseString(s string) (*Node, error) {
	return Parse(strings.NewReader(s))
}

// process templates in text nodes
func interpretText(v string) (string, error) {
	parts, err := expr.ExtractExpressions(v, '{', '}')
	if err != nil {
		return "", err
	}

	// for text all plain nodes are strings
	var args []ast.Expr
	for _, v := range parts {
		if v.Plain {
			a, err := v.QuoteExpr()
			if err != nil {
				return "", err
			}
			args = append(args, a)
		} else {
			a, err := v.Expr()
			if err != nil {
				return "", err
			}
			args = append(args, a)
		}
	}
	e := expr.Wrap(args...)
	var buf bytes.Buffer
	printer.Fprint(&buf, token.NewFileSet(), e)
	return buf.String(), nil
}

// interpret   attributes templates.
func interpret(v interface{}) (string, error) {
	switch e := v.(type) {
	case nil:
		return "nil", nil
	case string:
		e = strings.TrimSpace(e)
		if strings.HasPrefix(e, "{") {
			parts, err := expr.ExtractExpressions(e, '{', '}')
			if err != nil {
				return "", err
			}
			var args []ast.Expr
			for _, v := range parts {
				if v.Plain && strings.Contains(v.Text, "\"") {
					a, err := v.QuoteExpr()
					if err != nil {
						return "", err
					}
					args = append(args, a)
				} else {
					a, err := v.Expr()
					if err != nil {
						return "", err
					}
					args = append(args, a)
				}
			}
			var e ast.Expr
			if len(args) == 1 {
				e = args[0]
			} else {
				e = expr.Wrap(args...)
			}
			var buf bytes.Buffer
			printer.Fprint(&buf, token.NewFileSet(), e)
			return buf.String(), nil
		}
		return fmt.Sprintf("%q", e), nil
	default:
		return "nil", nil
	}
}

type GeneratorContext struct {
	StructType string
	Receiver   string
	Node       *Node
}

// Generate writes a g file that contains generated Render methods for struct
// defined in the GeneratorContext.
func Generate(w io.Writer, pkg string, ctx ...GeneratorContext) error {
	file := &ast.File{
		Name: &ast.Ident{
			Name: pkg,
		},
		Decls: []ast.Decl{
			importDecl(
				importSpec("context"),
			),
			importDecl(
				importSpec("fmt"),
			),
			importDecl(
				importSpec("github.com/gernest/vected"),
			),
			declareAlias("H", "vected", "NewNode"),
			declareAlias("HA", "vected", "Attr"),
			declareAlias("HAT", "vected", "Attrs"),
		},
	}
	for _, v := range ctx {
		e, err := render("Render", v.Receiver, v.StructType, v.Node)
		if err != nil {
			return err
		}
		file.Decls = append(file.Decls, e)
	}
	return format.Node(w, token.NewFileSet(), file)
}

func importSpec(pkg string) *ast.ImportSpec {
	return &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("%q", pkg),
		},
	}
}

func importDecl(pkg *ast.ImportSpec) *ast.GenDecl {
	return &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{pkg},
	}
}

func declareAlias(alias, pkg, selector string) *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			aliasVar(alias, pkg, selector),
		},
	}
}

func aliasVar(alias, pkg, selector string) *ast.ValueSpec {
	return &ast.ValueSpec{
		Names: []*ast.Ident{
			{
				Name: alias,
			},
		},
		Values: []ast.Expr{
			&ast.SelectorExpr{
				X: &ast.Ident{
					Name: pkg,
				},
				Sel: &ast.Ident{
					Name: selector,
				},
			},
		},
	}
}

func render(name, recv, typ string, node *Node) (*ast.FuncDecl, error) {
	e, err := h(node)
	if err != nil {
		return nil, err
	}
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: recv,
						},
					},
					Type: &ast.StarExpr{
						X: &ast.Ident{
							Name: typ,
						},
					},
				},
			},
		},
		Name: &ast.Ident{
			Name: name,
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							{
								Name: "ctx",
							},
						},
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: "context",
							},
							Sel: &ast.Ident{
								Name: "Context",
							},
						},
					},
					{
						Names: []*ast.Ident{
							{
								Name: "props",
							},
						},
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: "vected",
							},
							Sel: &ast.Ident{
								Name: "Props",
							},
						},
					},
					{
						Names: []*ast.Ident{
							{
								Name: "state",
							},
						},
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: "vected",
							},
							Sel: &ast.Ident{
								Name: "State",
							},
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: "vected",
								},
								Sel: &ast.Ident{
									Name: "Node",
								},
							},
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{e},
				},
			},
		},
	}, nil
}

func ha(ns, key string, val ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.Ident{
			Name: "HA",
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("%q", ns),
			},
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("%q", key),
			},
			val,
		},
	}
}

func hat(expr ...ast.Expr) ast.Expr {
	if len(expr) == 0 {
		return &ast.Ident{
			Name: "nil",
		}
	}
	return &ast.CallExpr{
		Fun: &ast.Ident{
			Name: "HAT",
		},
		Args: expr,
	}
}

func h(node *Node) (*ast.CallExpr, error) {
	args := []ast.Expr{
		&ast.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprint(uint32(node.Type)),
		},
		&ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("%q", node.Namespace),
		},
		&ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("%q", node.Data),
		},
	}
	var attrs []ast.Expr
	for _, v := range node.Attr {
		txt, err := interpret(v.Val)
		if err != nil {
			return nil, err
		}
		e, err := parser.ParseExpr(txt)
		if err != nil {
			return nil, err
		}
		attrs = append(attrs, ha(
			v.Namespace, v.Key, e,
		))
	}
	args = append(args, hat(attrs...))
	if len(node.Children) > 0 {
		for _, v := range node.Children {
			e, err := h(v)
			if err != nil {
				return nil, err
			}
			args = append(args, e)
		}
	}
	return &ast.CallExpr{
		Fun: &ast.Ident{
			Name: "H",
		},
		Args: args,
	}, nil
}
