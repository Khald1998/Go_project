package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatalln(err)
	}

	// Assigining string to byte array
	conn.Write([]byte("Hello Server!"))
	conn.Write([]byte("How Are You?"))

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Message Sent: Hello Server!")
	fmt.Println("Message Sent: How Are You?")

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
	}
}
