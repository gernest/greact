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

// Icon is ant design icon component which implements vecty.Component interface.
// The icons rendered are not styled so it is up to the user to add styles.
//
// Remember to add
//
// @font-face {
// 	font-family: 'anticon';
// 	src: url('https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i.eot');
// 	/* IE9*/
// 	src: url('https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i.woff') format('woff'), /* chrome、firefox、opera、Safari, Android, iOS 4.2+*/
// 	url('https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i.ttf') format('truetype'), /* iOS 4.1- */
// 	url('https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i.svg#iconfont') format('svg');
// }
//
// On the html page if you want the icons to be visible.
// TODO: figure out how to add that icon fontface using sheet.AddRules.
type Icon struct {
	vecty.Core

	// Kind is a string representing the name of the icon.
	Kind     Kind
	CSS      gs.CSSRule
	Style    vecty.Applyer
	Children func() vecty.MarkupOrChild
	sheet    *gs.Sheet
}

// Mount attches component stylest
func (i *Icon) Mount() {
	i.sheet.Attach()
}

func (i *Icon) Render() vecty.ComponentOrHTML {
	if i.sheet == nil {
		i.sheet = ui.Registry.NewSheet()
		i.sheet.AddRule(icon.Style1(string(i.Kind)))
	}
	var customStyles vecty.ClassMap
	if i.CSS != nil {
		customStyles = vecty.ClassMap(i.sheet.AddRule(i.CSS).Classes())
	}
	root := i.sheet.CLasses[icon.Root]
	k := join(icon.Root, "-", string(i.Kind))
	k = i.sheet.CLasses[k]
	c := vecty.ClassMap{
		toClass(root): true,
		toClass(k):    true,
	}
	return elem.Italic(
		vecty.Markup(c, customStyles, i.Style),
		i.getChildren(),
	)
}
func toClass(name string) string {
	if name == "" {
		return ""
	}
	if name[0] == '.' {
		return name[1:]
	}
	return name
}
func join(v ...string) string {
	o := ""
	for _, s := range v {
		o += s
	}
	return o
}

func (i *Icon) getChildren() vecty.MarkupOrChild {
	if i.Children != nil {
		return i.Children()
	}
	return nil
}

// Unmount detach component stylest
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
