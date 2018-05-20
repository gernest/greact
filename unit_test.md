# Unit testing

madtitan borrows from go's standard library's`testing` package and some ideas
from many javascript test packages/runners such as jasmine,jest, karma-runner
 etc.

Tests are defined as any function with the signature 

```go
func TestXxx() mad.T
```

You can the use `Describe` and `It` functions to compose your test suite.

## Describe

This is a container, meant to group/structure your tests.

Describing a test case is done with the `mad.Describe` function, this takes
the description string as first argument and arbitrary number of test cases as the
second argument. The description string is a human readable string stating whatis being tested.

Use this to describe what you are testing

For example , given we have a function called `Rainfall` which makes it rain
and we want to test it.

We can start describing our unit test like this

```go
func TestRainfall() mad.Test {
	return mad.Describe("Raining")
}
```

Here we are stating that we are testing Raining feature.

At the moment our test doesn't do anything. We are just describing a Rainy day.

Let's define our function

```go
// RainContext defines details about the rainfall.
type RainContext struct {

	//True when the day is cloudy and false otherwise.
	Cloudy bool

	// The current maonth, example january,february ...december.
	Month string
}

// Rainfall returns rainfall status
func Rainfall() *RainContext {
	return &RainContext{
		Cloudy: true,
		Month:  "november",
	}
}
```


## Expectation

This is where we verify the behavior of our functions. We use `It` for stating
our expectations.

It has the following signature
```go
func It(desc string, fn func(mad.T)) mad.Test {
``
What the user of the library does is supply the `fn` parameter. First parameter
is a string which tells what expectation we are trying to achieve.

Inside the passed `fn` body you can call `T.Error`, `T.Errorf` to signal failure. You
can call them as many times as you want inside your functions and they will be
included in the failed test report. If you want to halt the function execution
then you can use `T.Fatal` or `T.Falatf` as you can see the concept were
borrowed from the `testing` package.


We have seen how to describe what we are testing. Now we will see how to assert
expectation from our functions/piece of code.

We use `mad.It` to state our expectation. It is up for the one writing tests
to determine if the code meets expectation or not, I will explain this with
code in a moment.

 Let us say, it is only raining whenever there are clouds. We can now update our
 Rainfall test case to look like this.

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

```go
ctx := Rainfall()  #here we execute the function that we want to unit test 
```

```
			if !ctx.Cloudy {
				t.Error("expected to be cloudy")
            }
```

We are comparing the current behavior/output with our expectation. This test
will pass because `Rainfall` function sets `Cloudy` to be true, hence it will
meet expectations.

madtitan favors composition, so build your test suite to suit your needs from
smaller modular functional components.

__Good  luck__ and happy unit testing

