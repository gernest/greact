package generate

import "context"
import "fmt"
import "github.com/gernest/greact"

var vH = greact.NewNode
var vHA = greact.Attr
var vHAT = greact.Attrs
var _ = fmt.Print

func (t *Hello) Render(ctx context.Context, props greact.Props, state greact.State) *greact.Node {
	return vH(3, "", "div", vHAT(vHA("", "classname", props["classNames"]), vHA("", "key", "value")))
}
