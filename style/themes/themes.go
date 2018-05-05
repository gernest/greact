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

	//scaffolding
	BodyBackground         *color.Color
	ComponentBackground    *color.Color
	FontFamilyNoNumber     string
	FontFamily             string
	CodeFamily             string
	HeadingColor           *color.Color
	TextColor              *color.Color
	TextColorSecondary     *color.Color
	HeadingColorDark       *color.Color
	TextColorDark          *color.Color
	TextColorSecondaryDark *color.Color
	FontSizeBase           string
	FontSizeLG             string
	FontSizeSM             string
	LineHeightBase         string
	BorderRadiusBase       string
	BorderRadiusSM         string

	//vertical padding
	PaddingLG string
	PaddingMD string
	PaddingSM string
	PaddingXS string

	// vertical padding for all form controls
	ControlPaddingHorizontal   string
	ControlPaddingHorizontalSM string

	// The background colors for active and hover states for things like
	// list items or table cells.
	ActiveItemBG  *color.Color
	ActiveHoverBG *color.Color

	// ICONFONT
	IconFontPrefix string
	IconURL        string

	//  LINK
	LinkColor           *color.Color
	LinkHoverColor      *color.Color
	LinkActiveColor     *color.Color
	LinkDecoration      string
	LinkHoverDecoration string

	// Animation
	EaseOut        string
	EaseIn         string
	EaseInOut      string
	EaseOutBack    string
	EaseInBack     string
	EaseInOutBack  string
	EaseOutCirc    string
	EaseInCirc     string
	EaseInOutCirc  string
	EaseOutQuint   string
	EaseInQuint    string
	EaseInOutQuint string

	// Border color
	BorderColorBase  *color.Color
	BorderColorSplit string
	BorderWithBase   string
	BorderStyleBase  string

	// Outline
	OutlineBlurSize string
	OutlineWidth    string
	OutlineColor    *color.Color

	BackgroundColorLight *color.Color
	BackgroundColorBase  *color.Color

	// Disabled states
	DisabledColor     *color.Color
	DisabledBG        *color.Color
	DisabledColorDark *color.Color

	// Shadow
	ShadowColor   *color.Color
	BoxShadowBase string
	Shadow1Up     string
	Shadow1Down   string
	Shadow1Left   string
	Shadow1Right  string
	Shadow2       string

	// Buttons
	BtnFontWeight       string
	BtnBorderRadiusBase string
	BtnBorderRadiusSM   string

	BtnPrimaryColor  *color.Color
	BtnPrimaryBG     *color.Color
	BtnDefaultColor  *color.Color
	BtnDefaultBG     *color.Color
	BtnDefaultBorder *color.Color

	BtnDangerColor  *color.Color
	BtnDangerBG     *color.Color
	BtnDangerBorder *color.Color

	BtnDisabledColor  *color.Color
	BtnDisabledBG     *color.Color
	BtnDisabledBorder *color.Color

	BtnPaddingBase string
	BtnFontSizeLG  string
	BtnFontSizeSM  string
	BtnPaddingLG   string
	BtnPaddingSM   string

	BtnHeighBase string
	BtnHeighLG   string
	BtnHeighSM   string

	BtnCircleSize   string
	BtnCircleSizeLG string
	BtnCircleSizeSM string
	BtnGroupBorder  *color.Color

	// Checkbox
	CheckboxSize  string
	CheckboxColor *color.Color

	// Radio
	RadioSize     string
	RadioDotColor *color.Color

	// Radio buttons
	RadioBtnBG          *color.Color
	RadioBtnColor       *color.Color
	RadioBtnHoverColor  *color.Color
	RadioBtnActiveColor *color.Color

	// Media queries breakpoints
	// Extra small screen / phone
	ScreenXS    string
	ScreenXSMin string

	// Small screen / tablet
	ScreenSM    string
	ScreenSMMin string

	// Medium screen / desktop
	ScreenMD    string
	ScreenMDMin string

	// Large screen / wide desktop
	ScreenLG    string
	ScreenLGMin string

	// Extra large screen / full hd
	ScreenXL    string
	ScreenXLMin string

	// Extra extra large screen / large descktop
	ScreenXXL    string
	ScreenXXLMin string

	// provide a maximum
	ScreenXSMax string
	ScreenSMMax string
	ScreenMDMax string
	ScreenLGMax string
	ScreenXLMax string

	// Grid system
	GridColumns     int64
	GridGutterWidth int64

	// Layout
	LayoutBodyBackground    *color.Color
	LayoutHeaderBackground  *color.Color
	LayoutFooterBackground  *color.Color
	LayoutHeaderHeight      string
	LayoutHeaderPadding     string
	LayoutFooterPadding     string
	LayoutSiderBackground   *color.Color
	LayoutTriggerHeight     string
	LayoutTriggerBackground *color.Color
	LayoutTriggerColor      *color.Color
	LayoutZeroTriggerWidth  string
	LayoutZeroTriggerHeight string

	// z-index list
	ZIndexAffix        string
	ZIndexBackTop      string
	ZIndexModalMask    string
	ZIndexModal        string
	ZIndexNotification string
	ZIndexMessage      string
	ZIndexPopover      string
	ZIndexPicker       string
	ZIndexDropdown     string
	ZIndexTooltip      string

	// Animation
	AnimationDurationSlow string //Modal
	AnimationDurationBase string
	AnimationDurationFast string //Tooltip

	// Form
	// ---
	LabelRequiredColor       *color.Color
	LabelColor               *color.Color
	FormItemMarginBottom     string
	FormItemTrailingColon    string
	FormVerticalLabelPadding string
	FormVerticalLabelMargin  string

	// Input
	// --
	InputHeightBase            string
	InputHeightLG              string
	InputHeightSM              string
	InputPaddingHorizontal     string
	InputPaddingHorizontalBase string
	InputPaddingHorizontalSM   string
	InputPaddingHorizontalLG   string
	InputPaddingVerticalBase   string
	InputPaddingVerticalSM     string
	InputPaddingVerticalLG     string
	InputPlaceholderColor      *color.Color
	InputColor                 *color.Color
	InputBorderColor           *color.Color
	InputBG                    *color.Color
	InputAddonBG               *color.Color
	InputHoverBorderColor      *color.Color
	InputDisabledBG            *color.Color

	// Tooltip
	// ---
	//* Tooltip max width
	TooltipMaxWidth   string
	TooltipColor      *color.Color
	TooltipBG         *color.Color
	TooltipArrowWidth string
	TooltipDistance   string
	TooltipArrowColor *color.Color
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

		BodyBackground:      color.New("#fff"),
		ComponentBackground: color.New("#fff"),

		FontFamilyNoNumber: `"Chinese Quote", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif`,
		CodeFamily:         " Consolas, Menlo, Courier, monospace",

		HeadingColor:           color.Fade(color.New("#000"), 85),
		TextColor:              color.Fade(color.New("#000"), 65),
		TextColorSecondary:     color.Fade(color.New("#000"), 45),
		HeadingColorDark:       color.Fade(color.New("#fff"), 100),
		TextColorDark:          color.Fade(color.New("#fff"), 85),
		TextColorSecondaryDark: color.Fade(color.New("#fff"), 65),

		FontSizeBase:               "14px",
		FontSizeLG:                 "16px",
		FontSizeSM:                 "12px",
		LineHeightBase:             "1.5",
		BorderRadiusBase:           "4px",
		BorderRadiusSM:             "2px",
		PaddingLG:                  "24px",
		PaddingMD:                  "16px",
		PaddingSM:                  "12px",
		PaddingXS:                  "8px",
		BackgroundColorLight:       color.New("#FAFAFA"),
		BackgroundColorBase:        color.New("#F5F5F5"),
		DisabledColor:              color.Fade(color.New("#000"), 25),
		DisabledColorDark:          color.Fade(color.New("#fff"), 35),
		BtnFontWeight:              "400",
		BtnBorderRadiusBase:        "4px",
		BtnBorderRadiusSM:          "4px",
		BtnPrimaryColor:            color.New("#fff"),
		BtnDefaultBG:               color.New("#fff"),
		IconFontPrefix:             "anticon",
		IconURL:                    `"https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i"`,
		LinkDecoration:             "none",
		LinkHoverDecoration:        "none",
		EaseOut:                    "cubic-bezier(0.215, 0.61, 0.355, 1)",
		EaseIn:                     "cubic-bezier(0.55, 0.055, 0.675, 0.19)",
		EaseInOut:                  "cubic-bezier(0.645, 0.045, 0.355, 1)",
		EaseOutBack:                "cubic-bezier(0.12, 0.4, 0.29, 1.46)",
		EaseInBack:                 "cubic-bezier(0.71, -0.46, 0.88, 0.6)",
		EaseInOutBack:              "cubic-bezier(0.71, -0.46, 0.29, 1.46)",
		EaseOutCirc:                "cubic-bezier(0.08, 0.82, 0.17, 1)",
		EaseInCirc:                 "cubic-bezier(0.6, 0.04, 0.98, 0.34)",
		EaseInOutCirc:              "cubic-bezier(0.78, 0.14, 0.15, 0.86)",
		EaseOutQuint:               "cubic-bezier(0.23, 1, 0.32, 1)",
		EaseInQuint:                "cubic-bezier(0.755, 0.05, 0.855, 0.06)",
		EaseInOutQuint:             "cubic-bezier(0.86, 0, 0.07, 1)",
		OutlineBlurSize:            "0",
		OutlineWidth:               "2px",
		ShadowColor:                color.New("#3636BB"),
		CheckboxSize:               "16px",
		RadioSize:                  "16px",
		GridColumns:                24,
		GridGutterWidth:            0,
		ScreenXS:                   "480px",
		ScreenXSMin:                "480px",
		ScreenSM:                   "576px",
		ScreenSMMin:                "576px",
		ScreenMD:                   "768px",
		ScreenMDMin:                "768px",
		ScreenLG:                   "992px",
		ScreenLGMin:                "992px",
		ScreenXL:                   "1200px",
		ScreenXLMin:                "1200px",
		ScreenXXL:                  "1600px",
		ScreenXXLMin:               "1600px",
		ScreenXSMax:                "479px",
		ScreenSMMax:                "575px",
		ScreenMDMax:                "767px",
		ScreenLGMax:                "991px",
		ScreenXLMax:                "1199px",
		BorderColorBase:            color.New("#D9D9D9"), // base border outline a component
		BorderColorSplit:           "#E8E8E8",            // split border inside a component
		BorderWithBase:             "1px",                // width of the border for a component
		BorderStyleBase:            "solid",              // style of a components border
		LayoutBodyBackground:       color.New("#f0f2f5"),
		LayoutHeaderBackground:     color.New("#001529"),
		LayoutHeaderHeight:         "64px",
		LayoutHeaderPadding:        "0 50px",
		LayoutFooterPadding:        "24px 50px",
		LayoutTriggerHeight:        "48px",
		LayoutTriggerBackground:    color.New("#002140"),
		LayoutTriggerColor:         color.New("#fff"),
		LayoutZeroTriggerWidth:     "36px",
		LayoutZeroTriggerHeight:    "42px",
		ZIndexAffix:                "10",
		ZIndexBackTop:              "10",
		ZIndexModalMask:            "1000",
		ZIndexModal:                "1000",
		ZIndexNotification:         "1010",
		ZIndexMessage:              "1010",
		ZIndexPopover:              "1030",
		ZIndexPicker:               "1050",
		ZIndexDropdown:             "1050",
		ZIndexTooltip:              "1060",
		AnimationDurationSlow:      ".3s", //Modal
		AnimationDurationBase:      ".2s",
		AnimationDurationFast:      ".1s", //Tooltip
		FormItemMarginBottom:       "24px",
		FormItemTrailingColon:      "true",
		FormVerticalLabelPadding:   "0 0 8px",
		FormVerticalLabelMargin:    "0",
		InputHeightBase:            "32px",
		InputHeightLG:              "40px",
		InputHeightSM:              "24px",
		InputPaddingHorizontal:     "11px",
		InputPaddingHorizontalBase: "12px",
		InputPaddingHorizontalSM:   "7px",
		InputPaddingHorizontalLG:   "12px",
		InputPaddingVerticalBase:   "4px",
		InputPaddingVerticalSM:     "1px",
		InputPaddingVerticalLG:     "4px",
		InputPlaceholderColor:      color.New("#BFBFBF"),
		TooltipMaxWidth:            "250px",
		TooltipColor:               color.New("#fff"),
		TooltipBG:                  color.New("#970B97"),
		TooltipArrowWidth:          "5px",
		TooltipDistance:            "8px",
	}

	t.Primary = color.Generate(t.PrimaryColor)
	t.FontFamily = `"Monospaced Number",` + t.FontFamilyNoNumber
	t.DisabledBG = t.BackgroundColorBase
	t.BtnPrimaryBG = t.PrimaryColor
	t.BtnDefaultColor = t.TextColor
	t.BtnDefaultBorder = t.BorderColorBase
	t.BtnDangerColor = t.ErrorColor
	t.BtnDangerBG = t.BackgroundColorBase
	t.BtnDangerBorder = t.BorderColorBase
	t.BtnDisabledColor = t.DisabledColor
	t.BtnDisabledBG = t.DisabledBG
	t.BtnDisabledBorder = t.BorderColorBase
	t.ControlPaddingHorizontal = t.PaddingSM
	t.ControlPaddingHorizontalSM = t.PaddingXS
	t.ActiveItemBG = t.Primary[0]
	t.ActiveHoverBG = t.Primary[0]
	t.LinkColor = t.PrimaryColor
	t.LinkHoverColor = color.GenerateColor(t.LinkColor, 5)
	t.LinkActiveColor = color.GenerateColor(t.LinkColor, 7)
	t.OutlineColor = t.PrimaryColor
	t.Shadow1Up = "0 2px 8px " + t.ShadowColor.Hex()
	t.Shadow1Down = "0 2px 8px " + t.ShadowColor.Hex()
	t.Shadow1Left = "-2px 0 8px " + t.ShadowColor.Hex()
	t.Shadow1Right = " 2px 0 8px " + t.ShadowColor.Hex()
	t.Shadow2 = " 0 4px 12px  " + t.ShadowColor.Hex()
	t.BoxShadowBase = t.Shadow1Down
	t.CheckboxColor = t.PrimaryColor
	t.RadioBtnBG = t.BtnDefaultBG
	t.RadioBtnColor = t.BtnDefaultColor
	t.RadioBtnHoverColor = t.Primary[4]
	t.RadioBtnActiveColor = t.Primary[6]
	t.LayoutFooterBackground = t.LayoutBodyBackground
	t.LayoutSiderBackground = t.LayoutHeaderBackground
	t.LabelRequiredColor = t.HighlightColor
	t.LabelColor = t.HeadingColor
	t.InputPaddingHorizontal = t.ControlPaddingHorizontal
	t.InputColor = t.TextColor
	t.InputBorderColor = t.BorderColorBase
	t.InputBG = color.New("#fff")
	t.InputAddonBG = t.BackgroundColorLight
	t.InputHoverBorderColor = t.PrimaryColor
	t.InputDisabledBG = t.DisabledBG
	t.TooltipArrowColor = t.TooltipBG
	return t
}
