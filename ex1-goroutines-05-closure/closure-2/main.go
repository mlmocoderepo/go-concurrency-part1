package main

import (
	"fmt"
	"sync"
)

// Objective: to pass value to the goroutine
// so that the the goroutine can operate on the inputs passed to it
func main() {

	var wg = sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
