// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"time"
// )

// type result struct {
// 	url     string
// 	err     error
// 	latency time.Duration
// }

// func get(ctx context.Context, url string, ch chan<- result) {
// 	start := time.Now()
// 	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

// 	if resp, err := http.DefaultClient.Do(req); err != nil {
// 		ch <- result{url, err, 0}
// 	} else {
// 		t := time.Since(start).Round(time.Millisecond)
// 		ch <- result{url, nil, t}
// 		resp.Body.Close()
// 	}
// }

// func main() {
// 	stopper := time.After(3 * time.Second)
// 	results := make(chan result)
// 	list := []string{
// 		"https://amazon.com",
// 		"https://microsoft.com",
// 		"https://google.com",
// 		"https://nytimes.com",
// 		"http://localhost:8080",
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

// 	defer cancel()

// 	for _, url := range list {
// 		go get(ctx, url, results)
// 	}

// 	for range list {

// 		select {
// 		case r := <-results:
// 			if r.err != nil {
// 				log.Printf("%-20s %s\n", r.url, r.err)
// 			} else {
// 				log.Printf("%-20s %s\n", r.url, r.latency)
// 			}
// 		case <-stopper:
// 			log.Fatal("timeout!")

// 		}
// 	}
// }

package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {

	var r result
	start := time.Now()
	ticker := time.NewTicker(1 * time.Second).C
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		r = result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		r = result{url, nil, t}
		resp.Body.Close()
	}
	for {
		select {
		case ch <- r:
			return
		case <-ticker:
			log.Println("tick", r)
		}
	}
}

func first(ctx context.Context, urls []string) (*result, error) {
	results := make(chan result, len(urls)) // Buffer to avoid leaking!

	// context
	ctx, cancel := context.WithCancel(ctx)

	// defer

	defer cancel()

	for _, url := range urls {
		go get(ctx, url, results)
	}
	select {
	case r := <-results:
		return &r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	list := []string{
		"https://amazon.com",
		"https://microsoft.com",
		"https://google.com",
		"https://nytimes.com",
		//"http://localhost:8080",
	}

	r, _ := first(context.Background(), list)
	if r.err != nil {
		log.Printf("%-20s %s\n", r.url, r.err)
	} else {
		log.Printf("%-20s %s\n", r.url, r.latency)
	}

	time.Sleep(9 * time.Millisecond)
	log.Println("quit anyway", runtime.NumGoroutine(), "still running!")

}
