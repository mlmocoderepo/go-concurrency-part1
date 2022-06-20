package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Objective: to use a RWMutex to allow exclusive write and multiple reads

func main() {

	// 1. Set the number of cores to run the goroutines in concurrence
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	var balance int
	var count int = 100
	var rwm sync.RWMutex // IMPORTANT: Use RWMutex to control writing to variable 'balance'
	var wg = sync.WaitGroup{}

	// 2. Set up a function as a variable 'deposit' that applies a lock on the memory when accessed
	deposit := func(amount int) {
		rwm.Lock()
		defer rwm.Unlock()
		balance += amount
	}

	// 3. Set up a function as a variable 'read' that allows concurrent read access to the variable 'balance'
	// Note that an RLock and RUnlocked is used - RLock allows mutiple, concurrent read access
	read := func() int {
		rwm.RLock()
		defer rwm.RUnlock()
		return balance
	}

	// 4. Perform a deposit on the number of times stated in count
	// During the process, no other process has access to the memory location of 'balance'
	wg.Add(count)
	for i := 0; i < count; i++ {
		time.Sleep(100 * time.Millisecond)
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	// 5. After the write is performed, send 'count' number of goroutines to read from 'balance' concurrently
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("Balance: ", read())
		}()
	}

	wg.Wait()
	fmt.Println("Done processing.")
}
