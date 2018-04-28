package complete

import (
	"github.com/blevesearch/bleve"
)

func find(idx bleve.Index, prefix string) (*bleve.SearchResult, error) {
	q := bleve.NewPrefixQuery(prefix)
	sq := bleve.NewSearchRequest(q)
	return idx.Search(sq)
}
