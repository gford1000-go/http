package http

import (
	"log"
	nt "net/http"
)

// Wrapper allows us to create standard http Handlers that wrap extended variants reporting errors
func Wrapper(h HandlerFuncWithError, logger *log.Logger) func(nt.ResponseWriter, *nt.Request) {
	return func(w nt.ResponseWriter, req *nt.Request) {

		serverError := func(w nt.ResponseWriter, i interface{}) {
			logger.Printf("HTTP unknown - %v\n", i)
			nt.Error(w, nt.StatusText(nt.StatusInternalServerError), nt.StatusInternalServerError)
		}

		defer func() {
			if r := recover(); r != nil {
				serverError(w, r)
			}
		}()

		// Only continue if the context is valid
		select {
		case <-req.Context().Done():
			serverError(w, req.Context().Err())
			return
		default:
		}

		if err := h(w, req); err != nil {
			switch e := err.(type) {
			case Error:
				// We can retrieve the status here and write out a specific
				// HTTP status code.
				logger.Printf("HTTP %d - %s\n", e.Status(), e)
				nt.Error(w, e.Error(), e.Status())
			default:
				// Any error types we don't specifically look out for default
				// to serving a HTTP 500
				serverError(w, err)
			}
		}
	}
}
