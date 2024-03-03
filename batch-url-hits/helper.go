package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func checkURLs(method string, urls []string, json_data string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go checkURL(method, url, &wg, json_data)
	}

	wg.Wait()
}

func checkURL(method string, url string, wg *sync.WaitGroup, json_data string) {
	defer (*wg).Done()
	var req *http.Request
	var err error

	if method == "POST" {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(json_data)))
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	defer resp.Body.Close()
	fmt.Println("Status Code: ", resp.Status)
}
