package themes

import (
	"github.com/gernest/vected/style/color"
)

var Default = &Theme{
	FontFamily:      `"Monospaced Number", @font-family-no-number`,
	FontSizeBase:    "14px",
	LineHeightBase:  "1.5",
	TextColor:       color.Fade(color.New("#000"), 65),
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

	BorderColorBase:  "hsv(0, 0, 85%)", // base border outline a component
	BorderColorSplit: "hsv(0, 0, 91%",  // split border inside a component
	BorderWithBase:   "1px",            // width of the border for a component
	BorderStyleBase:  "solid",          // style of a components border
}

type Theme struct {
	FontFamily     string
	FontSizeBase   string
	LineHeightBase string
	TextColor      *color.Color

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

	BorderColorBase  string
	BorderColorSplit string
	BorderWithBase   string
	BorderStyleBase  string
}
