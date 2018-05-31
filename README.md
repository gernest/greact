![](logo.png)
# MadTitan

Inter galactic test runner and test framework for  Go frontend projects.

The name madtitan originate from nickname of Thanos (comics character), he
is obsessed with collecting infinity stones, this project is obsessed with
collecting all the good parts of javascript testing eco system with the go ones
to bring smooth developer experience for frontend works using the Go
programming language.

__warning__ I have actively tested this on darwin (os x), if you have issue
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


## What is this? and why?
I have been interested in frontend web development using the Go programming
language. One challenge that I faced was the way to test my code. 

Things that I wanted in any test solution
 - Go 100% (yep I believe Go is much cooler that js)
 - must look familiar
 - must be fast
 - must run on the browser for tests that need access to the dom
 - must be friendly with go tool chain

Available tools 

 spec                  |go test (`testing` package)| js based (karma & co)
 ----------------------|---------------------------|-----------------------
100% go                | Yes                       | No
Familiar               | Yes                       | Yes
Fast                   | Yes                       | ?
Dom mangling tests     | No                        | Yes
nice with go tools     | Yes                       | No

It is obvious now, no such tool existed to suit my needs so I built such tool
suit my need.

This is the table again with madtitan in it.

 spec                  |go test  | madtitan | js based (karma & co)
 ----------------------|---------|--------- |-----------------------
100% go                | Yes     | Yes      | No
Familiar               | Yes     | Yes      | Yes
Fast                   | Yes     | Yes      | ?
Dom mangling tests     | No      | Yes      | Yes
nice with go tools     | Yes     | Yes      | No


## Installing

We need both the library and the `mad` command

```
go get github.com/gernest/madtitan/cmd/mad
```



## Show me the code 


```go
// Unit test
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

// Integration test
func TestRenderBody() mad.Integration {
	txt := "hello,world"
	return mad.RenderBody("mad.RenderBody",
		func() interface{} {
			return elem.Body(
				vecty.Text(txt),
			)
		},
		mad.It("must have text node", func(t mad.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Error(err)
				}
			}()
			o := js.Global.Get("document").Get("body").Get("textContent").String()
			if o != txt {
				t.Errorf("expected %s got %s", txt, o)
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


# FAQ

## what is unit test ?

For mad tests, unit tests are tests which cover a small chunk of functionality.
They must not include any use of code that requires rendering/interacting with
the browser `dom`

## what is integration test ?

For mad tests. Integration tests are tests which cover a component/components
rendered on the `dom`.

# Credits

This would have never been possible without these projects

- [karma](https://github.com/karma-runner/karma)
- [jasmine](https://github.com/jasmine/jasmine.github.io)
- [chrome-launcher](https://github.com/GoogleChrome/chrome-launcher)
- [gopjerjs](https://github.com/gopherjs/vecty)
- `testing` package from the go standard library.
- And all the projects I picked ideas from that I can't remember.
