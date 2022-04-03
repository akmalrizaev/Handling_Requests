package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// We just created a handler and attached it to our server, so we’re no longer using any multiplexers. This
// means there’s no longer any URL matching to route the request to a particular handler, so all requests going into the server will go to this handler.

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
