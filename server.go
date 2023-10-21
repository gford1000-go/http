package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	nt "net/http"
	"sync"
	"time"
)

var errNoContext = errors.New("no context provided")
var errNoPaths = errors.New("no paths provided")
var errEmptyPath = errors.New("path must be specified")
var errNoHandler = errors.New("no handler provided with path")
var errContextEnded = errors.New("context has already ended")
var errStarted = errors.New("server already started")

// HandlerFuncWithError allows errors to be returned by http funcs
type HandlerFuncWithError func(nt.ResponseWriter, *nt.Request) Error

// New creates an instance of Server, ready to have paths and HandlerFuncs assigned
func New(port string, readTimeout time.Duration, writeTimeout time.Duration, logger *log.Logger) *Server {
	return &Server{
		log:     logger,
		port:    port,
		read:    readTimeout,
		write:   writeTimeout,
		paths:   make(map[string]nt.HandlerFunc),
		started: false,
	}
}

type Server struct {
	log     *log.Logger
	port    string
	read    time.Duration
	write   time.Duration
	paths   map[string]nt.HandlerFunc
	started bool
}

// AddPath adds a single path to the server, if it has not yet started
func (s *Server) AddPath(path string, handler HandlerFuncWithError) *Server {
	if len(path) == 0 {
		panic(errEmptyPath)
	}
	if handler == nil {
		panic(errNoHandler)
	}
	if s.started {
		panic(errStarted)
	}

	s.paths[path] = Wrapper(handler, s.log)
	return s
}

// Start begins serving of requests on the specified port with the
// specified context
func (s *Server) Start(ctx context.Context, wg *sync.WaitGroup) error {
	if ctx == nil {
		return errNoContext
	}
	if len(s.paths) == 0 {
		return errNoPaths
	}
	if s.started {
		return errStarted
	}
	select {
	case <-ctx.Done():
		return errContextEnded
	default:
	}

	wg.Add(1)
	s.started = true

	go func() {
		defer wg.Done()
		defer func() {
			recover() // A panic must not cause a fatal error
		}()

		mux := nt.NewServeMux()
		for path, handler := range s.paths {
			mux.Handle(path, handler)
		}

		srv := &nt.Server{
			Addr: fmt.Sprintf(":%v", s.port),
			BaseContext: func(l net.Listener) context.Context {
				return ctx
			},
			Handler:        mux,
			ReadTimeout:    s.read,
			WriteTimeout:   s.write,
			MaxHeaderBytes: nt.DefaultMaxHeaderBytes,
		}

		srv.ListenAndServe()
	}()

	return nil
}
