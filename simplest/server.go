package main

import (
	"net/http"
)

func main() {
	// If the network address is an empty string, the default is all network interfaces at port 80.
	// If the handler parameter is nil, the default multiplexer, DefaultServeMux, is used.

	http.ListenAndServe("", nil)
}
