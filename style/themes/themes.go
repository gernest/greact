package themes

import (
	"github.com/gernest/vected/style/color"
)

var palette = color.NewPalette()
var Default = New()

type Theme struct {
	//colors
	PrimaryColor    *color.Color
	InfoColor       *color.Color
	SuccessColor    *color.Color
	ProcessingColor *color.Color
	ErrorColor      *color.Color
	HighlightColor  *color.Color
	WarningColor    *color.Color
	NormalColor     *color.Color

	// Color used by default to control hover and active backgrounds and for
	// alert info backgrounds.
	Primary [10]*color.Color

	BodyBackground      string
	ComponentBackground string

	FontFamilyNoNumber string
	FontFamily         string
	CodeFamily         string

	HeadingColor           *color.Color
	TextColor              *color.Color
	TextColorSecondary     *color.Color
	HeadingColorDark       *color.Color
	TextColorDark          *color.Color
	TextColorSecondaryDark *color.Color

	FontSizeBase     string
	FontSizeLG       string
	FontSizeSM       string
	LineHeightBase   string
	BorderRadiusBase string
	BorderRadiusSM   string

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

// New returns a default theme.
func New() *Theme {
	t := &Theme{
		PrimaryColor:    palette.Blue[5],
		InfoColor:       palette.Blue[5],
		SuccessColor:    palette.Green[5],
		ProcessingColor: palette.Blue[5],
		ErrorColor:      palette.Red[5],
		HighlightColor:  palette.Red[5],
		WarningColor:    palette.Gold[5],
		NormalColor:     color.New("#d9d9d9"),

		BodyBackground:      "#fff",
		ComponentBackground: "#fff",

		FontFamilyNoNumber: `"Chinese Quote", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif`,
		CodeFamily:         " Consolas, Menlo, Courier, monospace",

		HeadingColor:           color.Fade(color.New("#000"), 85),
		TextColor:              color.Fade(color.New("#000"), 65),
		TextColorSecondary:     color.Fade(color.New("#000"), 45),
		HeadingColorDark:       color.Fade(color.New("#fff"), 100),
		TextColorDark:          color.Fade(color.New("#fff"), 85),
		TextColorSecondaryDark: color.Fade(color.New("#fff"), 65),

		FontSizeBase:     "14px",
		FontSizeLG:       "16px",
		FontSizeSM:       "12px",
		LineHeightBase:   "1.5",
		BorderRadiusBase: "4px",
		BorderRadiusSM:   "2px",

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
		BorderColorSplit: "#E8E8E8",        // split border inside a component
		BorderWithBase:   "1px",            // width of the border for a component
		BorderStyleBase:  "solid",          // style of a components border
	}
	t.Primary = color.Generate(t.PrimaryColor)
	t.FontFamily = `"Monospaced Number",` + t.FontFamilyNoNumber
	return t
}
