// 使用context控制goroutine退出.
package main

import (
	"context"
	"fmt"
	"time"
)


func worker1(ctx context.Context)  {
	fmt.Println("start working1")
	go worker2(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到退出通知，退出worker1 goroutine")
			return
		default:
			fmt.Println("working1")
		}
	}
}

func worker2(ctx context.Context)  {
	fmt.Println("start working2")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到退出通知，退出worker2 goroutine")
			return
		default:
			fmt.Println("working2")
		}
	}
}

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	go worker1(ctx)

	time.Sleep(3*time.Second)
	cancel()

	time.Sleep(2*time.Second)
	fmt.Println("finish")
}
