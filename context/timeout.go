package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context)  {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			fmt.Println("working")
		}

		time.Sleep(1*time.Second)
	}
}


func main()  {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go doWork(ctx)
	time.Sleep(10*time.Second)
	// 释放资源
	cancel()
}
