package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var testNum int32
var wgp sync.WaitGroup

func addNum() {
	for i := 0; i < 1000; i++ {
		// 原子操作 int 性能比加锁高 加锁需要上下文切换
		atomic.AddInt32(&testNum, 1)
		//testNum++
		time.Sleep(time.Millisecond)
	}

	wgp.Done()
}

func main() {
	wgp.Add(2)
	go addNum()
	go addNum()
	wgp.Wait()

	fmt.Println(testNum)
}
