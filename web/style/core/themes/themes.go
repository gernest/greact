package themes

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/color"
)

// Palette is the default color palette.
var Palette = color.NewPalette()

// Default is the default Theme.
var Default = New()

// Theme antd theme object. This defines variables that controls styles.
type Theme struct {
	// The prefix to use on all css classes from ant.
	AntPrefix string
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
	// In antd the color are labelled from 1-10. So actual maping will be form
	// index 0-9. So @primary-1 becomes Primary[0],@primary-5 becomes Primary[4]
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
	FontSizeBase           gs.U
	FontSizeLG             gs.U
	FontSizeSM             gs.U
	LineHeightBase         gs.U
	BorderRadiusBase       gs.U
	BorderRadiusSM         gs.U

	//vertical padding
	PaddingLG gs.U
	PaddingMD gs.U
	PaddingSM gs.U
	PaddingXS gs.U

	// vertical padding for all form controls
	ControlPaddingHorizontal   gs.U
	ControlPaddingHorizontalSM gs.U

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
	BorderColorSplit *color.Color
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

	BtnHeighBase gs.U
	BtnHeighLG   gs.U
	BtnHeighSM   gs.U

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
	InputPaddingHorizontal     gs.U
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

	// Popover
	// ---
	PopoverBG              *color.Color
	PopoverColor           *color.Color
	PopoverMinWidth        string
	PopoverArrowWidth      string
	PopoverArrowColor      *color.Color
	PopoverArrowOuterColor *color.Color
	PopoverDistance        string

	// Modal
	// --
	ModalMaskBG *color.Color

	// Progress
	// --
	ProgressDefaultColor   *color.Color
	ProgressRemainingColor *color.Color

	// Menu
	// ---
	MenuInlineTopLevelItem  string
	MenuItemHeight          string
	MenuCollapsedWidth      string
	MenuBG                  *color.Color
	MenuItemColor           *color.Color
	MenuHighlightColor      *color.Color
	MenuItemActiveBG        *color.Color
	MenuItemGroupTitleColor *color.Color
	MenuDarkColor           *color.Color
	MenuDarkBG              *color.Color
	MenuDarkArrowColor      *color.Color
	MenuDarkSubmenuBG       *color.Color
	MenuDarkHighlightColor  *color.Color
	MenuDarkItemActiveBG    *color.Color
	MenuDarkItemSelectedBG  *color.Color

	// Spin
	// ---
	SpinDotSizeSM string
	SpinDotSize   string
	SpinDotSizeLG string

	// Table
	// --
	TableHeaderBG          *color.Color
	TableHeaderSortBG      *color.Color
	TableRowHoverBG        *color.Color
	TableSelectedRowBG     *color.Color
	TableExpandedRowBG     *color.Color
	TablePaddingVertical   string
	TablePaddingHorizontal string

	// Tag
	// --
	TagDefaultBG    *color.Color
	TagDefaultColor *color.Color
	TagFontSize     gs.U

	// TimePicker
	// ---
	TimePIckerPanelColumnWidth string
	TimePIckerPanelWidth       string
	TimePIckerSelectedBG       *color.Color

	// Carousel
	// ---
	CarouselDotWidth       string
	CarouselDotHeight      string
	CarouselDotActiveWidth string

	// Badge
	// ---
	BadgeHeight     gs.U
	BadgeDotSize    gs.U
	BadgeFontSize   gs.U
	BadgeFontWeight string
	BadgeStatusSize string

	// Rate
	// ---
	RateStarColor *color.Color
	RateStarBG    *color.Color

	// Card
	// ---
	CardHeadColor         *color.Color
	CardHeadBackground    *color.Color
	CardHeadPadding       string
	CardInnerHeadPadding  string
	CardPaddingBase       string
	CardPaddingWider      string
	CardActionsBackground *color.Color
	CardShadow            string

	// Tabs
	// ---
	TabsCardHeadBackground *color.Color
	TabsCardHeight         string
	TabsCardActiveColor    *color.Color
	TabsTitleFontSize      gs.U
	TabsTitleFontSizeLG    gs.U
	TabsTitleFontSizeSM    gs.U
	TabsInkBarBGColor      *color.Color
	TabBarMargin           string
	TabHorizontalMargin    string
	TabVerticalMargin      string
	TabHorizontalPadding   string
	TabVerticalPadding     string
	TabScrollingSize       string
	TabHighlightColor      *color.Color
	TabHoverColor          *color.Color
	TabActiveColor         *color.Color

	// BackTop
	// ---
	BackTopColor   *color.Color
	BackTopBG      *color.Color
	BackTopHoverBG *color.Color

	// Avatar
	// ---
	AvatarSizeBase     string
	AvatarSizeLG       string
	AvatarSizeSM       string
	AvatarFontSizeBase string
	AvatarFontSizeLG   string
	AvatarFontSizeSM   string
	AvatarBG           *color.Color
	AvatarColor        *color.Color
	AvatarBorderRadius gs.U

	// Switch
	// ---
	SwitchHeight              string
	SwitchSMHeight            string
	SwitchSMCheckedMarginLeft string
	SwitchDisabledOpacity     string
	SwitchColor               *color.Color

	// Pagination
	// ---
	PaginationItemSize         string
	PaginationItemSizeSM       string
	PaginationFontFamily       string
	PaginationFontWeightActive string

	// Breadcrumb
	// ---
	BreadcrumbBaseColor       *color.Color
	BreadcrumbLastItemColor   *color.Color
	BreadcrumbFontSize        gs.U
	BreadcrumbIconFontSize    gs.U
	BreadcrumbLinkColor       *color.Color
	BreadcrumbLinkColorHover  *color.Color
	BreadcrumbSeparatorColor  *color.Color
	BreadcrumbSeparatorMargin string

	// Slider
	// ---
	SliderMargin                    string
	SliderRailBackgroundColor       *color.Color
	SliderRailBackgroundColorHover  *color.Color
	SliderTrackBackgroundColor      *color.Color
	SliderTrackBackgroundColorHover *color.Color
	SliderHandleColor               *color.Color
	SliderHandleColorHover          *color.Color
	SliderHandleColorFocus          *color.Color
	SliderHandleColorFocusShadow    *color.Color
	SliderHandleColorTooltipOpen    *color.Color
	SliderDotBorderColor            *color.Color
	SliderDotBorderColorActive      *color.Color
	SliderDisabledColor             *color.Color
	SliderDisabledBackgroundColor   *color.Color

	// Collapse
	// ---
	CollapseHeaderPadding  string
	CollapseHeaderBG       *color.Color
	CollapseContentPadding gs.U
	CollapseContentBG      *color.Color
}

// New returns a default theme.
func New() *Theme {
	t := &Theme{
		AntPrefix:       ".ant",
		PrimaryColor:    Palette.Blue[5],
		InfoColor:       Palette.Blue[5],
		SuccessColor:    Palette.Green[5],
		ProcessingColor: Palette.Blue[5],
		ErrorColor:      Palette.Red[5],
		HighlightColor:  Palette.Red[5],
		WarningColor:    Palette.Gold[5],
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

		FontSizeBase:                   "14px",
		FontSizeLG:                     "16px",
		FontSizeSM:                     "12px",
		LineHeightBase:                 "1.5",
		BorderRadiusBase:               "4px",
		BorderRadiusSM:                 "2px",
		PaddingLG:                      "24px",
		PaddingMD:                      "16px",
		PaddingSM:                      "12px",
		PaddingXS:                      "8px",
		BackgroundColorLight:           color.New("#FAFAFA"),
		BackgroundColorBase:            color.New("#F5F5F5"),
		DisabledColor:                  color.Fade(color.New("#000"), 25),
		DisabledColorDark:              color.Fade(color.New("#fff"), 35),
		BtnFontWeight:                  "400",
		BtnBorderRadiusBase:            "4px",
		BtnBorderRadiusSM:              "4px",
		BtnPrimaryColor:                color.New("#fff"),
		BtnDefaultBG:                   color.New("#fff"),
		IconFontPrefix:                 "anticon",
		IconURL:                        "https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i",
		LinkDecoration:                 "none",
		LinkHoverDecoration:            "none",
		EaseOut:                        "cubic-bezier(0.215, 0.61, 0.355, 1)",
		EaseIn:                         "cubic-bezier(0.55, 0.055, 0.675, 0.19)",
		EaseInOut:                      "cubic-bezier(0.645, 0.045, 0.355, 1)",
		EaseOutBack:                    "cubic-bezier(0.12, 0.4, 0.29, 1.46)",
		EaseInBack:                     "cubic-bezier(0.71, -0.46, 0.88, 0.6)",
		EaseInOutBack:                  "cubic-bezier(0.71, -0.46, 0.29, 1.46)",
		EaseOutCirc:                    "cubic-bezier(0.08, 0.82, 0.17, 1)",
		EaseInCirc:                     "cubic-bezier(0.6, 0.04, 0.98, 0.34)",
		EaseInOutCirc:                  "cubic-bezier(0.78, 0.14, 0.15, 0.86)",
		EaseOutQuint:                   "cubic-bezier(0.23, 1, 0.32, 1)",
		EaseInQuint:                    "cubic-bezier(0.755, 0.05, 0.855, 0.06)",
		EaseInOutQuint:                 "cubic-bezier(0.86, 0, 0.07, 1)",
		OutlineBlurSize:                "0",
		OutlineWidth:                   "2px",
		ShadowColor:                    color.New("#3636BB"),
		CheckboxSize:                   "16px",
		RadioSize:                      "16px",
		GridColumns:                    24,
		GridGutterWidth:                0,
		ScreenXS:                       "480px",
		ScreenXSMin:                    "480px",
		ScreenSM:                       "576px",
		ScreenSMMin:                    "576px",
		ScreenMD:                       "768px",
		ScreenMDMin:                    "768px",
		ScreenLG:                       "992px",
		ScreenLGMin:                    "992px",
		ScreenXL:                       "1200px",
		ScreenXLMin:                    "1200px",
		ScreenXXL:                      "1600px",
		ScreenXXLMin:                   "1600px",
		ScreenXSMax:                    "479px",
		ScreenSMMax:                    "575px",
		ScreenMDMax:                    "767px",
		ScreenLGMax:                    "991px",
		ScreenXLMax:                    "1199px",
		BorderColorBase:                color.New("#D9D9D9"), // base border outline a component
		BorderColorSplit:               color.New("#E8E8E8"), // split border inside a component
		BorderWithBase:                 "1px",                // width of the border for a component
		BorderStyleBase:                "solid",              // style of a components border
		LayoutBodyBackground:           color.New("#f0f2f5"),
		LayoutHeaderBackground:         color.New("#001529"),
		LayoutHeaderHeight:             "64px",
		LayoutHeaderPadding:            "0 50px",
		LayoutFooterPadding:            "24px 50px",
		LayoutTriggerHeight:            "48px",
		LayoutTriggerBackground:        color.New("#002140"),
		LayoutTriggerColor:             color.New("#fff"),
		LayoutZeroTriggerWidth:         "36px",
		LayoutZeroTriggerHeight:        "42px",
		ZIndexAffix:                    "10",
		ZIndexBackTop:                  "10",
		ZIndexModalMask:                "1000",
		ZIndexModal:                    "1000",
		ZIndexNotification:             "1010",
		ZIndexMessage:                  "1010",
		ZIndexPopover:                  "1030",
		ZIndexPicker:                   "1050",
		ZIndexDropdown:                 "1050",
		ZIndexTooltip:                  "1060",
		AnimationDurationSlow:          ".3s", //Modal
		AnimationDurationBase:          ".2s",
		AnimationDurationFast:          ".1s", //Tooltip
		FormItemMarginBottom:           "24px",
		FormItemTrailingColon:          "true",
		FormVerticalLabelPadding:       "0 0 8px",
		FormVerticalLabelMargin:        "0",
		InputHeightBase:                "32px",
		InputHeightLG:                  "40px",
		InputHeightSM:                  "24px",
		InputPaddingHorizontal:         "11px",
		InputPaddingHorizontalBase:     "12px",
		InputPaddingHorizontalSM:       "7px",
		InputPaddingHorizontalLG:       "12px",
		InputPaddingVerticalBase:       "4px",
		InputPaddingVerticalSM:         "1px",
		InputPaddingVerticalLG:         "4px",
		InputPlaceholderColor:          color.New("#BFBFBF"),
		TooltipMaxWidth:                "250px",
		TooltipColor:                   color.New("#fff"),
		TooltipBG:                      color.New("#970B97"),
		TooltipArrowWidth:              "5px",
		TooltipDistance:                "8px",
		PopoverBG:                      color.New("#fff"),
		PopoverMinWidth:                "177px",
		PopoverArrowWidth:              "5px",
		PopoverDistance:                "9px",
		ModalMaskBG:                    color.New("#1E2324"),
		MenuInlineTopLevelItem:         "40px",
		MenuItemHeight:                 "40px",
		MenuCollapsedWidth:             "80px",
		MenuDarkArrowColor:             color.New("#fff"),
		MenuDarkSubmenuBG:              color.New("#000c17"),
		MenuDarkHighlightColor:         color.New("#fff"),
		SpinDotSizeSM:                  "14px",
		SpinDotSize:                    "20px",
		SpinDotSizeLG:                  "32px",
		TableSelectedRowBG:             color.New("#fafafa"),
		TableExpandedRowBG:             color.New("#fbfbfb"),
		TablePaddingVertical:           "16px",
		TablePaddingHorizontal:         "16px",
		TimePIckerPanelColumnWidth:     "56px",
		TimePIckerPanelWidth:           "168px",
		CarouselDotWidth:               "16px",
		CarouselDotHeight:              "3px",
		CarouselDotActiveWidth:         "24px",
		BadgeHeight:                    "20px",
		BadgeDotSize:                   "6px",
		BadgeFontWeight:                "normal",
		BadgeStatusSize:                "6px",
		CardHeadPadding:                "16px",
		CardInnerHeadPadding:           "12px",
		CardPaddingBase:                "24px",
		CardPaddingWider:               "32px",
		CardShadow:                     "0 2px 8px rgba(0, 0, 0, .09)",
		TabsCardHeight:                 "40px",
		TabBarMargin:                   "0 0 16px 0",
		TabHorizontalMargin:            "0 32px 0 0",
		TabVerticalMargin:              "0 0 16px 0",
		TabHorizontalPadding:           "12px 16px",
		TabVerticalPadding:             "8px 24px",
		TabScrollingSize:               "32px",
		BackTopColor:                   color.New("#fff"),
		AvatarSizeBase:                 "32px",
		AvatarSizeLG:                   "40px",
		AvatarSizeSM:                   "24px",
		AvatarFontSizeBase:             "18px",
		AvatarFontSizeLG:               "24px",
		AvatarFontSizeSM:               "14px",
		AvatarBG:                       color.New("#ccc"),
		AvatarColor:                    color.New("#fff"),
		SwitchHeight:                   "22px",
		SwitchSMHeight:                 "16px",
		SwitchSMCheckedMarginLeft:      "-13px",
		SwitchDisabledOpacity:          " 0.4",
		PaginationItemSize:             "32px",
		PaginationItemSizeSM:           "24px",
		PaginationFontFamily:           "Arial",
		PaginationFontWeightActive:     "500",
		SliderMargin:                   "14px 6px 10px",
		SliderRailBackgroundColorHover: color.New("#e1e1e1"),
		CollapseHeaderPadding:          "12px 0 12px 40px",
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
	t.PopoverColor = t.TextColor
	t.PopoverArrowColor = t.PopoverBG
	t.PopoverArrowOuterColor = t.PopoverBG
	t.ProgressDefaultColor = t.ProcessingColor
	t.ProgressRemainingColor = t.BackgroundColorBase
	t.MenuBG = t.ComponentBackground
	t.MenuItemColor = t.TextColor
	t.MenuHighlightColor = t.PrimaryColor
	t.MenuItemActiveBG = t.ActiveItemBG
	t.MenuItemGroupTitleColor = t.TextColorSecondary
	t.MenuDarkColor = t.TextColorSecondaryDark
	t.MenuDarkBG = t.LayoutHeaderBackground
	t.MenuDarkItemActiveBG = t.PrimaryColor
	t.MenuDarkItemSelectedBG = t.PrimaryColor
	t.TableHeaderBG = t.BackgroundColorLight
	t.TableHeaderSortBG = t.BackgroundColorBase
	t.TableRowHoverBG = t.Primary[0]
	t.TagDefaultBG = t.BackgroundColorLight
	t.TagDefaultColor = t.TextColor
	t.TagFontSize = t.FontSizeSM
	t.TimePIckerSelectedBG = t.BackgroundColorBase
	t.BadgeFontSize = t.FontSizeSM
	t.RateStarColor = Palette.Yellow[5]
	t.RateStarBG = t.BorderColorSplit
	t.CardHeadColor = t.HeadingColor
	t.CardHeadBackground = t.ComponentBackground
	t.CardActionsBackground = t.BackgroundColorLight
	t.TabsCardHeadBackground = t.BackgroundColorLight
	t.TabsCardActiveColor = t.PrimaryColor
	t.TabsTitleFontSize = t.FontSizeBase
	t.TabsTitleFontSizeLG = t.FontSizeLG
	t.TabsTitleFontSizeSM = t.FontSizeBase
	t.TabsInkBarBGColor = t.PrimaryColor
	t.TabHighlightColor = t.PrimaryColor
	t.TabHoverColor = t.Primary[4]
	t.TabActiveColor = t.Primary[6]
	t.BackTopBG = t.TextColorSecondary
	t.BackTopHoverBG = t.TextColor
	t.AvatarBorderRadius = t.BorderRadiusBase
	t.SwitchColor = t.PrimaryColor
	t.BreadcrumbBaseColor = t.TextColorSecondary
	t.BreadcrumbLastItemColor = t.TextColor
	t.BreadcrumbFontSize = t.FontSizeBase
	t.BreadcrumbIconFontSize = t.FontSizeSM
	t.BreadcrumbLinkColor = t.TextColorSecondary
	t.BreadcrumbLinkColorHover = t.Primary[4]
	t.BreadcrumbSeparatorColor = t.TextColorSecondary
	t.BreadcrumbSeparatorMargin = "0 " + t.PaddingXS.String()
	t.SliderRailBackgroundColor = t.BackgroundColorBase
	t.SliderTrackBackgroundColor = t.Primary[2]
	t.SliderTrackBackgroundColorHover = t.Primary[3]
	t.SliderHandleColor = t.Primary[2]
	t.SliderHandleColorHover = t.Primary[3]
	t.SliderHandleColorFocus = color.Tint(t.PrimaryColor, 20)
	t.SliderHandleColorFocusShadow = color.Tint(t.PrimaryColor, 50)
	t.SliderHandleColorTooltipOpen = t.PrimaryColor
	t.SliderDotBorderColor = t.BorderColorSplit
	t.SliderDotBorderColorActive = color.Tint(t.PrimaryColor, 50)
	t.SliderDisabledColor = t.DisabledColor
	t.SliderDisabledBackgroundColor = t.ComponentBackground
	t.CollapseHeaderBG = t.BackgroundColorLight
	t.CollapseContentPadding = t.PaddingMD
	t.CollapseContentBG = t.ComponentBackground
	return t
}
