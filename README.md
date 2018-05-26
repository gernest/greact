![](logo.png)
# MadTitan

Inter galactic test runner and test framework for  Go frontend projects.

The name madtitan originate from nickname of Thanos (comics character), he
is obsessed with collecting infinity stones, this project is obsessed with
collecting all the good parts of javascript testing eco system with the go ones
to bring smooth developer experience for frontend works using the Go
programming language.

 __ warning__ I have actively tested this on darwin (os x), if you have issue
with a different os open it here and probably someone else who uses your os
might help.

## Features

- [x] Fast 
- [x] Combines good parts from the go testing package, karma,jest,jasmine etc
- [x] Write unit tests.
- [x] Write integration tests (vecty only), yep now you can test your vecty components.
- [x] Code coverage
 - [x] Like thanos who aims to destroy half of the universe, we will destroy half
 of your frontend project's problems by ensuring it meets your expectations.

## Supported browsers

Note that theoretically all browser's supporting the chrome debug protocol must
work. Unfortunate I'm only interested in chrome so if you want to take a stab
at testing with another browser you are welcome.

- [x] chrome
- [ ] Add a browser



## Show me the code 

```go
 func TestRainfall() mad.Test {
	return mad.Describe("Raining",
		mad.It("must be cloudy", func(t mad.T) {
			ctx := Rainfall()
			if !ctx.Cloudy {
				t.Error("expected to be cloudy")
			}
		}),
	)
}
```

## I want to write unit tests

We got you covered [Take a look at this page](unit_test.md)

## I want to write integration tests 

__NOTE__ This only works with vecty projects

We got you covered [Take a look at this page](integration_test.md)


# Credits

This would have never been possible without these projects

- [karma](https://github.com/karma-runner/karma)
- [jasmine](https://github.com/jasmine/jasmine.github.io)
- [chrome-launcher](https://github.com/GoogleChrome/chrome-launcher)
- [gopjerjs](https://github.com/gopherjs/vecty)
- `testing` package from the go standard library.
- And all the projects I picked ideas from that I can't remember.
