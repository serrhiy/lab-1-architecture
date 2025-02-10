package main

import (
	"lab-1-architecture/handlers"
	"log"
	"net"
	"net/http"
)

const host = "localhost"
const port = "8795"

func main() {
	address := net.JoinHostPort(host, port)
	for path, handler := range handlers.Handlers {
		http.HandleFunc(path, handler.Controller)
	}
	fail := http.ListenAndServe(address, nil)
	log.Fatal(fail)
}
