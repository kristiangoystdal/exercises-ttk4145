package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println("Handling new connection")

	defer func() {
		fmt.Println("Closing connection")
		conn.Close()
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		receivedText := scanner.Text()
		fmt.Printf("Received: %s\n", receivedText)

		_, writeErr := conn.Write([]byte("69420.\n"))
		if writeErr != nil {
			fmt.Printf("Error writing to connection: %s\n", writeErr)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from connection: %s\n", err)
	}

	fmt.Println("Finished handling connection")
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	defer ln.Close()

	fmt.Println("Hei!")

	for {
		conn, err := ln.Accept()
		
		if err != nil {
			fmt.Println("AcceptError")
			continue
		}

		fmt.Println("Hei2")

		go handleConnection(conn)
	}

}
