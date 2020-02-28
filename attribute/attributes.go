package attribute

// Attribute represents htm attributes
type Attribute struct {
	Name     string
	Elements []string
}

// Map maps html attribute name to the Attribute object.
var Map = map[string]Attribute{
	"accept": Attribute{
		Name:     "accept",
		Elements: []string{"form", "input"},
	},
	"accept-charset": Attribute{
		Name:     "accept-charset",
		Elements: []string{"form"},
	},
	"accesskey": Attribute{
		Name:     "accesskey",
		Elements: []string{"Globalattribute"},
	},
	"action": Attribute{
		Name:     "action",
		Elements: []string{"form"},
	},
	"align": Attribute{
		Name:     "align",
		Elements: []string{"applet", "caption", "col", "colgroup", "hr", "iframe", "img", "table", "tbody", "td", "tfoot", "th", "thead", "tr"},
	},
	"allow": Attribute{
		Name:     "allow",
		Elements: []string{"iframe"},
	},
	"alt": Attribute{
		Name:     "alt",
		Elements: []string{"applet", "area", "img", "input"},
	},
	"async": Attribute{
		Name:     "async",
		Elements: []string{"script"},
	},
	"autocapitalize": Attribute{
		Name:     "autocapitalize",
		Elements: []string{"Globalattribute"},
	},
	"autocomplete": Attribute{
		Name:     "autocomplete",
		Elements: []string{"form", "input", " select", "textarea"},
	},
	"autofocus": Attribute{
		Name:     "autofocus",
		Elements: []string{"button", "input", "keygen", "select", "textarea"},
	},
	"autoplay": Attribute{
		Name:     "autoplay",
		Elements: []string{"audio", "video"},
	},
	"background": Attribute{
		Name:     "background",
		Elements: []string{"body", "table", "td", "th"},
	},
	"bgcolor": Attribute{
		Name:     "bgcolor",
		Elements: []string{"body", "col", "colgroup", "marquee", "table", "tbody", "tfoot", "td", "th", "tr"},
	},
	"border": Attribute{
		Name:     "border",
		Elements: []string{"img", "object", "table"},
	},
	"buffered": Attribute{
		Name:     "buffered",
		Elements: []string{"audio", "video"},
	},
	"capture": Attribute{
		Name:     "capture",
		Elements: []string{"input"},
	},
	"challenge": Attribute{
		Name:     "challenge",
		Elements: []string{"keygen"},
	},
	"charset": Attribute{
		Name:     "charset",
		Elements: []string{"meta", "script"},
	},
	"checked": Attribute{
		Name:     "checked",
		Elements: []string{"command", "input"},
	},
	"cite": Attribute{
		Name:     "cite",
		Elements: []string{"blockquote", "del", "ins", "q"},
	},
	"class": Attribute{
		Name:     "class",
		Elements: []string{"Globalattribute"},
	},
	"code": Attribute{
		Name:     "code",
		Elements: []string{"applet"},
	},
	"codebase": Attribute{
		Name:     "codebase",
		Elements: []string{"applet"},
	},
	"color": Attribute{
		Name:     "color",
		Elements: []string{"basefont", "font", "hr"},
	},
	"cols": Attribute{
		Name:     "cols",
		Elements: []string{"textarea"},
	},
	"colspan": Attribute{
		Name:     "colspan",
		Elements: []string{"td", "th"},
	},
	"content": Attribute{
		Name:     "content",
		Elements: []string{"meta"},
	},
	"contenteditable": Attribute{
		Name:     "contenteditable",
		Elements: []string{"Globalattribute"},
	},
	"contextmenu": Attribute{
		Name:     "contextmenu",
		Elements: []string{"Globalattribute"},
	},
	"controls": Attribute{
		Name:     "controls",
		Elements: []string{"audio", "video"},
	},
	"coords": Attribute{
		Name:     "coords",
		Elements: []string{"area"},
	},
	"crossorigin": Attribute{
		Name:     "crossorigin",
		Elements: []string{"audio", "img", "link", "script", "video"},
	},
	"csp  ": Attribute{
		Name:     "csp  ",
		Elements: []string{"iframe"},
	},
	"data": Attribute{
		Name:     "data",
		Elements: []string{"object"},
	},
	"data-*": Attribute{
		Name:     "data-*",
		Elements: []string{"Globalattribute"},
	},
	"datetime": Attribute{
		Name:     "datetime",
		Elements: []string{"del", "ins", "time"},
	},
	"decoding": Attribute{
		Name:     "decoding",
		Elements: []string{"img"},
	},
	"default": Attribute{
		Name:     "default",
		Elements: []string{"track"},
	},
	"defer": Attribute{
		Name:     "defer",
		Elements: []string{"script"},
	},
	"dir": Attribute{
		Name:     "dir",
		Elements: []string{"Globalattribute"},
	},
	"dirname": Attribute{
		Name:     "dirname",
		Elements: []string{"input", "textarea"},
	},
	"disabled": Attribute{
		Name:     "disabled",
		Elements: []string{"button", "command", "fieldset", "input", "keygen", "optgroup", "option", "select", "textarea"},
	},
	"download": Attribute{
		Name:     "download",
		Elements: []string{"a", "area"},
	},
	"draggable": Attribute{
		Name:     "draggable",
		Elements: []string{"Globalattribute"},
	},
	"dropzone": Attribute{
		Name:     "dropzone",
		Elements: []string{"Globalattribute"},
	},
	"enctype": Attribute{
		Name:     "enctype",
		Elements: []string{"form"},
	},
	"enterkeyhint  ": Attribute{
		Name:     "enterkeyhint  ",
		Elements: []string{"textarea", "contenteditable"},
	},
	"for": Attribute{
		Name:     "for",
		Elements: []string{"label", "output"},
	},
	"form": Attribute{
		Name:     "form",
		Elements: []string{"button", "fieldset", "input", "keygen", "label", "meter", "object", "output", "progress", "select", "textarea"},
	},
	"formaction": Attribute{
		Name:     "formaction",
		Elements: []string{"input", "button"},
	},
	"formenctype": Attribute{
		Name:     "formenctype",
		Elements: []string{"button", "input"},
	},
	"formmethod": Attribute{
		Name:     "formmethod",
		Elements: []string{"button", "input"},
	},
	"formnovalidate": Attribute{
		Name:     "formnovalidate",
		Elements: []string{"button", "input"},
	},
	"formtarget": Attribute{
		Name:     "formtarget",
		Elements: []string{"button", "input"},
	},
	"headers": Attribute{
		Name:     "headers",
		Elements: []string{"td", "th"},
	},
	"height": Attribute{
		Name:     "height",
		Elements: []string{"canvas", "embed", "iframe", "img", "input", "object", "video"},
	},
	"hidden": Attribute{
		Name:     "hidden",
		Elements: []string{"Globalattribute"},
	},
	"high": Attribute{
		Name:     "high",
		Elements: []string{"meter"},
	},
	"href": Attribute{
		Name:     "href",
		Elements: []string{"a", "area", "base", "link"},
	},
	"hreflang": Attribute{
		Name:     "hreflang",
		Elements: []string{"a", "area", "link"},
	},
	"http-equiv": Attribute{
		Name:     "http-equiv",
		Elements: []string{"meta"},
	},
	"icon": Attribute{
		Name:     "icon",
		Elements: []string{"command"},
	},
	"id": Attribute{
		Name:     "id",
		Elements: []string{"Globalattribute"},
	},
	"importance  ": Attribute{
		Name:     "importance  ",
		Elements: []string{"iframe", "img", "link", "script"},
	},
	"integrity": Attribute{
		Name:     "integrity",
		Elements: []string{"link", "script"},
	},
	"intrinsicsize  ": Attribute{
		Name:     "intrinsicsize  ",
		Elements: []string{"img"},
	},
	"inputmode": Attribute{
		Name:     "inputmode",
		Elements: []string{"textarea", "contenteditable"},
	},
	"ismap": Attribute{
		Name:     "ismap",
		Elements: []string{"img"},
	},
	"itemprop": Attribute{
		Name:     "itemprop",
		Elements: []string{"Globalattribute"},
	},
	"keytype": Attribute{
		Name:     "keytype",
		Elements: []string{"keygen"},
	},
	"kind": Attribute{
		Name:     "kind",
		Elements: []string{"track"},
	},
	"label": Attribute{
		Name:     "label",
		Elements: []string{"optgroup", " option", " track"},
	},
	"lang": Attribute{
		Name:     "lang",
		Elements: []string{"Globalattribute"},
	},
	"language": Attribute{
		Name:     "language",
		Elements: []string{"script"},
	},
	"loading  ": Attribute{
		Name:     "loading  ",
		Elements: []string{"img", "iframe"},
	},
	"list": Attribute{
		Name:     "list",
		Elements: []string{"input"},
	},
	"loop": Attribute{
		Name:     "loop",
		Elements: []string{"audio", "bgsound", "marquee", "video"},
	},
	"low": Attribute{
		Name:     "low",
		Elements: []string{"meter"},
	},
	"manifest": Attribute{
		Name:     "manifest",
		Elements: []string{"html"},
	},
	"max": Attribute{
		Name:     "max",
		Elements: []string{"input", "meter", "progress"},
	},
	"maxlength": Attribute{
		Name:     "maxlength",
		Elements: []string{"input", "textarea"},
	},
	"minlength": Attribute{
		Name:     "minlength",
		Elements: []string{"input", "textarea"},
	},
	"media": Attribute{
		Name:     "media",
		Elements: []string{"a", "area", "link", "source", "style"},
	},
	"method": Attribute{
		Name:     "method",
		Elements: []string{"form"},
	},
	"min": Attribute{
		Name:     "min",
		Elements: []string{"input", "meter"},
	},
	"multiple": Attribute{
		Name:     "multiple",
		Elements: []string{"input", "select"},
	},
	"muted": Attribute{
		Name:     "muted",
		Elements: []string{"audio", "video"},
	},
	"name": Attribute{
		Name:     "name",
		Elements: []string{"button", "form", "fieldset", "iframe", "input", "keygen", "object", "output", "select", "textarea", "map", "meta", "param"},
	},
	"novalidate": Attribute{
		Name:     "novalidate",
		Elements: []string{"form"},
	},
	"open": Attribute{
		Name:     "open",
		Elements: []string{"details"},
	},
	"optimum": Attribute{
		Name:     "optimum",
		Elements: []string{"meter"},
	},
	"pattern": Attribute{
		Name:     "pattern",
		Elements: []string{"input"},
	},
	"ping": Attribute{
		Name:     "ping",
		Elements: []string{"a", "area"},
	},
	"placeholder": Attribute{
		Name:     "placeholder",
		Elements: []string{"input", "textarea"},
	},
	"poster": Attribute{
		Name:     "poster",
		Elements: []string{"video"},
	},
	"preload": Attribute{
		Name:     "preload",
		Elements: []string{"audio", "video"},
	},
	"radiogroup": Attribute{
		Name:     "radiogroup",
		Elements: []string{"command"},
	},
	"readonly": Attribute{
		Name:     "readonly",
		Elements: []string{"input", "textarea"},
	},
	"referrerpolicy": Attribute{
		Name:     "referrerpolicy",
		Elements: []string{"a", "area", "iframe", "img", "link", "script"},
	},
	"rel": Attribute{
		Name:     "rel",
		Elements: []string{"a", "area", "link"},
	},
	"required": Attribute{
		Name:     "required",
		Elements: []string{"input", "select", "textarea"},
	},
	"reversed": Attribute{
		Name:     "reversed",
		Elements: []string{"ol"},
	},
	"rows": Attribute{
		Name:     "rows",
		Elements: []string{"textarea"},
	},
	"rowspan": Attribute{
		Name:     "rowspan",
		Elements: []string{"td", "th"},
	},
	"sandbox": Attribute{
		Name:     "sandbox",
		Elements: []string{"iframe"},
	},
	"scope": Attribute{
		Name:     "scope",
		Elements: []string{"th"},
	},
	"scoped": Attribute{
		Name:     "scoped",
		Elements: []string{"style"},
	},
	"selected": Attribute{
		Name:     "selected",
		Elements: []string{"option"},
	},
	"shape": Attribute{
		Name:     "shape",
		Elements: []string{"a", "area"},
	},
	"size": Attribute{
		Name:     "size",
		Elements: []string{"input", "select"},
	},
	"sizes": Attribute{
		Name:     "sizes",
		Elements: []string{"link", "img", "source"},
	},
	"slot": Attribute{
		Name:     "slot",
		Elements: []string{"Globalattribute"},
	},
	"span": Attribute{
		Name:     "span",
		Elements: []string{"col", "colgroup"},
	},
	"spellcheck": Attribute{
		Name:     "spellcheck",
		Elements: []string{"Globalattribute"},
	},
	"src": Attribute{
		Name:     "src",
		Elements: []string{"audio", "embed", "iframe", "img", "input", "script", "source", "track", "video"},
	},
	"srcdoc": Attribute{
		Name:     "srcdoc",
		Elements: []string{"iframe"},
	},
	"srclang": Attribute{
		Name:     "srclang",
		Elements: []string{"track"},
	},
	"srcset": Attribute{
		Name:     "srcset",
		Elements: []string{"img", "source"},
	},
	"start": Attribute{
		Name:     "start",
		Elements: []string{"ol"},
	},
	"step": Attribute{
		Name:     "step",
		Elements: []string{"input"},
	},
	"style": Attribute{
		Name:     "style",
		Elements: []string{"Globalattribute"},
	},
	"summary": Attribute{
		Name:     "summary",
		Elements: []string{"table"},
	},
	"tabindex": Attribute{
		Name:     "tabindex",
		Elements: []string{"Globalattribute"},
	},
	"target": Attribute{
		Name:     "target",
		Elements: []string{"a", "area", "base", "form"},
	},
	"title": Attribute{
		Name:     "title",
		Elements: []string{"Globalattribute"},
	},
	"translate": Attribute{
		Name:     "translate",
		Elements: []string{"Globalattribute"},
	},
	"type": Attribute{
		Name:     "type",
		Elements: []string{"button", "input", "command", "embed", "object", "script", "source", "style", "menu"},
	},
	"usemap": Attribute{
		Name:     "usemap",
		Elements: []string{"img", "input", "object"},
	},
	"value": Attribute{
		Name:     "value",
		Elements: []string{"button", "data", "input", "li", "meter", "option", "progress", "param"},
	},
	"width": Attribute{
		Name:     "width",
		Elements: []string{"canvas", "embed", "iframe", "img", "input", "object", "video"},
	},
	"wrap": Attribute{
		Name:     "wrap",
		Elements: []string{"textarea"},
	}}
