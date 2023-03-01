package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func inputWriter(inCh chan<- string) {
	for i := 0; i < 1000; i++ {
		time.Sleep(10 * time.Millisecond)
		inCh <- fmt.Sprintf("String %d", i)
	}
	wg.Done()
}

func reader(outCh <-chan string) {
	defer wg.Done()
	for v := range outCh {
		wg.Add(1)
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}
}

func main() {
	start := time.Now()

	ch := make(chan string, 4)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go reader(ch)
	}
	inputWriter(ch)
	close(ch)
	wg.Wait()
	end := time.Since(start)
	fmt.Printf("Time elapsed: %s", end)
}
