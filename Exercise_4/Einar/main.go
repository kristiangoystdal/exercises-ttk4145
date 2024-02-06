package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
	"os"
	"strconv"
)

func Broadcast(counter int) {
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

		counterStr := strconv.Itoa(counter)

		_, err := conn.Write([]byte(counterStr))
		if err != nil {
			panic(err)
		}

		fmt.Println("Broadcasting: ", counter)

		counter ++

		time.Sleep(time.Second)
	}
}

func Listen1() int{

	var counter int = 69

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
		// Set a deadline for the read operation
		err := conn.SetReadDeadline(time.Now().Add(3 * time.Second))
		if err != nil {
			fmt.Println("Error setting deadline:", err)
			continue
		}

		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				fmt.Println("Read timed out")
				break
			}
			fmt.Println("Error receiving:", err)
			continue
		}

		fmt.Printf("Received: %s\n", string(buffer[:n]))

		counter, err = strconv.Atoi(string(buffer[:n]))
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
	}
	return counter
}

func Primary(counter int){
	fmt.Println("Hello from primary!")

	cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "main.exe 1")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	
	time.Sleep(time.Second)

	Broadcast(counter)	
}

func Secondary() int {

	var counter int = 69

	fmt.Println("Hello from secondary!")

	counter = Listen1()

	return counter
}



func main() {

	var counter int = 0

	if len(os.Args) > 1{
		counter = Secondary()

		counter ++
	}

	Primary(counter)
	
}
