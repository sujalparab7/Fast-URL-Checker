package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL     string
	Status  int
	Latency time.Duration
	Error   error
}

func checkUrl(url string) Result {
	start := time.Now()
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	latency := time.Since(start)

	if err != nil {
		return Result{URL: url, Error: err, Latency: latency}
	}
	defer resp.Body.Close()
	return Result{URL: url, Status: resp.StatusCode, Latency: latency}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.reddit.com",
		"https://invalid-url-example.com",
	}

	var wg sync.WaitGroup
	
	resultsChan := make(chan Result, len(urls))

	fmt.Printf("Checking %d URLs with WaitGroup...\n\n", len(urls))

	for _, url := range urls {
		wg.Add(1)
		
		go func(u string) {
			defer wg.Done()
			
			resultsChan <- checkUrl(u)
		}(url) 
	}

	go func() {
		wg.Wait() 
		close(resultsChan) 
	}()

	for res := range resultsChan {
		if res.Error != nil {
			fmt.Printf("[Error] %s - %s\n", res.URL, res.Error)
		} else {
			fmt.Printf("[%d] %s - took %v\n", res.Status, res.URL, res.Latency)
		}
	}
	
	fmt.Println("\nAll checks completed.")
}