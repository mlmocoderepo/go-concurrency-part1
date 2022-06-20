package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Objective: to use mutex locks and unlocks
// Note: there is a difference between Mutex and RWMutex

func main() {

	// 1. find the number of cores on the system
	// Use the available cores to split the processes
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	var balance int
	var count int = 5000
	var m sync.Mutex // IMPORTANT: Use a Mutex to control reading and writing to the variable 'balance'
	var wg = sync.WaitGroup{}

	// 2. create a function as a variable that takes in a deposit
	deposit := func(amount int) {
		m.Lock()
		balance += amount
		m.Unlock()
	}

	// 3. Create a function as a variable that accepts a withdrawl
	withdrawal := func(amount int) {
		m.Lock()
		balance -= amount
		m.Unlock()
	}

	// 4. Run a loop to deposit and withdraw from balance via goroutines
	wg.Add(count)
	for i := 0; i < count; i++ {
		// create a goroutine that makes a DEPOSIT of 1
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(count)
	for i := 0; i < count; i++ {
		// create a goroutine that makes a WITHDRAWAL of 1
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()

	fmt.Println("total deposited:", balance)
}
