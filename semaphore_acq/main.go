package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

const (
	n          = 10
	goroutines = 2
)

var sm = semaphore.NewWeighted(goroutines)

func main() {
	ctx := context.Background()
	for i := 0; i < n; i++ {
		sm.Acquire(ctx, 1)
		go func() {
			defer sm.Release(1)
			time.Sleep(time.Second)
		}()
		fmt.Printf("currently running: %d goroutines\n", runtime.NumGoroutine()-1)
	}
}
