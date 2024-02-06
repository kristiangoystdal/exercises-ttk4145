package pack

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// ListenForMaster listens for broadcasts and decides if it should become the master
// It returns two values: a bool indicating whether to become the master, and an int with the last received number.
func ListenForMaster() (bool, int) {
	const timeoutDuration = 1500 * time.Millisecond
	const listenDuration = 3 * time.Second // Total time to listen for messages before assuming master role

	lastReceivedNumber := 0 // Initialize the last received number to 0

	// Listen on all interfaces
	addr := net.UDPAddr{
		Port: 9999,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error setting up UDP listener:", err)
		return false, lastReceivedNumber
	}
	defer conn.Close()

	endTime := time.Now().Add(listenDuration)
	for time.Now().Before(endTime) {
		// Update the deadline for each iteration of the loop
		err = conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		if err != nil {
			fmt.Println("Error setting read deadline:", err)
			return false, lastReceivedNumber
		}

		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			if e, ok := err.(net.Error); ok && e.Timeout() {
				fmt.Println("Listen timed out, checking for master...")
				continue // Continue listening until the listenDuration has passed
			}
			fmt.Println("Error receiving:", err)
			return false, lastReceivedNumber
		}

		// Attempt to convert the received message to an integer
		message := string(buffer[:n])
		number, err := strconv.Atoi(message)
		if err != nil {
			fmt.Println("Error converting received message to number:", err)
			// Optionally handle the error, e.g., by continuing to listen
			continue
		}

		// Update the last received number
		lastReceivedNumber = number

		// Print the received number
		fmt.Printf("Broadcast received, master is alive. Number: %d\n", number)
		endTime = time.Now().Add(listenDuration) // Reset the end time to wait for another period
	}

	// If the loop exits, it means no broadcasts have been received for the duration of listenDuration
	fmt.Println("No broadcast received for the duration, assuming master role")
	return false, lastReceivedNumber // Indicate it's time to become the master, and return the last received number
}
