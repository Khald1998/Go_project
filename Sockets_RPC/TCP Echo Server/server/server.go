package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		log.Fatalln(err)
	}

	// Localhost IP address is 127.0.0.1
	fmt.Println("Server Listening on LocalHost:8888")

	defer listener.Close()

	// Accept Connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("New Connection Request")

		// Concurrent Server
		go listenConnection(conn)
	}
}

// Listen to messages coming from connection

func listenConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			log.Fatalln(err)
			fmt.Println("Connection Closed!")
			return
		}
		data := buffer[:dataSize]
		fmt.Println("Received Message: ", string(data))

		// Echo back the message
		_, err = conn.Write(data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message Sent: ", string(data))
	}
}
