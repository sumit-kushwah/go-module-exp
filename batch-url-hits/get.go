package main

import (
	"fmt"
	"time"
)

func main() {
	requests := 205
	start := time.Now()
	var urls []string
	for i := 0; i < requests; i++ {
		urls = append(urls, "https://www.indiamart.com")
	}
	checkURLs("GET", urls, "")
	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}
