package gen

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"strings"

	"github.com/gernest/greact/expr"
	"github.com/gernest/greact/node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const packageName = "greact"
const packageImport = "github.com/gernest/greact"

const (
	newNode  = "createNode"
	newAttr  = "createAttr"
	newAttrs = "createAttrs"
)

// ToNode recursively transform n to a *Node.
func ToNode(n *html.Node) *node.Node {
	nd := &node.Node{
		Type:      node.NodeType(uint32(n.Type)),
		Data:      n.Data,
		Namespace: n.Namespace,
	}
	for _, v := range n.Attr {
		if v.Key == "" {
			continue
		}
		nd.Attr = append(nd.Attr, node.Attribute{
			Namespace: v.Namespace,
			Key:       v.Key,
			Val:       v.Val,
		})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode && strings.TrimSpace(c.Data) == "" {
			continue
		}
		nd.Children = append(nd.Children, ToNode(c))
	}
	return nd
}

// Parse parses src as html component definition and returns their *Node
// representation. r must be reading from a subset of xml/html document that is
// going to processed and compiled to *Node.
func Parse(r io.Reader) (*node.Node, error) {
	base := root()
	n, err := html.ParseFragment(r, base)
	if err != nil {
		return nil, err
	}
	var rst []*node.Node
	for _, v := range n {
		nd := ToNode(v)
		if nd.Type == node.TextNode && strings.TrimSpace(nd.Data) == "" {
			continue
		}
		rst = append(rst, nd)
	}
	container := &node.Node{
		Type: node.ElementNode,
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
func ParseString(s string) (*node.Node, error) {
	return Parse(strings.NewReader(s))
}

// process templates in text nodes
func interpretText(v string) (string, error) {
	parts, err := expr.ExtractExpressions(v, '{', '}')
	if err != nil {
		return "", err
	}
	return expr.WrapString(parts...)

}

// interpret   attributes templates.
func interpret(v interface{}) (string, error) {
	switch e := v.(type) {
	case nil:
		return "nil", nil
	case string:
		exprs, err := expr.ExtractExpressions(e, '{', '}')
		if err != nil {
			return "", err
		}
		return expr.WrapString(exprs...)
	default:
		return "nil", nil
	}
}

func pickExpressions(src string) []expr.Expression {
	return nil
}

// GeneratorContext stores info about the node that we want to generate the
// Render function for.
type GeneratorContext struct {
	// StructName the name of the struct that implements the Component interface.
	StructName string

	// This is a receiver name that will be used for the generated output. This is
	// important because we intend to generate code that the linting tools will be
	// happy with.
	//
	// For instance if StructName is Hello and Recv is h the generated method will
	// have signature like
	//	func (h *Hello)Render
	Recv string

	// The actual node we want to generate go ast for.
	Node *node.Node
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
				importSpec("github.com/gernest/greact"),
				importSpec("github.com/gernest/greact/expr"),
				importSpec("github.com/gernest/greact/node"),
			),
			declareAlias(newNode, "node", "New"),
			declareAlias(newAttr, "node", "Attr"),
			declareAlias(newAttrs, "node", "Attrs"),
			declareAlias("_", "expr", "Eval"),
		},
	}
	for _, v := range ctx {
		e, err := renderNode("Render", v.Recv, v.StructName, v.Node)
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

func importDecl(pkg ...ast.Spec) *ast.GenDecl {
	return &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: pkg,
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

func renderNode(name, recv, typ string, node *node.Node) (*ast.FuncDecl, error) {
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
								Name: packageName,
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
								Name: packageName,
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
									Name: "node",
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
			Name: newAttr,
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
			Name: newAttrs,
		},
		Args: expr,
	}
}

func h(nd *node.Node) (*ast.CallExpr, error) {
	args := []ast.Expr{
		&ast.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprint(uint32(nd.Type)),
		},
		&ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("%q", nd.Namespace),
		},
	}
	if nd.Type == node.TextNode {
		e, err := interpretText(nd.Data)
		if err != nil {
			return nil, err
		}
		x, err := parser.ParseExpr(e)
		if err != nil {
			return nil, err
		}
		args = append(args, x)
	} else {
		args = append(args, &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("%q", nd.Data),
		})
	}
	var attrs []ast.Expr
	for _, v := range nd.Attr {
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
	if len(nd.Children) > 0 {
		for _, v := range nd.Children {
			e, err := h(v)
			if err != nil {
				return nil, err
			}
			args = append(args, e)
		}
	}
	return &ast.CallExpr{
		Fun: &ast.Ident{
			Name: newNode,
		},
		Args: args,
	}, nil
}
