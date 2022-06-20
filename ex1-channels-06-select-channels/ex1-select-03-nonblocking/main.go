package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)

	go func(ch chan<- string) {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- fmt.Sprintf("Message %v", i)
		}
	}(ch1)

	// as the loop takes place, adding the default statement serves as a non-blocking statement
	// that allows the iteration to continue till the channel receives a message and meets the case's condition.
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		default:
			fmt.Println("no message received")
		}

		fmt.Printf("%v Processing...\n", i)
		time.Sleep(1500 * time.Millisecond)
	}
}
