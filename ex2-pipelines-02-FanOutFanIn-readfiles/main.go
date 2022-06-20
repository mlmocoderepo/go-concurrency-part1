package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

// 1. Open the file sent over
// 2. Reads the file
// 3. Run a goroutines, continuously read and return when reading reaches EOF
// 4. Store each read []string slice into a chan []string and return it
func read(file string) (<-chan []string, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to open %v", file)
	}

	ch := make(chan []string)
	cr := csv.NewReader(f)

	go func() {
		for {
			records, err := cr.Read()
			if err == io.EOF {
				close(ch)
				return
			}
			ch <- records
		}

	}()

	return ch, nil
}

// 1. Take in variadic number of receiver channels of []string slices
// 2. Runs a loop to invoke goroutine send() taking in the channel []string slice
// 3. Send() loops throug the channel of []string slice and stores the values into the merged channel []string
// 4. Return the channel of []string slices
func merge1(cs ...<-chan []string) <-chan []string {

	var wg sync.WaitGroup = sync.WaitGroup{}

	out := make(chan []string)

	send := func(c <-chan []string) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, v := range cs {
		go send(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	// This is the FanOut stage
	ch1, err := read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("Unable to read file %v", err))
	}
	ch2, err := read("file2.csv")
	if err != nil {
		panic(fmt.Errorf("Unable to read file %v", err))
	}

	// This is teh FanIn stage
	chm := merge1(ch1, ch2)

	for v := range chm {
		fmt.Println(v)
	}

	fmt.Println("Merge completed. ")

}
