package main

import "fmt"

// Objective: to create 2 functions ping and pong WITHOUT channel direction

// ping sends a message to ch1
func ping(out chan string) {

	// sent data to ch1 that has been passed in
	out <- "ping"
}

// pong receives a message from ch1 and passes it to ch2
func pong(out, in chan string) {
	msg := <-out + " pong"
	in <- msg
}

// pong receives a message from ch1

// Objective: Implement a ping pong with channel direction
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go ping(ch1)
	go pong(ch1, ch2)

	fmt.Printf("The results coming from ch2: %s", <-ch2)

}
