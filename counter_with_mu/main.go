package main

import (
	"fmt"
	"sync"
)

var (
	n       = 10000
	counter = 0
	wg      = &sync.WaitGroup{}
	mutex   = &sync.Mutex{}
)

func main() {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer mutex.Unlock()
			defer wg.Done()
			mutex.Lock()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
