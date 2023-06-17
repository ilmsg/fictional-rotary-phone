package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

const word = "the"

var urls = []string{
	"https://code-pilot.me/how-ive-made-this-platform",
	"https://code-pilot.me/making-a-beautiful-error-handler-in-go",
	"https://code-pilot.me/mastering-goroutines-and-channels",
	"https://code-pilot.me/why-should-you-curry",
	"https://code-pilot.me/not-a-real-page",
	"not@validURL",
}

func get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if res.StatusCode == 404 {
		return "", fmt.Errorf("not found: %s", url)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	resChan := make(chan int)
	errChan := make(chan error)
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)

	for _, url := range urls {
		go func(url string) {
			body, err := get(url)
			if err != nil {
				errChan <- err
				return
			}

			totalWord := strings.Count(strings.ToLower(body), word)
			fmt.Printf("found %d occurrences of the word \"%s\" in %s\n", totalWord, word, url)
			resChan <- totalWord
		}(url)
	}

	sum := 0
	for {
		select {
		case x := <-resChan:
			sum += x
		case err := <-errChan:
			fmt.Println(err.Error())
		case <-sigChan:
			fmt.Printf("\nfound %d occurrences of the word \"%s\" in total\n", sum, word)
			os.Exit(0)
		}
	}
}
