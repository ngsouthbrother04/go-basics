package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"nnama.com/excercise/models"
	"nnama.com/excercise/monitors"
	"nnama.com/excercise/processors"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	monitorList := []models.Monitor{
		&monitors.CPUMonitor{},
		&monitors.MemoryMonitor{},
		&monitors.NetMonitor{},
		&monitors.DiskMonitor{},
	}
	var wg sync.WaitGroup

	statCh := make(chan models.SystemStat)

	for _, m := range monitorList {
		wg.Add(1)
		go processors.RunMonitor(ctx, &wg, statCh, m)
	}

	go func() {
		for stat := range statCh {
			models.Stats.Set(stat)
		}
	}()

	printTicker := time.NewTicker(1200 * time.Millisecond)
	defer printTicker.Stop()

	go func() {
		for range printTicker.C {
			fmt.Println("================System Stats================")
			for k, v := range models.Stats.Snapshot() {
				fmt.Printf("[%s] %s\n", k, v)
			}

			fmt.Println(processors.GetTopProcesses(ctx))
		}
	}()

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
	close(statCh)
}
