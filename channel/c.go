package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup
var ch = make(chan struct{}, 1)

func add() {
	//ptr := (*int64)(unsafe.Pointer(&x))
	//atomic.AddInt64(ptr, 1)
	// 原子锁
	ch <- struct{}{}
	x++
	wg.Done()
	<-ch
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 1
	close(ch)

	num, ok := <-ch
	fmt.Println(num, ok)

	num, ok = <-ch
	fmt.Println(num, ok)

	num, ok = <-ch
	fmt.Println(num, ok)

	for x := 0; x < 100; x++ {
		wg.Add(1)
		go add()
	}

	wg.Wait()
	fmt.Println(x)

}
