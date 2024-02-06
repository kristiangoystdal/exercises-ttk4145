package pack

import (
	"fmt"
	"net"
)

func Listen1() {
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
