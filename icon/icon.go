package icon

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/style/icon"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Kind string

const (
	Addfile            Kind = "addfile"
	Addfolder          Kind = "addfolder"
	Alipay             Kind = "alipay"
	AlipayCircle       Kind = "alipay-circle"
	Aliwangwang        Kind = "aliwangwang"
	AliwangwangO       Kind = "aliwangwang-o"
	Aliyun             Kind = "aliyun"
	Amazon             Kind = "amazon"
	Android            Kind = "android"
	AndroidO           Kind = "android-o"
	AntDesign          Kind = "ant-design"
	Api                Kind = "api"
	Apple              Kind = "apple"
	AppleO             Kind = "apple-o"
	Appstore           Kind = "appstore"
	AppstoreO          Kind = "appstore-o"
	AreaChart          Kind = "area-chart"
	ArrowDown          Kind = "arrow-down"
	ArrowLeft          Kind = "arrow-left"
	ArrowRight         Kind = "arrow-right"
	ArrowSalt          Kind = "arrow-salt"
	ArrowUp            Kind = "arrow-up"
	ArrowsAlt          Kind = "arrows-alt"
	Backward           Kind = "backward"
	Bank               Kind = "bank"
	BarChart           Kind = "bar-chart"
	Barcode            Kind = "barcode"
	Bars               Kind = "bars"
	Behance            Kind = "behance"
	BehanceSquare      Kind = "behance-square"
	Bell               Kind = "bell"
	Book               Kind = "book"
	Bulb               Kind = "bulb"
	Calculator         Kind = "calculator"
	Calendar           Kind = "calendar"
	Camera             Kind = "camera"
	CameraO            Kind = "camera-o"
	Car                Kind = "car"
	CaretCircleDown    Kind = "caret-circle-down"
	CaretCircleLeft    Kind = "caret-circle-left"
	CaretCircleODown   Kind = "caret-circle-o-down"
	CaretCircleOLeft   Kind = "caret-circle-o-left"
	CaretCircleORight  Kind = "caret-circle-o-right"
	CaretCircleOUp     Kind = "caret-circle-o-up"
	CaretCircleRight   Kind = "caret-circle-right"
	CaretCircleUp      Kind = "caret-circle-up"
	CaretDown          Kind = "caret-down"
	CaretLeft          Kind = "caret-left"
	CaretRight         Kind = "caret-right"
	CaretUp            Kind = "caret-up"
	Check              Kind = "check"
	CheckCircle        Kind = "check-circle"
	CheckCircleO       Kind = "check-circle-o"
	CheckSquare        Kind = "check-square"
	CheckSquareO       Kind = "check-square-o"
	Chrome             Kind = "chrome"
	CircleDown         Kind = "circle-down-"
	CircleDownO        Kind = "circle-down-o"
	CircleLeft         Kind = "circle-left"
	CircleODown        Kind = "circle-o-down"
	CircleOLeft        Kind = "circle-o-left"
	CircleORight       Kind = "circle-o-right"
	CircleOUp          Kind = "circle-o-up"
	CircleRight        Kind = "circle-right"
	CircleUp           Kind = "circle-up"
	ClockCircle        Kind = "clock-circle"
	ClockCircleO       Kind = "clock-circle-o"
	Close              Kind = "close"
	CloseCircle        Kind = "close-circle"
	CloseCircleO       Kind = "close-circle-o"
	CloseSquare        Kind = "close-square"
	CloseSquareO       Kind = "close-square-o"
	Cloud              Kind = "cloud"
	CloudDownload      Kind = "cloud-download"
	CloudDownloadO     Kind = "cloud-download-o"
	CloudO             Kind = "cloud-o"
	CloudUpload        Kind = "cloud-upload"
	CloudUploadO       Kind = "cloud-upload-o"
	Code               Kind = "code"
	CodeO              Kind = "code-o"
	Codepen            Kind = "codepen"
	CodepenCircle      Kind = "codepen-circle"
	Coffee             Kind = "coffee"
	Compass            Kind = "compass"
	Contacts           Kind = "contacts"
	Copy               Kind = "copy"
	Copyright          Kind = "copyright"
	CreditCard         Kind = "credit-card"
	Cross              Kind = "cross"
	CrossCircle        Kind = "cross-circle"
	CrossCircleO       Kind = "cross-circle-o"
	CustomerService    Kind = "customer-service"
	Customerservice    Kind = "customerservice"
	Dashboard          Kind = "dashboard"
	Database           Kind = "database"
	Delete             Kind = "delete"
	Desktop            Kind = "desktop"
	Dingding           Kind = "dingding"
	DingdingO          Kind = "dingding-o"
	Disconnect         Kind = "disconnect"
	Dislike            Kind = "dislike"
	DislikeO           Kind = "dislike-o"
	DotChart           Kind = "dot-chart"
	DoubleLeft         Kind = "double-left"
	DoubleRight        Kind = "double-right"
	Down               Kind = "down"
	DownCircle         Kind = "down-circle"
	DownCircleO        Kind = "down-circle-o"
	DownSquare         Kind = "down-square"
	DownSquareO        Kind = "down-square-o"
	Download           Kind = "download"
	Dribbble           Kind = "dribbble"
	DribbbleSquare     Kind = "dribbble-square"
	Dropbox            Kind = "dropbox"
	Edit               Kind = "edit"
	Ellipsis           Kind = "ellipsis"
	Enter              Kind = "enter"
	Environment        Kind = "environment"
	EnvironmentO       Kind = "environment-o"
	Exception          Kind = "exception"
	Exclamation        Kind = "exclamation"
	ExclamationCircle  Kind = "exclamation-circle"
	ExclamationCircleO Kind = "exclamation-circle-o"
	Export             Kind = "export"
	Eye                Kind = "eye"
	EyeO               Kind = "eye-o"
	Facebook           Kind = "facebook"
	FastBackward       Kind = "fast-backward"
	FastForward        Kind = "fast-forward"
	File               Kind = "file"
	FileAdd            Kind = "file-add"
	FileExcel          Kind = "file-excel"
	FileJpg            Kind = "file-jpg"
	FileMarkdown       Kind = "file-markdown"
	FilePdf            Kind = "file-pdf"
	FilePpt            Kind = "file-ppt"
	FileText           Kind = "file-text"
	FileUnknown        Kind = "file-unknown"
	FileWord           Kind = "file-word"
	Filter             Kind = "filter"
	Flag               Kind = "flag"
	Folder             Kind = "folder"
	FolderAdd          Kind = "folder-add"
	FolderOpen         Kind = "folder-open"
	Fork               Kind = "fork"
	Form               Kind = "form"
	Forward            Kind = "forward"
	Frown              Kind = "frown"
	FrownCircle        Kind = "frown-circle"
	FrownO             Kind = "frown-o"
	Gift               Kind = "gift"
	Github             Kind = "github"
	Gitlab             Kind = "gitlab"
	Global             Kind = "global"
	Google             Kind = "google"
	GooglePlus         Kind = "google-plus"
	Hdd                Kind = "hdd"
	Heart              Kind = "heart"
	HeartO             Kind = "heart-o"
	Home               Kind = "home"
	Hourglass          Kind = "hourglass"
	Html5              Kind = "html5"
	Idcard             Kind = "idcard"
	Ie                 Kind = "ie"
	Inbox              Kind = "inbox"
	Info               Kind = "info"
	InfoCircle         Kind = "info-circle"
	InfoCircleO        Kind = "info-circle-o"
	Instagram          Kind = "instagram"
	Key                Kind = "key"
	Laptop             Kind = "laptop"
	Layout             Kind = "layout"
	Left               Kind = "left"
	LeftCircle         Kind = "left-circle"
	LeftCircleO        Kind = "left-circle-o"
	LeftSquare         Kind = "left-square"
	LeftSquareO        Kind = "left-square-o"
	Like               Kind = "like"
	LikeO              Kind = "like-o"
	LineChart          Kind = "line-chart"
	Link               Kind = "link"
	Linkedin           Kind = "linkedin"
	Loading            Kind = "loading"
	Loading3Quarters   Kind = "loading-3-quarters"
	Lock               Kind = "lock"
	Login              Kind = "login"
	Logout             Kind = "logout"
	Mail               Kind = "mail"
	Man                Kind = "man"
	MedicineBox        Kind = "medicine-box"
	Medium             Kind = "medium"
	MediumWorkmark     Kind = "medium-workmark"
	Meh                Kind = "meh"
	MehCircle          Kind = "meh-circle"
	MehO               Kind = "meh-o"
	MenuFold           Kind = "menu-fold"
	MenuUnfold         Kind = "menu-unfold"
	Message            Kind = "message"
	Minus              Kind = "minus"
	MinusCircle        Kind = "minus-circle"
	MinusCircleO       Kind = "minus-circle-o"
	MinusSquare        Kind = "minus-square"
	MinusSquareO       Kind = "minus-square-o"
	Mobile             Kind = "mobile"
	Notification       Kind = "notification"
	PaperClip          Kind = "paper-clip"
	Pause              Kind = "pause"
	PauseCircle        Kind = "pause-circle"
	PauseCircleO       Kind = "pause-circle-o"
	PayCircle          Kind = "pay-circle"
	PayCircleO         Kind = "pay-circle-o"
	Phone              Kind = "phone"
	Picture            Kind = "picture"
	PieChart           Kind = "pie-chart"
	PlayCircle         Kind = "play-circle"
	PlayCircleO        Kind = "play-circle-o"
	Plus               Kind = "plus"
	PlusCircle         Kind = "plus-circle"
	PlusCircleO        Kind = "plus-circle-o"
	PlusSquare         Kind = "plus-square"
	PlusSquareO        Kind = "plus-square-o"
	Poweroff           Kind = "poweroff"
	Printer            Kind = "printer"
	Profile            Kind = "profile"
	Pushpin            Kind = "pushpin"
	PushpinO           Kind = "pushpin-o"
	Qq                 Kind = "qq"
	Qrcode             Kind = "qrcode"
	Question           Kind = "question"
	QuestionCircle     Kind = "question-circle"
	QuestionCircleO    Kind = "question-circle-o"
	RedEnvelope        Kind = "red-envelope"
	Reload             Kind = "reload"
	Retweet            Kind = "retweet"
	Right              Kind = "right"
	RightCircle        Kind = "right-circle"
	RightCircleO       Kind = "right-circle-o"
	RightSquare        Kind = "right-square"
	RightSquareO       Kind = "right-square-o"
	Rocket             Kind = "rocket"
	Rollback           Kind = "rollback"
	Safety             Kind = "safety"
	Save               Kind = "save"
	Scan               Kind = "scan"
	Schedule           Kind = "schedule"
	Search             Kind = "search"
	Select             Kind = "select"
	Setting            Kind = "setting"
	Shake              Kind = "shake"
	ShareAlt           Kind = "share-alt"
	Shop               Kind = "shop"
	ShoppingCart       Kind = "shopping-cart"
	Shrink             Kind = "shrink"
	Skin               Kind = "skin"
	Skype              Kind = "skype"
	Slack              Kind = "slack"
	SlackSquare        Kind = "slack-square"
	Smile              Kind = "smile"
	SmileCircle        Kind = "smile-circle"
	SmileO             Kind = "smile-o"
	Solution           Kind = "solution"
	Sound              Kind = "sound"
	Spin               Kind = "spin"
	Star               Kind = "star"
	StarO              Kind = "star-o"
	StepBackward       Kind = "step-backward"
	StepForward        Kind = "step-forward"
	Swap               Kind = "swap"
	SwapLeft           Kind = "swap-left"
	SwapRight          Kind = "swap-right"
	Switcher           Kind = "switcher"
	Sync               Kind = "sync"
	Table              Kind = "table"
	Tablet             Kind = "tablet"
	Tag                Kind = "tag"
	TagO               Kind = "tag-o"
	Tags               Kind = "tags"
	TagsO              Kind = "tags-o"
	Taobao             Kind = "taobao"
	TaobaoCircle       Kind = "taobao-circle"
	Team               Kind = "team"
	ToTop              Kind = "to-top"
	Tool               Kind = "tool"
	Trademark          Kind = "trademark"
	Trophy             Kind = "trophy"
	Twitter            Kind = "twitter"
	Unlock             Kind = "unlock"
	Up                 Kind = "up"
	UpCircle           Kind = "up-circle"
	UpCircleO          Kind = "up-circle-o"
	UpSquare           Kind = "up-square"
	UpSquareO          Kind = "up-square-o"
	Upload             Kind = "upload"
	Usb                Kind = "usb"
	User               Kind = "user"
	UserAdd            Kind = "user-add"
	UserDelete         Kind = "user-delete"
	UsergroupAdd       Kind = "usergroup-add"
	UsergroupDelete    Kind = "usergroup-delete"
	VerticleLeft       Kind = "verticle-left"
	VerticleRight      Kind = "verticle-right"
	VideoCamera        Kind = "video-camera"
	Wallet             Kind = "wallet"
	Warning            Kind = "warning"
	Wechat             Kind = "wechat"
	Weibo              Kind = "weibo"
	WeiboCircle        Kind = "weibo-circle"
	WeiboSquare        Kind = "weibo-square"
	Wifi               Kind = "wifi"
	Windows            Kind = "windows"
	WindowsO           Kind = "windows-o"
	Woman              Kind = "woman"
	Youtube            Kind = "youtube"
	Yuque              Kind = "yuque"
	Zhihu              Kind = "zhihu"
)

var icons = map[string]string{
	"step-forward":         `"\e600"`,
	"step-backward":        `"\e601"`,
	"forward":              `"\e602"`,
	"backward":             `"\e603"`,
	"caret-right":          `"\e604"`,
	"caret-left":           `"\e605"`,
	"caret-down":           `"\e606"`,
	"caret-up":             `"\e607"`,
	"right-circle":         `"\e608"`,
	"circle-right":         `"\e608"`, // antd@1.x compatibility alias: right-circle
	"caret-circle-right":   `"\e608"`, // antd@1.x compatibility alias: right-circle
	"left-circle":          `"\e609"`,
	"circle-left":          `"\e609"`, // antd@1.x compatibility alias: left-circle
	"caret-circle-left":    `"\e609"`, // antd@1.x compatibility alias: left-circle
	"up-circle":            `"\e60a"`,
	"circle-up":            `"\e60a"`, // antd@1.x compatibility alias: up-circle
	"caret-circle-up":      `"\e60a"`, // antd@1.x compatibility alias: up-circle
	"down-circle":          `"\e60b"`,
	"circle-down":          `"\e60b"`, // antd@1.x compatibility alias: down-circle
	"caret-circle-down":    `"\e60b"`, // antd@1.x compatibility alias: down-circle
	"right-circle-o":       `"\e60c"`,
	"circle-o-right":       `"\e60c"`, // antd@1.x compatibility alias: right-circle-o
	"caret-circle-o-right": `"\e60c"`, // antd@1.x compatibility alias: right-circle-o
	"left-circle-o":        `"\e60d"`,
	"circle-o-left":        `"\e60d"`, // antd@1.x compatibility alias: left-circle-o
	"caret-circle-o-left":  `"\e60d"`, // antd@1.x compatibility alias: left-circle-o
	"up-circle-o":          `"\e60e"`,
	"circle-o-up":          `"\e60e"`, // antd@1.x compatibility alias: up-circle-o
	"caret-circle-o-up":    `"\e60e"`, // antd@1.x compatibility alias: up-circle-o
	"down-circle-o":        `"\e60f"`,
	"circle-o-down":        `"\e60f"`, // antd@1.x compatibility alias: down-circle-o
	"caret-circle-o-down":  `"\e60f"`, // antd@1.x compatibility alias: down-circle-o
	"verticle-left":        `"\e610"`,
	"verticle-right":       `"\e611"`,
	"rollback":             `"\e612"`,
	"retweet":              `"\e613"`,
	"shrink":               `"\e614"`,
	"arrows-alt":           `"\e615"`,
	"arrow-salt":           `"\e615"`, // antd@1.x compatibility alias: arrows-alt
	"reload":               `"\e616"`,
	"double-right":         `"\e617"`,
	"double-left":          `"\e618"`,
	"arrow-down":           `"\e619"`,
	"arrow-up":             `"\e61a"`,
	"arrow-right":          `"\e61b"`,
	"arrow-left":           `"\e61c"`,
	"down":                 `"\e61d"`,
	"up":                   `"\e61e"`,
	"right":                `"\e61f"`,
	"left":                 `"\e620"`,
	"minus-square-o":       `"\e621"`,
	"minus-circle":         `"\e622"`,
	"minus-circle-o":       `"\e623"`,
	"minus":                `"\e624"`,
	"plus-circle-o":        `"\e625"`,
	"plus-circle":          `"\e626"`,
	"plus":                 `"\e627"`,
	"info-circle":          `"\e628"`,
	"info-circle-o":        `"\e629"`,
	"info":                 `"\e62a"`,
	"exclamation":          `"\e62b"`,
	"exclamation-circle":   `"\e62c"`,
	"exclamation-circle-o": `"\e62d"`,
	"close-circle":         `"\e62e"`,
	"cross-circle":         `"\e62e"`, // antd@1.x compatibility alias: close-circle
	"close-circle-o":       `"\e62f"`,
	"cross-circle-o":       `"\e62f"`, // antd@1.x compatibility alias: close-circle-o
	"check-circle":         `"\e630"`,
	"check-circle-o":       `"\e631"`,
	"check":                `"\e632"`,
	"close":                `"\e633"`,
	"cross":                `"\e633"`, // antd@1.x compatibility alias: close
	"customer-service":     `"\e634"`,
	"customerservice":      `"\e634"`, // antd@1.x compatibility alias: customer-service
	"credit-card":          `"\e635"`,
	"code-o":               `"\e636"`,
	"book":                 `"\e637"`,
	"bars":                 `"\e639"`,
	"question":             `"\e63a"`,
	"question-circle":      `"\e63b"`,
	"question-circle-o":    `"\e63c"`,
	"pause":                `"\e63d"`,
	"pause-circle":         `"\e63e"`,
	"pause-circle-o":       `"\e63f"`,
	"clock-circle":         `"\e640"`,
	"clock-circle-o":       `"\e641"`,
	"swap":                 `"\e642"`,
	"swap-left":            `"\e643"`,
	"swap-right":           `"\e644"`,
	"plus-square-o":        `"\e645"`,
	"frown":                `"\e646"`,
	"frown-circle":         `"\e646"`, // antd@1.x compatibility alias: frown
	"ellipsis":             `"\e647"`,
	"copy":                 `"\e648"`,
	"menu-fold":            `"\e9ac"`,
	"mail":                 `"\e659"`,
	"logout":               `"\e65a"`,
	"link":                 `"\e65b"`,
	"area-chart":           `"\e65c"`,
	"line-chart":           `"\e65d"`,
	"home":                 `"\e65e"`,
	"laptop":               `"\e65f"`,
	"star":                 `"\e660"`,
	"star-o":               `"\e661"`,
	"folder":               `"\e662"`,
	"filter":               `"\e663"`,
	"file":                 `"\e664"`,
	"exception":            `"\e665"`,
	"meh":                  `"\e666"`,
	"meh-circle":           `"\e666"`, // antd@1.x compatibility alias: meh
	"meh-o":                `"\e667"`,
	"shopping-cart":        `"\e668"`,
	"save":                 `"\e669"`,
	"user":                 `"\e66a"`,
	"video-camera":         `"\e66b"`,
	"to-top":               `"\e66c"`,
	"team":                 `"\e66d"`,
	"tablet":               `"\e66e"`,
	"solution":             `"\e66f"`,
	"search":               `"\e670"`,
	"share-alt":            `"\e671"`,
	"setting":              `"\e672"`,
	"poweroff":             `"\e6d5"`,
	"picture":              `"\e674"`,
	"phone":                `"\e675"`,
	"paper-clip":           `"\e676"`,
	"notification":         `"\e677"`,
	"mobile":               `"\e678"`,
	"menu-unfold":          `"\e9ad"`,
	"inbox":                `"\e67a"`,
	"lock":                 `"\e67b"`,
	"qrcode":               `"\e67c"`,
	"play-circle":          `"\e6d0"`,
	"play-circle-o":        `"\e6d1"`,
	"tag":                  `"\e6d2"`,
	"tag-o":                `"\e6d3"`,
	"tags":                 `"\e67d"`,
	"tags-o":               `"\e67e"`,
	"cloud-o":              `"\e67f"`,
	"cloud":                `"\e680"`,
	"cloud-upload":         `"\e681"`,
	"cloud-download":       `"\e682"`,
	"cloud-download-o":     `"\e683"`,
	"cloud-upload-o":       `"\e684"`,
	"environment":          `"\e685"`,
	"environment-o":        `"\e686"`,
	"eye":                  `"\e687"`,
	"eye-o":                `"\e688"`,
	"camera":               `"\e689"`,
	"camera-o":             `"\e68a"`,
	"windows":              `"\e68b"`,
	"apple":                `"\e68c"`,
	"apple-o":              `"\e6d4"`,
	"android":              `"\e938"`,
	"android-o":            `"\e68d"`,
	"aliwangwang":          `"\e68e"`,
	"aliwangwang-o":        `"\e68f"`,
	"export":               `"\e691"`,
	"edit":                 `"\e692"`,
	"circle-down-o":        `"\e693"`,
	"circle-down-":         `"\e694"`,
	"appstore-o":           `"\e695"`,
	"appstore":             `"\e696"`,
	"scan":                 `"\e697"`,
	"file-text":            `"\e698"`,
	"folder-open":          `"\e699"`,
	"hdd":                  `"\e69a"`,
	"ie":                   `"\e69b"`,
	"file-jpg":             `"\e69c"`,
	"like":                 `"\e64c"`,
	"like-o":               `"\e69d"`,
	"dislike":              `"\e64b"`,
	"dislike-o":            `"\e69e"`,
	"delete":               `"\e69f"`,
	"enter":                `"\e6a0"`,
	"pushpin-o":            `"\e6a1"`,
	"pushpin":              `"\e6a2"`,
	"heart":                `"\e6a3"`,
	"heart-o":              `"\e6a4"`,
	"pay-circle":           `"\e6a5"`,
	"pay-circle-o":         `"\e6a6"`,
	"smile":                `"\e6a7"`,
	"smile-circle":         `"\e6a7"`, // antd@1.x compatibility alias: smile
	"smile-o":              `"\e6a8"`,
	"frown-o":              `"\e6a9"`,
	"calculator":           `"\e6aa"`,
	"message":              `"\e6ab"`,
	"chrome":               `"\e6ac"`,
	"github":               `"\e6ad"`,
	"file-unknown":         `"\e6af"`,
	"file-excel":           `"\e6b0"`,
	"file-ppt":             `"\e6b1"`,
	"file-word":            `"\e6b2"`,
	"file-pdf":             `"\e6b3"`,
	"desktop":              `"\e6b4"`,
	"upload":               `"\e6b6"`,
	"download":             `"\e6b7"`,
	"pie-chart":            `"\e6b8"`,
	"unlock":               `"\e6ba"`,
	"calendar":             `"\e6bb"`,
	"windows-o":            `"\e6bc"`,
	"dot-chart":            `"\e6bd"`,
	"bar-chart":            `"\e6be"`,
	"code":                 `"\e6bf"`,
	"api":                  `"\e951"`,
	"plus-square":          `"\e6c0"`,
	"minus-square":         `"\e6c1"`,
	"close-square":         `"\e6c2"`,
	"close-square-o":       `"\e6c3"`,
	"check-square":         `"\e6c4"`,
	"check-square-o":       `"\e6c5"`,
	"fast-backward":        `"\e6c6"`,
	"fast-forward":         `"\e6c7"`,
	"up-square":            `"\e6c8"`,
	"down-square":          `"\e6c9"`,
	"left-square":          `"\e6ca"`,
	"right-square":         `"\e6cb"`,
	"right-square-o":       `"\e6cc"`,
	"left-square-o":        `"\e6cd"`,
	"down-square-o":        `"\e6ce"`,
	"up-square-o":          `"\e6cf"`,
	"loading":              `"\e64d"`,
	"loading-3-quarters":   `"\e6ae"`,
	"bulb":                 `"\e649"`,
	"select":               `"\e64a"`,
	"addfile":              `"\e910"`,
	"file-add":             `"\e910"`,
	"addfolder":            `"\e914"`,
	"folder-add":           `"\e914"`,
	"switcher":             `"\e913"`,
	"rocket":               `"\e90f"`,
	"dingding":             `"\e923"`,
	"dingding-o":           `"\e925"`,
	"bell":                 `"\e64e"`,
	"disconnect":           `"\e64f"`,
	"database":             `"\e650"`,
	"compass":              `"\e6db"`,
	"barcode":              `"\e652"`,
	"hourglass":            `"\e653"`,
	"key":                  `"\e654"`,
	"flag":                 `"\e655"`,
	"layout":               `"\e656"`,
	"login":                `"\e657"`,
	"printer":              `"\e673"`,
	"sound":                `"\e6e9"`,
	"usb":                  `"\e6d7"`,
	"skin":                 `"\e6d8"`,
	"tool":                 `"\e6d9"`,
	"sync":                 `"\e6da"`,
	"wifi":                 `"\e6d6"`,
	"car":                  `"\e6dc"`,
	"copyright":            `"\e6de"`,
	"schedule":             `"\e6df"`,
	"user-add":             `"\e6ed"`,
	"user-delete":          `"\e6e0"`,
	"usergroup-add":        `"\e6dd"`,
	"usergroup-delete":     `"\e6e1"`,
	"man":                  `"\e6e2"`,
	"woman":                `"\e6ec"`,
	"shop":                 `"\e6e3"`,
	"gift":                 `"\e6e4"`,
	"idcard":               `"\e6e5"`,
	"medicine-box":         `"\e6e6"`,
	"red-envelope":         `"\e6e7"`,
	"coffee":               `"\e6e8"`,
	"trademark":            `"\e651"`,
	"safety":               `"\e6ea"`,
	"wallet":               `"\e6eb"`,
	"bank":                 `"\e6ee"`,
	"trophy":               `"\e6ef"`,
	"contacts":             `"\e6f0"`,
	"global":               `"\e6f1"`,
	"shake":                `"\e94f"`,
	"fork":                 `"\e6f2"`,
	"dashboard":            `"\e99a"`,
	"profile":              `"\e999"`,
	"table":                `"\e998"`,
	"warning":              `"\e997"`,
	"form":                 `"\e996"`,
	"weibo-square":         `"\e6f5"`,
	"weibo-circle":         `"\e6f4"`,
	"taobao-circle":        `"\e6f3"`,
	"html5":                `"\e9c7"`,
	"weibo":                `"\e9c6"`,
	"twitter":              `"\e9c5"`,
	"wechat":               `"\e9c4"`,
	"youtube":              `"\e9c3"`,
	"alipay-circle":        `"\e9c2"`,
	"taobao":               `"\e9c1"`,
	"skype":                `"\e9c0"`,
	"qq":                   `"\e9bf"`,
	"medium-workmark":      `"\e9be"`,
	"gitlab":               `"\e9bd"`,
	"medium":               `"\e9bc"`,
	"linkedin":             `"\e9bb"`,
	"google-plus":          `"\e9ba"`,
	"dropbox":              `"\e9b9"`,
	"facebook":             `"\e9b8"`,
	"codepen":              `"\e9b7"`,
	"amazon":               `"\e9b6"`,
	"google":               `"\e9b5"`,
	"codepen-circle":       `"\e9b4"`,
	"alipay":               `"\e9b3"`,
	"ant-design":           `"\e9b2"`,
	"aliyun":               `"\e9f4"`,
	"zhihu":                `"\e703"`,
	"file-markdown":        `"\e704"`,
	"slack":                `"\e705"`,
	"slack-square":         `"\e706"`,
	"behance":              `"\e707"`,
	"behance-square":       `"\e708"`,
	"dribbble":             `"\e709"`,
	"dribbble-square":      `"\e70a"`,
	"instagram":            `"\e70b"`,
	"yuque":                `"\e70c"`,
}

type Icon struct {
	vecty.Core
	Kind     Kind
	Spin     bool
	CSS      gs.CSSRule
	Style    vecty.Applyer
	Children vecty.MarkupOrChild
	sheet    *gs.Sheet
}

func (i *Icon) Mount() {
	i.sheet.Attach()
}

func (i *Icon) Render() vecty.ComponentOrHTML {
	if i.sheet == nil {
		i.sheet = ui.Registry.NewSheet()
		key := string(i.Kind)
		class := "." + key
		println(icons[key])
		i.sheet.AddRule(icon.Style(class, icons[key], i.Kind == Spin))
		if i.CSS != nil {
			i.sheet.AddRule(i.CSS)
		}
	}
	c := vecty.ClassMap(i.sheet.CLasses.Classes())
	return elem.Italic(
		vecty.Markup(c, i.Style),
		i.Children,
	)
}

func (i *Icon) Unmount() {
	i.sheet.Detach()
}

//All returns a slice of all antd icons.
func All() []Kind {
	return []Kind{
		Addfile, Addfolder, Alipay, AlipayCircle, Aliwangwang, AliwangwangO, Aliyun,
		Amazon, Android, AndroidO, AntDesign, Api, Apple, AppleO, Appstore,
		AppstoreO, AreaChart, ArrowDown, ArrowLeft, ArrowRight, ArrowSalt, ArrowUp,
		ArrowsAlt, Backward, Bank, BarChart, Barcode, Bars, Behance, BehanceSquare,
		Bell, Book, Bulb, Calculator, Calendar, Camera, CameraO, Car,
		CaretCircleDown, CaretCircleLeft, CaretCircleODown, CaretCircleOLeft,
		CaretCircleORight, CaretCircleOUp, CaretCircleRight, CaretCircleUp,
		CaretDown, CaretLeft, CaretRight, CaretUp, Check, CheckCircle, CheckCircleO,
		CheckSquare, CheckSquareO, Chrome, CircleDown, CircleDownO, CircleLeft,
		CircleODown, CircleOLeft, CircleORight, CircleOUp, CircleRight, CircleUp,
		ClockCircle, ClockCircleO, Close, CloseCircle, CloseCircleO, CloseSquare,
		CloseSquareO, Cloud, CloudDownload, CloudDownloadO, CloudO, CloudUpload,
		CloudUploadO, Code, CodeO, Codepen, CodepenCircle, Coffee, Compass, Contacts,
		Copy, Copyright, CreditCard, Cross, CrossCircle, CrossCircleO,
		CustomerService, Customerservice, Dashboard, Database, Delete, Desktop,
		Dingding, DingdingO, Disconnect, Dislike, DislikeO, DotChart, DoubleLeft,
		DoubleRight, Down, DownCircle, DownCircleO, DownSquare, DownSquareO,
		Download, Dribbble, DribbbleSquare, Dropbox, Edit, Ellipsis, Enter,
		Environment, EnvironmentO, Exception, Exclamation, ExclamationCircle,
		ExclamationCircleO, Export, Eye, EyeO, Facebook, FastBackward, FastForward,
		File, FileAdd, FileExcel, FileJpg, FileMarkdown, FilePdf, FilePpt, FileText,
		FileUnknown, FileWord, Filter, Flag, Folder, FolderAdd, FolderOpen, Fork,
		Form, Forward, Frown, FrownCircle, FrownO, Gift, Github, Gitlab, Global,
		Google, GooglePlus, Hdd, Heart, HeartO, Home, Hourglass, Html5, Idcard, Ie,
		Inbox, Info, InfoCircle, InfoCircleO, Instagram, Key, Laptop, Layout, Left,
		LeftCircle, LeftCircleO, LeftSquare, LeftSquareO, Like, LikeO, LineChart,
		Link, Linkedin, Loading, Loading3Quarters, Lock, Login, Logout, Mail, Man,
		MedicineBox, Medium, MediumWorkmark, Meh, MehCircle, MehO, MenuFold,
		MenuUnfold, Message, Minus, MinusCircle, MinusCircleO, MinusSquare,
		MinusSquareO, Mobile, Notification, PaperClip, Pause, PauseCircle,
		PauseCircleO, PayCircle, PayCircleO, Phone, Picture, PieChart, PlayCircle,
		PlayCircleO, Plus, PlusCircle, PlusCircleO, PlusSquare, PlusSquareO,
		Poweroff, Printer, Profile, Pushpin, PushpinO, Qq, Qrcode, Question,
		QuestionCircle, QuestionCircleO, RedEnvelope, Reload, Retweet, Right,
		RightCircle, RightCircleO, RightSquare, RightSquareO, Rocket, Rollback,
		Safety, Save, Scan, Schedule, Search, Select, Setting, Shake, ShareAlt, Shop,
		ShoppingCart, Shrink, Skin, Skype, Slack, SlackSquare, Smile, SmileCircle,
		SmileO, Solution, Sound, Star, StarO, StepBackward, StepForward, Swap,
		SwapLeft, SwapRight, Switcher, Sync, Table, Tablet, Tag, TagO, Tags, TagsO,
		Taobao, TaobaoCircle, Team, ToTop, Tool, Trademark, Trophy, Twitter, Unlock,
		Up, UpCircle, UpCircleO, UpSquare, UpSquareO, Upload, Usb, User, UserAdd,
		UserDelete, UsergroupAdd, UsergroupDelete, VerticleLeft, VerticleRight,
		VideoCamera, Wallet, Warning, Wechat, Weibo, WeiboCircle, WeiboSquare, Wifi,
		Windows, WindowsO, Woman, Youtube, Yuque, Zhihu,
	}
}
