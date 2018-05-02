package color

import (
	"math"
)

var (
	Blue     = New("#1890ff")
	Purple   = New("#722ed1")
	Cyan     = New("#13c2c2")
	Green    = New("#52c41a")
	Magenta  = New("#eb2f96")
	Pink     = New("#eb2f96")
	Red      = New("#f5222d")
	Orange   = New("#fa8c16")
	Yellow   = New("#fadb14")
	Volcano  = New("#fa541c")
	Geekblue = New("#2f54eb")
	Lime     = New("#a0d911")
	Gold     = New("#faad14")
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

func NewPalette() *Palette {
	return &Palette{
		Blue:     Generate(Blue),
		Purple:   Generate(Purple),
		Cyan:     Generate(Cyan),
		Green:    Generate(Green),
		Magenta:  Generate(Magenta),
		Pink:     Generate(Pink),
		Red:      Generate(Red),
		Orange:   Generate(Orange),
		Yellow:   Generate(Yellow),
		Volcano:  Generate(Volcano),
		GeekBlue: Generate(Geekblue),
		Lime:     Generate(Lime),
		Gold:     Generate(Gold),
	}
}

func Generate(base *Color) [10]*Color {
	var c [10]*Color
	for i := 0; i < 10; i++ {
		c[i] = GenerateColor(base, i+1)
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

func GenerateColor(base *Color, index int) *Color {
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
