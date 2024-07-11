package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response from %s: %v\n", url, err)
		return
	}

	fmt.Printf("Response from %s: %s\n", url, body[:100]) // Print only the first 100 characters
}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"http://localhost:8080/hello",
		"http://localhost:8080/hello",
		"http://localhost:8080/hello",
	}

	wg.Add(len(urls))

	for _, url := range urls {
		go fetchURL(url, &wg)
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}
