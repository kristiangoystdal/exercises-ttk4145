package pack

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

// ListenForMaster listens for connections and decides if it should become the master
// It returns two values: a bool indicating whether to become the master, and an int with the last received number.
func ListenForMaster() (bool, int) {
	const listenDuration = 3 * time.Second

	lastReceivedNumber := 0

	// Listen on all interfaces
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error setting up TCP listener:", err)
		return false, lastReceivedNumber
	}
	defer listener.Close()

	// Create a channel to signal the acceptance of a new connection
	connChan := make(chan net.Conn)
	errChan := make(chan error)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			errChan <- err
			return
		}
		connChan <- conn
	}()

	// Set up a timer for listen duration
	timer := time.NewTimer(listenDuration)

	for {
		select {
		case conn := <-connChan:
			// Handle the connection
			reader := bufio.NewReader(conn)
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from connection:", err)
				conn.Close()
				continue
			}

			number, err := strconv.Atoi(message)
			if err != nil {
				fmt.Println("Error converting received message to number:", err)
				conn.Close()
				continue
			}

			lastReceivedNumber = number
			fmt.Printf("Message received, master is alive. Number: %d\n", number)
			conn.Close()
			// Reset the timer if you want to keep listening after receiving a message
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(listenDuration)
		case <-timer.C:
			// Time out
			fmt.Println("No connections received for the duration, assuming master role")
			return false, lastReceivedNumber
		case err := <-errChan:
			if !timer.Stop() {
				<-timer.C
			}
			fmt.Println("Error accepting connection:", err)
			return false, lastReceivedNumber
		}
	}
}
