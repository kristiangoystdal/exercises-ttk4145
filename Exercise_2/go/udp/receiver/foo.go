package main

import (
    "fmt"
    "net"
    "os"
)

const (
    PORT    = ":30000"
    BUFLEN  = 1024
)

func main() {
    // Resolve UDP address
    addr, err := net.ResolveUDPAddr("udp", PORT)
    if err != nil {
        fmt.Println("Error resolving UDP address:", err)
        os.Exit(1)
    }

    // Create a UDP socket
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Error creating UDP socket:", err)
        os.Exit(1)
    }
    defer conn.Close()

    fmt.Println("Waiting for data...")

    buf := make([]byte, BUFLEN)

    for {
        // Try to receive some data, this is a blocking call
        n, remoteAddr, err := conn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("Error receiving data:", err)
            continue
        }

        // Print details of the client/peer and the data received
        fmt.Printf("Received packet from %s\n", remoteAddr.String())
        fmt.Printf("Data: %s\n", string(buf[:n]))
    }
}
