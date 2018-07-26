// Code generated by vected DO NOT EDIT.
package sample

import (
	vp "github.com/gernest/vected/lib/props"
	"github.com/gernest/vected/lib/vdom"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Render implements vected.Renderer interface.
func (Component) Render(props vp.Props) *vdom.Node {
	return &vdom.Node{
		Children: []*vdom.Node{
			&vdom.Node{
				Type:     html.ElementNode,
				DataAtom: atom.Html,
				Data:     "html",
				Children: []*vdom.Node{
					&vdom.Node{
						Type:     html.ElementNode,
						DataAtom: atom.Head,
						Data:     "head",
					},
					&vdom.Node{
						Type:     html.ElementNode,
						DataAtom: atom.Body,
						Data:     "body",
						Children: []*vdom.Node{
							&vdom.Node{
								Type:     html.ElementNode,
								DataAtom: atom.Div,
								Data:     "div",
								Attr: []html.Attribute{
									{Namespace: "", Key: "name", Val: "value"},
								},
								Children: []*vdom.Node{
									&vdom.Node{
										Type: html.TextNode,
										Data: " hello, world",
									},
								},
							},
						},
					},
				},
			},
		},
	}

}
