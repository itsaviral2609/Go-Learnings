package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	stopper := time.After(3 * time.Second)
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://microsoft.com",
		"https://google.com",
		"https://nytimes.com",
		"http://localhost:8080",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {

		select {
		case r := <-results:
			if r.err != nil {
				log.Printf("%-20s %s\n", r.url, r.err)
			} else {
				log.Printf("%-20s %s\n", r.url, r.latency)
			}
		case <-stopper:
			log.Fatal("tiemout!")

		}
	}
}

/*

//var nextID = make(chan int)

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1> You got this %d</h1>", <-ch)
}

func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {
	var nextID nextCh = make(chan int)
	go counter(nextID)
	http.HandleFunc("/", nextID.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


*/

// PRIME SEIVE!!

/*
func generator(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i
	}

	close(ch)
}

func filter(src <-chan int, dest chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dest <- i
		}
	}
	close(dest)
}

func seive(limit int) {
	ch := make(chan int)

	go generator(limit, ch)

	for {
		prime, ok := <-ch

		if !ok {
			break
		}

		ch1 := make(chan int)

		go filter(ch, ch1, prime)

		ch = ch1

		fmt.Print(prime, " ")
	}
}

func main() {
	seive(100)
}


*/
