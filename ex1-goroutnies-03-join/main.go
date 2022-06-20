package main

import (
	"fmt"
	"sync"
)

func main() {

	// program to print the value of 1 deterministically
	// using a Waitgroup

	var wg = sync.WaitGroup{}
	var data int

	wg.Add(1)
	go func() {
		data++
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(data)
	fmt.Println("Done...")
}
