package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Define the server address and port
	addr := net.UDPAddr{
		Port: 30000,
		IP:   net.ParseIP(""),
	}

	// Create a UDP socket
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Listening on %s\n", conn.LocalAddr().String())

	// Buffer to store received data
	buffer := make([]byte, 1024)

	for {
		// Read from UDP connection
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Received %d bytes from %s: %s\n", n, remoteAddr, string(buffer[:n]))

		// Process the data (just printing it here)
		// ...
	}
}
