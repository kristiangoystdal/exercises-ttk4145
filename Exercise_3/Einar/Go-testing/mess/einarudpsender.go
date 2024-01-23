package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Define the server address and port
	serverAddr := "127.0.0.1:30000" // Replace with the destination IP and port

	// Resolve the server address
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Dial UDP
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	// Message to send
	message := []byte("Hello, UDP server!")

	// Send the message
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Sent message to %s: %s\n", serverAddr, string(message))
}
