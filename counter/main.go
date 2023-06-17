package main

import "fmt"

var (
	n       = 1000
	counter = 0
)

func main() {
	for i := 0; i < n; i++ {
		go func() {
			counter++
		}()
	}

	fmt.Println(counter)
}
