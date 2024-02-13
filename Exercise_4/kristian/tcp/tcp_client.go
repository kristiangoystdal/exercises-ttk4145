package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	serverAddr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Connected to server at", serverAddr)

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Print("Message from server: ", message)
}
