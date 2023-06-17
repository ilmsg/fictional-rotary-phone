package main

import (
	"fmt"
	"sync"
)

var (
	n       = 1000
	counter = 0
	wg      = &sync.WaitGroup{}
)

func main() {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
