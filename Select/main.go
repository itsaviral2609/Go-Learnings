/*

package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])

	}

	for i := 0; i < 12; i++ {
		select {
		case m0 := <-chans[0]:
			log.Println("Received from m0", m0)
		case m1 := <-chans[1]:
			log.Println("Received from m0", m1)
		}
	}

}

*/

package main

import (
	"log"
	"time"
)

func main() {

	log.Println("start!!!")
	const tickRate = 2 * time.Second
	stopper := time.After(5 * tickRate)
	ticker := time.NewTicker(tickRate).C

loop:
	for {
		select {
		case <-ticker:
			log.Println("tick!")
		case <-stopper:
			break loop
		}
	}
	log.Println("finsih!!!")

}
