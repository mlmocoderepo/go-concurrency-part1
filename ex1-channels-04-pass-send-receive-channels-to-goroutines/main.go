package main

import "fmt"

// Objective: create 2 functions ping and pong WITH channel directions

// ping takes in a sender channel that sends string data
func ping(out chan<- string) {
	out <- "ping"
}

// pong takes in a receiver channel that receives string data
func pong(in <-chan string, out chan<- string) {
	msg := <-in
	out <- msg + " pong"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go ping(ch1)
	go pong(ch1, ch2)

	fmt.Printf("The output of ch2 after using directional channels in functions: %s", <-ch2)
}
