package themes

import (
	"github.com/gernest/vected/style/tinycolor"
)

var Default = &Theme{
	FontFamily:      `"Monospaced Number", @font-family-no-number`,
	FontSizeBase:    "14px",
	LineHeightBase:  "1.5",
	TextColor:       tinycolor.Fade("#000", 65),
	IconURL:         `"https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i"`,
	IconFontPrefix:  "anticon",
	GridColumns:     24,
	GridGutterWidth: 0,

	ScreenXS:    "480px",
	ScreenXSMin: "480px",

	ScreenSM:    "576px",
	ScreenSMMin: "576px",

	ScreenMD:    "768px",
	ScreenMDMin: "768px",

	ScreenLG:    "992px",
	ScreenLGMin: "992px",

	ScreenXL:    "1200px",
	ScreenXLMin: "1200px",

	ScreenXXL:    "1600px",
	ScreenXXLMin: "1600px",
}

type Theme struct {
	FontFamily     string
	FontSizeBase   string
	LineHeightBase string
	TextColor      *tinycolor.Color

	IconURL         string
	IconFontPrefix  string
	GridColumns     int64
	GridGutterWidth int64
	ScreenXS        string
	ScreenXSMin     string
	ScreenSM        string
	ScreenSMMin     string
	ScreenMD        string
	ScreenMDMin     string
	ScreenLG        string
	ScreenLGMin     string
	ScreenXL        string
	ScreenXLMin     string
	ScreenXXL       string
	ScreenXXLMin    string
}
