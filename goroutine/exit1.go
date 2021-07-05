// 使用channel控制goroutine退出.
package main

import (
	"fmt"
	"time"
)

func workerG(ch <-chan struct{}) {
	fmt.Println("start working")
	for {
		select {
		case <-ch:
			fmt.Println("收到退出通知，退出worker goroutine")
			return
		default:
			fmt.Println("working")
		}
	}
}

func main() {
	// 缺点 需要维护一个共用的channel
	ch := make(chan struct{})
	go workerG(ch)

	time.Sleep(3 * time.Second)
	ch <- struct{}{}

	fmt.Println("回到主goroutine")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
}
