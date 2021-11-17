package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	limit = make(chan struct{}, 3)
)

func worker(i int, wg *sync.WaitGroup) {
	limit <- struct{}{}
	defer func() {
		<-limit
		wg.Done()
	}()
	fmt.Printf("%d working...\n", i)
	time.Sleep(3 * time.Second)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
