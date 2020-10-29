package main

import (
	"fmt"
	"sync"
)

var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex
)

// 互斥锁
func add() {
	for i := 0; i < 50000; i++ {
		// 加锁
		lock.Lock()
		x++
		// 解锁
		lock.Unlock()
	}

	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
