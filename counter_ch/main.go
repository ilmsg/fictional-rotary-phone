package main

import "fmt"

const n = 1000

var counter = 0

func main() {
	ch := make(chan int)
	defer close(ch)

	for i := 0; i < n; i++ {
		go func() {
			ch <- 1
		}()
	}

	for counter < n {
		counter += <-ch
	}

	fmt.Println(counter)
}
