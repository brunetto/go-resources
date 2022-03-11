package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type work struct{ data int }
	fanOut := make(chan work, 15)

	activityDuration := 10 * time.Second

	nWkrs := 10
	wg := &sync.WaitGroup{}
	wg.Add(nWkrs)

	for i := 0; i < nWkrs; i++ {
		go func(ctx context.Context, wg *sync.WaitGroup, sleep time.Duration, in chan work) {
			defer wg.Done()

			i := 0
			for item := range in {
				if ShouldLeaveNow(ctx) {
					fmt.Println("timed out")
					return
				}

				fmt.Println(i, item.data)
				time.Sleep(sleep)

				i++
			}

			fmt.Println("work finished")
		}(ctx, wg, activityDuration, fanOut)
	}

	for i := 0; i < 20; i++ {
		fanOut <- work{i}
	}
	close(fanOut)

	cancel()

	wg.Wait()
}

func ShouldLeaveNow(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
