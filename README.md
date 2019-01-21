# Go Hello World

This is a simple Golang program with propose of studies.

What will be done herer:

* Run a src code and see what it does;
* Run some tests;
* Se the coverage of it test;
* Make program bin from its code;

## Dependencies

* Golang 1.11.2 or futher, see [install](https://golang.org/doc/install);

## No docker for you

Sorry this is a simple example, so not much advanced stuffs herer.

# Running Test

The command bellow is for run tests file existent on this project, the `-cover` is used for ask for coverage.

`go test ./... -cover`

## Retrieving Coverage Pretty Way

To retrieve a nice coverage report, it's possible with `go tool cover`, see bellow how.

1. `go test ./... -coverprofile=./system.cov`
2. `go tool cover -html=./system.cov`

If want save the html result just add `-o coverage.html` at the end of second command like this:

`go tool cover -html=./system.cov -o coverage.html`

# Building project

`go build -o ./build/hello ./cmd/hello/main.go`
