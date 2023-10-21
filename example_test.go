package http

import (
	"context"
	"fmt"
	"io"
	"log"
	nt "net/http"
	"sync"
	"time"
)

func panicFn(nt.ResponseWriter, *nt.Request) Error {
	panic("This should create an Internal Server error")
}

func startServer(port int) {
	logger := log.New(io.Discard, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

	s := New(fmt.Sprint(port), 10*time.Second, 10*time.Second, logger)
	s.AddPath("/panic", panicFn)

	var wg sync.WaitGroup
	s.Start(context.Background(), &wg)
	wg.Wait()
}

func ExampleWrapper() {

	var port = 8098
	go startServer(port)
	time.Sleep(100 * time.Millisecond)

	resp, _ := nt.Get(fmt.Sprintf("http://localhost:%v/panic", port))

	fmt.Println(resp.StatusCode)
	// Output: 500
}
