package pack

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func Broadcast(number int) {
	// Specify the server address and port
	serverAddr := "127.0.0.1:9999"

	// Establish a connection to the server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Optionally, set a write buffer size
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.SetWriteBuffer(1024)
	}

	newNumber := number
	time.Sleep(time.Second)
	for {
		fmt.Printf("Sending %v\n", newNumber)
		_, err := conn.Write([]byte(strconv.Itoa(newNumber) + "\n")) // Ensure the message ends with a newline character
		if err != nil {
			panic(err)
		}
		newNumber = Counter(newNumber)
		time.Sleep(time.Second)
	}
}

func Counter(a int) int {
	return a + 1
}
