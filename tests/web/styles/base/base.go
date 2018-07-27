package base

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/style/core/base"
)

func TestBase() mad.Test {
	return mad.It("generates normalize css", func(t mad.T) {
		css := gs.ToString(base.Base())
		if css != expectedBaseStyle {
			t.Error("got wrong styles")
		}
	})
}

const expectedBaseStyle = `@font-face {
  font-family:"Monospaced Number";
  src:local("Tahoma");
  unicode-range:U+30-39;
}
@font-face {
  font-family:"Monospaced Number";
  font-weight:bold;
  src:local("Tahoma-Bold");
  unicode-range:U+30-39;
}
@font-face {
  font-family:"Chinese Quote";
  font-weight:bold;
  src:local("PingFang SC"), local("SimSun");
  unicode-range:U+2018, U+2019, U+201c, U+201d;
}
html, body {
  width:100%;
  height:100%;
}
input::-ms-clear, input::-ms-reveal {
  display:none;
}
*, *:::before, *::after {
  box-sizing:border-box;
}
html {
  font-family:sans-serif;
  line-height:1.15;
  -webkit-text-size-adjust:100%;
  -ms-text-size-adjust:100%;
  -ms-overflow-style:scrollbar;
  -webkit-tap-highlight-color:rgba(0, 0, 0, 0);
}
@-ms-viewport  {
  width:device-width;
}
article, aside, dialog, figcaption, figure, footer, header, hgroup, main, nav, section {
  display:block;
}
body {
  margin:0;
  font-family:"Monospaced Number","Chinese Quote", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size:14px;
  line-height:1.5;
  color:rgba(0,0,0,0.65);
  background-color:rgb(255,255,255);
}
[tabindex="-1"]:focus {
  outline:one !important;
}
hr {
  box-sizing:content-box;
  height:0;
  overflow:visible;
}
h1, h2, h3, h4, h5, h6 {
  margin-top:0;
  margin-bottom::.5em;
  color:rgba(0,0,0,0.85);
  font-weight:500;
}
p {
  margin-top:0;
  margin-bottom:1em;
}
abbr[title], abbr[data-original-title] {
  text-decoration:underline;
  text-decoration:underline dotted;
  cursor:help;
  border-bottom:0;
}
address {
  margin-bottom:1em;
  font-style:normal;
  line-height:inherit;
}
input[type="text"], input[type="password"], input[type="number"], textarea {
  -webkit-appearance:none;
}
ol, ul, dl, {
  margin-top:0;
  margin-bottom:1em;
}
ol ol, ul ul, ol ul, ul ol {
  margin-bottom:1em;
}
dt {
  font-weight:500;
}
dd {
  margin-bottom:.5em;
  margin-left:0;
}
blockquote {
  margin: 0 1em;
}
dfn {
  font-style:italic;
}
b, strong {
  font-weight:bolder;
}
small {
  font-size:80%;
}
sub, sup {
  position:relative;
  font-size:75%;
  line-height:0;
  vertical-align:baseline;
}
sub {
  bottom:-.25em;
}
sup {
  top:-.5em;
}
a {
  color:rgb(24,144,255);
  background-color:transparent;
  text-decoration:none;
  outline:none;
  cursor:pointer;
  transition:color .3s;
  -webkit-text-decoration-skip:objects;
}
a:focus {
  text-decoration:underline;
  text-decoration-skip:ink;
}
a:hover {
  color:#40a9ff;
}
a:active {
  color:#096dd9;
}
a:active, a:hover {
  outline:0;
  text-decoration:none;
}
a[disabled] {
  color:rgba(0,0,0,0.25);
  cursor:not-allowed;
  pointer-events:none;
}
pre, code, kbd, samp {
  font-family:: Consolas, Menlo, Courier, monospace;
  font-size:1em;
}
pre {
  margin-to:0;
  margin-bottom:1em;
  overflow:auto;
}
figure {
  margin:0 0 1em;
}
img {
  vertical-align:middle;
  border-style:none;
}
svg:not(:root) {
  overflow:hidden;
}
a, area, button, [role="button"], input:not([type=range]), label, select, summary, textarea {
  touch-action:manipulation;
}
table {
  border-collapse:collapse;
}
caption {
  padding-top:.75em;
  padding-bottom:.3em;
  color:rgba(0,0,0,0.45);
  text-align:left;
  caption-side:bottom;
}
th {
  text-align:inherit;
}
input, button, select, optgroup, textarea {
  margin:0;
  font-family:inherit;
  font-size::inherit;
  line-height:inherit;
}
button, input {
  overflow:visible;
}
button, select {
  text-transform:none;
}
button, html [type="button"], [type="reset"], [type="submit"] {
  -webkit-appearance:button;
}
button::-moz-focus-inner, [type="button"]::-moz-focus-inner, [type="reset"]::-moz-focus-inner, [type="submit"]::-moz-focus-inner {
  padding:0;
  border-style::none;
}
input[type="radio"], input[type="checkbox"]  {
  box-sizing:border-box;
  padding:0;
}
input[type="date"], input[type="datetime-local"], input[type="month"], {
  -webkit-appearance:listbox;
}
textarea {
  overflow:auto;
  resize:vertical;
}
fieldset {
  min-width:0;
  padding:0;
  margin:0;
  border:0;
}
legend {
  display:block;
  width:100%;
  max-width:100%;
  padding:0;
  margin-bottom:.5em;
  font-size:1.5em;
  line-height:inherit;
  white-space:normal;
}
progress {
  vertical-align:baseline;
}
[type="number"]::-webkit-inner-spin-button, [type="number"]::-webkit-outer-spin-button {
  height:auto;
}
[type="search"] {
  outline-offset:-2px;
  -webkit-appearance:none;
}
[type="search"]::-webkit-search-cancel-button, [type="search"]::-webkit-search-decoration {
  -webkit-appearance:none;
}
::-webkit-file-upload-button {
  font:inherit;
  -webkit-appearance:button;
}
output {
  display:inline-block;
}
summary {
  display:list-item;
}
template {
  display:none;
}
[hidden] {
  display:none !important;
}
mark {
  padding:.2em;
  background-color:rgb(254,255,230);
}
::selection {
  background:rgb(24,144,255);
  color:#fff;
}
.clearfix {
  zoom:1;
}
.clearfix:before {
  content:"";
  display:table;
}
.clearfix:after {
  content:"";
  display:table;
  clear:both;
  visibility:hidden;
  font-size:0;
  height:0;
}`
