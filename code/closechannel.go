package main

import (
	"fmt"
	"os"
	"time"
)

var done = make(chan struct{})

func canceled() bool {
	for {
		time.Sleep(1 * time.Second)

		select {
		case <-done:
			fmt.Println("收到退出信号")
			return true
		default:
			return false
		}

	}
}

func main() {

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done) // close channel后继续读取channel的值不阻塞 可以实现广播效果，退出其他goroutine
	}()

	for {
		bo := canceled()
		if bo {
			return
		}
		fmt.Println(bo)
	}
}
