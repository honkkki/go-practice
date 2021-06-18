package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

// 使用context实现goroutine的退出
func main()  {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for  {
			select {
			case <-ctx.Done():
				ch<- struct{}{}
				return
			default:
				fmt.Println("waiting done")
			}

			time.Sleep(300 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3*time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("finish")
}
