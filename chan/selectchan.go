package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		// 类似switch 但会阻塞等待
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("default")
		}
	}
}
