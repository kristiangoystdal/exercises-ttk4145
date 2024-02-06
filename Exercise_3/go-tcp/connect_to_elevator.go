package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	serverIP   = "127.0.0.1" // Localhost IP address
	serverPort = "15657"     // Local server's port
)

func main() {
	address := serverIP + ":" + serverPort
	fmt.Println("Connecting to server at", address)

	// Connect to the server
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected. You can now send a 4-byte message.")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a 4-byte message: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		// Ensure the message is exactly 4 bytes
		if len(input) != 4 {
			fmt.Println("Error: Message must be exactly 4 bytes.")
			continue
		}

		msg:= []byte{0x04, 0x01,0x00,0x00}
		fmt.Println("Sending message:", input)
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("Error sending:", err)
			break
		}

		// Read the echoed message from the server
		buffer := make([]byte, 4) // buffer for 4-byte message
		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println("Error receiving:", err)
			break
		}
		fmt.Println("Echoed:", string(buffer))
	}
}

