package complete

import (
	"context"
	"fmt"

	"github.com/blevesearch/bleve"
	"github.com/urfave/cli"
)

type Completer struct {
	htmlIndex bleve.Index
	cssIndex  bleve.Index
}

func New() (*Completer, error) {
	m := bleve.NewIndexMapping()
	hdx, err := bleve.NewMemOnly(m)
	if err != nil {
		return nil, err
	}
	cdx, err := bleve.NewMemOnly(m)
	if err != nil {
		return nil, err
	}
	AddHTMLTags(hdx)
	AddCSSProps(cdx)
	return &Completer{htmlIndex: hdx, cssIndex: cdx}, nil
}

func find(idx bleve.Index, prefix string) (*bleve.SearchResult, error) {
	q := bleve.NewPrefixQuery(prefix)
	sq := bleve.NewSearchRequest(q)
	return idx.Search(sq)
}

func (c *Completer) FindHTML(prefix string) ([]string, error) {
	rs, err := find(c.htmlIndex, prefix)
	if err != nil {
		return nil, err
	}
	var o []string
	for _, hit := range rs.Hits {
		o = append(o, hit.ID)
	}
	return o, nil
}

func (c *Completer) FindCSS(prefix string) ([]string, error) {
	rs, err := find(c.cssIndex, prefix)
	if err != nil {
		return nil, err
	}
	var o []string
	for _, hit := range rs.Hits {
		o = append(o, hit.ID)
	}
	return o, nil
}

func (c *Completer) MultiSearch(prefix string) (matches []string, err error) {
	q := bleve.NewPrefixQuery(prefix)
	sq := bleve.NewSearchRequest(q)
	rs, err := bleve.MultiSearch(context.Background(), sq, c.htmlIndex, c.cssIndex)
	if err != nil {
		return nil, err
	}
	var o []string
	for _, hit := range rs.Hits {
		o = append(o, hit.ID)
	}
	return o, nil
}

func Command() cli.Command {
	return cli.Command{
		Name:  "complete",
		Usage: "autocompletes html and css tags for gss",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name: "css",
			},
			cli.BoolFlag{
				Name: "html",
			},
		},
	}
}

func complete(ctx *cli.Context) error {
	c, err := New()
	if err != nil {
		return err
	}
	a := ctx.Args().First()
	if a == "" {
		return nil
	}
	css := ctx.Bool("css")
	html := ctx.Bool("html")
	if css {
		rs, err := c.FindCSS(a)
		if err != nil {
			return err
		}
		fmt.Println(rs)
		return nil
	}
	if html {
		rs, err := c.FindHTML(a)
		if err != nil {
			return err
		}
		fmt.Println(rs)
		return nil
	}
	rs, err := c.MultiSearch(a)
	if err != nil {
		return err
	}
	fmt.Println(rs)
	return nil
}
