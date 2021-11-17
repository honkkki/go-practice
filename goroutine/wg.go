package main

import (
	"fmt"
	"sync"
)

// 同步锁 等待执行完成
var wg1 sync.WaitGroup

func main() {
	wg1.Add(1000)

	for i := 0; i < 1000; i++ {
		go func(x int) {
			fmt.Println("num", x)
			wg1.Done()
		}(i)
	}

	wg1.Wait()
}
