package main

import (
	"fmt"
	"time"
)

func selectChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "Hello from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from channel 2"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
