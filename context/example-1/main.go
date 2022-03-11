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

	activityDuration := 10 * time.Second

	nWkrs := 10
	wg := &sync.WaitGroup{}
	wg.Add(nWkrs)

	for i := 0; i < nWkrs; i++ {
		go func(ctx context.Context, wg *sync.WaitGroup, sleep time.Duration) {
			defer wg.Done()

			c := make(chan struct{})

			go func() {
				time.Sleep(sleep)
				fmt.Println("elapsed", sleep)
				c <- struct{}{}
			}()

			select {
			case <-ctx.Done():
				fmt.Println("timed out")
				return
			case <-c:
			}

		}(ctx, wg, activityDuration)
	}

	cancel()

	wg.Wait()
}
