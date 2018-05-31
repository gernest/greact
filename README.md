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

This inspired me to build such tool.

## Installing

We need both the library and the `mad` command

```
go get github.com/gernest/madtitan/cmd/mad
```


## Getting started

Create a tests directory in the root of your go package. Then write test
functions. There is no convention over the filenames.


## Show me the code 


```go
// Unit tests
func TestBrowser() mad.Test {
	return mad.List{
		mad.Describe("Prefixes",
			mad.It("must contain all browser vendor prefixes", func(t mad.T) {
				b := prefix.NewBrowser()
				expect := []string{"-moz-", "-ms-", "-o-", "-webkit-"}
				g := b.Prefixcache
				for k, v := range expect {
					if g[k] != v {
						t.Fatalf("expected %v got %v", expect, g)
					}
				}
			}),
		),
		mad.Describe("WithPrefix",
			mad.It("must return true when the prefix exist", func(t mad.T) {
				s := "1 -o-calc(1)"
				b := prefix.NewBrowser()
				if !b.WithPrefix(s) {
					t.Error("expected to be true")
				}
			}),
			mad.It("must return false when the prefix does not exist", func(t mad.T) {
				s := "1 calc(1)"
				b := prefix.NewBrowser()
				if b.WithPrefix(s) {
					t.Error("expected to be false")
				}
			}),
		),
	}
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

check [this  project tests for inspiration](https://github.com/gernest/gs)

## I want to write unit tests

We got you covered [Take a look at this page](unit_test.md)

## I want to write integration tests 

__NOTE__ This only works with vecty projects

We got you covered [Take a look at this page](integration_test.md)


# FAQ

## what is unit test ?

For `mad` , unit tests are tests which cover a small chunk of functionality.
They must not include any use of code that requires rendering/interacting with
the browser `dom`

## what is integration test ?

For `mad` . Integration tests are tests which cover a component/components
rendered on the `dom`.

# Credits

This would have never been possible without these projects

- [karma](https://github.com/karma-runner/karma)
- [jasmine](https://github.com/jasmine/jasmine.github.io)
- [chrome-launcher](https://github.com/GoogleChrome/chrome-launcher)
- [gopjerjs](https://github.com/gopherjs/vecty)
- `testing` package from the go standard library.
- And all the projects I picked ideas from that I can't remember.
