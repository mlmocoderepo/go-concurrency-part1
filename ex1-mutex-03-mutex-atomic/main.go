package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Objective: to implement concurrency safe counter

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	count := 50
	var counter int64
	var wg sync.WaitGroup = sync.WaitGroup{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				atomic.AddInt64(&counter, 1) //Implement a concurrency safe counter to increment counter
				// counter += 1 //when this is done, the results of the addition are inconsistent
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter value is:", counter)
}
