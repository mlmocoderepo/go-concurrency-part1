package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var count = 0

// Objective: to create a tcp listener to handling incoming server
// requests coming through tcp protocol address: localhost:9010
func main() {
	// 1. Create a listener to accept tcp connections via localhost:9010
	listener, err := net.Listen("tcp", "localhost:9010")
	if err != nil {
		log.Fatal(err)
	}

	for { // 2. use a for loop to continuously listen for incoming tcp request
		fmt.Println("Server is listening and accepting requests")
		conn, err := listener.Accept() //3. use the listener to accept requests

		if err != nil {
			log.Fatal()
		}

		returnResponse(conn) // 4. handle with a response
	}
}

//Utility function  handles the requests accepted from the connection
func returnResponse(c net.Conn) {
	// 5. Gracefully close the connection when the main function exits
	defer c.Close()

	for { // 6. use a for loop to write a response as the client continues to connect
		fmt.Println("Server is responding to the client")
		serverfeedback := fmt.Sprintf("response from the server %d\n", count)
		count++
		_, err := io.WriteString(c, serverfeedback)
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
