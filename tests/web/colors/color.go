package color

import (
	"reflect"
	"strings"

	"github.com/gernest/mad"
	"github.com/gernest/vected/web/color"
)

func TestNew() mad.Test {
	return mad.It("initialize new color ", func(t mad.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Error(err)
			}
		}()
		s := []struct {
			src    interface{}
			expect color.Color
			err    string
		}{
			{
				src:    []uint8{255, 255, 0},
				expect: color.Color{RGB: []uint8{255, 255, 0}},
			},
			{
				src: 255,
				err: "unsupported type",
			},
			{
				src:    "#fff",
				expect: color.Color{RGB: []uint8{255, 255, 255}},
			},
		}

		for _, v := range s {
			if v.err != "" {
				checkError(t, v.err, func() {
					color.New(v.src)
				})
			} else {
				g := color.New(v.src)
				compareColors(t, g, &v.expect)
			}
		}
	})

}

func checkError(t mad.T, msg string, fn func()) {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			if !strings.Contains(e.Error(), msg) {
				t.Errorf("expected %v to contain %v", e, msg)
			}
		} else {
			t.Error("expected an error")
		}
	}()
	fn()
}

func compareColors(t mad.T, a, b *color.Color) {
	if !reflect.DeepEqual(a.RGB, b.RGB) {
		t.Errorf("rgb: expected %v to equal %v", a.RGB, b.RGB)
	}
}

func TestColor_Hex() mad.Test {
	return mad.It("must generate correct hex color", func(t mad.T) {
		for k, v := range color.CommonColors {
			n := color.New(v)
			g := n.Hex()
			if g != v {
				t.Errorf("%s: expected %s got %s", k, v, g)
			}
		}
	})
}

func TestPrintColors() mad.Test {
	sample := []struct {
		name, src, hex, rgb, hsl, hsv string
	}{

		{name: "aliceblue", src: "#f0f8ff", hex: "#f0f8ff", rgb: "rgb(240,248,255)", hsl: "hsl(208,100%,97%)", hsv: "hsv(208,6%,100%)"},
		{name: "antiquewhite", src: "#faebd7", hex: "#faebd7", rgb: "rgb(250,235,215)", hsl: "hsl(34,78%,91%)", hsv: "hsv(34,14%,98%)"},
		{name: "aqua", src: "#00ffff", hex: "#00ffff", rgb: "rgb(0,255,255)", hsl: "hsl(180,100%,50%)", hsv: "hsv(180,100%,100%)"},
		{name: "aquamarine", src: "#7fffd4", hex: "#7fffd4", rgb: "rgb(127,255,212)", hsl: "hsl(159,100%,75%)", hsv: "hsv(159,50%,100%)"},
		{name: "azure", src: "#f0ffff", hex: "#f0ffff", rgb: "rgb(240,255,255)", hsl: "hsl(180,100%,97%)", hsv: "hsv(180,6%,100%)"},
		{name: "beige", src: "#f5f5dc", hex: "#f5f5dc", rgb: "rgb(245,245,220)", hsl: "hsl(60,56%,91%)", hsv: "hsv(60,10%,96%)"},
		{name: "bisque", src: "#ffe4c4", hex: "#ffe4c4", rgb: "rgb(255,228,196)", hsl: "hsl(32,100%,88%)", hsv: "hsv(32,23%,100%)"},
		{name: "black", src: "#000000", hex: "#000000", rgb: "rgb(0,0,0)", hsl: "hsl(0,0%,0%)", hsv: "hsv(0,0%,0%)"},
		{name: "blanchedalmond", src: "#ffebcd", hex: "#ffebcd", rgb: "rgb(255,235,205)", hsl: "hsl(35,100%,90%)", hsv: "hsv(35,20%,100%)"},
		{name: "blue", src: "#0000ff", hex: "#0000ff", rgb: "rgb(0,0,255)", hsl: "hsl(240,100%,50%)", hsv: "hsv(240,100%,100%)"},
		{name: "blueviolet", src: "#8a2be2", hex: "#8a2be2", rgb: "rgb(138,43,226)", hsl: "hsl(271,76%,53%)", hsv: "hsv(271,81%,89%)"},
		{name: "brown", src: "#a52a2a", hex: "#a52a2a", rgb: "rgb(165,42,42)", hsl: "hsl(0,59%,41%)", hsv: "hsv(0,75%,65%)"},
		{name: "burlywood", src: "#deb887", hex: "#deb887", rgb: "rgb(222,184,135)", hsl: "hsl(33,57%,70%)", hsv: "hsv(33,39%,87%)"},
		{name: "cadetblue", src: "#5f9ea0", hex: "#5f9ea0", rgb: "rgb(95,158,160)", hsl: "hsl(181,25%,50%)", hsv: "hsv(181,41%,63%)"},
		{name: "chartreuse", src: "#7fff00", hex: "#7fff00", rgb: "rgb(127,255,0)", hsl: "hsl(90,100%,50%)", hsv: "hsv(90,100%,100%)"},
		{name: "chocolate", src: "#d2691e", hex: "#d2691e", rgb: "rgb(210,105,30)", hsl: "hsl(24,75%,47%)", hsv: "hsv(24,86%,82%)"},
		{name: "coral", src: "#ff7f50", hex: "#ff7f50", rgb: "rgb(255,127,80)", hsl: "hsl(16,100%,66%)", hsv: "hsv(16,69%,100%)"},
		{name: "cornflowerblue", src: "#6495ed", hex: "#6495ed", rgb: "rgb(100,149,237)", hsl: "hsl(218,79%,66%)", hsv: "hsv(218,58%,93%)"},
		{name: "cornsilk", src: "#fff8dc", hex: "#fff8dc", rgb: "rgb(255,248,220)", hsl: "hsl(47,100%,93%)", hsv: "hsv(47,14%,100%)"},
		{name: "crimson", src: "#dc143c", hex: "#dc143c", rgb: "rgb(220,20,60)", hsl: "hsl(348,83%,47%)", hsv: "hsv(348,91%,86%)"},
		{name: "cyan", src: "#00ffff", hex: "#00ffff", rgb: "rgb(0,255,255)", hsl: "hsl(180,100%,50%)", hsv: "hsv(180,100%,100%)"},
		{name: "darkblue", src: "#00008b", hex: "#00008b", rgb: "rgb(0,0,139)", hsl: "hsl(240,100%,27%)", hsv: "hsv(240,100%,55%)"},
		{name: "darkcyan", src: "#008b8b", hex: "#008b8b", rgb: "rgb(0,139,139)", hsl: "hsl(180,100%,27%)", hsv: "hsv(180,100%,55%)"},
		{name: "darkgoldenrod", src: "#b8860b", hex: "#b8860b", rgb: "rgb(184,134,11)", hsl: "hsl(42,89%,38%)", hsv: "hsv(42,94%,72%)"},
		{name: "darkgray", src: "#a9a9a9", hex: "#a9a9a9", rgb: "rgb(169,169,169)", hsl: "hsl(0,0%,66%)", hsv: "hsv(0,0%,66%)"},
		{name: "darkgreen", src: "#006400", hex: "#006400", rgb: "rgb(0,100,0)", hsl: "hsl(120,100%,20%)", hsv: "hsv(120,100%,39%)"},
		{name: "darkgrey", src: "#a9a9a9", hex: "#a9a9a9", rgb: "rgb(169,169,169)", hsl: "hsl(0,0%,66%)", hsv: "hsv(0,0%,66%)"},
		{name: "darkkhaki", src: "#bdb76b", hex: "#bdb76b", rgb: "rgb(189,183,107)", hsl: "hsl(55,38%,58%)", hsv: "hsv(55,43%,74%)"},
		{name: "darkmagenta", src: "#8b008b", hex: "#8b008b", rgb: "rgb(139,0,139)", hsl: "hsl(300,100%,27%)", hsv: "hsv(300,100%,55%)"},
		{name: "darkolivegreen", src: "#556b2f", hex: "#556b2f", rgb: "rgb(85,107,47)", hsl: "hsl(82,39%,30%)", hsv: "hsv(82,56%,42%)"},
		{name: "darkorange", src: "#ff8c00", hex: "#ff8c00", rgb: "rgb(255,140,0)", hsl: "hsl(32,100%,50%)", hsv: "hsv(32,100%,100%)"},
		{name: "darkorchid", src: "#9932cc", hex: "#9932cc", rgb: "rgb(153,50,204)", hsl: "hsl(280,61%,50%)", hsv: "hsv(280,75%,80%)"},
		{name: "darkred", src: "#8b0000", hex: "#8b0000", rgb: "rgb(139,0,0)", hsl: "hsl(0,100%,27%)", hsv: "hsv(0,100%,55%)"},
		{name: "darksalmon", src: "#e9967a", hex: "#e9967a", rgb: "rgb(233,150,122)", hsl: "hsl(15,72%,70%)", hsv: "hsv(15,48%,91%)"},
		{name: "darkseagreen", src: "#8fbc8f", hex: "#8fbc8f", rgb: "rgb(143,188,143)", hsl: "hsl(120,25%,65%)", hsv: "hsv(120,24%,74%)"},
		{name: "darkslateblue", src: "#483d8b", hex: "#483d8b", rgb: "rgb(72,61,139)", hsl: "hsl(248,39%,39%)", hsv: "hsv(248,56%,55%)"},
		{name: "darkslategray", src: "#2f4f4f", hex: "#2f4f4f", rgb: "rgb(47,79,79)", hsl: "hsl(180,25%,25%)", hsv: "hsv(180,41%,31%)"},
		{name: "darkslategrey", src: "#2f4f4f", hex: "#2f4f4f", rgb: "rgb(47,79,79)", hsl: "hsl(180,25%,25%)", hsv: "hsv(180,41%,31%)"},
		{name: "darkturquoise", src: "#00ced1", hex: "#00ced1", rgb: "rgb(0,206,209)", hsl: "hsl(180,100%,41%)", hsv: "hsv(180,100%,82%)"},
		{name: "darkviolet", src: "#9400d3", hex: "#9400d3", rgb: "rgb(148,0,211)", hsl: "hsl(282,100%,41%)", hsv: "hsv(282,100%,83%)"},
		{name: "deeppink", src: "#ff1493", hex: "#ff1493", rgb: "rgb(255,20,147)", hsl: "hsl(327,100%,54%)", hsv: "hsv(327,92%,100%)"},
		{name: "deepskyblue", src: "#00bfff", hex: "#00bfff", rgb: "rgb(0,191,255)", hsl: "hsl(195,100%,50%)", hsv: "hsv(195,100%,100%)"},
		{name: "dimgray", src: "#696969", hex: "#696969", rgb: "rgb(105,105,105)", hsl: "hsl(0,0%,41%)", hsv: "hsv(0,0%,41%)"},
		{name: "dimgrey", src: "#696969", hex: "#696969", rgb: "rgb(105,105,105)", hsl: "hsl(0,0%,41%)", hsv: "hsv(0,0%,41%)"},
		{name: "dodgerblue", src: "#1e90ff", hex: "#1e90ff", rgb: "rgb(30,144,255)", hsl: "hsl(209,100%,56%)", hsv: "hsv(209,88%,100%)"},
		{name: "firebrick", src: "#b22222", hex: "#b22222", rgb: "rgb(178,34,34)", hsl: "hsl(0,68%,42%)", hsv: "hsv(0,81%,70%)"},
		{name: "floralwhite", src: "#fffaf0", hex: "#fffaf0", rgb: "rgb(255,250,240)", hsl: "hsl(39,100%,97%)", hsv: "hsv(39,6%,100%)"},
		{name: "forestgreen", src: "#228b22", hex: "#228b22", rgb: "rgb(34,139,34)", hsl: "hsl(120,61%,34%)", hsv: "hsv(120,76%,55%)"},
		{name: "fuchsia", src: "#ff00ff", hex: "#ff00ff", rgb: "rgb(255,0,255)", hsl: "hsl(300,100%,50%)", hsv: "hsv(300,100%,100%)"},
		{name: "gainsboro", src: "#dcdcdc", hex: "#dcdcdc", rgb: "rgb(220,220,220)", hsl: "hsl(0,0%,86%)", hsv: "hsv(0,0%,86%)"},
		{name: "ghostwhite", src: "#f8f8ff", hex: "#f8f8ff", rgb: "rgb(248,248,255)", hsl: "hsl(240,100%,99%)", hsv: "hsv(240,3%,100%)"},
		{name: "gold", src: "#ffd700", hex: "#ffd700", rgb: "rgb(255,215,0)", hsl: "hsl(50,100%,50%)", hsv: "hsv(50,100%,100%)"},
		{name: "goldenrod", src: "#daa520", hex: "#daa520", rgb: "rgb(218,165,32)", hsl: "hsl(42,74%,49%)", hsv: "hsv(42,85%,85%)"},
		{name: "gray", src: "#808080", hex: "#808080", rgb: "rgb(128,128,128)", hsl: "hsl(0,0%,50%)", hsv: "hsv(0,0%,50%)"},
		{name: "green", src: "#008000", hex: "#008000", rgb: "rgb(0,128,0)", hsl: "hsl(120,100%,25%)", hsv: "hsv(120,100%,50%)"},
		{name: "greenyellow", src: "#adff2f", hex: "#adff2f", rgb: "rgb(173,255,47)", hsl: "hsl(83,100%,59%)", hsv: "hsv(83,82%,100%)"},
		{name: "grey", src: "#808080", hex: "#808080", rgb: "rgb(128,128,128)", hsl: "hsl(0,0%,50%)", hsv: "hsv(0,0%,50%)"},
		{name: "honeydew", src: "#f0fff0", hex: "#f0fff0", rgb: "rgb(240,255,240)", hsl: "hsl(120,100%,97%)", hsv: "hsv(120,6%,100%)"},
		{name: "hotpink", src: "#ff69b4", hex: "#ff69b4", rgb: "rgb(255,105,180)", hsl: "hsl(330,100%,71%)", hsv: "hsv(330,59%,100%)"},
		{name: "indianred", src: "#cd5c5c", hex: "#cd5c5c", rgb: "rgb(205,92,92)", hsl: "hsl(0,53%,58%)", hsv: "hsv(0,55%,80%)"},
		{name: "indigo", src: "#4b0082", hex: "#4b0082", rgb: "rgb(75,0,130)", hsl: "hsl(274,100%,25%)", hsv: "hsv(274,100%,51%)"},
		{name: "ivory", src: "#fffff0", hex: "#fffff0", rgb: "rgb(255,255,240)", hsl: "hsl(60,100%,97%)", hsv: "hsv(60,6%,100%)"},
		{name: "khaki", src: "#f0e68c", hex: "#f0e68c", rgb: "rgb(240,230,140)", hsl: "hsl(54,77%,75%)", hsv: "hsv(54,42%,94%)"},
		{name: "lavender", src: "#e6e6fa", hex: "#e6e6fa", rgb: "rgb(230,230,250)", hsl: "hsl(240,67%,94%)", hsv: "hsv(240,8%,98%)"},
		{name: "lavenderblush", src: "#fff0f5", hex: "#fff0f5", rgb: "rgb(255,240,245)", hsl: "hsl(339,100%,97%)", hsv: "hsv(339,6%,100%)"},
		{name: "lawngreen", src: "#7cfc00", hex: "#7cfc00", rgb: "rgb(124,252,0)", hsl: "hsl(90,100%,49%)", hsv: "hsv(90,100%,99%)"},
		{name: "lemonchiffon", src: "#fffacd", hex: "#fffacd", rgb: "rgb(255,250,205)", hsl: "hsl(53,100%,90%)", hsv: "hsv(53,20%,100%)"},
		{name: "lightblue", src: "#add8e6", hex: "#add8e6", rgb: "rgb(173,216,230)", hsl: "hsl(194,53%,79%)", hsv: "hsv(194,25%,90%)"},
		{name: "lightcoral", src: "#f08080", hex: "#f08080", rgb: "rgb(240,128,128)", hsl: "hsl(0,79%,72%)", hsv: "hsv(0,47%,94%)"},
		{name: "lightcyan", src: "#e0ffff", hex: "#e0ffff", rgb: "rgb(224,255,255)", hsl: "hsl(180,100%,94%)", hsv: "hsv(180,12%,100%)"},
		{name: "lightgoldenrodyellow", src: "#fafad2", hex: "#fafad2", rgb: "rgb(250,250,210)", hsl: "hsl(60,80%,90%)", hsv: "hsv(60,16%,98%)"},
		{name: "lightgray", src: "#d3d3d3", hex: "#d3d3d3", rgb: "rgb(211,211,211)", hsl: "hsl(0,0%,83%)", hsv: "hsv(0,0%,83%)"},
		{name: "lightgreen", src: "#90ee90", hex: "#90ee90", rgb: "rgb(144,238,144)", hsl: "hsl(120,73%,75%)", hsv: "hsv(120,39%,93%)"},
		{name: "lightgrey", src: "#d3d3d3", hex: "#d3d3d3", rgb: "rgb(211,211,211)", hsl: "hsl(0,0%,83%)", hsv: "hsv(0,0%,83%)"},
		{name: "lightpink", src: "#ffb6c1", hex: "#ffb6c1", rgb: "rgb(255,182,193)", hsl: "hsl(350,100%,86%)", hsv: "hsv(350,29%,100%)"},
		{name: "lightsalmon", src: "#ffa07a", hex: "#ffa07a", rgb: "rgb(255,160,122)", hsl: "hsl(17,100%,74%)", hsv: "hsv(17,52%,100%)"},
		{name: "lightseagreen", src: "#20b2aa", hex: "#20b2aa", rgb: "rgb(32,178,170)", hsl: "hsl(176,70%,41%)", hsv: "hsv(176,82%,70%)"},
		{name: "lightskyblue", src: "#87cefa", hex: "#87cefa", rgb: "rgb(135,206,250)", hsl: "hsl(202,92%,75%)", hsv: "hsv(202,46%,98%)"},
		{name: "lightslategray", src: "#778899", hex: "#778899", rgb: "rgb(119,136,153)", hsl: "hsl(210,14%,53%)", hsv: "hsv(210,22%,60%)"},
		{name: "lightslategrey", src: "#778899", hex: "#778899", rgb: "rgb(119,136,153)", hsl: "hsl(210,14%,53%)", hsv: "hsv(210,22%,60%)"},
		{name: "lightsteelblue", src: "#b0c4de", hex: "#b0c4de", rgb: "rgb(176,196,222)", hsl: "hsl(213,41%,78%)", hsv: "hsv(213,21%,87%)"},
		{name: "lightyellow", src: "#ffffe0", hex: "#ffffe0", rgb: "rgb(255,255,224)", hsl: "hsl(60,100%,94%)", hsv: "hsv(60,12%,100%)"},
		{name: "lime", src: "#00ff00", hex: "#00ff00", rgb: "rgb(0,255,0)", hsl: "hsl(120,100%,50%)", hsv: "hsv(120,100%,100%)"},
		{name: "limegreen", src: "#32cd32", hex: "#32cd32", rgb: "rgb(50,205,50)", hsl: "hsl(120,61%,50%)", hsv: "hsv(120,76%,80%)"},
		{name: "linen", src: "#faf0e6", hex: "#faf0e6", rgb: "rgb(250,240,230)", hsl: "hsl(30,67%,94%)", hsv: "hsv(30,8%,98%)"},
		{name: "magenta", src: "#ff00ff", hex: "#ff00ff", rgb: "rgb(255,0,255)", hsl: "hsl(300,100%,50%)", hsv: "hsv(300,100%,100%)"},
		{name: "maroon", src: "#800000", hex: "#800000", rgb: "rgb(128,0,0)", hsl: "hsl(0,100%,25%)", hsv: "hsv(0,100%,50%)"},
		{name: "mediumaquamarine", src: "#66cdaa", hex: "#66cdaa", rgb: "rgb(102,205,170)", hsl: "hsl(159,51%,60%)", hsv: "hsv(159,50%,80%)"},
		{name: "mediumblue", src: "#0000cd", hex: "#0000cd", rgb: "rgb(0,0,205)", hsl: "hsl(240,100%,40%)", hsv: "hsv(240,100%,80%)"},
		{name: "mediumorchid", src: "#ba55d3", hex: "#ba55d3", rgb: "rgb(186,85,211)", hsl: "hsl(288,59%,58%)", hsv: "hsv(288,60%,83%)"},
		{name: "mediumpurple", src: "#9370d8", hex: "#9370d8", rgb: "rgb(147,112,216)", hsl: "hsl(260,57%,64%)", hsv: "hsv(260,48%,85%)"},
		{name: "mediumseagreen", src: "#3cb371", hex: "#3cb371", rgb: "rgb(60,179,113)", hsl: "hsl(146,50%,47%)", hsv: "hsv(146,66%,70%)"},
		{name: "mediumslateblue", src: "#7b68ee", hex: "#7b68ee", rgb: "rgb(123,104,238)", hsl: "hsl(248,80%,67%)", hsv: "hsv(248,56%,93%)"},
		{name: "mediumspringgreen", src: "#00fa9a", hex: "#00fa9a", rgb: "rgb(0,250,154)", hsl: "hsl(156,100%,49%)", hsv: "hsv(156,100%,98%)"},
		{name: "mediumturquoise", src: "#48d1cc", hex: "#48d1cc", rgb: "rgb(72,209,204)", hsl: "hsl(177,60%,55%)", hsv: "hsv(177,66%,82%)"},
		{name: "mediumvioletred", src: "#c71585", hex: "#c71585", rgb: "rgb(199,21,133)", hsl: "hsl(322,81%,43%)", hsv: "hsv(322,89%,78%)"},
		{name: "midnightblue", src: "#191970", hex: "#191970", rgb: "rgb(25,25,112)", hsl: "hsl(240,64%,27%)", hsv: "hsv(240,78%,44%)"},
		{name: "mintcream", src: "#f5fffa", hex: "#f5fffa", rgb: "rgb(245,255,250)", hsl: "hsl(149,100%,98%)", hsv: "hsv(149,4%,100%)"},
		{name: "mistyrose", src: "#ffe4e1", hex: "#ffe4e1", rgb: "rgb(255,228,225)", hsl: "hsl(6,100%,94%)", hsv: "hsv(6,12%,100%)"},
		{name: "moccasin", src: "#ffe4b5", hex: "#ffe4b5", rgb: "rgb(255,228,181)", hsl: "hsl(38,100%,85%)", hsv: "hsv(38,29%,100%)"},
		{name: "navajowhite", src: "#ffdead", hex: "#ffdead", rgb: "rgb(255,222,173)", hsl: "hsl(35,100%,84%)", hsv: "hsv(35,32%,100%)"},
		{name: "navy", src: "#000080", hex: "#000080", rgb: "rgb(0,0,128)", hsl: "hsl(240,100%,25%)", hsv: "hsv(240,100%,50%)"},
		{name: "oldlace", src: "#fdf5e6", hex: "#fdf5e6", rgb: "rgb(253,245,230)", hsl: "hsl(39,85%,95%)", hsv: "hsv(39,9%,99%)"},
		{name: "olive", src: "#808000", hex: "#808000", rgb: "rgb(128,128,0)", hsl: "hsl(60,100%,25%)", hsv: "hsv(60,100%,50%)"},
		{name: "olivedrab", src: "#6b8e23", hex: "#6b8e23", rgb: "rgb(107,142,35)", hsl: "hsl(79,60%,35%)", hsv: "hsv(79,75%,56%)"},
		{name: "orange", src: "#ffa500", hex: "#ffa500", rgb: "rgb(255,165,0)", hsl: "hsl(38,100%,50%)", hsv: "hsv(38,100%,100%)"},
		{name: "orangered", src: "#ff4500", hex: "#ff4500", rgb: "rgb(255,69,0)", hsl: "hsl(16,100%,50%)", hsv: "hsv(16,100%,100%)"},
		{name: "orchid", src: "#da70d6", hex: "#da70d6", rgb: "rgb(218,112,214)", hsl: "hsl(302,59%,65%)", hsv: "hsv(302,49%,85%)"},
		{name: "palegoldenrod", src: "#eee8aa", hex: "#eee8aa", rgb: "rgb(238,232,170)", hsl: "hsl(54,67%,80%)", hsv: "hsv(54,29%,93%)"},
		{name: "palegreen", src: "#98fb98", hex: "#98fb98", rgb: "rgb(152,251,152)", hsl: "hsl(120,93%,79%)", hsv: "hsv(120,39%,98%)"},
		{name: "paleturquoise", src: "#afeeee", hex: "#afeeee", rgb: "rgb(175,238,238)", hsl: "hsl(180,65%,81%)", hsv: "hsv(180,26%,93%)"},
		{name: "palevioletred", src: "#d87093", hex: "#d87093", rgb: "rgb(216,112,147)", hsl: "hsl(339,57%,64%)", hsv: "hsv(339,48%,85%)"},
		{name: "papayawhip", src: "#ffefd5", hex: "#ffefd5", rgb: "rgb(255,239,213)", hsl: "hsl(37,100%,92%)", hsv: "hsv(37,16%,100%)"},
		{name: "peachpuff", src: "#ffdab9", hex: "#ffdab9", rgb: "rgb(255,218,185)", hsl: "hsl(28,100%,86%)", hsv: "hsv(28,27%,100%)"},
		{name: "peru", src: "#cd853f", hex: "#cd853f", rgb: "rgb(205,133,63)", hsl: "hsl(29,59%,53%)", hsv: "hsv(29,69%,80%)"},
		{name: "pink", src: "#ffc0cb", hex: "#ffc0cb", rgb: "rgb(255,192,203)", hsl: "hsl(349,100%,88%)", hsv: "hsv(349,25%,100%)"},
		{name: "plum", src: "#dda0dd", hex: "#dda0dd", rgb: "rgb(221,160,221)", hsl: "hsl(300,47%,75%)", hsv: "hsv(300,28%,87%)"},
		{name: "powderblue", src: "#b0e0e6", hex: "#b0e0e6", rgb: "rgb(176,224,230)", hsl: "hsl(186,52%,80%)", hsv: "hsv(186,23%,90%)"},
		{name: "purple", src: "#800080", hex: "#800080", rgb: "rgb(128,0,128)", hsl: "hsl(300,100%,25%)", hsv: "hsv(300,100%,50%)"},
		{name: "rebeccapurple", src: "#663399", hex: "#663399", rgb: "rgb(102,51,153)", hsl: "hsl(270,50%,40%)", hsv: "hsv(270,67%,60%)"},
		{name: "red", src: "#ff0000", hex: "#ff0000", rgb: "rgb(255,0,0)", hsl: "hsl(0,100%,50%)", hsv: "hsv(0,100%,100%)"},
		{name: "rosybrown", src: "#bc8f8f", hex: "#bc8f8f", rgb: "rgb(188,143,143)", hsl: "hsl(0,25%,65%)", hsv: "hsv(0,24%,74%)"},
		{name: "royalblue", src: "#4169e1", hex: "#4169e1", rgb: "rgb(65,105,225)", hsl: "hsl(225,73%,57%)", hsv: "hsv(225,71%,88%)"},
		{name: "saddlebrown", src: "#8b4513", hex: "#8b4513", rgb: "rgb(139,69,19)", hsl: "hsl(24,76%,31%)", hsv: "hsv(24,86%,55%)"},
		{name: "salmon", src: "#fa8072", hex: "#fa8072", rgb: "rgb(250,128,114)", hsl: "hsl(6,93%,71%)", hsv: "hsv(6,54%,98%)"},
		{name: "sandybrown", src: "#f4a460", hex: "#f4a460", rgb: "rgb(244,164,96)", hsl: "hsl(27,87%,67%)", hsv: "hsv(27,61%,96%)"},
		{name: "seagreen", src: "#2e8b57", hex: "#2e8b57", rgb: "rgb(46,139,87)", hsl: "hsl(146,50%,36%)", hsv: "hsv(146,67%,55%)"},
		{name: "seashell", src: "#fff5ee", hex: "#fff5ee", rgb: "rgb(255,245,238)", hsl: "hsl(24,100%,97%)", hsv: "hsv(24,7%,100%)"},
		{name: "sienna", src: "#a0522d", hex: "#a0522d", rgb: "rgb(160,82,45)", hsl: "hsl(19,56%,40%)", hsv: "hsv(19,72%,63%)"},
		{name: "silver", src: "#c0c0c0", hex: "#c0c0c0", rgb: "rgb(192,192,192)", hsl: "hsl(0,0%,75%)", hsv: "hsv(0,0%,75%)"},
		{name: "skyblue", src: "#87ceeb", hex: "#87ceeb", rgb: "rgb(135,206,235)", hsl: "hsl(197,71%,73%)", hsv: "hsv(197,43%,92%)"},
		{name: "slateblue", src: "#6a5acd", hex: "#6a5acd", rgb: "rgb(106,90,205)", hsl: "hsl(248,53%,58%)", hsv: "hsv(248,56%,80%)"},
		{name: "slategray", src: "#708090", hex: "#708090", rgb: "rgb(112,128,144)", hsl: "hsl(210,13%,50%)", hsv: "hsv(210,22%,56%)"},
		{name: "slategrey", src: "#708090", hex: "#708090", rgb: "rgb(112,128,144)", hsl: "hsl(210,13%,50%)", hsv: "hsv(210,22%,56%)"},
		{name: "snow", src: "#fffafa", hex: "#fffafa", rgb: "rgb(255,250,250)", hsl: "hsl(0,100%,99%)", hsv: "hsv(0,2%,100%)"},
		{name: "springgreen", src: "#00ff7f", hex: "#00ff7f", rgb: "rgb(0,255,127)", hsl: "hsl(149,100%,50%)", hsv: "hsv(149,100%,100%)"},
		{name: "steelblue", src: "#4682b4", hex: "#4682b4", rgb: "rgb(70,130,180)", hsl: "hsl(207,44%,49%)", hsv: "hsv(207,61%,71%)"},
		{name: "tan", src: "#d2b48c", hex: "#d2b48c", rgb: "rgb(210,180,140)", hsl: "hsl(34,44%,69%)", hsv: "hsv(34,33%,82%)"},
		{name: "teal", src: "#008080", hex: "#008080", rgb: "rgb(0,128,128)", hsl: "hsl(180,100%,25%)", hsv: "hsv(180,100%,50%)"},
		{name: "thistle", src: "#d8bfd8", hex: "#d8bfd8", rgb: "rgb(216,191,216)", hsl: "hsl(300,24%,80%)", hsv: "hsv(300,12%,85%)"},
		{name: "tomato", src: "#ff6347", hex: "#ff6347", rgb: "rgb(255,99,71)", hsl: "hsl(9,100%,64%)", hsv: "hsv(9,72%,100%)"},
		{name: "turquoise", src: "#40e0d0", hex: "#40e0d0", rgb: "rgb(64,224,208)", hsl: "hsl(174,72%,56%)", hsv: "hsv(174,71%,88%)"},
		{name: "violet", src: "#ee82ee", hex: "#ee82ee", rgb: "rgb(238,130,238)", hsl: "hsl(300,76%,72%)", hsv: "hsv(300,45%,93%)"},
		{name: "wheat", src: "#f5deb3", hex: "#f5deb3", rgb: "rgb(245,222,179)", hsl: "hsl(39,77%,83%)", hsv: "hsv(39,27%,96%)"},
		{name: "white", src: "#ffffff", hex: "#ffffff", rgb: "rgb(255,255,255)", hsl: "hsl(0,0%,100%)", hsv: "hsv(0,0%,100%)"},
		{name: "whitesmoke", src: "#f5f5f5", hex: "#f5f5f5", rgb: "rgb(245,245,245)", hsl: "hsl(0,0%,96%)", hsv: "hsv(0,0%,96%)"},
		{name: "yellow", src: "#ffff00", hex: "#ffff00", rgb: "rgb(255,255,0)", hsl: "hsl(60,100%,50%)", hsv: "hsv(60,100%,100%)"},
		{name: "yellowgreen", src: "#9acd32", hex: "#9acd32", rgb: "rgb(154,205,50)", hsl: "hsl(79,61%,50%)", hsv: "hsv(79,76%,80%)"},
	}
	return mad.Describe("PrintColor",
		mad.It("prints hex", func(t mad.T) {
			for _, v := range sample {
				c := color.New(v.src)
				hex := color.PrintColor(c, "hex")
				if hex != v.hex {
					t.Errorf("%s: hex expected %s got %s", v.name, v.hex, hex)
				}
				rgb := color.PrintColor(c, "rgb")
				if rgb != v.rgb {
					t.Errorf("%s: rgb expected %s got %s", v.name, v.rgb, rgb)
				}
				hsl := color.PrintColor(c, "hsl")
				if hsl != v.hsl {
					t.Errorf("%s: hsl expected %s got %s", v.name, v.hsl, hsl)
				}
				hsv := color.PrintColor(c, "hsv")
				if hsv != v.hsv {
					t.Errorf("%s: hsv expected %s got %s", v.name, v.hsv, hsv)
				}
			}
		}),
		mad.It("prints rgb", func(t mad.T) {
			for _, v := range sample {
				c := color.New(v.rgb)
				hex := color.PrintColor(c, "hex")
				if hex != v.hex {
					t.Errorf("%s: hex expected %s got %s", v.name, v.hex, hex)
				}
				rgb := color.PrintColor(c, "rgb")
				if rgb != v.rgb {
					t.Errorf("%s: rgb expected %s got %s", v.name, v.rgb, rgb)
				}
				hsl := color.PrintColor(c, "hsl")
				if hsl != v.hsl {
					t.Errorf("%s: hsl expected %s got %s", v.name, v.hsl, hsl)
				}
				hsv := color.PrintColor(c, "hsv")
				if hsv != v.hsv {
					t.Errorf("%s: hsv expected %s got %s", v.name, v.hsv, hsv)
				}
			}
		}),
		mad.It("prints HSLA", func(t mad.T) {
			for _, v := range sample {
				c := color.HSLA(color.New(v.rgb).HSLA())
				hex := color.PrintColor(c, "hex")
				if hex != v.hex {
					t.Errorf("%s: hex expected %s got %s", v.name, v.hex, hex)
				}
				rgb := color.PrintColor(c, "rgb")
				if rgb != v.rgb {
					t.Errorf("%s: rgb expected %s got %s", v.name, v.rgb, rgb)
				}
				hsl := color.PrintColor(c, "hsl")
				if hsl != v.hsl {
					t.Errorf("%s: hsl expected %s got %s", v.name, v.hsl, hsl)
				}
				hsv := color.PrintColor(c, "hsv")
				if hsv != v.hsv {
					t.Errorf("%s: hsv expected %s got %s", v.name, v.hsv, hsv)
				}
			}
		}),
		mad.It("prints HSVA", func(t mad.T) {
			for _, v := range sample {
				c := color.HSVA(color.New(v.rgb).HSVA())
				hex := color.PrintColor(c, "hex")
				if hex != v.hex {
					t.Errorf("%s: hex expected %s got %s", v.name, v.hex, hex)
				}
				rgb := color.PrintColor(c, "rgb")
				if rgb != v.rgb {
					t.Errorf("%s: rgb expected %s got %s", v.name, v.rgb, rgb)
				}
				hsl := color.PrintColor(c, "hsl")
				if hsl != v.hsl {
					t.Errorf("%s: hsl expected %s got %s", v.name, v.hsl, hsl)
				}
				hsv := color.PrintColor(c, "hsv")
				if hsv != v.hsv {
					t.Errorf("%s: hsv expected %s got %s", v.name, v.hsv, hsv)
				}
			}
		}),
	)
}

func TestHSLA() mad.Test {
	return mad.It("can covert between hsla and hex", func(t mad.T) {
		e := color.New("#c71585")
		n := color.HSLA(e.HSLA())
		if n.Hex() != e.Hex() {
			t.Errorf("expected %s got %s", e.Hex(), n.Hex())
		}
	})

}
