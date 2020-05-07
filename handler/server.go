package handler

import (
	"log"
	"net/http"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// Wrapper allows us to create standard http Handlers that wrap extended variants reporting errors
func Wrapper(h func(http.ResponseWriter, *http.Request) Error, logger *log.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

		serverError := func(w http.ResponseWriter, i interface{}) {
			logger.Printf("HTTP unknown - %v\n", i)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		defer func() {
			if r := recover(); r != nil {
				serverError(w, r)
			}
		}()

		err := h(w, req)
		if err != nil {
			switch e := err.(type) {
			case Error:
				// We can retrieve the status here and write out a specific
				// HTTP status code.
				logger.Printf("HTTP %d - %s\n", e.Status(), e)
				http.Error(w, e.Error(), e.Status())
			default:
				// Any error types we don't specifically look out for default
				// to serving a HTTP 500
				serverError(w, err)
			}
		}
	}
}
