package main

import (
	"io"
	"log"
	"net"
	"os"
)

// Objective: to create a client that dials a tcp connection
// sending a request to network address: localhost:9010
func main() {

	// 1. Dial a tcp connection to localhost:9010
	conn, err := net.Dial("tcp", "localhost:9010")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close() // 2. gracefully close the conn when main() is exited

	captureResponse(os.Stdout, conn) // 3. use os.Stdout and conn (display dst and source)

}

// Utility function handles the display returned from the server
func captureResponse(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
