package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:15657")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Heihei")

	// Prepare a 4-byte message

	// Create a 4-byte message
	msg := []byte{0x07, 0x00, 0x00, 0x00} // Example 4-byte message

	conn.Write(msg)

	// Send the message
	// _, err = conn.Write(msg)
	// if err != nil {
	// 	fmt.Println("Error sending message:", err)
	// 	return
	// }

	fmt.Println("Message sent successfully")

	// Read the server's response
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Server response:", response)
}
