package style

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

var prefix = themes.Default.AntPrefix + "-back-top"

const iconBackground = `url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAoCAYAAACWwljjAAAABGdBTUEAALGPC/xhBQAAAbtJREFUWAntmMtKw0AUhhMvS5cuxILgQlRUpIggIoKIIoigG1eC+AA+jo+i6FIXBfeuXIgoeKVeitVWJX5HWhhDksnUpp3FDPyZk3Nm5nycmZKkXhAEOXSA3lG7muTeRzmfy6HneUvIhnYkQK+Q9NhAA0Opg0vBEhjBKHiyb8iGMyQMOYuK41BcBSypAL+MYXSKjtFAW7EAGEO3qN4uMQbbAkXiSfRQJ1H6a+yhlkKRcAoVFYiweYNjtCVQJJpBz2GCiPt7fBOZQpFgDpUikse5HgnkM4Fi4QX0Fpc5wf9EbLqpUCy4jMoJSXWhFwbMNgWKhVbRhy5jirhs9fy/oFhgHVVTJEs7RLZ8sSEoJm6iz7SZDMbJ+/OKERQTttCXQRLToRUmrKWCYuA2+jbN0MB4OQobYShfdTCgn/sL1K36M7TLrN3n+758aPy2rrpR6+/od5E8tf/A1uLS9aId5T7J3CNYihkQ4D9PiMdMC7mp4rjB9kjFjZp8BlnVHJBuO1yFXIV0FdDF3RlyFdJVQBdv5AxVdIsq8apiZ2PyYO1EVykesGfZEESsCkweyR8MUW+V8uJ1gkYipmpdP1pm2aJVPEGzAAAAAElFTkSuQmCC) ~"100%/100%" no-repeat
`

func BackTop() gs.CSSRule {
	return gs.CSS(
		gs.S(prefix,
			mixins.ResetComponent(),
			gs.P(" z-index", themes.Default.ZIndexBackTop),
			gs.P("position", "fixed"),
			gs.P("right", "100px"),
			gs.P("bottom", "50px"),
			gs.P("height", "40px"),
			gs.P("width", "40px"),
			gs.P("cursor", "pointer"),
			gs.S("&-content",
				gs.P("height", "40px"),
				gs.P("width", "40px"),
				gs.P("border-radius", "20px"),
				gs.P("background-color", themes.Default.BackTopBG.Hex()),
				gs.P("color", themes.Default.BackTopColor.Hex()),
				gs.P("text-align", "center"),
				gs.P("transition", "all .3s "+themes.Default.EaseInOut),
				gs.P("overflow", "hidden"),
				gs.S("&:hover",
					gs.P("background-color", themes.Default.BackTopHoverBG.Hex()),
					gs.P("transition", "all .3s "+themes.Default.EaseInOut),
				),
			),
			gs.S("&-icon",
				gs.P("margin", "12px auto"),
				gs.P("width", "14px"),
				gs.P("height", "16px"),
				gs.P("background", iconBackground),
			),
		),
		gs.Media("screen and (max-width: @screen-md) ",
			gs.S(prefix,
				gs.P("right", "60px"),
			),
		),
		gs.Media("screen and (max-width: @screen-xs)",
			gs.S(prefix,
				gs.P("right", "20px"),
			),
		),
	)
}
