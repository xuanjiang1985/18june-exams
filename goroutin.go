package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err()) // prints "context deadline exceeded"
				return
			case <-time.NewTicker(10 * time.Second).C:
				//time.Sleep(10 * time.Second)
				//time.Sleep(1 * time.Second)
				fmt.Println("Done after 10 seconds")
			}
		}
	}(ctx)
	wg.Wait()
}
