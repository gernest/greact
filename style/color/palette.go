package color

import (
	"github.com/gernest/vected/style/tinycolor"
)

var (
	blue     = tinycolor.InputToRGB("#1890ff")
	purple   = tinycolor.InputToRGB("#722ed1")
	cyan     = tinycolor.InputToRGB("#13c2c2")
	green    = tinycolor.InputToRGB("#52c41a")
	magenta  = tinycolor.InputToRGB("#eb2f96")
	pink     = tinycolor.InputToRGB("#eb2f96")
	red      = tinycolor.InputToRGB("#f5222d")
	orange   = tinycolor.InputToRGB("#fa8c16")
	yellow   = tinycolor.InputToRGB("#fadb14")
	volcano  = tinycolor.InputToRGB("#fa541c")
	geekblue = tinycolor.InputToRGB("#2f54eb")
	lime     = tinycolor.InputToRGB("#a0d911")
	gold     = tinycolor.InputToRGB("#faad14")
)

type Palette struct {
	Blue     [10]*tinycolor.Color
	Purple   [10]*tinycolor.Color
	Cyan     [10]*tinycolor.Color
	Green    [10]*tinycolor.Color
	Magenta  [10]*tinycolor.Color
	Pink     [10]*tinycolor.Color
	Red      [10]*tinycolor.Color
	Orange   [10]*tinycolor.Color
	Yellow   [10]*tinycolor.Color
	Volcano  [10]*tinycolor.Color
	GeekBlue [10]*tinycolor.Color
	Lime     [10]*tinycolor.Color
	Gold     [10]*tinycolor.Color
}

func New() *Palette {
	return &Palette{
		Blue:     Generate(blue),
		Purple:   Generate(purple),
		Cyan:     Generate(cyan),
		Green:    Generate(green),
		Magenta:  Generate(magenta),
		Pink:     Generate(pink),
		Red:      Generate(red),
		Orange:   Generate(orange),
		Yellow:   Generate(yellow),
		Volcano:  Generate(volcano),
		GeekBlue: Generate(geekblue),
		Lime:     Generate(lime),
		Gold:     Generate(gold),
	}
}

func Generate(base *tinycolor.Color) [10]*tinycolor.Color {
	var c [10]*tinycolor.Color
	for i := 0; i < 10; i++ {
		println(i)
		if i == 5 {
			c[i] = base
		} else {
			c[i] = tinycolor.Palette(base, i)
		}
	}
	return c
}
