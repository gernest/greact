package color

import (
	"math"
)

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
	for i := 0; i < 10; i++ {
		c[i] = generate(base, i+1)
	}
	return c
}

const (
	hueStep         = 2
	saturationStep  = 16
	saturationStep2 = 5
	brightnessStep1 = 5
	brightnessStep2 = 15
	lightColorCount = 5
	darkColorCount  = 4
)

func generate(base *Color, index int) *Color {
	isLight := index < 6
	h, s, v, _ := base.HSVA()
	var i int
	if isLight {
		i = lightColorCount + 1 - index
	} else {
		i = index - lightColorCount - 1
	}
	// calculate hue
	var hue float64
	if h >= 60 && h <= 240 {
		if isLight {
			hue = h - hueStep*float64(i)
		} else {
			hue = h + hueStep*float64(i)
		}
	} else {
		if isLight {
			hue = h + hueStep*float64(i)
		} else {
			hue = h - hueStep*float64(i)
		}
	}
	if hue < 0 {
		hue += 360
	} else if hue > 360 {
		hue -= 360
	}
	hue = math.Round(hue)
	// calculate saturation
	var sat float64
	if isLight {
		sat = math.Round(s*100) - float64(saturationStep*i)
	} else if i == darkColorCount {
		sat = math.Round(s*100) + float64(saturationStep)
	} else {
		sat = math.Round(s*100) + float64(saturationStep2*i)
	}
	if sat > 100 {
		sat = 100
	}
	if isLight && i == lightColorCount && sat > 10 {
		sat = 10
	}
	if sat < 6 {
		sat = 6
	}
	// calculate value
	var value float64
	if isLight {
		value = math.Round(v*100) + float64(brightnessStep1*i)
	} else {
		value = math.Round(v*100) - float64(brightnessStep2*i)
	}
	if value > 100 {
		value = 100
	}
	return HSV(hue, sat/100, value/100)
}
