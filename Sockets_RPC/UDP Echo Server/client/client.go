package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", "localhost:9999")
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}

	// Assigining string to byte array
	conn.Write([]byte("Hello Server!"))
	fmt.Println("Message Sent: Hello Server!")
	conn.Write([]byte("How Are You?"))
	fmt.Println("Message Sent: How Are You?")

	if err != nil {
		log.Fatalln(err)
	}

	for {
		buffer := make([]byte, 1400)
		dataSize, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalln(err)
			fmt.Println("Connection Closed!")
			return
		}
		data := buffer[:dataSize]
		fmt.Println("Received Message: ", string(data))
	}
}
