package media

import (
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/media"
)

type queryList struct {
	listeners map[string]*media.Listener
}

func (q *queryList) AddListener(ls *media.Listener) {
	if q.listeners == nil {
		q.listeners = make(map[string]*media.Listener)
	}
	q.listeners[ls.Name] = ls
}

func (q *queryList) RemoveListener(ls *media.Listener) {
	delete(q.listeners, ls.Name)
}

func (q *queryList) Matches() bool {
	return len(q.listeners) > 0
}

func TestMediaQuery() mad.Test {
	return mad.List{
		mad.It("will add media query listener when constructed", func(t mad.T) {
			ql := &queryList{}
			query := "(max-width: 1000px)"
			m := media.NewMediaQuery(ql, query, false)
			v, ok := ql.listeners[query]
			if !ok {
				t.Fatal("expected listener to be added")
			}
			if v != m.Listener {
				t.Errorf("expected listeners to match")
			}
		}),
		mad.It("will turn on handler when added if query is already matching", func(t mad.T) {
			ql := &queryList{}
			query := "(max-width: 1000px)"
			m := media.NewMediaQuery(ql, query, false)
			called := false
			m.AddHandler(&media.Options{
				Match: func() {
					called = true
				},
			})
			if !called {
				t.Errorf("expected options.Match to be called")
			}
		}),
	}
}
