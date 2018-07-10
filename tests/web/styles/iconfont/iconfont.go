package iconfont

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/style/iconfont"
)

func TestFont() mad.Test {
	return mad.It("generates  iconfont  class style", func(t mad.T) {
		css := gs.ToString(iconfont.Font())
		expect := `.anticon {
  display:inline-block;
  font-style:normal;
  vertical-align:baseline;
  text-align:center;
  text-transform::none;
  line-height:1;
  text-rendering:optimizeLegibility;
  -webkit-font-smoothing:antialiased;
  -moz-osx-font-smoothing:grayscale;
}
.anticon:before {
  display:block;
  font-family:"anticon" !important;
}`
		if css != expect {
			t.Errorf("expected %s got %s", expect, css)
		}
	})
}

func TestFontStyles() mad.Test {
	return mad.It("generates css for all antd icons", func(t mad.T) {
		css := gs.ToString(iconfont.FontStyles())
		expect := `.anticon-addfile:before {
  content:"\e910";
}
.anticon-addfolder:before {
  content:"\e914";
}
.anticon-alipay:before {
  content:"\e9b3";
}
.anticon-alipay-circle:before {
  content:"\e9c2";
}
.anticon-aliwangwang:before {
  content:"\e68e";
}
.anticon-aliwangwang-o:before {
  content:"\e68f";
}
.anticon-aliyun:before {
  content:"\e9f4";
}
.anticon-amazon:before {
  content:"\e9b6";
}
.anticon-android:before {
  content:"\e938";
}
.anticon-android-o:before {
  content:"\e68d";
}
.anticon-ant-design:before {
  content:"\e9b2";
}
.anticon-api:before {
  content:"\e951";
}
.anticon-apple:before {
  content:"\e68c";
}
.anticon-apple-o:before {
  content:"\e6d4";
}
.anticon-appstore:before {
  content:"\e696";
}
.anticon-appstore-o:before {
  content:"\e695";
}
.anticon-area-chart:before {
  content:"\e65c";
}
.anticon-arrow-down:before {
  content:"\e619";
}
.anticon-arrow-left:before {
  content:"\e61c";
}
.anticon-arrow-right:before {
  content:"\e61b";
}
.anticon-arrow-salt:before {
  content:"\e615";
}
.anticon-arrow-up:before {
  content:"\e61a";
}
.anticon-arrows-alt:before {
  content:"\e615";
}
.anticon-backward:before {
  content:"\e603";
}
.anticon-bank:before {
  content:"\e6ee";
}
.anticon-bar-chart:before {
  content:"\e6be";
}
.anticon-barcode:before {
  content:"\e652";
}
.anticon-bars:before {
  content:"\e639";
}
.anticon-behance:before {
  content:"\e707";
}
.anticon-behance-square:before {
  content:"\e708";
}
.anticon-bell:before {
  content:"\e64e";
}
.anticon-book:before {
  content:"\e637";
}
.anticon-bulb:before {
  content:"\e649";
}
.anticon-calculator:before {
  content:"\e6aa";
}
.anticon-calendar:before {
  content:"\e6bb";
}
.anticon-camera:before {
  content:"\e689";
}
.anticon-camera-o:before {
  content:"\e68a";
}
.anticon-car:before {
  content:"\e6dc";
}
.anticon-caret-circle-down:before {
  content:"\e60b";
}
.anticon-caret-circle-left:before {
  content:"\e609";
}
.anticon-caret-circle-o-down:before {
  content:"\e60f";
}
.anticon-caret-circle-o-left:before {
  content:"\e60d";
}
.anticon-caret-circle-o-right:before {
  content:"\e60c";
}
.anticon-caret-circle-o-up:before {
  content:"\e60e";
}
.anticon-caret-circle-right:before {
  content:"\e608";
}
.anticon-caret-circle-up:before {
  content:"\e60a";
}
.anticon-caret-down:before {
  content:"\e606";
}
.anticon-caret-left:before {
  content:"\e605";
}
.anticon-caret-right:before {
  content:"\e604";
}
.anticon-caret-up:before {
  content:"\e607";
}
.anticon-check:before {
  content:"\e632";
}
.anticon-check-circle:before {
  content:"\e630";
}
.anticon-check-circle-o:before {
  content:"\e631";
}
.anticon-check-square:before {
  content:"\e6c4";
}
.anticon-check-square-o:before {
  content:"\e6c5";
}
.anticon-chrome:before {
  content:"\e6ac";
}
.anticon-circle-down:before {
  content:"\e60b";
}
.anticon-circle-down-:before {
  content:"\e694";
}
.anticon-circle-down-o:before {
  content:"\e693";
}
.anticon-circle-left:before {
  content:"\e609";
}
.anticon-circle-o-down:before {
  content:"\e60f";
}
.anticon-circle-o-left:before {
  content:"\e60d";
}
.anticon-circle-o-right:before {
  content:"\e60c";
}
.anticon-circle-o-up:before {
  content:"\e60e";
}
.anticon-circle-right:before {
  content:"\e608";
}
.anticon-circle-up:before {
  content:"\e60a";
}
.anticon-clock-circle:before {
  content:"\e640";
}
.anticon-clock-circle-o:before {
  content:"\e641";
}
.anticon-close:before {
  content:"\e633";
}
.anticon-close-circle:before {
  content:"\e62e";
}
.anticon-close-circle-o:before {
  content:"\e62f";
}
.anticon-close-square:before {
  content:"\e6c2";
}
.anticon-close-square-o:before {
  content:"\e6c3";
}
.anticon-cloud:before {
  content:"\e680";
}
.anticon-cloud-download:before {
  content:"\e682";
}
.anticon-cloud-download-o:before {
  content:"\e683";
}
.anticon-cloud-o:before {
  content:"\e67f";
}
.anticon-cloud-upload:before {
  content:"\e681";
}
.anticon-cloud-upload-o:before {
  content:"\e684";
}
.anticon-code:before {
  content:"\e6bf";
}
.anticon-code-o:before {
  content:"\e636";
}
.anticon-codepen:before {
  content:"\e9b7";
}
.anticon-codepen-circle:before {
  content:"\e9b4";
}
.anticon-coffee:before {
  content:"\e6e8";
}
.anticon-compass:before {
  content:"\e6db";
}
.anticon-contacts:before {
  content:"\e6f0";
}
.anticon-copy:before {
  content:"\e648";
}
.anticon-copyright:before {
  content:"\e6de";
}
.anticon-credit-card:before {
  content:"\e635";
}
.anticon-cross:before {
  content:"\e633";
}
.anticon-cross-circle:before {
  content:"\e62e";
}
.anticon-cross-circle-o:before {
  content:"\e62f";
}
.anticon-customer-service:before {
  content:"\e634";
}
.anticon-customerservice:before {
  content:"\e634";
}
.anticon-dashboard:before {
  content:"\e99a";
}
.anticon-database:before {
  content:"\e650";
}
.anticon-delete:before {
  content:"\e69f";
}
.anticon-desktop:before {
  content:"\e6b4";
}
.anticon-dingding:before {
  content:"\e923";
}
.anticon-dingding-o:before {
  content:"\e925";
}
.anticon-disconnect:before {
  content:"\e64f";
}
.anticon-dislike:before {
  content:"\e64b";
}
.anticon-dislike-o:before {
  content:"\e69e";
}
.anticon-dot-chart:before {
  content:"\e6bd";
}
.anticon-double-left:before {
  content:"\e618";
}
.anticon-double-right:before {
  content:"\e617";
}
.anticon-down:before {
  content:"\e61d";
}
.anticon-down-circle:before {
  content:"\e60b";
}
.anticon-down-circle-o:before {
  content:"\e60f";
}
.anticon-down-square:before {
  content:"\e6c9";
}
.anticon-down-square-o:before {
  content:"\e6ce";
}
.anticon-download:before {
  content:"\e6b7";
}
.anticon-dribbble:before {
  content:"\e709";
}
.anticon-dribbble-square:before {
  content:"\e70a";
}
.anticon-dropbox:before {
  content:"\e9b9";
}
.anticon-edit:before {
  content:"\e692";
}
.anticon-ellipsis:before {
  content:"\e647";
}
.anticon-enter:before {
  content:"\e6a0";
}
.anticon-environment:before {
  content:"\e685";
}
.anticon-environment-o:before {
  content:"\e686";
}
.anticon-exception:before {
  content:"\e665";
}
.anticon-exclamation:before {
  content:"\e62b";
}
.anticon-exclamation-circle:before {
  content:"\e62c";
}
.anticon-exclamation-circle-o:before {
  content:"\e62d";
}
.anticon-export:before {
  content:"\e691";
}
.anticon-eye:before {
  content:"\e687";
}
.anticon-eye-o:before {
  content:"\e688";
}
.anticon-facebook:before {
  content:"\e9b8";
}
.anticon-fast-backward:before {
  content:"\e6c6";
}
.anticon-fast-forward:before {
  content:"\e6c7";
}
.anticon-file:before {
  content:"\e664";
}
.anticon-file-add:before {
  content:"\e910";
}
.anticon-file-excel:before {
  content:"\e6b0";
}
.anticon-file-jpg:before {
  content:"\e69c";
}
.anticon-file-markdown:before {
  content:"\e704";
}
.anticon-file-pdf:before {
  content:"\e6b3";
}
.anticon-file-ppt:before {
  content:"\e6b1";
}
.anticon-file-text:before {
  content:"\e698";
}
.anticon-file-unknown:before {
  content:"\e6af";
}
.anticon-file-word:before {
  content:"\e6b2";
}
.anticon-filter:before {
  content:"\e663";
}
.anticon-flag:before {
  content:"\e655";
}
.anticon-folder:before {
  content:"\e662";
}
.anticon-folder-add:before {
  content:"\e914";
}
.anticon-folder-open:before {
  content:"\e699";
}
.anticon-fork:before {
  content:"\e6f2";
}
.anticon-form:before {
  content:"\e996";
}
.anticon-forward:before {
  content:"\e602";
}
.anticon-frown:before {
  content:"\e646";
}
.anticon-frown-circle:before {
  content:"\e646";
}
.anticon-frown-o:before {
  content:"\e6a9";
}
.anticon-gift:before {
  content:"\e6e4";
}
.anticon-github:before {
  content:"\e6ad";
}
.anticon-gitlab:before {
  content:"\e9bd";
}
.anticon-global:before {
  content:"\e6f1";
}
.anticon-google:before {
  content:"\e9b5";
}
.anticon-google-plus:before {
  content:"\e9ba";
}
.anticon-hdd:before {
  content:"\e69a";
}
.anticon-heart:before {
  content:"\e6a3";
}
.anticon-heart-o:before {
  content:"\e6a4";
}
.anticon-home:before {
  content:"\e65e";
}
.anticon-hourglass:before {
  content:"\e653";
}
.anticon-html5:before {
  content:"\e9c7";
}
.anticon-idcard:before {
  content:"\e6e5";
}
.anticon-ie:before {
  content:"\e69b";
}
.anticon-inbox:before {
  content:"\e67a";
}
.anticon-info:before {
  content:"\e62a";
}
.anticon-info-circle:before {
  content:"\e628";
}
.anticon-info-circle-o:before {
  content:"\e629";
}
.anticon-instagram:before {
  content:"\e70b";
}
.anticon-key:before {
  content:"\e654";
}
.anticon-laptop:before {
  content:"\e65f";
}
.anticon-layout:before {
  content:"\e656";
}
.anticon-left:before {
  content:"\e620";
}
.anticon-left-circle:before {
  content:"\e609";
}
.anticon-left-circle-o:before {
  content:"\e60d";
}
.anticon-left-square:before {
  content:"\e6ca";
}
.anticon-left-square-o:before {
  content:"\e6cd";
}
.anticon-like:before {
  content:"\e64c";
}
.anticon-like-o:before {
  content:"\e69d";
}
.anticon-line-chart:before {
  content:"\e65d";
}
.anticon-link:before {
  content:"\e65b";
}
.anticon-linkedin:before {
  content:"\e9bb";
}
.anticon-loading:before {
  content:"\e64d";
}
.anticon-loading-3-quarters:before {
  content:"\e6ae";
}
.anticon-lock:before {
  content:"\e67b";
}
.anticon-login:before {
  content:"\e657";
}
.anticon-logout:before {
  content:"\e65a";
}
.anticon-mail:before {
  content:"\e659";
}
.anticon-man:before {
  content:"\e6e2";
}
.anticon-medicine-box:before {
  content:"\e6e6";
}
.anticon-medium:before {
  content:"\e9bc";
}
.anticon-medium-workmark:before {
  content:"\e9be";
}
.anticon-meh:before {
  content:"\e666";
}
.anticon-meh-circle:before {
  content:"\e666";
}
.anticon-meh-o:before {
  content:"\e667";
}
.anticon-menu-fold:before {
  content:"\e9ac";
}
.anticon-menu-unfold:before {
  content:"\e9ad";
}
.anticon-message:before {
  content:"\e6ab";
}
.anticon-minus:before {
  content:"\e624";
}
.anticon-minus-circle:before {
  content:"\e622";
}
.anticon-minus-circle-o:before {
  content:"\e623";
}
.anticon-minus-square:before {
  content:"\e6c1";
}
.anticon-minus-square-o:before {
  content:"\e621";
}
.anticon-mobile:before {
  content:"\e678";
}
.anticon-notification:before {
  content:"\e677";
}
.anticon-paper-clip:before {
  content:"\e676";
}
.anticon-pause:before {
  content:"\e63d";
}
.anticon-pause-circle:before {
  content:"\e63e";
}
.anticon-pause-circle-o:before {
  content:"\e63f";
}
.anticon-pay-circle:before {
  content:"\e6a5";
}
.anticon-pay-circle-o:before {
  content:"\e6a6";
}
.anticon-phone:before {
  content:"\e675";
}
.anticon-picture:before {
  content:"\e674";
}
.anticon-pie-chart:before {
  content:"\e6b8";
}
.anticon-play-circle:before {
  content:"\e6d0";
}
.anticon-play-circle-o:before {
  content:"\e6d1";
}
.anticon-plus:before {
  content:"\e627";
}
.anticon-plus-circle:before {
  content:"\e626";
}
.anticon-plus-circle-o:before {
  content:"\e625";
}
.anticon-plus-square:before {
  content:"\e6c0";
}
.anticon-plus-square-o:before {
  content:"\e645";
}
.anticon-poweroff:before {
  content:"\e6d5";
}
.anticon-printer:before {
  content:"\e673";
}
.anticon-profile:before {
  content:"\e999";
}
.anticon-pushpin:before {
  content:"\e6a2";
}
.anticon-pushpin-o:before {
  content:"\e6a1";
}
.anticon-qq:before {
  content:"\e9bf";
}
.anticon-qrcode:before {
  content:"\e67c";
}
.anticon-question:before {
  content:"\e63a";
}
.anticon-question-circle:before {
  content:"\e63b";
}
.anticon-question-circle-o:before {
  content:"\e63c";
}
.anticon-red-envelope:before {
  content:"\e6e7";
}
.anticon-reload:before {
  content:"\e616";
}
.anticon-retweet:before {
  content:"\e613";
}
.anticon-right:before {
  content:"\e61f";
}
.anticon-right-circle:before {
  content:"\e608";
}
.anticon-right-circle-o:before {
  content:"\e60c";
}
.anticon-right-square:before {
  content:"\e6cb";
}
.anticon-right-square-o:before {
  content:"\e6cc";
}
.anticon-rocket:before {
  content:"\e90f";
}
.anticon-rollback:before {
  content:"\e612";
}
.anticon-safety:before {
  content:"\e6ea";
}
.anticon-save:before {
  content:"\e669";
}
.anticon-scan:before {
  content:"\e697";
}
.anticon-schedule:before {
  content:"\e6df";
}
.anticon-search:before {
  content:"\e670";
}
.anticon-select:before {
  content:"\e64a";
}
.anticon-setting:before {
  content:"\e672";
}
.anticon-shake:before {
  content:"\e94f";
}
.anticon-share-alt:before {
  content:"\e671";
}
.anticon-shop:before {
  content:"\e6e3";
}
.anticon-shopping-cart:before {
  content:"\e668";
}
.anticon-shrink:before {
  content:"\e614";
}
.anticon-skin:before {
  content:"\e6d8";
}
.anticon-skype:before {
  content:"\e9c0";
}
.anticon-slack:before {
  content:"\e705";
}
.anticon-slack-square:before {
  content:"\e706";
}
.anticon-smile:before {
  content:"\e6a7";
}
.anticon-smile-circle:before {
  content:"\e6a7";
}
.anticon-smile-o:before {
  content:"\e6a8";
}
.anticon-solution:before {
  content:"\e66f";
}
.anticon-sound:before {
  content:"\e6e9";
}
.anticon-star:before {
  content:"\e660";
}
.anticon-star-o:before {
  content:"\e661";
}
.anticon-step-backward:before {
  content:"\e601";
}
.anticon-step-forward:before {
  content:"\e600";
}
.anticon-swap:before {
  content:"\e642";
}
.anticon-swap-left:before {
  content:"\e643";
}
.anticon-swap-right:before {
  content:"\e644";
}
.anticon-switcher:before {
  content:"\e913";
}
.anticon-sync:before {
  content:"\e6da";
}
.anticon-table:before {
  content:"\e998";
}
.anticon-tablet:before {
  content:"\e66e";
}
.anticon-tag:before {
  content:"\e6d2";
}
.anticon-tag-o:before {
  content:"\e6d3";
}
.anticon-tags:before {
  content:"\e67d";
}
.anticon-tags-o:before {
  content:"\e67e";
}
.anticon-taobao:before {
  content:"\e9c1";
}
.anticon-taobao-circle:before {
  content:"\e6f3";
}
.anticon-team:before {
  content:"\e66d";
}
.anticon-to-top:before {
  content:"\e66c";
}
.anticon-tool:before {
  content:"\e6d9";
}
.anticon-trademark:before {
  content:"\e651";
}
.anticon-trophy:before {
  content:"\e6ef";
}
.anticon-twitter:before {
  content:"\e9c5";
}
.anticon-unlock:before {
  content:"\e6ba";
}
.anticon-up:before {
  content:"\e61e";
}
.anticon-up-circle:before {
  content:"\e60a";
}
.anticon-up-circle-o:before {
  content:"\e60e";
}
.anticon-up-square:before {
  content:"\e6c8";
}
.anticon-up-square-o:before {
  content:"\e6cf";
}
.anticon-upload:before {
  content:"\e6b6";
}
.anticon-usb:before {
  content:"\e6d7";
}
.anticon-user:before {
  content:"\e66a";
}
.anticon-user-add:before {
  content:"\e6ed";
}
.anticon-user-delete:before {
  content:"\e6e0";
}
.anticon-usergroup-add:before {
  content:"\e6dd";
}
.anticon-usergroup-delete:before {
  content:"\e6e1";
}
.anticon-verticle-left:before {
  content:"\e610";
}
.anticon-verticle-right:before {
  content:"\e611";
}
.anticon-video-camera:before {
  content:"\e66b";
}
.anticon-wallet:before {
  content:"\e6eb";
}
.anticon-warning:before {
  content:"\e997";
}
.anticon-wechat:before {
  content:"\e9c4";
}
.anticon-weibo:before {
  content:"\e9c6";
}
.anticon-weibo-circle:before {
  content:"\e6f4";
}
.anticon-weibo-square:before {
  content:"\e6f5";
}
.anticon-wifi:before {
  content:"\e6d6";
}
.anticon-windows:before {
  content:"\e68b";
}
.anticon-windows-o:before {
  content:"\e6bc";
}
.anticon-woman:before {
  content:"\e6ec";
}
.anticon-youtube:before {
  content:"\e9c3";
}
.anticon-yuque:before {
  content:"\e70c";
}
.anticon-zhihu:before {
  content:"\e703";
}
.anticon-spin:before {
  display:inline-block;;
  animation:loadingCircle 1s infinite linear;
}`
		if css != expect {
			t.Errorf("expected %s got %s", expect, css)
		}
	})
}

func TestFontFace() mad.Test {
	return mad.It("generates css for icons fontface", func(t mad.T) {
		css := gs.ToString(iconfont.FontFace())
		expect := `@font-face {
  font-family:'anticon';
  src:url('https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i.woff')format('woff'),url('https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i.ttf')format('truetype'),url('https://at.alicdn.com/t/font_148784_v4ggb6wrjmkotj4i.iconfont')format('svg');
}`
		if css != expect {
			t.Error("expected %s got %s", expect, css)
		}
	})
}
