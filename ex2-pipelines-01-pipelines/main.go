package main

import "fmt"

func generateNumbers(n ...int) <-chan int {

	out := make(chan int)

	// call the goroutine that passes each value from the n integer slice to the out channel
	go func() {
		for _, v := range n {
			out <- v
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {

	out := make(chan int)

	// call the goroutine that passes each squared value from the in channel to the out channel
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func main() {

	/*
		ch := make(<-chan int)
		// generate values to be passed to a channel
		ch = generateNumbers(2, 3)
		// squre the values from channel ch
		out := square(ch)
		// run a loop to retrieve the values from the out channel
		for v := range out {
			fmt.Println(v)
		}
	*/

	// shorter way of the above, and you can stack the call
	for out := range square(square(generateNumbers(2, 3))) {
		fmt.Println(out)
	}

}
