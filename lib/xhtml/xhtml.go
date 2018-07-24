package xhtml

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/gernest/vected/lib/html/template"
	"golang.org/x/net/html"
)

var globalCache = &NodeCache{nodes: make(map[string]*Node)}

type NodeCache struct {
	nodes map[string]*Node
	mu    sync.RWMutex
}

func (n *NodeCache) Set(name string, v *Node) {
	n.mu.Lock()
	n.nodes[name] = v
	n.mu.RUnlock()
}

func (n *NodeCache) Get(name string) (*Node, bool) {
	n.mu.RLock()
	v, ok := n.nodes[name]
	n.mu.RUnlock()
	return v, ok
}

type Node struct {
	Tree          *html.Node
	TemplateCache map[*html.Node]map[string]*template.Template
}

// Compile  parses src as html.
func Compile(src io.Reader, name string) (*Node, error) {
	name = strings.ToLower(name)
	n, err := compile(src, name)
	if err != nil {
		return nil, err
	}
	globalCache.Set(name, n)
	return n, nil
}

func compile(src io.Reader, name string) (*Node, error) {
	doc, err := html.Parse(src)
	if err != nil {
		return nil, err
	}
	idx := 0
	cache := make(map[*html.Node]map[string]*template.Template)
	var f func(*html.Node) error
	f = func(n *html.Node) error {
		m := make(map[string]*template.Template)
		switch n.Type {
		case html.TextNode:
			if strings.Contains(n.Data, "{") {
				tplName := fmt.Sprintf("__%s__%d__", name, idx)
				t, err := template.New(tplName).Delims("{", "}").Parse(n.Data)
				if err != nil {
					return err
				}
				m[name] = t
				n.Data = name
			}

		}
		cache[n] = m
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			err := f(c)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err = f(doc)
	if err != nil {
		return nil, err
	}
	return &Node{Tree: doc, TemplateCache: cache}, nil
}
