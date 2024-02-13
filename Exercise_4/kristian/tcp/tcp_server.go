package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listenAddr := "127.0.0.1:9999"
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("TCP Server listening on", listenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Connection established")

		// Send message to client
		_, err = conn.Write([]byte("Hello from Server!\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			conn.Close()
			continue
		}

		// Wait for client's confirmation
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading confirmation from client:", err)
			conn.Close()
			continue
		}
		fmt.Printf("Client says: %s", response)

		conn.Close()
	}
}
