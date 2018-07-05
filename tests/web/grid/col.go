package grid

import (
	"github.com/gernest/mad"
	"github.com/gernest/vected"
	"github.com/gernest/vected/props"
	"github.com/gernest/vected/web/grid"
)

func TestColStyle() mad.Test {
	return mad.It("generates css for grid columns", func(t mad.T) {
	})
}

func TestRenderCol() mad.Test {
	return mad.List{
		mad.It("must render spans with given context", func(t mad.T) {
			r := vected.NewTemplateCache("ui")
			r.Register(grid.Col{})
			v, err := r.RenderHTML(`{Col .}`, props.Props{
				"span": 2,
			})
			if err != nil {
				t.Fatal(err)
			}
			expect := `<div  class=".ant-col-2"></div>`
			g := string(v)
			if g != expect {
				t.Errorf("expected %s got %s", expect, g)
			}
		}),
		mad.It("renders column with extra attribute", func(t mad.T) {
			r := vected.NewTemplateCache("ui")
			r.Register(grid.Col{})
			v, err := r.RenderHTML(`{Col .}`, props.Props{
				"span":    2,
				"onClick": "handleClickEvent",
			})
			if err != nil {
				t.Fatal(err)
			}
			expect := `<div onClick="handleClickEvent" class=".ant-col-2"></div>`
			g := string(v)
			if g != expect {
				t.Errorf("expected %s got %s", expect, g)
			}
		}),
	}
}
