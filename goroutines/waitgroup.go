package main

import (
	"fmt"
	"sync"
	"time"
)

func wgTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Printf("Task %d ket thuc\n", id)
	defer fmt.Println()

	fmt.Printf("Task %d bat dau\n", id)
	time.Sleep(1 * time.Second)
}

func waitGroup() {
	fmt.Println("Bat dau waitGroup")
	start := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go wgTask(i, &wg)
	}
	wg.Wait()

	fmt.Println("Duoc thuc hien sau khi bo dem bang 0")
	fmt.Println("Tong thoi gian:", time.Since(start))
}
