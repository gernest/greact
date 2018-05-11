# Unit testing

Promnight borrows from go's standard `testing` package and some ideas from behavior driven development(BDD).

Tests are defined as any function with the signature 

```go
func TestXxx() prom.Test
```

So, tests are just functions which can be composed in very interesting ways.

## Describe
Describing a test case is done with the `prom.Describe` function, this takes
the description as first argument and arbitrary number of test cases as the
second argument.

Use this to describe what you are testing

For example , given we have a function callled `Rainfall` which makes it rain
and we want to test it.

We can start describing our unit test like this

```go
func TestRainfall() prom.Test {
	return prom.Describe("Raining")
}
```

At the moment out test doesn't do anything. We are just describing a Rainy day.

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
We have seen how to describe what we are testing. Now we will see how to assert
expectation from our functions/piece of code.

We use `prom.It` to state our expectation. It is up for the one writing tests
to determine if the code meets expectation or now, I will explain this with
code in a moment.

 Let's say, it is only raining when there are clouds. We can now update our
 Rainfall test case to look like this.

 ```go
 func TestRainfall() prom.Test {
	return prom.Describe("Raining",
		prom.It("must be cloudy", func(rs prom.Result) {
			ctx := Rainfall()
			if !ctx.Cloudy {
				rs.Error("expected to be cloudy")
			}
		}),
	)
}
```

`ctx := Rainfall()`  here we run the code that we want to unit test and here

```go
			if !ctx.Cloudy {
				rs.Error("expected to be cloudy")
            }
```
We are comparing the current behavior/output with our expectation. This test will pass because `Rainfall` function sets `Cloudy` to be true, hence it will meet expectations.

Promnight favors composition, so build your test suite to suit your needs from
smaller modular functional components.

__Good  luck__ and happy unit testing

