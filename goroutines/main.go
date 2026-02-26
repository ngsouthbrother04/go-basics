package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for i := 1; i <= 100e8; i++ {
		sum += i
	}
}

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPU cores: %d\n", numCPU)

	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i <= 20; i++ {
		wg.Add(1)
		go heavyTask(&wg)
	}

	wg.Wait()
	fmt.Println("Time taken:", time.Since(start))
}
