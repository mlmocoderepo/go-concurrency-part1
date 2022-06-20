package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateNumbers(max int) []int {

	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

func Add(numbers []int) int64 {
	var sum int64
	for _, v := range numbers {
		sum += int64(v)
	}
	return sum
}

func AddConcurrency(numbers []int) int64 {

	numOfCores := runtime.NumCPU() //check the number of cores on the system (e.g. 6)
	runtime.GOMAXPROCS(numOfCores) //set the number of cores to be used for the application

	var sum int64                  //create a placeholder variable to store the sum of values
	max := len(numbers)            //det. the length of the slice (e.g 48)
	numOfParts := max / numOfCores //det. the num parts processed per core max/numOfCores (e.g. 48/6 = 8)

	var wg = sync.WaitGroup{}

	for i := 0; i < numOfCores; i++ { //use a for loop to prog

		start := i * numOfParts    //at index 5 (max), start = 40
		end := start + numOfParts  //at index 5 (max), end = 48
		part := numbers[start:end] //at index 5 (max), part would part[40:48]

		wg.Add(1)          //indicate that a goroutine is added
		go func(n []int) { //run each part in a separate goroutine
			defer wg.Done()

			var result int64 //create a placeholder variable to store the sum of parts
			for _, v := range n {
				result += int64(v) //add the values to results to store the sum of parts
			}

			atomic.AddInt64(&sum, result) //add the sum of parts to the main sum to be returned

		}(part)
	}
	wg.Wait()

	return int64(sum) //return the sum
}
