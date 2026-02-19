package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Phân biệt sự khác nhau giữa <-chan và chan<- khi truyền vào hàm:
  - <-chan: Đây là một channel chỉ có thể nhận dữ liệu (receive-only channel).
  - chan<-: Đây là một channel chỉ có thể gửi dữ liệu (send-only channel).
*/
func channelTask(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %d bat dau\n", id)
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Task %d ket thuc\n", id)
	ch <- fmt.Sprintf("Ending %dth\n", id)
}

func waitGroupChannel() {
	fmt.Println("Bat dau waitGroup")
	start := time.Now()

	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go channelTask(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Print(v)
	}

	fmt.Println("Tong thoi gian:", time.Since(start))
}
