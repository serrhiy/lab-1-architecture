package main

import (
	"net"
	"fmt"
)

const host = "localhost"
const port = "8795"

func main() {
	address := net.JoinHostPort(host, port)
	fmt.Println(address)
}
