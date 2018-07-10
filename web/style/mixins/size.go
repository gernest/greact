package mixins

import "github.com/gernest/vected/lib/gs"

func Size(height, width string) gs.CSSRule {
	return gs.CSS(
		gs.P("width", width),
		gs.P("height", height),
	)
}

func Square(size string) gs.CSSRule {
	return Size(size, size)
}
