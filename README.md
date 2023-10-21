[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Build Status](https://travis-ci.org/gford1000-go/http.svg?branch=master)](https://travis-ci.org/gford1000-go/http)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/gford1000-go/http)


HTTP
====

HTTP provides helper functions when using the standard golang `http` package.

An example of use is available in GoDocs.

The `Wrapper` function wrappers the extended `HandleFuncWithError` type, which allows request handlers to return errors that are handled in a standard manner.  Additionally panics are captured and managed as 500 Internal Server errors.

The `Server` type uses the `HandleFuncWithError` and simplifies starting an http server.  Any number of `Server` instances can be started, each running in a separate goroutine.

`http/json` replaces boilerplate code needed to read from an `io.ReadCloser` and construct an object from the
JSON byte array.

Installing and building the library
===================================

This project requires Go 1.21.1 or later.

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

