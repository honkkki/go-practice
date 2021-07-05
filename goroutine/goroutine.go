package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello()  {
	fmt.Println("hello golang")
	// 任务完成
	wg.Done()
}

func main()  {
	// 计数
	wg.Add(1)
	go hello()
	fmt.Println("main")
	// 等待子goroutine完成 计数器为0
	wg.Wait()
}
