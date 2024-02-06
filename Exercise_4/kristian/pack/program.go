package pack

import (
	"fmt"
	"net"
	"os"
	"time"
)

func Program() {
	// Retrieve the terminal ID environment variable
	terminalID := os.Getenv("TERMINAL_ID")

	switch terminalID {
	case "1":
		fmt.Println("Starting Broadcast...")
		go Broadcast()
	case "2":
		fmt.Println("Starting Listen...")
		go Listen()
	default:
		fmt.Println("Terminal ID not recognized. Neither Broadcasting nor Listening.")
	}

	// Common functionality here; for instance:
	fmt.Println("Hello, world!")
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Running...")
	}
}

func Broadcast() {
	// Create a UDP address for broadcasting
	broadcastAddr := "255.255.255.255:9999"
	conn, err := net.Dial("udp", broadcastAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Enable broadcast mode
	if udpConn, ok := conn.(*net.UDPConn); ok {
		udpConn.SetWriteBuffer(1024)
	}

	for {
		_, err := conn.Write([]byte("Hello, network!"))
		if err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)
	}
}

func Listen() {
	// Listen on all interfaces
	addr := net.UDPAddr{
		Port: 9999,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error receiving:", err)
			continue
		}
		fmt.Printf("Received '%s' from %s\n", string(buffer[:n]), remoteAddr)
	}
}
