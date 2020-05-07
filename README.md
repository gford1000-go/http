[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Build Status](https://travis-ci.org/gford1000-go/http.svg?branch=master)](https://travis-ci.org/gford1000-go/http)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/gford1000-go/http)


HTTP | Geoff Ford
=================

HTTP provides helper functions when using the standard golang `http` package.

An example of use is available in GoDocs.

`http/handler` provides the `Wrapper` function, which allows `http` handler functions to return errors
and have these interpreted into HTTP return codes.  Panics are handled as 500 internal server errors.

`http/json` replaces boilerplate code needed to read from an `io.ReadCloser` and construct an object from the
JSON byte array.

Installing and building the library
===================================

This project requires Go 1.14.2

To use this package in your own code, install it using `go get`:

    go get github.com/gford1000-go/http

Then, you can include it in your project:

	import "github.com/gford1000-go/http"

Alternatively, you can clone it yourself:

    git clone https://github.com/gford1000-go/http.git

Testing and benchmarking
========================

To run all tests, `cd` into the directory and use:

	go test -v

