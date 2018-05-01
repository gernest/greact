package color

var (
	blue     = New("#1890ff")
	purple   = New("#722ed1")
	cyan     = New("#13c2c2")
	green    = New("#52c41a")
	magenta  = New("#eb2f96")
	pink     = New("#eb2f96")
	red      = New("#f5222d")
	orange   = New("#fa8c16")
	yellow   = New("#fadb14")
	volcano  = New("#fa541c")
	geekblue = New("#2f54eb")
	lime     = New("#a0d911")
	gold     = New("#faad14")
)

type Palette struct {
	Blue     [10]*Color
	Purple   [10]*Color
	Cyan     [10]*Color
	Green    [10]*Color
	Magenta  [10]*Color
	Pink     [10]*Color
	Red      [10]*Color
	Orange   [10]*Color
	Yellow   [10]*Color
	Volcano  [10]*Color
	GeekBlue [10]*Color
	Lime     [10]*Color
	Gold     [10]*Color
}

func NewPaletter() *Palette {
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

func Generate(base *Color) [10]*Color {
	var c [10]*Color
	return c
}
