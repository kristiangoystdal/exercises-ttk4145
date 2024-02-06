package pack

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func Broadcast(number int) {
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
	newNumber := number
	time.Sleep(time.Second)
	for {
		fmt.Printf("Broadcasting %v\n", newNumber)
		_, err := conn.Write([]byte(strconv.Itoa(newNumber)))
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
