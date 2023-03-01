package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func inputWriter(inCh *chan string) {
	for i := 0; i < 1000; i++ {
		time.Sleep(10 * time.Millisecond)
		*inCh <- fmt.Sprintf("String %d", i)
	}
}

func send(msg string) {
	fmt.Println(msg)
}

func main() {
	start := time.Now()

	ch := make(chan string)
	go inputWriter(&ch)

	close(ch)
	wg.Wait()
	end := time.Since(start)
	fmt.Printf("Time elapsed: %s", end)
}
