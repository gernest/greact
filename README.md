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
- [ ] firefox (not tested)
- [ ] Add a browser


## How it works

This is what happens when you run `mad test`

- Files in the test packages are processed, identifying test functions. After
 processing the output is saved into a temporary location, by default this is
the directory `madness` in the root of the project. Check `madness/main.go`
file to see what is happening.

The generated package is self contained. Capable of running the test suite.
and collecting results.

- `gopherjs` is used to compile the generated package from the previous step.

- a browser is launched. All unit tests are executed in a single tab. Each
 integration test is done on separate tab, so be warned about resource limits
 for integration tests, this will soon be resolved by reusing the tabs but that
 is not the top priority.

- test results are then streamed back to the user's console via websocket.
- we close the browser and free any resources acquired while running the steps.
