package complete

import (
	"github.com/blevesearch/bleve"
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
