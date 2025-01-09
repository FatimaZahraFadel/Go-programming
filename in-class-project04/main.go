package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func fetchURL(ctx context.Context, url string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("failed to create request: %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			fmt.Printf("request canceled or timed out: %v\n", ctx.Err())
		default:
			fmt.Printf("failed to fetch URL: %v\n", err)
		}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v\n", err)
		return
	}

	fmt.Printf("fetched URL: %s\n", url)
	fmt.Printf("response: %s\n", body)
}
func main() {

	urls := []string{
		"http://example.com",
		"http://httpbin.org/delay/5",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fetchURL(ctx, url)
		}(url)
	}

	wg.Wait()

}
