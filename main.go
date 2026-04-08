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
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://chatgpt.com",
		"https://www.wikipedia.org",
		"https://www.quora.com",
		"https://x.com",
		"https://www.whatsapp.com",
		"https://www.bing.com",
		"https://www.tiktok.com",
		"https://www.yahoo.com",
		"https://www.amazon.com",
		"https://gemini.google.com",
		"https://www.linkedin.com",
		"https://www.baidu.com",
		"https://www.naver.com",
		"https://www.netflix.com",
		"https://www.pinterest.com",
		"https://live.com",
		"https://www.bilibili.com",
		"https://www.temu.com",
		"https://dzen.ru",
		"https://www.office.com",
		"https://www.microsoft.com",
		"https://www.walmart.com",
		"https://www.twitch.tv",
		"https://www.espn.com",
		"https://www.canva.com",
		"https://weather.com",
		"https://vk.com",
		"https://www.globo.com",
		"https://www.fandom.com",
		"https://www.samsung.com",
		"https://mail.ru",
		"https://duckduckgo.com",
		"https://www.nytimes.com",
		"https://www.bbc.com",
		"https://www.cnn.com",
		"https://www.ebay.com",
		"https://zoom.us",
		"https://cloud.microsoft",
		"https://www.paypal.com",
		"https://github.com",
		"https://www.apple.com",
		"https://www.imdb.com",
		"https://invalid-url-example1.com",
		"https://invalid-url-example2.com",
		"https://invalid-url-example3.com",
		"https://invalid-url-example4.com",
		"https://invalid-url-example5.com",
		"https://invalid-url-example6.com",
		"https://invalid-url-example7.com",
		"https://invalid-url-example8.com",
		"https://invalid-url-example9.com",
		"https://invalid-url-example10.com",
		"https://invalid-url-example11.com",
		"https://invalid-url-example12.com",
		"https://invalid-url-example13.com",
		"https://invalid-url-example14.com",
		"https://invalid-url-example15.com",
		"https://invalid-url-example16.com",
		"https://invalid-url-example17.com",
		"https://invalid-url-example18.com",
		"https://invalid-url-example19.com",
		"https://invalid-url-example20.com",
		"https://invalid-url-example21.com",
		"https://invalid-url-example22.com",
		"https://invalid-url-example23.com",
		"https://invalid-url-example24.com",
		"https://invalid-url-example25.com",
		"https://invalid-url-example26.com",
		"https://invalid-url-example27.com",
		"https://invalid-url-example28.com",
		"https://invalid-url-example29.com",
		"https://invalid-url-example30.com",
		"https://invalid-url-example31.com",
		"https://invalid-url-example32.com",
		"https://invalid-url-example33.com",
		"https://invalid-url-example34.com",
		"https://invalid-url-example35.com",
		"https://invalid-url-example36.com",
		"https://invalid-url-example37.com",
		"https://invalid-url-example38.com",
		"https://invalid-url-example39.com",
		"https://invalid-url-example40.com",
		"https://invalid-url-example41.com",
		"https://invalid-url-example42.com",
		"https://invalid-url-example43.com",
		"https://invalid-url-example44.com",
		"https://invalid-url-example45.com",
		"https://invalid-url-example46.com",
		"https://invalid-url-example47.com",
		"https://invalid-url-example48.com",
		"https://invalid-url-example49.com",
		"https://invalid-url-example50.com",
		"https://invalid-url-example51.com",
		"https://invalid-url-example52.com",
		"https://invalid-url-example53.com",
		"https://invalid-url-example54.com",
		"https://invalid-url-example55.com",
		"https://invalid-url-example56.com",
		"https://invalid-url-example57.com",
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