package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var num int64
	go func() {
		num = 2
	}()

	for {
		b := atomic.CompareAndSwapInt64(&num, 2, 1)
		fmt.Println(b)
		if b {
			break
		}
	}

	fmt.Println(num)
}
