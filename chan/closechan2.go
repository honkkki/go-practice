package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 6) // 有缓冲区的管道
	done := make(chan struct{})

	go func() {
		for {
			select {
			case ch <- "✌":
			case <-done:
				// 不知道什么时候关闭
				close(ch)
				return
			}

		}
	}()

	// 3秒后关闭ch 退出上面的goroutine
	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	for v := range ch {
		fmt.Println("get data: ", v)
	}

	fmt.Println("finish")
}
