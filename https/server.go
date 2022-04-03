package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	// the cert.pem file is the SSL certificate whereas key.pem is the private key for the server.
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
