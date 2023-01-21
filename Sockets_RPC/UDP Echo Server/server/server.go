package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", "localhost:9999")
	if err != nil {
		log.Fatalln(err)
	}

	// Localhost IP address is 127.0.0.1
	fmt.Println("\nServer Listening on LocalHost:9999")

	defer conn.Close()

	for {
		buffer := make([]byte, 1400)
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handleRequest(conn, addr, buffer[:n])
	}
}

func handleRequest(conn net.PacketConn, addr net.Addr, buffer []byte) {
	fmt.Println("Received Message: ", string(buffer))

	// Echo back the message
	conn.WriteTo(buffer, addr)
}
