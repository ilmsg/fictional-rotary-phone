package main

import (
	"fmt"
	"runtime"
	"time"
)

const n = 10

func main() {
	for i := 0; i < n; i++ {
		go func() {
			time.Sleep(time.Second)
		}()
		fmt.Printf("Currently running: %d goroutines\n", runtime.NumGoroutine()-1)
	}
}
