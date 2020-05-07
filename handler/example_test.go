package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func panicFn(http.ResponseWriter, *http.Request) Error {
	panic("This should create an Internal Server error")
}

func startServer(port int) {
	logger := log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

	http.HandleFunc("/panic", Wrapper(panicFn, logger))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func ExampleWrapper() {

	var port = 8081
	go startServer(port)

	resp, _ := http.Get(fmt.Sprintf("http://localhost:%v/panic", port))

	fmt.Println(resp.StatusCode)
	// Output: 500
}
