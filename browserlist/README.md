port of [browserlist](https://github.com/browserslist/browserslist)

support

- `> 5%` versions selected by global usage statistics. >=, < and <= work too.
- `cover 99.5%`: most popular browsers that provide coverage.
- `Firefox > 20`: versions of Firefox newer than 20. >=, < and <= work too.
- `iOS 7`: the iOS browser version 7 directly.
- `unreleased versions` or `unreleased Chrome versions`: alpha and beta versions.
- `last 2 major versions` or `last 2 iOS major versions`: all minor/patch releases of last 2 major versions.
- `since 2015` or `last 2 years`: all versions released since year 2015 (also since 2015-03 and since 2015-03-10).
- `dead`: browsers from `last 2 version` query, but with less than 0.5% in global usage statistics and without official support or updates for 24 months. Right now it is IE 10, IE_Mob 10, BlackBerry 10, BlackBerry 7, and OperaMobile 12.1.
- `last 2 versions`: the last 2 versions for each browser.
- `last 2 Chrome versions`: the last 2 versions of Chrome browser.
- `defaults`: Browserslist’s default browsers (> 0.5%, last 2 versions, Firefox ESR, not dead).
- `not ie <= 8`: exclude browsers selected by previous queries
