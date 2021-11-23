package main

import (
	"fmt"
	"sync"
)

var (
	lockx sync.Mutex
	xx    int
	yy    int
	w     sync.WaitGroup
)

func main() {
	w.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			lockx.Lock()
			xx++
			yy++
			lockx.Unlock()
			w.Done()
		}()
	}

	w.Wait()
	fmt.Println(xx)
	fmt.Println(yy)

}
