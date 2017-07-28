package main

import (
	"net/http"
	"sync"
	"time"
)

func fetchAll(urls []string) {
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			resp, err := http.Get("http://example.com/")
			if (err != nil) {
				println(url, "Error")
			} else {
				println(url, resp.Status, time.Since(start).String())
			}
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	fetchAll([]string{
		"https://www.fastly.com/",
		"https://pwa.fastlydemo.net/",
		"https://www.cloudflare.com/",
	})
}