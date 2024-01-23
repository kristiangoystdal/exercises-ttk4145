package main

import (
	"bufio"
	"fmt"
	"net"
	// "os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Hello, server!\n")
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		// handle error
	}
	fmt.Println("Server response:", response)
}
