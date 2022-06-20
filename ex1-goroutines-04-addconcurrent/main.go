package main

import (
	"01-goroutine-add/counting"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	numbers := counting.GenerateNumbers(1e7)
	fmt.Printf("Add: %v, completed in - %s\n", counting.Add(numbers), time.Since(t))
	fmt.Printf("AddConcurrency: %v, completed in - %s\n", counting.AddConcurrency(numbers), time.Since(t))
}
