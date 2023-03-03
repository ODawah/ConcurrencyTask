package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	workerCount = 10
)

func inputWriter(inCh *chan string) {
	defer close(*inCh)
	for i := 1; i <= 10000; i++ {
		time.Sleep(10 * time.Millisecond)
		*inCh <- fmt.Sprintf("String %d", i)
	}
}

func send(msg string) {
	fmt.Println(msg)
}

func Worker(inCh *chan string, wg *sync.WaitGroup) {
	for msg := range *inCh {
		send(msg)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	ch := make(chan string)
	go inputWriter(&ch)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go Worker(&ch, &wg)
	}

	wg.Wait()
	T := time.Since(start)
	fmt.Printf("Time elapsed: %s", T)
}
